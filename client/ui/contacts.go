// contacts
package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

////binding to qml
var qmlContacts *QmlContacts

//

//represent contacts as QObject
type QmlContacts struct {
	core.QObject
	_ func(newIndex int)                    `slot:"changeCurrentContact"`
	_ func(searchedUser string)             `slot:"searchUser"`
	_ func(newUsername, newNickname string) `slot:"addUser"`
	_ func(lastMessage string, index int)   `signal:"sendLastMessage"`
	_ func(history string, index int)       `signal:"sendHistory"`
}

func initQmlContacts(quickWidget *quick.QQuickWidget) {
	qmlContacts = NewQmlContacts(nil)
	quickWidget.RootContext().SetContextProperty("qmlContacts", qmlContacts)
	//change current user, when it changed on UI
	qmlContacts.ConnectChangeCurrentContact(func(newIndex int) {

	})
	//send search request to server
	qmlContacts.ConnectSearchUser(func(searchedUser string) {

	})
	//add finded user to contacts list
	qmlContacts.ConnectAddUser(func(newUsername, newNickname string) {

	})
	//
}

func GetQmlContacts() *QmlContacts {
	return qmlContacts
}
