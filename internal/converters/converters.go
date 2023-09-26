package converters

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"golang-api/internal/models"
)

func ApiModelToServiceModel(req models.Request) *models.User {
	hash := sha256.New()
	hash.Write([]byte(req.AuthParams.Password))
	hashB := hash.Sum(nil)
	hashedPassword := hex.EncodeToString(hashB)
	
	res := &models.User{
		Login: req.AuthParams.Login,
		HashedPassword: hashedPassword,
		RegistrationDate: time.Now().String(),
	}

	return res
}

func UserUpdate(req models.UserUpdateRequest) *models.User {
	res := &models.User{
		Firstname: req.Firstname,
		Lastname: req.Lastname,
	}

	return res
}