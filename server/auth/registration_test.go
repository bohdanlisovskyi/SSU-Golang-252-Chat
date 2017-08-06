package auth

import (
	"testing"

	"fmt"

	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
)

func TestRegisterNewUser(t *testing.T) {

	test_user := messageService.User{
		UserName: "vasya6644",
		Password: "vasya2",
		NickName: "voron",
	}
	user, err := RegisterNewUser(&test_user)
	fmt.Println(user, err)
}
