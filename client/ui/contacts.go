// contacts
package ui

import (
	"encoding/json"

	"github.com/8tomat8/SSU-Golang-252-Chat/client/config"
	"github.com/8tomat8/SSU-Golang-252-Chat/client/contacts"
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
var listOfContactsModel *core.QAbstractListModel
var helper *modelHelper

//ContactObject struct needed to build model for listView
type ContactObject struct {
	UserName  string
	NickName  string
	IsBlocked bool
	History   string
}

//QmlContacts represent contacts as QObject
type QmlContacts struct {
	core.QObject

	_ func(status bool, index int) `slot:"blockContact"`
	_ func(index int)              `signal:"sendLastMessage"`
	_ func(status bool, index int) `signal:"contactBlocked"`
}

type modelHelper struct {
	core.QObject

	_ func(row int, role string) *core.QVariant `slot:"getData"`
	_ func()                                    `signal:"dataChanged"`
}

func initQmlContacts(quickWidget *quick.QQuickWidget) {
	qmlContacts = NewQmlContacts(nil)
	quickWidget.RootContext().SetContextProperty("qmlContacts", qmlContacts)
	//set isBlocked field
	qmlContacts.ConnectBlockContact(func(status bool, index int) {
		var iData, exists = listOfContacts.Get(index)
		if !exists {
			return
		}
		var data = iData.(*ContactObject)
		currentContactUserName := data.UserName //contacts.ContactsList.ContactsList[index].UserName
		ContactCantBlocked := "Contact " + currentContactUserName + " can't be blocked"
		newMessageHeader := messageService.MessageHeader{
			Type_:    config.GetConfig().MessageType.Contacts,
			Command:  config.GetConfig().MessageCommand.BlockContact,
			UserName: userinfo.CurrentUserInfo.UserName,
			Token:    userinfo.CurrentUserInfo.Token,
		}
		var intStatus int
		if status {
			intStatus = 1
		} else {
			intStatus = 0
		}
		newMessageBody := messageService.ClientContact{
			UserName:  currentContactUserName,
			IsBlocked: intStatus,
		}
		newRawMessageBody, err := json.Marshal(newMessageBody)
		if err != nil {
			qmlContacts.ContactBlocked(false, index)
			loger.Log.Warningf("Can`t marshal contact-block body. %s", err)
			qmlStatus.SendStatus(ContactCantBlocked)
			return
		}
		newMessage := messageService.Message{
			Header: newMessageHeader,
			Body:   newRawMessageBody,
		}
		marshaledMessage, err := messageService.MarshalMessage(&newMessage)
		if err != nil {
			qmlContacts.ContactBlocked(false, index)
			loger.Log.Warningf("Can`t marshal contact-block message. %s", err)
			qmlStatus.SendStatus(ContactCantBlocked)
			return
		}
		if err := connection.WriteMessage(websocket.TextMessage, marshaledMessage); err != nil {
			qmlContacts.ContactBlocked(false, index)
			loger.Log.Warningf("Can not send contact-block message. %s", err)
			qmlStatus.SendStatus(ContactCantBlocked)
			return
		}
		loger.Log.Infof("Message for block %s sent", currentContactUserName)
		//now we are waiting for server response
	})
}

func initContactObject(quickWidget *quick.QQuickWidget) {
	contacts.ContactsList = &messageService.ClientContacts{}
	listOfContacts = arraylist.New()

	listOfContactsModel = core.NewQAbstractListModel(nil)
	listOfContactsModel.ConnectData(contactsData)
	listOfContactsModel.ConnectRowCount(rowCount)
	quickWidget.RootContext().SetContextProperty("ContactModel", listOfContactsModel)

	helper = NewModelHelper(nil)
	helper.ConnectGetData(getData)
	quickWidget.RootContext().SetContextProperty("ModelHelper", helper)
}

func getData(row int, role string) *core.QVariant {

	var iData, exists = listOfContacts.Get(row)
	if !exists {
		return core.NewQVariant()
	}

	var data = iData.(*ContactObject)
	if role == "username" {
		return core.NewQVariant14(data.UserName)
	}
	if role == "nickname" {
		return core.NewQVariant14(data.NickName)
	}
	if role == "history" {
		return core.NewQVariant14(data.History)
	}
	return core.NewQVariant11(data.IsBlocked)
}

//this function provide connect between listView (on UI) and modelView (in code)
func contactsData(index *core.QModelIndex, role int) *core.QVariant {
	if role == 0 && index.IsValid() {
	}
	return core.NewQVariant()
}

//we need manually provide function for count of row in model
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
		return err
	}
	if err := connection.WriteMessage(websocket.TextMessage, marshaledMessage); err != nil {
		loger.Log.Warningf("Can not send contacts request. %s", err)
		qmlStatus.SendStatus("Can not load contacts.")
		return err
	}
	return err
}

func updateContactsList() {
	//check if local contacts still online
	for i := 0; i < listOfContacts.Size(); {
		var contains bool = false
		var iData, exists = listOfContacts.Get(i)
		if !exists {
			break
		}
		var data = iData.(*ContactObject)
		for _, contact := range contacts.ContactsList.ContactsList {
			if data.UserName == contact.UserName {
				contains = true
				break
			}
		}
		//if not - remove from local list
		if !contains {
			listOfContacts.Remove(i)
		} else {
			i++
		}
	}
	//now check for new contacts online
	for _, contact := range contacts.ContactsList.ContactsList {
		var contains bool = false
		for i := 0; i < listOfContacts.Size(); i++ {
			var iData, exists = listOfContacts.Get(i)
			if !exists {
				break
			}
			var data = iData.(*ContactObject)
			if data.UserName == contact.UserName {
				contains = true
				break
			}
		}
		if !contains {
			listOfContacts.Add(&ContactObject{
				UserName:  contact.UserName,
				NickName:  contact.NickName,
				IsBlocked: contact.IsBlocked != 0,
			})
		}
	}
	//send signal to update model
	helper.DataChanged()
}

//return index, if contact exists, -1 otherwise
func addToHistory(userName, message string, owner bool) int {
	for i := 0; i < listOfContacts.Size(); i++ {
		var iData, exists = listOfContacts.Get(i)
		if !exists {
			break
		}
		var data = iData.(*ContactObject)
		if data.UserName == userName {
			if owner {
				data.History += "<div style=\"background-color: lightblue; text-align: right;\">" + message + "</div>"
			} else {
				data.History += "<div style=\"background-color: lightblue; text-align: left;\">" + message + "</div>"
			}
			return i
		}
	}
	return -1
}
