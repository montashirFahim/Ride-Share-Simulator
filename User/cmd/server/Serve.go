package server

import (
	"User/internal/setup"

	"github.com/spf13/cobra"
)

var (
	configPath string
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start User-app Server",
	Long:  `Start the chi-based User-app server with postgreSQL, Redis and basic authentication.`,
	RunE:  setup.RunServer,
}

func init() {
	ServerCmd.Flags().StringVar(&configPath, "config", ".", "Path to config file")
}
