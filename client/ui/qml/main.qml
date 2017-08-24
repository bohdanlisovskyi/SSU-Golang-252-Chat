
import QtQuick 2.2
import QtQuick.Controls 2.1
import QtQuick.Layouts 1.1
import QtQml.Models 2.2
import "content"

ApplicationWindow {
    id: mainWindow
    visible: true
    width: 400
    height: 320
    title: "SendMeMessage"
    property bool canToolButtonBeClicked
    signal openSettings()
    signal logout()
    signal quit()

    onQuit: {
        qmlLogin.logOut()
    }

    canToolButtonBeClicked: true

    header: ToolBar {
        id: toolBar
        Layout.fillWidth: true
        width: mainWindow.width
        height: 20
        RowLayout {
            anchors.fill: parent
            ToolButton {
                id: fileMenuButton
                text: qsTr("File")
                height: 20
                checkable: true
                onCheckedChanged: {
                    if(fileMenuButton.checked&&canToolButtonBeClicked){
                        menu.open()
                    } else {
                        menu.close()
                    }
                }

                Layout.maximumHeight: 20
                Menu {
                    id: menu
                    visible: false
                    y: fileMenuButton.height
                    Layout.maximumWidth: 100
                    width: 100
                    closePolicy: Popup.NoAutoClose
                    MenuItem {
                        id: settingsMenu
                        text: "Settings"
                        height: 0
                        visible: false
                        Layout.maximumHeight: 20
                        onTriggered: {
                            fileMenuButton.checked = false
                            openSettings()
                        }

                        background: Rectangle {
                            color: "white"
                            border.color: "#585594"
                            border.width: 1
                        }
                    }
                    MenuItem {
                        id: logoutMenu
                        text: "Log out"
                        height: 0
                        visible: false
                        Layout.maximumHeight: 20
                        onTriggered: {
                            fileMenuButton.checked = false
                            logout()
                        }

                        background: Rectangle {
                            color: "white"
                            border.color: "#585594"
                            border.width: 1
                        }
                    }
                    MenuSeparator {}
                    MenuItem {
                        id: quitMenu
                        text: "Quit"
                        height: 20
                        Layout.maximumHeight: 20
                        onTriggered: {
                            fileMenuButton.checked = false
                            mainWindow.quit()
                        }
                        background: Rectangle {
                            color: "white"
                            border.color: "#585594"
                            border.width: 1
                        }
                    }
                    MenuSeparator {}
                }
            }
        }
    }


    Rectangle {
        id: backgroundRect
        color: "transparent"
        z: -1
        border.color: "transparent"
        anchors.fill: parent
    }

    Login {
        id: loginWindowLocal
        visible: false
        onLogin: {
            qmlLogin.checkLoginData(userInput.text, passInput.text)
        }
        onGoBackToRegister: {
            registerWindowLocal.visible = true
            visible = false
        }
        onInputFieldsChanged: {
            errText.text = ""
        }
    }

    Register {
        id: registerWindowLocal
        visible: true
        onRegister: {
            qmlRegister.checkRegisterData(userInput.text, nickInput.text, passInput.text)
        }
        onGoNextToLogin: {
            loginWindowLocal.visible = true
            visible = false
        }
        onInputFieldsChanged: {
            errText.text = ""
        }
    }

    Messager {
        id: messagerWindowLocal
        visible: false
        onSend: {
            qmlMessage.sendMessage(message, index)
        }
        onBlock: {
            qmlContacts.blockContact(status, index)
        }
    }

    Settings {
        id: settingsWindowLocal
        visible: false
        onBackToMessager: {
            messagerWindowLocal.visible = true
            visible = false
        }
        onApplyChanges: {
            qmlSettings.applySettings(newPassword, oldPassword, newNickname, birthday, about)
        }
    }

    Connections {
        target: mainWindow
        onQuit: {
            mainWindow.close()
        }
    }

    Connections {
        target: mainWindow
        onOpenSettings: {
            messagerWindowLocal.visible = false
            settingsWindowLocal.visible = true
        }
    }

    Connections {
        target: mainWindow
        onLogout: {
            qmlLogin.logOut()
            messagerWindowLocal.visible = false
            loginWindowLocal.visible = true
            mainWindow.width = loginWindowLocal.width
            mainWindow.height = loginWindowLocal.height + toolBar.height + statusBar.height
        }
    }

    Connections {
        target: qmlRegister
        onRegisterDataIsValid: {
            if (isRegisterValid == true) {
                registerWindowLocal.errText.text = ""
                loginWindowLocal.visible = true
                registerWindowLocal.visible = false

                registerWindowLocal.userInput.text = ""
                registerWindowLocal.nickInput.text = ""
                registerWindowLocal.passInput.text = ""
            } else {
                registerWindowLocal.errText.text = "Something wrong. Check fields"
            }
        }
    }

    Connections {
        target: qmlLogin
        onLoginDataIsValid: {
            loginWindowLocal.busyInd.running = false
            if (isLoginValid == true){
                loginWindowLocal.errText.text = ""
                //resize window to size of messanger
                mainWindow.width = messagerWindowLocal.width
                mainWindow.height = messagerWindowLocal.height + toolBar.height + statusBar.height
                messagerWindowLocal.visible = true
                loginWindowLocal.visible = false
                //restore height of hiden menus
                logoutMenu.height = 20
                logoutMenu.visible = true
                settingsMenu.height = 20
                settingsMenu.visible = true

                loginWindowLocal.userInput.text = ""
                loginWindowLocal.passInput.text = ""
            } else {
                loginWindowLocal.errText.text = "Something wrong. Check fields"
            }
        }
        onStartBusyIndicator: {
            loginWindowLocal.busyInd.running = true
        }
    }

    Connections {
        target: qmlMessage
        onMessageSent: {
            if(messageWasSent) {
                messagerWindowLocal.historyText.text += messagerWindowLocal.messageText.text
                messagerWindowLocal.messageText.text = ""
            }
        }
    }

    Connections {
        target: qmlStatus
        onSendStatus: {
            statusText.text = statusMessage
        }
    }

    footer: ToolBar {
        id: statusBar
        height: 20
        RowLayout {
            Text {
                id: statusText
                text: "Status: ok"
                horizontalAlignment: Text.AlignHCenter
            }
        }
    }
}
