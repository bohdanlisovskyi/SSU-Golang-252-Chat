package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

//binding to qml
var qmlStatus *QmlStatus
//represent status structure as QObject
type QmlStatus struct{
	core.QObject

	_ func(statusMessage string)	`signal:"sendStatus"`
}

func initQmlStatus(quickWidget *quick.QQuickWidget){
	qmlStatus = NewQmlStatus(nil)
	quickWidget.RootContext().SetContextProperty("qmlStatus", qmlStatus)
}