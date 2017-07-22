// contacts
package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	//"github.com/8tomat8/SSU-Golang-252-Chat/client/contacts"
)

var qmlContacts *QmlContacts
//var contactsList contacts.Contacts

type QmlContacts struct {
	core.QObject
	_ func(newIndex int)        			`slot:"changeCurrentContact"`
	_ func(searchedUser string) 			`slot:"searchUser"`
	_ func(newUsername, newNickname string) `slot:"addUser"`
	_ func(lastMessage string)  			`signal:"sendLastMessage"`
	_ func(history string)      			`signal:"sendHistory"`
}

func initQmlContacts(quickWidget *quick.QQuickWidget) {
	qmlContacts = NewQmlContacts(nil)
	quickWidget.RootContext().SetContextProperty("qmlContacts", qmlContacts)
	qmlContacts.ConnectChangeCurrentContact(func(newIndex int) {

	})
	qmlContacts.ConnectSearchUser(func(searchedUser string){

	})
	qmlContacts.ConnectAddUser(func(newUsername, newNickname string){

	})

	//contactsList.LoadDataFromDB()
}
