package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var cmddummy = &cobra.Command{
	Use:   "dummy",
	Short: "To populate 5 million dummy user data in db",
	Run: func(cmd *cobra.Command, args []string) {

		// Open the file for writing
		file, err := os.Create("./load-testing/users_data.json")
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		// Create a slice to hold 5 million users
		var users []User

		// Create 5 million users and add them to the slice
		for i := 1; i <= 10000; i++ {
			user := User{
				UserName: fmt.Sprintf("user%d", i),
				Password: fmt.Sprintf("user%dpassword", i),
			}
			users = append(users, user)
		}

		// Convert the slice to JSON
		usersJSON, err := json.Marshal(users)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}

		// Write JSON data to the file
		_, err = file.Write(usersJSON)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		fmt.Println("Users data has been generated and written to users_data.json.")
	},
}

type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func init() {
	rootCmd.AddCommand(cmddummy)
}
