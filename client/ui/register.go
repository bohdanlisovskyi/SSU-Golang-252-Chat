// register
package ui

import (
	"fmt"

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
	fmt.Print("Reg")
	qmlReg = NewQmlReg(nil)
	fmt.Print("After New")
	quickWidget.RootContext().SetContextProperty("qmlReg", qmlReg)
	fmt.Print("Set properties")
	qmlReg.ConnectSendRegD(func(userName, nickName, password string) {
		if userName != "" && nickName != "" && password != "" {
			qmlReg.SendRegIsValid(true)
		} else {
			qmlReg.SendRegIsValid(false)
		}
	})
}
