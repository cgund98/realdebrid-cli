package downloads

import (
	"fmt"

	initialize "github.com/cgund98/realdebrid-cli/internal/init"
	"github.com/schollz/progressbar/v3"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Delete all downloads.",
	Run:   runCleanCmd,
}

func runCleanCmd(cmd *cobra.Command, args []string) {

	c := initialize.RealDebridController()

	downloads := c.DownloadsList()

	bar := progressbar.NewOptions(len(downloads),
		progressbar.OptionSetDescription("Deleting downloads..."))

	for _, d := range downloads {
		bar.Describe(fmt.Sprintf("Deleting download %s...", d.Id))
		c.DownloadDelete(d.Id)
		bar.Add(1)
	}
}
