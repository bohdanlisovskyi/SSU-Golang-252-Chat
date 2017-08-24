import QtQuick 2.7
import QtQuick.Controls 2.1
import QtQuick.Layouts 1.3

Item {
    id: registerWindow
    property TextInput userInput: userNameInput
    property TextInput nickInput: nickNameInput
    property TextInput passInput: passwordInput
    property Text errText: errorText
    signal register(string userName, string nickName, string password)
    signal goNextToLogin()
    signal inputFieldsChanged()
    width: 400
    height: 280

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
            id: nickNameInput
            y: 120
            width: 220
            height: 40
            color: "#363377"
            font.pointSize: 20
            Layout.maximumHeight: 40
            Layout.maximumWidth: 220
            horizontalAlignment: Text.AlignHCenter
            verticalAlignment: Text.AlignVCenter
            Layout.minimumHeight: 40
            Layout.minimumWidth: 220
            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
            onTextChanged: inputFieldsChanged()
            property string placeholderText: "Enter first- and lastname here..."
            Rectangle {
                id: rectangleNickNameInput
                color: "#827fb2"
                radius: 0
                border.width: 2
                border.color: "#585594"
                z: -1
                anchors.fill: parent
            }
            Text{
                text: nickNameInput.placeholderText
                font.family: "Arial"
                anchors.fill: parent
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
                color: "#585594"
                visible: !nickNameInput.text
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
            id: registerButton
            y: 190
            height: 30
            text: qsTr("Register")
            Layout.maximumHeight: 30
            Layout.maximumWidth: 100
            spacing: 10
            autoRepeat: false
            autoExclusive: false
            checked: false
            highlighted: false
            enabled: true
            Layout.minimumHeight: 30
            Layout.minimumWidth: 100
            checkable: false
            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
            onClicked: registerWindow.register(userNameInput.text, nickNameInput.text, passwordInput.text)
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
            id: toLoginRectangle
            y: 230
            width: 300
            height: 40
            color: "#ffffff"
            border.color: "#827fb2"
            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
            Layout.fillWidth: false
            Text {
                id: toLoginText
                text: qsTr("Click here to swith to login page")
                anchors.fill: parent
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
                font.pixelSize: 12
            }
            MouseArea {
                anchors.fill: parent
                onClicked: {
                    goNextToLogin()
                }
            }
        }
    }
}
