package main

import (
	"TlacuachesAsesinos/constants"
	"TlacuachesAsesinos/database"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"net/http"
	_ "net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

var (
	encText string
	bytes   = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
)

const (
	key   = "JHZ697heza258641"
	test  = "JuanHeza"
	limit = 20
)

func main() {
	now := fmt.Sprintf("%v%s", time.Now().Unix(), test)
	fmt.Println(now)
	constants.GenerateCredentials()
	database.Connect()
	encText, err := Encrypt(now, key)
	fmt.Println(encText)
	if err != nil {
		fmt.Println("error encrypting your classified text: ", err)
	}
	err = qrcode.WriteFile("https://example.org", qrcode.Medium, 256, "qr.png")
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	fmt.Println(botInit())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user/profile/:id", func(c *gin.Context) {
		name := c.Param("id")
		decText, err := Decrypt(name, key)
		if err != nil || !IsValid(decText) {

			c.AbortWithError(http.StatusInternalServerError, nil)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": decText,
			})
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
func IsValid(text string) bool {
	if strings.Contains(text, test) {
		strin := strings.ReplaceAll(text, test, "")
		aux, err := strconv.ParseInt(strin, 10, 64)
		if err != nil {
			return false
		}
		return time.Now().Sub(time.Unix(aux, 0)).Seconds() < limit
	} else {
		return false
	}
}

func generateCode() {
	return
}
