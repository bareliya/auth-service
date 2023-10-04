package api

import (
	"errors"
	"github.com/auth-service/db"
	"github.com/auth-service/types"
	"github.com/auth-service/utils"
)

func UserLogin(user types.UserRegister) (string, error) {
	if user.UserName == types.EMPTY_STRING {
		return "", errors.New("please provide user name")
	}
	if user.Password == types.EMPTY_STRING {
		return "", errors.New("please provide the password")
	}
	userCred, err := db.GetUserCredentialsByUserName(user.UserName)
	if err != nil {
		return "", err
	}

	ok := utils.CheckPassword(user.Password, userCred.Hash)
	if ok {
		return userCred.AccessToken, nil
	}

	return "", errors.New("invalid password")
}

func RegisterNewUser(user types.UserRegister) error {
	if user.FirstName == types.EMPTY_STRING {
		return errors.New("please provide first name")
	}
	if user.LastName == types.EMPTY_STRING {
		return errors.New("please provide last name")
	}
	if user.UserName == types.EMPTY_STRING {
		return errors.New("please provide user name")
	}
	if user.Password == types.EMPTY_STRING {
		return errors.New("please create a password")
	}

	if len(user.Password) < 8 {
		return errors.New("password is too short, should have at least 8 character")
	}
	if len(user.Password) > 16 {
		return errors.New("password is too long, max 16 char")
	}

	exist, err := db.IsUserExists(user.UserName)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("user already exist please login instead")
	}

	passHash := utils.GetHashFromPassword(user.Password)
	accessToken := utils.GenerateRandomPassword(16)

	err = db.CreateNewUSer(user, passHash, accessToken)
	return err

}
