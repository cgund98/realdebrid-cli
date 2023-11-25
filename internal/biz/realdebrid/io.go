package realdebrid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/cgund98/realdebrid-cli/internal/logging"
	"github.com/fatih/structs"
)

/*
 * Helper functions
 */

// serializeBody will serialize a struct into a valid HTTP url key-set
func serializeBody(body any) *bytes.Buffer {
	bodyReader := bytes.NewBuffer([]byte(""))
	if body != nil {
		m := structs.Map(body)

		values := url.Values{}
		for key, value := range m {
			values.Set(key, fmt.Sprintf("%v", value))
		}

		bodyStr := values.Encode()
		bodyReader = bytes.NewBuffer([]byte(bodyStr))
	}

	return bodyReader
}

// handleStatusError will log the response and exit the application
func handleStatusError(resp *http.Response) {
	fmt.Printf("Received response with error status code: %d\n", resp.StatusCode)
	if resp.StatusCode == http.StatusPartialContent {
		logging.Fatalf("please delete destination file and try again.")
	}

	respStr, err := io.ReadAll(resp.Body)
	if err != nil {
		logging.Fatalf("unable to parse response body: %v", err)
	}

	logging.Fatalf("%s", respStr)
}

// parseResponse will parse an HTTP response into a valid struct
func parseResponse(response *http.Response, v any) error {
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("io.ReadAll: %v", err)
	}

	// Response is empty
	if len(data) == 0 {
		return nil
	}

	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("json.Unmarshal: %v", err)
	}

	return nil
}
