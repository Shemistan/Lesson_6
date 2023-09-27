package converters

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/Shemistan/Lesson_6/internal/models"
)

const (
	defaultStatus = "Activate"
	defaultRole   = "User"
)

func ApiToServiceModel(req *models.AddRequest) *models.User {
	hash := sha256.New()
	hash.Write([]byte(req.AuthParams.Password))
	hashB := hash.Sum(nil)
	hashPass := hex.EncodeToString(hashB)

	today := time.Now()
	cur_date := fmt.Sprintf("%d-%d-%d", today.Year(), today.Month(), today.Day())

	res := &models.User{
		Login:            req.AuthParams.Login,
		Password:         hashPass,
		Name:             req.Data.Name,
		Surname:          req.Data.Surname,
		Status:           defaultStatus,
		Role:             defaultRole,
		RegistrationDate: cur_date,
		UpdateDate:       cur_date,
	}

	return res
}

func GetServiceToApiModel(user *models.User) *models.GetUserData {
	res := &models.GetUserData{
		Id:               user.Id,
		Login:            user.Login,
		Name:             user.Name,
		Surname:          user.Surname,
		Status:           user.Status,
		Role:             user.Role,
		RegistrationDate: user.RegistrationDate,
		UpdateDate:       user.UpdateDate,
	}

	return res
}
