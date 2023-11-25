package downloads

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"

	initialize "github.com/cgund98/realdebrid-cli/internal/init"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all downloads.",
	Run:   runListCmd,
}

func runListCmd(cmd *cobra.Command, args []string) {

	c := initialize.RealDebridController()

	response := c.DownloadsList()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "File Name", "File Size", "Generated At"})

	for _, d := range response {
		t.AppendRow(table.Row{d.Id, d.Filename, d.FileSize, d.Generated})
	}

	t.Render()
}
