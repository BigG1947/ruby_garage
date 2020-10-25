package models

import (
	"RubyGarage_v.2.0/utils"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strings"
)

type Token struct {
	UserId int64
	jwt.StandardClaims
}

type Account struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func (ac *Account) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(ac.Email, "@") {
		return utils.Message(false, "Email address is required"), false
	}

	if len(ac.Password) < 6 {
		return utils.Message(false, "Password must be a least 6 characters"), false
	}

	var temp Account

	err := db.QueryRow("SELECT id, email, password FROM account WHERE email = ?", ac.Email).Scan(&temp.Id, &temp.Email, &temp.Password)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("%s\n", err)
		return utils.Message(false, "DB connection error. Please retry.."), false
	}

	if err == nil {
		return utils.Message(false, "This email already in use by another account"), false
	}

	return utils.Message(true, "Validate passed"), true
}

func (ac *Account) Login() map[string]interface{} {
	var hashPwd string
	err := db.QueryRow("SELECT id, password FROM account WHERE email = ?", ac.Email).Scan(&ac.Id, &hashPwd)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.Message(false, "Email address not found")
		}
		log.Printf("%s\n", err)
		return utils.Message(false, "DB connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(ac.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return utils.Message(false, "Invalid login or password. Please try again")
		}
		log.Printf("%s\n", err)
		return utils.Message(false, "Internal server error. Please try later.")
	}

	ac.Password = ""

	var tk Token
	tk.UserId = ac.Id
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(os.Getenv("secret-key")))
	if err != nil {
		log.Printf("%s\n", err)
		return utils.Message(false, "Internal server error. Please try later.")
	}

	ac.Token = tokenString

	response := utils.Message(true, "Success login!")
	response["account"] = ac
	return response
}

func (ac *Account) Create() map[string]interface{} {
	if response, ok := ac.Validate(); !ok {
		return response
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ac.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("%s\n", err)
		return utils.Message(false, "Internal server error. Please try later.")
	}
	ac.Password = string(hashedPassword)

	res, err := db.Exec("INSERT INTO account(email, password) VALUES (?, ?)", ac.Email, ac.Password)
	if err != nil {
		log.Printf("%s\n", err)
		return utils.Message(false, "DB connection error. Please retry later")
	}
	if ac.Id, err = res.LastInsertId(); err != nil {
		log.Printf("%s\n", err)
		return utils.Message(false, "Internal server error. Please try later.")
	}

	var tk Token
	tk.UserId = ac.Id
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(os.Getenv("secret-key")))
	if err != nil {
		log.Printf("%s\n", err)
		return utils.Message(false, "Internal server error. Please try later.")
	}
	ac.Token = tokenString

	ac.Password = ""

	response := utils.Message(true, "Account has been created")
	response["account"] = ac
	return response
}

func (ac *Account) Get() map[string]interface{} {
	if err := db.QueryRow("SELECT email, password FROM account WHERE id = ?", ac.Id).Scan(&ac.Email, &ac.Password); err != nil {
		if err == sql.ErrNoRows {
			return utils.Message(false, "User not found!")
		}
		log.Printf("%s\n", err)
		return utils.Message(false, "DB connection Error. Please try later")
	}

	ac.Password = ""

	response := utils.Message(true, "Account found")
	response["user"] = ac
	return response
}
