package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "auth-service",
	Short: "basic authentication service to provide login and registration functionality",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello world")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
