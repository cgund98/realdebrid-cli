package downloads

import (
	"fmt"

	initialize "github.com/cgund98/realdebrid-cli/internal/init"
	"github.com/cgund98/realdebrid-cli/internal/logging"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagOutputPath = "output"
)

var DownloadCmd = &cobra.Command{
	Use:   "downloads",
	Short: "Download a file from a restricted link.",
	Run:   runDownloadCmd,
}

func AddDownloadChildren() {
	DownloadCmd.AddCommand(checkCmd)

	DownloadCmd.PersistentFlags().StringP(flagOutputPath, "o", "", "Real Debrid API token. https://real-debrid.com/apitoken")
	viper.BindPFlag("downloads.output_path", DownloadCmd.PersistentFlags().Lookup(flagOutputPath))
}

func runDownloadCmd(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		logging.Fatalf("must provide a download url")
	}

	outputPath := viper.GetString("downloads.output_path")
	fmt.Printf("Saving files to %s...\n", outputPath)

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

/*
go run ./cmd downloads \
	https://rg.to/file/7a4edfd102de0139c5955878cec31dd9/ArtStationEnvironmentDesignGraphicSketchingwithGradyFrederick.part1.rar.html \
    -o ~/Downloads/rbt/cowboy

*/
