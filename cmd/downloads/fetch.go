package downloads

import (
	"errors"
	"fmt"
	"os"

	initialize "github.com/cgund98/realdebrid-cli/internal/init"
	"github.com/cgund98/realdebrid-cli/internal/logging"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Download a file from a restricted link.",
	Run:   runCheckCmd,
}

func createOutputDirIfNotExists(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			logging.Fatalf("error creating output directory: %v", err)
		}
	}
}

func runFetchCmd(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		logging.Fatalf("must provide a download url")
	}

	outputPath := viper.GetString("downloads.output_path")
	fmt.Printf("Saving files to %s...\n", outputPath)
	createOutputDirIfNotExists(outputPath)

	c := initialize.RealDebridController()

	// Unpack folders
	links := []string{}
	for _, arg := range args {
		sublinks := c.FolderUnrestrict(arg)

		if len(sublinks) == 0 {
			links = append(links, arg)
		} else {
			links = append(links, sublinks...)
		}
	}

	// Download files
	for _, link := range links {
		c.LinkDownload(link, outputPath)
	}
}
