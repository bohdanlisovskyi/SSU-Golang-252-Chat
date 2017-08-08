package auth

import (
	"fmt"
	"testing"

	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
)

func TestLogin(t *testing.T) {
	test_data := messageService.Authentification{
		UserName: "vasya",
		Password: "vasya2vasya444",
	}
	user, token, err := Login(test_data.UserName, test_data.Password)
	fmt.Println(user, token, err)
}
