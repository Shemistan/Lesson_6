package converters

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/Shemistan/Lesson_6/internal/models"
)

func ApiModelToServiceModel(apiModel models.Account) *models.User {
	hash := sha256.New()
	hash.Write([]byte(apiModel.AuthData.Password))
	hashB := hash.Sum(nil)
	hashPassword := hex.EncodeToString(hashB)
	res := &models.User{
		Login:    apiModel.AuthData.Login,
		Password: hashPassword,
		Name:     apiModel.UserInfo.Name,
		Surname:  apiModel.UserInfo.Surname,
		Active:   apiModel.UserInfo.Active,
		Role:     apiModel.UserInfo.Role,
	}
	return res
}
