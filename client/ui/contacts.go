// contacts
package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

var qmlContacts *QmlContacts

type QmlContacts struct {
	core.QObject
	_ func(newIndex int)        `slot:"changeCurrentContact"`
	_ func(searchedUser string) `slot:"searchUser"`
	_ func(lastMessage string)  `signal:"sendLastMessage"`
	_ func(history string)      `signal:"sendHistory"`
}

func initQmlContacts(quickWidget *quick.QQuickWidget) {
	qmlContacts = NewQmlContacts(nil)
	quickWidget.RootContext().SetContextProperty("qmlContacts", qmlContacts)
	qmlContacts.ConnectChangeCurrentContact(func(newIndex int) {

	})

}
