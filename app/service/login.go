package service

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"perScore/app/model"
	"perScore/perScoreProto/user"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

// Login ...
func Login(body []byte) (string, error) {
	const key = "fkzfgk0FY2CaYJhyXbshnPJaRrFtCwfj"
	userDetails := new(user.GetSessionRequest)
	json.Unmarshal([]byte(body), userDetails)
	fmt.Println("userDetails", userDetails)
	conn, err := grpc.Dial("192.168.100.88:6060", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial gRPC connection: %v", err)
	}

	defer conn.Close()
	ctx := context.Background()
	fmt.Println("conn")
	loginUserClientConnection := user.NewUserClient(conn)
	fmt.Println(loginUserClientConnection)
	response, err := loginUserClientConnection.GetSession(ctx, userDetails)
	fmt.Println(response)
	token := Decrypt([]byte(key), response.Token)
	saveToken(token, response.Token)
	return token, err
}

func saveToken(token string, uid string) {
	s := strings.Split(token, ",")
	email := s[0]
	exprationTime, _ := strconv.Atoi(s[1])

	genToken := new(model.Token)
	genToken.Token = uid
	genToken.ExpirationTime = exprationTime
	genToken.Email = email

	db, err := gorm.Open("mysql", "root:root@/library?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Create(&genToken).Error
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}

// Decrypt ...
func Decrypt(key []byte, cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}
