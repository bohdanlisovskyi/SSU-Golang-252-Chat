import QtQuick 2.7
import QtQuick.Controls 2.1
import QtQuick.Layouts 1.3

Item {
    id: loginWindow
    width: 400
    height: 280
    property TextInput userInput: userNameInput
    property TextInput passInput: passwordInput
    property Text errText: errorText
    signal login(string userNameLogin, string passwordLogin)
    signal goBackToRegister()
    signal inputFieldsChanged()

    Rectangle {
        color: "#ffffff"
        border.color: "#363377"
        z: -1
        anchors.fill: parent
    }

    ColumnLayout {
        id: columnLayout
        anchors.fill: parent

        TextInput {
            id: userNameInput
            y: 65
            width: 220
            height: 40
            color: "#363377"
            Layout.maximumHeight: 40
            Layout.maximumWidth: 220
            font.pointSize: 20
            horizontalAlignment: Text.AlignHCenter
            verticalAlignment: Text.AlignVCenter
            Layout.minimumHeight: 40
            Layout.minimumWidth: 220
            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
            onTextChanged: inputFieldsChanged()
            property string placeholderText: "Enter username here..."
            Rectangle {
                id: rectangleUserNameInput
                color: "#827fb2"
                radius: 0
                border.width: 2
                border.color: "#585594"
                z: -1
                anchors.fill: parent
            }

            Text{
                text: userNameInput.placeholderText
                font.family: "Arial"
                anchors.fill: parent
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
                color: "#585594"
                visible: !userNameInput.text
            }

        }


        TextInput {
            id: passwordInput
            y: 175
            width: 220
            height: 40
            color: "#363377"
            echoMode: TextInput.Password
            Layout.maximumHeight: 40
            Layout.maximumWidth: 220
            horizontalAlignment: Text.AlignHCenter
            verticalAlignment: Text.AlignVCenter
            Layout.minimumHeight: 40
            Layout.minimumWidth: 220
            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
            onTextChanged: inputFieldsChanged()
            property string placeholderText: "Enter password here..."
            Rectangle {
                id: rectanglePasswordInput
                color: "#827fb2"
                radius: 0
                border.width: 2
                border.color: "#585594"
                z: -1
                anchors.fill: parent
            }
            Text{
                text: passwordInput.placeholderText
                font.family: "Arial"
                anchors.fill: parent
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
                color: "#585594"
                visible: !passwordInput.text
            }
        }

        Button {
            id: loginButton
            y: 190
            height: 20
            text: qsTr("Login")
            highlighted: false
            Layout.minimumHeight: 20
            Layout.minimumWidth: 100
            checkable: false
            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
            onClicked: loginWindow.login(userNameInput.text, passwordInput.text)
        }

        Text {
            id: errorText
            width: 220
            text: qsTr("")
            Layout.maximumHeight: 13
            Layout.maximumWidth: 220
            Layout.minimumHeight: 13
            Layout.minimumWidth: 220
            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
            verticalAlignment: Text.AlignVCenter
            horizontalAlignment: Text.AlignHCenter
            font.pixelSize: 12
        }

        Rectangle {
            id: toRegisterRectangle
            y: 230
            width: 300
            height: 40
            color: "#ffffff"
            border.color: "#827fb2"
            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
            Layout.fillWidth: false
            Text {
                id: toLoginText
                text: qsTr("Click here to swith to register page")
                anchors.fill: parent
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
                font.pixelSize: 12
            }
            MouseArea {
                anchors.fill: parent
                onClicked: {
                    goBackToRegister()
                }
            }
        }

    }


}
