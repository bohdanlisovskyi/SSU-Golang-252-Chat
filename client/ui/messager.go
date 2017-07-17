// messager
package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

var qmlMessage *QmlMessage

type QmlMessage struct {
	core.QObject

	_ func(message string) `slot:"sendMessage"`
}

func initQmlMessage(quickWidget *quick.QQuickWidget) {
	qmlMessage = NewQmlMessage(nil)
	quickWidget.RootContext().SetContextProperty("qmlMessage", qmlMessage)
	qmlMessage.ConnectSendMessage(func(message string) {

	})
}
