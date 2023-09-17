package converters

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/Shemistan/Lesson_6/internal/models"
)

func ApiAuthModelToServiceUserModel(req models.AddRequest) *models.User{
	hash := sha256.New()
	hash.Write([]byte(req.AuthParams.Password))
	hashB := hash.Sum(nil)
	hashPassword := hex.EncodeToString(hashB)
	res := &models.User{
		Login: req.AuthParams.Login,
		HashPassword: hashPassword,
		Name:             req.Date.Name,
		Surname:          req.Date.Surname,
		Status:           req.Date.Status,
		Role:             req.Date.Role,
		RegistrationDate: req.Date.RegistrationDate,
		UpdateDate:       req.Date.UpdateDate,
	}

	return res
}
