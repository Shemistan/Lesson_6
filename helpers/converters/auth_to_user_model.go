package converters

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/Shemistan/Lesson_6/api/dtos"
	"github.com/Shemistan/Lesson_6/models"
)

func AuthToUserModel(dto dtos.AuthRequest) models.User {
	return models.User{
		Login:    dto.Login,
		Password: hashPassword(dto.Password),
		Status:   models.StatusActive,
		Role:     models.RoleUser,
	}
}

func hashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	hashB := hash.Sum(nil)

	return hex.EncodeToString(hashB)
}
