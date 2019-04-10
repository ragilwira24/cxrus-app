package impl

import (
	"cxrus-app/service/models"
	auth "cxrus-app/util/auth/models"
	"cxrus-app/util/db"
	handler "cxrus-app/util/handler"
	"os"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"
)

// Register function to register new user
func Register(user *models.User) map[string]interface{} {

	if resp, ok := ValidatingUser(user); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	db.GetDB().Create(user)
	if user.ID <= 0 {
		return handler.Message(false, "Failed to create account, connection error")
	}

	user.Token = SignJWT(user.ID)
	user.Password = ""

	response := handler.Message(true, "User Has Been Created")
	response["users"] = user
	return response

}

// ValidatingUser is function for checking user
func ValidatingUser(user *models.User) (map[string]interface{}, bool) {

	if !strings.Contains(user.Email, "@") {
		return handler.Message(false, "Invalid Email"), false
	}

	if len(user.Password) < 6 {
		return handler.Message(false, "Minimum Character For Password is 6"), false
	}

	userTemp := &models.User{}
	err := db.GetDB().Table("users").Where("email = ?", user.Email).First(userTemp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return handler.Message(false, "Connection Error to DB, Please try again"), false
	}

	if userTemp.Email != "" {
		return handler.Message(false, "Email Address already created with: "+userTemp.Email), false
	}

	return handler.Message(false, "There No Error When Validating this user"), true

}

// Login function for user to sign in
func Login(email string, password string) map[string]interface{} {

	user := &models.User{}
	err := db.GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return handler.Message(false, "User not found")
		}
		return handler.Message(false, "Error when connecting to DB")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return handler.Message(false, "The Password Was Mismatch")
	}

	user.Password = ""
	user.Token = SignJWT(user.ID)

	response := handler.Message(true, "Logged In")
	response["user"] = user

	return response

}

// SignJWT function for create a JSONWebToken
func SignJWT(userID uint) string {

	tk := &auth.Token{UserID: userID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token.password")))

	return tokenString

}
