import QtQuick 2.7
import QtQuick.Controls 2.1
import QtQuick.Layouts 1.3

Item {
    id: settingsWindow
    visible: true
    width: 720
    height: 460
    Layout.fillHeight: true
    Layout.fillWidth: true
    signal backToMessager()

    RowLayout {
        id: rowLayout
        x: 8
        y: 8
        width: 704
        height: 444

        ColumnLayout {
            id: columnLayout
            width: 100
            height: 100
            Layout.fillHeight: true
            Layout.fillWidth: true

            RowLayout {
                id: topRowLayout
                width: 100
                height: 100

                Button {
                    id: backButton
                    text: qsTr("Back")
                    onClicked: backToMessager()
                }

                Text {
                    id: nickNameText
                    text: "Your name"
                    Layout.fillWidth: true
                    font.pixelSize: 30
                    style: Text.Normal
                    verticalAlignment: Text.AlignTop
                    horizontalAlignment: Text.AlignHCenter
                }

            }

            RowLayout {
                id: newPasswordRowLayout
                width: 100
                height: 100

                Text {
                    id: newPasswordText
                    text: qsTr("New password")
                    font.pixelSize: 16
                }

                TextInput {
                    id: newPasswordTextInput
                    width: 200
                    height: 20
                    text: qsTr("")
                    echoMode: TextInput.Password
                    font.pixelSize: 12
                    Layout.minimumHeight: 25
                    Layout.minimumWidth: 100
                    property string placeholderText: "Enter new password here..."
                    Rectangle {
                        id: rectangleNewPasswordInput
                        color: "#827fb2"
                        radius: 0
                        border.width: 2
                        border.color: "#585594"
                        z: -1
                        anchors.fill: parent
                    }

                    Text{
                        text: newPasswordTextInput.placeholderText
                        font.family: "Arial"
                        anchors.fill: parent
                        verticalAlignment: Text.AlignVCenter
                        horizontalAlignment: Text.AlignHCenter
                        color: "#585594"
                        visible: !newPasswordTextInput.text
                    }
                }
            }

            RowLayout {
                id: oldPasswordColumnLayout
                width: 100
                height: 100

                Text {
                    id: oldPasswordText
                    text: qsTr("Old password")
                    font.pixelSize: 16
                }

                TextInput {
                    id: oldPasswordTextInput
                    width: 200
                    height: 20
                    text: qsTr("")
                    echoMode: TextInput.Password
                    Layout.minimumHeight: 25
                    Layout.minimumWidth: 100
                    font.pixelSize: 12
                    property string placeholderText: "Enter old password here..."
                    Rectangle {
                        id: rectangleOldPasswordInput
                        color: "#827fb2"
                        radius: 0
                        border.width: 2
                        border.color: "#585594"
                        z: -1
                        anchors.fill: parent
                    }

                    Text{
                        text: oldPasswordTextInput.placeholderText
                        font.family: "Arial"
                        anchors.fill: parent
                        verticalAlignment: Text.AlignVCenter
                        horizontalAlignment: Text.AlignHCenter
                        color: "#585594"
                        visible: !oldPasswordTextInput.text
                    }
                }
            }


            RowLayout {
                id: textColorRowLayout
                width: 100
                height: 100
            }


        }
    }

}
