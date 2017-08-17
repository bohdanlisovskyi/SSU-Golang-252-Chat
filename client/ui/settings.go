package ui

import (
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

var qmlSettings *QmlSettings

type QmlSettings struct {
	core.QObject

	_ func(nickName string, date time.Time, about string)                              `signal:"sendSettingsData"`
	_ func(oldPassword, newPassword, newNickname string, date time.Time, about string) `slot:"applySettings"`
}

func initQmlSettings(quickWidget *quick.QQuickWidget) {
	qmlSettings = NewQmlSettings(nil)
	quickWidget.RootContext().SetContextProperty("qmlSettings", qmlSettings)

	//TODO: fix difference between c++ date and golang time
	qmlSettings.ConnectApplySettings(func(oldPassword, newPassword, newNickname string, date time.Time, about string) {
		//we need to form settings message and sent it to server

		//after send message we just waiting for response

	})
}
