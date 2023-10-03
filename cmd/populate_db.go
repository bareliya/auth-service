package cmd

import (
	"fmt"
	"github.com/auth-service/db"
	"github.com/spf13/cobra"
)

var cmdDb = &cobra.Command{
	Use:   "db",
	Short: "To populate 500 million dummy user data in db",
	Run: func(cmd *cobra.Command, args []string) {
		db.InitPostgress()
		db.PopulateDB(5000000)
		fmt.Println("hello world 2")
	},
}

func init() {
	rootCmd.AddCommand(cmdDb)
}
