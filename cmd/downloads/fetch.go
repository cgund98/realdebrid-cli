package downloads

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"

	initialize "github.com/cgund98/realdebrid-cli/internal/init"
	"github.com/cgund98/realdebrid-cli/internal/logging"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Download a file from a restricted link.",
	Run:   runFetchCmd,
}

func createOutputDirIfNotExists(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			logging.Fatalf("error creating output directory: %v", err)
		}
	}
}

func isUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func runFetchCmd(cmd *cobra.Command, args []string) {
	listFile := viper.GetString("downloads.link_file")

	if listFile == "" && len(args) < 1 {
		logging.Fatalf("must provide a download url")
	}

	outputPath := viper.GetString("downloads.output_path")
	fmt.Printf("Saving files to %s...\n", outputPath)
	createOutputDirIfNotExists(outputPath)

	c := initialize.RealDebridController()

	// Parse links file
	inputLinks := args
	if listFile != "" {
		content, err := os.ReadFile(listFile)
		if err != nil {
			logging.Fatalf("unable to parse links file: %v", err)
		}
		contentSplit := strings.Fields(string(content))
		for _, link := range contentSplit {
			if isUrl(link) {
				inputLinks = append(inputLinks, link)
			}
		}

		if len(inputLinks) > 0 {
			fmt.Printf("Parsed %d links from input file.\n", len(inputLinks))
		}
	}

	// Parse folders
	skipFolders := viper.GetBool("downloads.skip_folders")
	links := []string{}
	if skipFolders {
		fmt.Println("Skipping folder checks...")
		links = inputLinks
	} else {
		fmt.Println("Checking if links are folders...")
		for _, inputLink := range inputLinks {
			sublinks := c.FolderUnrestrict(inputLink)

			if len(sublinks) == 0 {
				links = append(links, inputLink)
			} else {
				links = append(links, sublinks...)
			}
		}
	}

	// Download files
	fmt.Println("Downloading links...")
	for _, link := range links {
		c.LinkDownload(link, outputPath)
	}
}
