package api

import (
	"errors"
	"github.com/auth-service/db"
	"github.com/auth-service/types"
	"github.com/auth-service/utils"
	"github.com/rs/zerolog/log"
)

func RegisterNewAdmin(admin types.AdminCredential) error {
	if admin.AdminUser == types.EMPTY_STRING {
		return errors.New("please provide admin user name")
	}
	if admin.AdminPassword == types.EMPTY_STRING {
		return errors.New("please create a password")
	}

	if len(admin.AdminPassword) < 8 {
		return errors.New("password is too short, should have at least 8 character")
	}
	if len(admin.AdminPassword) > 16 {
		return errors.New("password is too long, max 16 char")
	}

	exist, err := db.IsAdminExists(admin.AdminUser)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("admin already exist please login instead")
	}

	passHash := utils.GetHashFromPassword(admin.AdminPassword)
	accessToken := utils.GenerateRandomPassword(16)

	admin.IsSuperAdmin = false
	admin.IsApproved = false
	err = db.CreateNewAdmin(admin, passHash, accessToken)
	return err
}

func AdminLogin(admin types.AdminCredential) (string, error) {
	if admin.AdminUser == types.EMPTY_STRING {
		return "", errors.New("please provide admin user name")
	}
	if admin.AdminPassword == types.EMPTY_STRING {
		return "", errors.New("please provide the password")
	}
	adminCred, err := db.GetAdminByUsername(admin.AdminUser)
	if err != nil {
		return "", err
	}

	ok := utils.CheckPassword(admin.AdminPassword, adminCred.AdminPasswordHash)
	if ok {
		if adminCred.IsApproved {
			return adminCred.AccessToken, nil
		} else {
			return "", errors.New("admin access not approved, please ask super admin to approve")
		}
	}

	return "", errors.New("invalid password")
}

func ApproveAdmin(admin_user string) error {
	if admin_user == types.EMPTY_STRING {
		return errors.New("please provide admin user name")
	}

	exist, err := db.IsAdminExists(admin_user)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("no admin found with given user name " + admin_user)
	}

	approve := true
	err = db.ApproveAdmin(admin_user, approve)
	return err

}
func IsAuthorisedAdminRequest(admin_user, access_token string, superAdmin bool) bool {
	if admin_user == "" || access_token == "" {
		return false
	}
	adminCred, err := db.GetAdminByUsername(admin_user)
	if err != nil {
		log.Err(err).Msgf("admin credential not found")
		return false
	}
	if adminCred.AccessToken == access_token && adminCred.IsApproved {
		if superAdmin {
			if adminCred.IsSuperAdmin {
				return true
			} else {
				return false
			}
		}

		return true
	}

	return false

}
