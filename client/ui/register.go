// register
package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

var qmlRegister *QmlRegister

type QmlRegister struct {
	core.QObject
	_ func(userName, nickName, password string)      `slot:"checkRegisterData"`
	_ func(isRegisterValid bool)                     `signal:"registerDataIsValid"`
}

func initQmlRegister(quickWidget *quick.QQuickWidget) {
	qmlRegister = NewQmlRegister(nil)
	quickWidget.RootContext().SetContextProperty("qmlRegister", qmlRegister)
	qmlRegister.ConnectCheckRegisterData(func(userName, nickName, password string) {
		if userName != "" && nickName != "" && password != "" {
			qmlRegister.RegisterDataIsValid(true)
		} else {
			qmlRegister.RegisterDataIsValid(false)
		}
	})
}
