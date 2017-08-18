// contacts
package ui

import (
	"strconv"

	"encoding/json"

	"github.com/8tomat8/SSU-Golang-252-Chat/client/config"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/8tomat8/SSU-Golang-252-Chat/userinfo"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gorilla/websocket"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

//binding to qml
var qmlContacts *QmlContacts
var listOfContacts *arraylist.List

//ContactObject struct needed to build model for listView
type ContactObject struct {
	core.QObject

	_ string `property:"username"`
	_ string `property:"nickname"`
	_ bool   `property:"isblocked"`
}

//QmlContacts represent contacts as QObject
type QmlContacts struct {
	core.QObject

	_ func(newIndex int)                    `slot:"changeCurrentContact"`
	_ func(searchedUser string)             `slot:"searchUser"`
	_ func(newUsername, newNickname string) `slot:"addUser"`
	_ func(status bool, index int)          `slot:"blockContact"`
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
	//set isBlocked field
	qmlContacts.ConnectBlockContact(func(status bool, index int) {
		//this code will be uncommented after Contacts will be born
		//contacts.ContactsList.ContactsList[index].IsBlocked = status
	})
}

func initContactObject(quickWidget *quick.QQuickWidget) {
	listOfContacts = arraylist.New()

	//test contacts. will be removed after implement Contacts service
	var testContact *ContactObject
	for i := 0; i < 5; i++ {
		testContact = NewContactObject(nil)
		testContact.SetUsername("Me " + strconv.Itoa(i))
		testContact.SetNickname("Volodymyr " + strconv.Itoa(i))
		testContact.SetIsblocked(i%2 == 0)
		listOfContacts.Add(testContact)
	}
	var test2Contact = NewContactObject(nil)
	test2Contact.SetUsername("Me " + "test2")
	test2Contact.SetNickname("Volodymyr " + "test2")
	test2Contact.SetIsblocked(false)
	listOfContacts.Add(test2Contact)
	var model = core.NewQAbstractListModel(nil)
	model.ConnectData(contactsData)
	model.ConnectRowCount(rowCount)
	quickWidget.RootContext().SetContextProperty("ContactModel", model)
}

//this function provide connect between listView (on UI) and modelView (in code)
func contactsData(index *core.QModelIndex, role int) *core.QVariant {
	//when we click on list not guaranteed that data wasn't change before update UI
	//if we out of range, or < 0 - return empty *QVariant
	if !index.IsValid() {
		return core.NewQVariant()
	}
	//iData - list element, i have no idea, why it is not already ContactObject
	//example of code taken from developers of therecipe/qt library code
	var iData, exists = listOfContacts.Get(index.Row())
	if !exists {
		return core.NewQVariant()
	}
	//if all right - get *ContactObject from list element
	resultContact := iData.(*ContactObject)

	return resultContact.ToVariant()
}

//we need manualy provide function for count of row in model
func rowCount(parent *core.QModelIndex) int {
	return listOfContacts.Size()
}

//
func sendContactsRequestToServer() error {
	const errorStatusString = "Can not load contacts."
	newMessageHeader := messageService.MessageHeader{
		Type_:    config.GetConfig().MessageType.Contacts,
		Command:  config.GetConfig().MessageCommand.ContactsRequest,
		UserName: userinfo.CurrentUserInfo.UserName,
		Token:    userinfo.CurrentUserInfo.Token,
	}
	//we have no data to send, so create empty Body
	newMessageBody := messageService.MessageBody{}
	newRawMessageBody, err := json.Marshal(&newMessageBody)
	if err != nil {
		loger.Log.Warningf("Can`t marshal contacts message body. %s", err)
		qmlStatus.SendStatus(errorStatusString)
		return err
	}
	newMessage := messageService.Message{
		Header: newMessageHeader,
		Body:   newRawMessageBody,
	}
	marshaledMessage, err := messageService.MarshalMessage(&newMessage)
	if err != nil {
		loger.Log.Warningf("Can not marshal contacts message. %s", err)
		qmlStatus.SendStatus("Can not load contacts.")
		return
	}
	if err := connection.WriteMessage(websocket.TextMessage, marshaledMessage); err != nil {
		loger.Log.Warningf("Can not send contacts request. %s", err)
		qmlStatus.SendStatus("Can not load contacts.")
		return
	}
	//connection.WriteMessage(websocket.TextMessage, marshaledMessage)
}
