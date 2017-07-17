// login
package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

var qmlLogin *QmlLogin

type QmlLogin struct {
	core.QObject
	_ func(userName, password string) `slot:"checkLoginData"`
	_ func(isLoginValid bool)           `signal:"loginDataIsValid"`
}

func initQmlLogin(quickWidget *quick.QQuickWidget) {
	qmlLogin = NewQmlLogin(nil)
	quickWidget.RootContext().SetContextProperty("qmlLogin", qmlLogin)
	qmlLogin.ConnectCheckLoginData(func(userName, password string) {
		if userName != "" && password != "" {
			qmlLogin.LoginDataIsValid(true)
		} else {
			qmlLogin.LoginDataIsValid(false)
		}
	})
}
