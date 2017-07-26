package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	//_ "github.com/mattn/go-sqlite3"
	//"database/sql"
)


var qmlStatus *QmlStatus

type QmlStatus struct{
	core.QObject

	_ func(statusMessage string)	`signal:"sendStatus"`
}

func initQmlStatus(quickWidget *quick.QQuickWidget){
	qmlStatus = NewQmlStatus(nil)
	quickWidget.RootContext().SetContextProperty("qmlStatus", qmlStatus)
	//db, _ := sql.Open("sqlite3", "./foo.db")
	//db.Ping()
}
