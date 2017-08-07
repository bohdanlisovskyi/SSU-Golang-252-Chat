package settingService

import (
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
)

var (
	Type_        = "settings"
	Command      = "block/unblock user"
	UserName     = "userOne"
	ReceiverName = "userTwo"
	CurrentTime  = 1500373713
	Token        = "token"
	IsBlocked    = 1
	byteRequest  = []byte(`{"header":{"type":"settings","command":"block/unblockUser","userName":"userName","token": "someToken"},"body":{"is_blocked":1, "name": "contactName"}}`)
	badRequest   = []byte(`{"header":{"type":"settings","command":"block/unblockUser","userName":"userName","token": "someToken"},"body":{"is_blocked":2, "name": "contactName"}}`)
	badRequest2  = []byte(``)
)

func Test1UnmarshalBlockUserRequest(t *testing.T) {
	IsBlocked, name, err := UnmarshalBlockUserRequestBody(byteRequest)
	if err != nil {
		t.Fatal("Error has occured: ", err)
	}
	t.Logf("IsBlocked and Name values after unmarshaling has types %T and %T \n", IsBlocked, name)
	t.Logf("IsBlocked and Name values after unmarshaling : %+v and %+v \n ", IsBlocked, name)
}

func Test2UnmarshalBlockUserRequest(t *testing.T) {
	if _, _, err := UnmarshalBlockUserRequestBody(badRequest); err == nil {
		t.Fatal("Error expected, but absent: ", err)
	} else {
		t.Logf("Expected error has occurred : %+v \n ", err)
	}
}

func Test3UnmarshalBlockUserRequest(t *testing.T) {
	if _, _, err := UnmarshalBlockUserRequestBody(badRequest2); err == nil {
		t.Fatal("Error expected, but absent: ", err)
	} else {
		t.Logf("Expected error has occurred : %+v \n ", err)
	}
}

func TestBlockUnblockUser(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	if ok, _ := BlockUnblockUser(byteRequest); !ok {
		t.Fatal("Error has occurred")
	}
	if 	_, err = BlockUnblockUser(byteRequest); err != nil{
		t.Fatal("Error has occurred")

	}


}
