package realdebrid

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/schollz/progressbar/v3"

	"github.com/cgund98/realdebrid-cli/internal/logging"
)

type Controller struct {
	httpClient *retryablehttp.Client
	grabClient *grab.Client
	apiToken   string
	baseUrl    string
}

func NewController(apiToken string, baseUrl string) *Controller {
	httpClient := &retryablehttp.Client{
		HTTPClient: cleanhttp.DefaultPooledClient(),
		Logger:     nil, // Set to nil to keep from writing debug statements
		// Logger:       retryablehttp.NewClient().Logger,
		RetryWaitMin: 1 * time.Second,
		RetryWaitMax: 30 * time.Second,
		RetryMax:     4,
		CheckRetry:   retryablehttp.DefaultRetryPolicy,
		Backoff:      retryablehttp.DefaultBackoff,
	}
	grabClient := grab.NewClient()

	return &Controller{
		httpClient,
		grabClient,
		apiToken,
		baseUrl,
	}
}

// request will send an HTTP request to the RealDebrid API
func (c *Controller) request(method, endpoint string, body any, params ...string) *http.Response {

	// serialize body
	bodyReader := serializeBody(body)

	req, err := retryablehttp.NewRequest(method, endpoint, bodyReader)
	if err != nil {
		logging.Fatalf("client: could not create request: %s\n", err)
	}

	// Append query parameters
	if len(params)%2 != 0 {
		logging.Fatalf("params argument must have an even length.")
	}

	q := req.URL.Query()
	i := 0
	for i < len(params) {
		key := params[i]
		value := params[i+1]
		q.Add(key, value)
		i += 2
	}
	req.URL.RawQuery = q.Encode()

	// Set authorization header with API Token
	authHeader := fmt.Sprintf("Bearer %s", c.apiToken)
	req.Header.Set("Authorization", authHeader)

	// Send request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		logging.Fatalf("error occured while making request to '%s': %v", endpoint, err)
	}

	// Handle error response
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		handleStatusError(resp)
	}

	return resp
}

func (c *Controller) download(url, path string) {
	// Create request
	req, err := grab.NewRequest(path, url)
	if err != nil {
		logging.Fatalf("unable to create grab request: %v", err)
	}

	// Start download
	resp := c.grabClient.Do(req)

	// Handle error response
	if resp.HTTPResponse.StatusCode != http.StatusOK {
		if resp.HTTPResponse.StatusCode == http.StatusPartialContent {
			fmt.Printf("Error code 206 encountered. Deleting file and trying again...\n")
			if err := os.Remove(resp.Filename); err != nil {
				logging.Fatalf("unable to remove file: %v", err)
			}
			c.download(url, path)
			return
		}
		handleStatusError(resp.HTTPResponse)
	}

	bar := progressbar.NewOptions(int(resp.Size()),
		progressbar.OptionSetDescription(fmt.Sprintf("Downloading %s...", filepath.Base(resp.Filename))),
		progressbar.OptionShowBytes(true),
	)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			bar.Set(int(resp.BytesComplete()))

		// download is complete
		case <-resp.Done:
			bar.Set(int(resp.BytesComplete()))
			fmt.Println()
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		logging.Fatalf("download failed: %v", err)
	}

}

func (c *Controller) LinkCheck(link string) *LinkCheckResponse {
	// Make request
	endpoint := fmt.Sprintf("%s/unrestrict/check", c.baseUrl)
	resp := c.request(http.MethodPost, endpoint, LinkCheckRequest{link})

	// Parse response
	var response LinkCheckResponse
	err := parseResponse(resp, &response)
	if err != nil {
		logging.Fatalf("unable to parse response: %v", err)
	}

	return &response
}

func (c *Controller) linkUnrestrict(link string) *LinkUnrestrictResponse {
	// Make request
	endpoint := fmt.Sprintf("%s/unrestrict/link", c.baseUrl)
	resp := c.request(http.MethodPost, endpoint, LinkUnrestrictRequest{link})

	// Parse response
	var response LinkUnrestrictResponse
	err := parseResponse(resp, &response)
	if err != nil {
		logging.Fatalf("unable to parse response: %v", err)
	}

	return &response
}

func (c *Controller) LinkDownload(link, outputPath string) {
	// Unrestrict link
	uResp := c.linkUnrestrict(link)

	// Download file
	c.download(uResp.Download, outputPath)
}

func (c *Controller) FolderUnrestrict(link string) FolderUnrestrictResponse {
	// Make request
	endpoint := fmt.Sprintf("%s/unrestrict/folder", c.baseUrl)
	resp := c.request(http.MethodPost, endpoint, FolderUnrestrictRequest{link})

	// Parse response
	response := FolderUnrestrictResponse{}
	err := parseResponse(resp, &response)
	if err != nil {
		logging.Fatalf("unable to parse response: %v", err)
	}

	return response
}

func (c *Controller) DownloadsList() DownloadsListResponse {
	// Make request
	endpoint := fmt.Sprintf("%s/downloads", c.baseUrl)

	response := DownloadsListResponse{}
	page := 0
	for {
		page += 1
		resp := c.request(http.MethodGet, endpoint, nil, "page", fmt.Sprint(page))

		// Parse response
		newPage := DownloadsListResponse{}
		err := parseResponse(resp, &newPage)
		if err != nil {
			logging.Fatalf("unable to parse response: %v", err)
		}

		if len(newPage) == 0 {
			break
		}

		response = append(response, newPage...)
	}

	return response
}

func (c *Controller) DownloadDelete(id string) {
	// Make request
	endpoint := fmt.Sprintf("%s/downloads/delete/%s", c.baseUrl, id)
	c.request(http.MethodDelete, endpoint, nil)
}
