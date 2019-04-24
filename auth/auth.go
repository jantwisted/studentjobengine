package auth


import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
	"os"
	jwt "github.com/dgrijalva/jwt-go"
)


func PasswordHash(pwd []byte) string {
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
        log.Println(err)
    }
    return string(hash)
}


func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
    byteHash := []byte(hashedPwd)
    err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
    if err != nil {
        log.Println(err)
        return false
    }
    return true
}


func GenerateJWT(user_name string)(string, error){
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = user_name
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	signing_key := string(os.Getenv("SIGNKEY"))
	tokenString, err := token.SignedString([]byte(signing_key))

	if err != nil{
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
