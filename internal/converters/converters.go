package converters

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"golang-api/internal/models"
)

func ApiAuthModelToServiceUserModel(req models.AddRequest) *models.User{
	hash := sha256.New()
	hash.Write([]byte(req.AuthParams.Password))
	hashB := hash.Sum(nil)
	hashPassword := hex.EncodeToString(hashB)
	res := &models.User{
		Login: req.AuthParams.Login,
		HashPassword: hashPassword,
		RegistrationDate: time.Now(),
	}

	return res
}
