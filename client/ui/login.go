// login
package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

var qmlLog *QmlLog

type QmlLog struct {
	core.QObject
	_ func(userName, password string) `slot:"sendLogD"`
	_ func(isLogValid bool)           `signal:"sendLogIsValid"`
}

func initQmlLog(quickWidget *quick.QQuickWidget) {
	qmlLog = NewQmlLog(nil)
	quickWidget.RootContext().SetContextProperty("qmlLog", qmlLog)
	qmlLog.ConnectSendLogD(func(userName, password string) {
		if userName != "" && password != "" {
			qmlLog.SendLogIsValid(true)
		} else {
			qmlLog.SendLogIsValid(false)
		}
	})
}
