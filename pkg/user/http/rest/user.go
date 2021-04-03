package rest

import (
	"errors"
	"github.com/GeorgVartanov/myWebApp/pkg/user/service/create"
	"github.com/GeorgVartanov/myWebApp/utils"
	"log"
	"strings"
	"time"
)

var (
	emailIsEmpty      = errors.New("email field is empty, should be filled")
	passwordIsEmpty   = errors.New("password field is empty, should be filled")
	emailContainsChar = errors.New("email doesn't contains @, please fill email field properly")
	passwordLength    = errors.New("password is less than 6 characters")
	passwordMatch     = errors.New("passwords doesn't match")
	//displayNameIsEmpty = errors.New("display name field is empty, should be filled")
)

//User ...
type User struct {
	UserName      string    `json:"userName"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	PasswordCheck string    `json:"passwordCheck"`
	Created       time.Time `json:"created"`
	Changed       time.Time `json:"changed"`
	IsAdmin       bool      `json:"isAdmin"`
}

func (u *User) ConvertToServiceUser() (*create.User, error) {

	hash, err := utils.HashGenerateFromPassword(&u.Password)
	if err != nil {
		return nil, err
	}
	MoscowTime, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatal(err)
	}
	return &create.User{
		UserName:  u.UserName,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  hash,
		Created:   time.Now().In(MoscowTime),
		Changed:   time.Now().In(MoscowTime),
		IsAdmin:   false,
	}, err
}

//ValidateFields All User fields
func (u *User) ValidateFields() error {
	if err := u.ValidateEmail(); err != nil {
		return err
	}
	if err := u.ValidatePassword(); err != nil {
		return err
	}
	u.ValidateDisplayName()

	return nil
}

func (u *User) ValidateEmail() error {
	if u.Email == "" {
		return emailIsEmpty
	}
	if !strings.Contains(u.Email, "@") {
		return emailContainsChar
	}
	return nil
}

func (u *User) ValidatePassword() error {
	if u.Password == "" {
		return passwordIsEmpty
	}
	if len(u.Password) < 6 {
		return passwordLength
	}
	if u.Password != u.PasswordCheck {
		return passwordMatch
	}
	return nil
}

func (u *User) ValidateDisplayName() {
	if u.UserName == "" {
		u.UserName = u.Email
	}
}
