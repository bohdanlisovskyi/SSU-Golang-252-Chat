
import QtQuick 2.2
import QtQuick.Controls 2.1
import QtQml.Models 2.2
import "content"

ApplicationWindow {
    id: mainWindow
    visible: true
    width: 400
    height: 300
    title: "SendMeMessage"

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
            //registerWindowLocal.userInput.text = userNameLogin
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
            qmlMessage.sendMessage(message)
        }
    }

    Connections {
        target: qmlRegister
        onRegisterDataIsValid: {
            if (isRegisterValid == true) {
                registerWindowLocal.errText.text = ""
                loginWindowLocal.visible = true
                registerWindowLocal.visible = false
            } else {
                registerWindowLocal.errText.text = "Something wrong. Check fields"
            }
        }
    }

    Connections {
        target: qmlLogin
        onLoginDataIsValid: {
            if (isLoginValid == true){
                loginWindowLocal.errText.text = ""
                mainWindow.width = messagerWindowLocal.width
                mainWindow.height = messagerWindowLocal.height
                messagerWindowLocal.visible = true
                loginWindowLocal.visible = false
            } else {
                loginWindowLocal.errText.text = "Something wrong. Check fields"
            }
        }
    }


}
