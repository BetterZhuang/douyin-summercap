package service

import (
	"time"

	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/golang-jwt/jwt"
)

var jwtSecret []byte = []byte("xxxxxx")

type Claims struct {
	Id uint64 `json:"id"`
	jwt.StandardClaims
}

func Register(name, password string) (uint64, bool) {
	//检查是否存在同名用户
	var count int64
	config.Db.Model(&model.UserDB{}).Where("username = ?", name).Count(&count)
	if count > 0 {
		return 0, false
	}

	//创建用户实例
	user := model.UserDB{
		Name:     name,
		Username: name,
		Password: password,
	}
	row := config.Db.Create(&user).RowsAffected
	if row == 0 {
		return 0, false
	}
	return user.Id, true
}

func GenerateToken(Id uint64) (string, error) {
	claims := Claims{Id, jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + 2*60*60,
		Issuer:    "Leospard",
	},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func Login(name string, password string) (model.User, bool) {
	var userDB model.UserDB
	var user model.User
	row := config.Db.Where("username = ? and password = ?", name, password).First(&userDB).RowsAffected
	if row == 0 {
		return model.User{}, false
	} else {
		config.Db.Model(&model.UserDB{}).Where("id = ?", userDB.Id).First(&user)
		return user, true
	}
}
