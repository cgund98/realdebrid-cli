package downloads

import (
	"fmt"
	"math"

	initialize "github.com/cgund98/realdebrid-cli/internal/init"
	"github.com/cgund98/realdebrid-cli/internal/logging"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Validate a download link.",
	Run:   runCheckCmd,
}

func runCheckCmd(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		logging.Fatalf("must provide a download url")
	}
	link := args[0]

	c := initialize.RealDebridController()

	response := c.LinkCheck(link)

	fmt.Printf("Host: %s\n", response.Host)
	fmt.Printf("Link: %s\n", response.Link)
	fmt.Printf("File name: %s\n", response.Filename)
	fmt.Printf("File size: %.1f MB\n", float64(response.FileSize)/math.Pow(1024, 2))
	fmt.Printf("Supported: %d\n", response.Supported)
}
