package cmd

import (
	"github.com/auth-service/api"
	"github.com/auth-service/db"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var cmdServer = &cobra.Command{
	Use:   "server",
	Short: "To start api server",
	Run: func(cmd *cobra.Command, args []string) {
		db.InitPostgress()
		go api.StartServer()
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		// Select on error channels from different modules
		for {
			select {
			case sig := <-sigs:
				log.Info().Msgf("Got signal, beginning shutdown", "signal", sig)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cmdServer)
}
