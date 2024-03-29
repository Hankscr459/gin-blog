package configs

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"gin-blog/plugins/dto"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gobeam/stringy"
	"golang.org/x/crypto/bcrypt"
)

var Email string
var IDUser string

func EncriptPassword(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}

func CheckUser(email string, password string) (dto.ReadUserWithPassword, error) {
	user, err := User().FindByEmail(email)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return dto.ReadUserWithPassword{}, errors.New("此會員不存在")
		} else {
			return *user, err
		}
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		if err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password" {
			return *user, errors.New("密碼錯誤")
		}
		return *user, err
	}
	return *user, nil
}

func GenerJWT(t dto.ReadUser) (string, error) {
	secret := []byte(os.Getenv("JwtSecretKey"))

	payload := jwt.MapClaims{
		"email": t.Email,
		"name":  t.Name,
		"_id":   t.ID,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	fmt.Println("t.ID: ", t)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

func ProccessToken(tk string) (*dto.Cliam, bool, string, error) {
	secret := []byte(os.Getenv("JwtSecretKey"))
	claims := &dto.Cliam{}
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Format token invalid")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err == nil {
		find, _ := User().FindByEmail(claims.Email)
		if find != nil {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, true, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token invalid")
	}
	return claims, false, string(""), err
}

func Body[T any](ctx *gin.Context) T {
	value, _ := ctx.Get("reqBody")
	body := value.(T)
	return body
}

func Profile(ctx *gin.Context) *dto.Cliam {
	value, _ := ctx.Get("claims")
	user := value.(*dto.Cliam)
	fmt.Println("decode: ", user)
	return user
}

func UcFirst(value string) string {
	container := stringy.New(value)
	return container.UcFirst()
}
