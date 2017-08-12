package ui

import (
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

var qmlSettings *QmlSettings

type QmlSettings struct {
	core.QObject

	_ func(oldPassword, newPassword, newNickname string, date time.Time, about string) `slot:"applySettings"`
}

func initQmlSettings(quickWidget *quick.QQuickWidget) {
	qmlSettings = NewQmlSettings(nil)
	quickWidget.RootContext().SetContextProperty("qmlSettings", qmlSettings)

	//TODO: fix difference between c++ date and golang time
	qmlSettings.ConnectApplySettings(func(oldPassword, newPassword, newNickname string, date time.Time, about string) {
		//we need to form settings message and sent it to server
		//code will be uncommented after fix Settings functionality
		/*newMessageHeader := messageService.MessageHeader{
			Type_:    config.GetConfig().MessageType.Settings,
			Command:  config.GetConfig().MessageCommand.SendSettings,
			UserName: userinfo.CurrentUserInfo.UserName,
			Token:    userinfo.CurrentUserInfo.Token,
		}*/

		//after send message we just waiting for response

	})
}
