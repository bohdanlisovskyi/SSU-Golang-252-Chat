// register
package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

var qmlReg *QmlReg

type QmlReg struct {
	core.QObject
	_ func(userName, nickName, password string) `slot:"sendRegD"`
	_ func(isRegValid bool)                     `signal:"sendRegIsValid"`
}

func initQmlReg(quickWidget *quick.QQuickWidget) {
	qmlReg = NewQmlReg(nil)
	quickWidget.RootContext().SetContextProperty("qmlReg", qmlReg)
	qmlReg.ConnectSendRegD(func(userName, nickName, password string) {
		if userName != "" && nickName != "" && password != "" {
			qmlReg.SendRegIsValid(true)
		} else {
			qmlReg.SendRegIsValid(false)
		}
	})
}
