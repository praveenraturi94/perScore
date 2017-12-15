package service

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"perScoreServer/app/model"
	pb "perScoreServer/perScoreProto/user"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

// Login ...
func Login(body []byte) (*pb.GetSessionResponse, error) {
	userDetails := new(pb.GetSessionRequest)
	json.Unmarshal([]byte(body), userDetails)
	conn, err := grpc.Dial(os.Getenv("PER_SCORE_AUTH_SERVICE_DIAL"), grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatalf("Failed to dial gRPC connection: %v", err)
	}

	ctx := context.Background()
	loginUserClientConnection := pb.NewUserClient(conn)
	response, err := loginUserClientConnection.GetSession(ctx, userDetails)
	fmt.Println("perScoreAuth Response:", response)
	return response, err
}

// Decrypt ...
func Decrypt(key []byte, cryptoText string) map[string]string {
	mappedResult := make(map[string]string)
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
	dataByte := ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	dataArray := strings.Split(fmt.Sprintf("%s", dataByte), ",")

	mappedResult["email"] = dataArray[0]
	mappedResult["role"] = dataArray[1]
	mappedResult["sessionTime"] = dataArray[2]

	saveToken(mappedResult, cryptoText)

	return mappedResult
}

func saveToken(mappedResult map[string]string, uid string) {
	sessionTime, _ := strconv.Atoi(mappedResult["sessionTime"])

	genToken := new(model.Token)
	genToken.Token = uid
	genToken.ExpirationTime = sessionTime
	genToken.Email = mappedResult["email"]

	dbString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=%s",
		os.Getenv("DEV_HOST"), os.Getenv("DEV_DBNAME"), os.Getenv("DEV_USERNAME"), os.Getenv("DEV_PASSWORD"), os.Getenv("DEV_SSLMODE"))
	db, err := gorm.Open(os.Getenv("DEV_DB_DRIVER"), dbString)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Create(&genToken).Error
	if err != nil {
		log.Fatal(err)
	}
}
