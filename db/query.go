package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/auth-service/types"
	"github.com/auth-service/utils"
	"github.com/rs/zerolog/log"
	"strings"
	"time"
)

func IsUserExists(username string) (bool, error) {
	var exists bool
	err := DBPool.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM user_credentials WHERE username = $1)", username).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking if user exists: %v", err)
	}

	return exists, nil
}
func IsAdminExists(admin string) (bool, error) {
	var exists bool
	err := DBPool.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM admin_credentials WHERE admin_user = $1)", admin).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking if admin exists: %v", err)
	}

	return exists, nil
}

func CreateNewUSer(user types.UserRegister, hash, access string) error {
	query := `INSERT INTO user_credentials (first_name, last_name, username, hash_id, access_token)
	VALUES ($1, $2, $3, $4, $5)`
	_, err := DBPool.Exec(context.Background(), query, user.FirstName, user.LastName, user.UserName, hash, access)
	if err != nil {
		return err
	}

	return nil
}
func CreateNewAdmin(admin types.AdminCredential, hash, access string) error {
	query := `INSERT INTO admin_credentials (admin_user, admin_password_hash, is_super_admin, access_token, is_approved)
	VALUES ($1, $2, $3, $4, $5)`
	_, err := DBPool.Exec(context.Background(), query, admin.AdminUser, hash, admin.IsSuperAdmin, access, admin.IsApproved)
	if err != nil {
		return fmt.Errorf("error creating new admin: %v", err)
	}

	return nil
}

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
	log.Info().Msgf("GetUserCredentialsByUserName tool total time %v", time.Since(t))
	return &user, nil

}

func GetUsersWithPagination(limit, page int) ([]*types.UserCred, error) {
	t := time.Now()
	offset := (page - 1) * limit

	query := "SELECT user_id, username, first_name, last_name, hash_id, access_token FROM user_credentials LIMIT $1 OFFSET $2"
	rows, err := DBPool.Query(context.Background(), query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error querying users: %v", err)
	}
	defer rows.Close()

	var users []*types.UserCred
	for rows.Next() {
		var user types.UserCred
		err := rows.Scan(&user.UserId, &user.UserName, &user.FirstName, &user.LastName, &user.Hash, &user.AccessToken)
		if err != nil {
			return nil, fmt.Errorf("error scanning user row: %v", err)
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating user rows: %v", err)
	}

	log.Info().Msgf("GetUsersWithPagination took time %v", time.Since(t))
	return users, nil
}

func GetAdminByUsername(userName string) (*types.AdminCredential, error) {
	t := time.Now()

	query := "SELECT admin_id, admin_user, admin_password_hash, is_super_admin, access_token, is_approved FROM admin_credentials WHERE admin_user = $1"
	admin := types.AdminCredential{}
	row := DBPool.QueryRow(context.Background(), query, userName)

	err := row.Scan(&admin.AdminID, &admin.AdminUser, &admin.AdminPasswordHash, &admin.IsSuperAdmin, &admin.AccessToken, &admin.IsApproved)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("admin with username %s not found", userName)
		}
		return nil, err
	}

	log.Info().Msgf("GetAdminByUsername took time %v", time.Since(t))

	return &admin, nil
}

func ApproveAdmin(adminUser string, isApproved bool) error {
	query := "UPDATE admin_credentials SET is_approved = $1 WHERE admin_user = $2"
	_, err := DBPool.Exec(context.Background(), query, isApproved, adminUser)
	if err != nil {
		return err
	}
	return nil
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
			panic("Error inserting data: " + err.Error())
		}
		fmt.Println(i)
	}
	elapsedTime := time.Since(startTime)
	fmt.Printf("Inserted %d users in %v\n", totalUsers, elapsedTime)
}
