package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/auth-service/types"
	"github.com/auth-service/utils"
	"log"
	"strings"
	"time"
)

func GetUserCredentialsByUserName(userName string) (*types.UserCred, error) {
	t := time.Now()
	query := "SELECT user_id, username, first_name, last_name, hash_id, access_token FROM user_credentials WHERE username = $1"
	user := types.UserCred{}
	row := DBPool.QueryRow(context.Background(), query, userName)

	err := row.Scan(&user.UserId, &user.UserName, &user.FirstName, &user.LastName, &user.Hash, &user.AccessToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with username %s not found", userName)
		}
		return nil, err
	}
	fmt.Println(time.Since(t))

	return &user, nil

}

func PopulateDB(totalUsers int) {
	startTime := time.Now()

	batchSize := 1000 // Number of users to insert in each batch
	for i := 0; i < totalUsers; i += batchSize {
		var values []string
		for j := 0; j < batchSize; j++ {
			username := fmt.Sprintf("user%d", i+j+1)
			firstName := fmt.Sprintf("fname%d", i+j+1)
			lastName := fmt.Sprintf("lname%d", i+j+1)
			password := username + "password"
			hashedPassword := utils.GetHashFromPassword(password)
			accessToken := utils.GenerateRandomPassword(50)

			values = append(values, fmt.Sprintf("('%s', '%s', '%s', '%s', '%s')", username, firstName, lastName, hashedPassword, accessToken))
		}

		query := fmt.Sprintf("INSERT INTO user_credentials (username, first_name, last_name, hash_id, access_token) VALUES %s", strings.Join(values, ","))
		_, err := DBPool.Exec(context.Background(), query)
		if err != nil {
			log.Fatal("Error inserting data: ", err)
		}
		fmt.Println(i)
	}
	elapsedTime := time.Since(startTime)
	fmt.Printf("Inserted %d users in %v\n", totalUsers, elapsedTime)
}
