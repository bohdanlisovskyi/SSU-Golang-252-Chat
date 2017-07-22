// uimain
package ui

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
	//"github.com/8tomat8/SSU-Golang-252-Chat/loger"
)

func CreateWindow() {
	widgets.NewQApplication(len(os.Args), os.Args)
	var layout = widgets.NewQHBoxLayout()
	layout.AddWidget(newQmlWidget(), 0, 0)
	var window = widgets.NewQMainWindow(nil, 0)
	var centralWidget = widgets.NewQWidget(window, 0)
	centralWidget.SetLayout(layout)
	window.SetCentralWidget(centralWidget)

	widgets.QApplication_Exec()
}

func newQmlWidget() *quick.QQuickWidget {
	var quickWidget = quick.NewQQuickWidget(nil)
	quickWidget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)

	initQmlRegister(quickWidget)
	initQmlLogin(quickWidget)
	initQmlMessage(quickWidget)
	quickWidget.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))

	return quickWidget
}
