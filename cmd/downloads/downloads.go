package downloads

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagOutputPath  = "output"
	flagFile        = "file"
	flagSkipFolders = "skip-folders"
)

var DownloadCmd = &cobra.Command{
	Use:   "downloads",
	Short: "Manage downloads",
	Run:   runDownloadCmd,
}

func AddDownloadChildren() {
	DownloadCmd.AddCommand(checkCmd, fetchCmd, listCmd, cleanCmd)

	fetchCmd.PersistentFlags().StringP(flagOutputPath, "o", "", "Path to write files to.")
	viper.BindPFlag("downloads.output_path", fetchCmd.PersistentFlags().Lookup(flagOutputPath))
	fetchCmd.PersistentFlags().StringP(flagFile, "f", "", "File containing list of links")
	viper.BindPFlag("downloads.link_file", fetchCmd.PersistentFlags().Lookup(flagFile))
	fetchCmd.PersistentFlags().Bool(flagSkipFolders, false, "Skip the parsing of folders")
	viper.BindPFlag("downloads.skip_folders", fetchCmd.PersistentFlags().Lookup(flagSkipFolders))
}

func runDownloadCmd(cmd *cobra.Command, args []string) {
	cmd.Help()
	os.Exit(0)
}
