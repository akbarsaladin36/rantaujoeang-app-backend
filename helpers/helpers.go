package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	mathrand "math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"golang.org/x/crypto/bcrypt"
)

func GenerateUUID(data string) (newUUIDString string) {
	namespace := uuid.NameSpaceDNS
	customUUID := uuid.NewSHA1(namespace, []byte(data))
	uuidString := customUUID.String()

	return uuidString
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPassword(user_password string, input_password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user_password), []byte(input_password))

	if err != nil {
		fmt.Println("Your password is not matched! Please try again!")
	}

	return err
}

func GenerateSessionToken(length int) string {
	token := make([]byte, length)

	_, err := rand.Read(token)
	if err != nil {
		return ""
	}

	customToken := base64.URLEncoding.EncodeToString(token)

	return customToken
}

func GenerateCode(processCode string) string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02")
	randomInt := strconv.Itoa(mathrand.Intn(100000))
	dateTime := strings.ReplaceAll(formattedTime, "-", "")
	if processCode == "top-up" {
		topUpCode := "TU"
		newProcessCode := fmt.Sprintf("%s-%s-%s", topUpCode, dateTime, randomInt)
		return newProcessCode
	} else if processCode == "transfer" {
		transferCode := "TF"
		newProcessCode := fmt.Sprintf("%s-%s-%s", transferCode, dateTime, randomInt)
		return newProcessCode
	} else {
		transferCode := strings.ToUpper(processCode)
		newProcessCode := fmt.Sprintf("%s-%s-%s", transferCode, dateTime, randomInt)
		return newProcessCode
	}
}

func GenerateSlug(input string) string {
	slug := slug.Make(input)
	return slug
}
