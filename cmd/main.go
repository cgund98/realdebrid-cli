package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cgund98/realdebrid-cli/cmd/downloads"
	"github.com/cgund98/realdebrid-cli/internal/logging"
)

const (
	flagToken  = "token"
	apiBaseUrl = "https://api.real-debrid.com/rest/1.0"
)

func init() {
	cobra.OnInitialize(initViper)

	// Commands
	downloads.AddDownloadChildren()
	rootCmd.AddCommand(downloads.DownloadCmd)

	// Flags
	rootCmd.PersistentFlags().String(flagToken, "", "Real Debrid API token. https://real-debrid.com/apitoken")
	viper.BindPFlag("api.token", rootCmd.PersistentFlags().Lookup(flagToken))
}

func initViper() {
	if err := viper.BindEnv("api.token", "REAL_DEBRID_API_TOKEN"); err != nil {
		log.Fatalf("viper.BindEnv: %v", err)
	}

	viper.SetDefault("api.base_url", apiBaseUrl)
	viper.SetDefault("api.base_url", apiBaseUrl)

	homedir, err := os.UserHomeDir()
	if err != nil {
		logging.Fatalf("unable to set home directory: %v", err)
	}
	viper.SetDefault("downloads.output_path", filepath.Join(homedir, "Downloads"))
}

func main() {
	rootCmd.Execute()
}
