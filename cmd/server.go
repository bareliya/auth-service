package cmd

import (
	"fmt"
	"github.com/auth-service/db"
	"github.com/spf13/cobra"
	"time"
)

var cmdServer = &cobra.Command{
	Use:   "server",
	Short: "To start api server",
	Run: func(cmd *cobra.Command, args []string) {
		db.InitPostgress()
		for i := 1; i < 10000; i++ {
			go func() {
				_, err := db.GetUserCredentialsByUserName(fmt.Sprintf("user%d", i*i))
				if err != nil {
					panic(err)
				}
			}()
		}

		time.Sleep(20 * time.Second)

		//fmt.Printf("%#v", user)
		//fmt.Println("hello world 3")
	},
}

func init() {
	rootCmd.AddCommand(cmdServer)
}
