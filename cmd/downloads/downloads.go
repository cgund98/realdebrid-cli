package downloads

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagOutputPath = "output"
)

var DownloadCmd = &cobra.Command{
	Use:   "downloads",
	Short: "Manage downloads",
	Run:   runDownloadCmd,
}

func AddDownloadChildren() {
	DownloadCmd.AddCommand(checkCmd, fetchCmd, listCmd, cleanCmd)

	fetchCmd.PersistentFlags().StringP(flagOutputPath, "o", "", "Real Debrid API token. https://real-debrid.com/apitoken")
	viper.BindPFlag("downloads.output_path", fetchCmd.PersistentFlags().Lookup(flagOutputPath))
}

func runDownloadCmd(cmd *cobra.Command, args []string) {
	cmd.Help()
	os.Exit(0)
}
