package ui

import "github.com/therecipe/qt/core"

var qmlSettings *QmlSettings

type QmlSettings struct {
	core.QObject

	_ func ()
}
