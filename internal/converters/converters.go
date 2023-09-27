package converters

import (
	"crypto/sha256"
	"encoding/hex"
	"golang-api/internal/models"
)

func ApiUserModelToServiceModel(req models.AddUserRequest) models.User {
	hash := sha256.New()
	hash.Write([]byte(req.Password))
	hashB := hash.Sum(nil)
	hashPassword := hex.EncodeToString(hashB)

	res := models.User{
		Login:        req.Login,
		PasswordHash: hashPassword,
		Name:         req.Name,
		Surname:      req.Surname,
	}
	return res
}

func ApiUpdateUserModelToServiceModel(req models.UpdateUserRequest) models.User {
	hash := sha256.New()
	hash.Write([]byte(req.Password))
	hashB := hash.Sum(nil)
	hashPassword := hex.EncodeToString(hashB)

	res := models.User{
		Login:        req.Login,
		PasswordHash: hashPassword,
		Name:         req.Name,
		Surname:      req.Surname,
	}
	return res
}
