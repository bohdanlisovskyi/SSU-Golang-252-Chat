package auth

import (
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// UserAuth is a structure for Messager
type UserAuth struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (ug *UserGorm) ByName(UserName string) *User {
	return ug.byQuery(ug.DB.Where("username=?", UserName))
}

func (ug *UserGorm) Logination(username, password string) *User {
	foundUser := ug.ByName(username)
	if foundUser == nil {
		loger.Log.Errorf("No user with that Username")
		return nil
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(foundUser.Password),
		[]byte(password))
	if err != nil {
		loger.Log.Errorf("Invalid password", err)
		return nil
	}

	return foundUser
}

func (ug *UserGorm) byQuery(query *gorm.DB) *User {
	ret := &User{}
	err := ug.DB.First(ret).Error
	switch err {
	case nil:
		return ret
	case gorm.ErrRecordNotFound:
		return nil
	default:
		loger.Log.Errorf("search failed", err)
	}
	return nil
}
