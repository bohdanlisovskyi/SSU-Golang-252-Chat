package auth

import (
	"github.com/8tomat8/SSU-Golang-252-Chat/database"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)


func ByName(UserName string) (*User, error) {
	db, err := database.GetStorage()
	if err != nil {
		loger.Log.Errorf("Failed to open db", err)
		return nil, err
	}
	res, err := byQuery(db.Where("username=?", UserName))
	if err != nil {
		loger.Log.Errorf("Search in db failed")
		return nil, err
	}
	return res, err
}

func Login(username, password string) error {
	foundUser, err := ByName(username)
	if err != nil {
		loger.Log.Errorf("No user with that Username")
		return err
	}

	err2 := bcrypt.CompareHashAndPassword(
		[]byte(foundUser.Password),
		[]byte(password))
	if err2 != nil {
		loger.Log.Errorf("Invalid password", err2)
		return err2
	}

	return nil
}

func byQuery(db *gorm.DB) (*User, error) {
	ret := &User{}
	err := db.First(ret).Error
	if err != nil {
		loger.Log.Errorf("Failed to find user in DB")
		return nil, err
	}
	return ret, err
}

//var state string
//
//func randToken() string {
//	b := make([]byte, 32)
//	rand.Read(b)
//	return base64.StdEncoding.EncodeToString(b)
//}
//
//func loginHandler(UserName, Password string) {
//	if UserName != "" && Password != "" {
//		state = randToken()
//		session := sessions.Default()
//		session.Set("state", state)
//		session.Save()
//	} else {
//		//print error("Empty fields")
//		return
//	}
//}
//
//func authHandler() {
//	// Check state validity.
//	session := sessions.Default()
//	retrievedState := session.Get("state")
//	if retrievedState != byQuery("state") {
//		loger.Log.Errorf("Invalid session state: %s", retrievedState)
//		return
//	}
//	// Handle the exchange code to initiate a transport.
//	tok, err := conf.Exchange(NoContext, ug.byQuery("code"))
//	if err != nil {
//		//AbortWithError(http.StatusBadRequest, err)
//		return
//	}
//}
