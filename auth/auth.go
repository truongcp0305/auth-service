package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"math/rand"
	"os"
	"time"

	"auth-service/model"

	"github.com/golang-jwt/jwt"
)

type claims struct {
	Email       string
	DisplayName string
	jwt.StandardClaims
}

func GetPrivate() []byte {
	file, err := os.Open("crypt/private.pem")
	if err != nil {
		return nil
	}
	defer file.Close()
	fileByte, _ := io.ReadAll(file)
	return fileByte
}

func CreateJwt(model model.User) (string, error) {
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &claims{
		Email:       model.UserName,
		DisplayName: model.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	keyBys := GetPrivate()
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyBys)
	if err != nil {
		return "", err
	}
	return token.SignedString(privateKey)
}

func HashString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func GenerateUUID() string {
	// Sử dụng thời gian hiện tại để tạo một seed ngẫu nhiên
	rand.Seed(time.Now().UnixNano())

	// Tạo một chuỗi ngẫu nhiên có độ dài 32 ký tự
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	uuid := make([]byte, 32)
	for i := range uuid {
		uuid[i] = charset[rand.Intn(len(charset))]
	}

	// Thêm một dấu gạch ngang vào vị trí thứ 8 và 13 để tạo định dạng UUID
	uuid[8] = '-'
	uuid[13] = '-'

	return string(uuid)
}
