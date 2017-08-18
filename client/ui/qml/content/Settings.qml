import QtQuick 2.7
import QtQuick.Controls 1.4
import QtQuick.Layouts 1.3
import QtQuick.Controls.Private 1.0
import QtQuick.Controls.Styles 1.1

Item {
    id: settingsWindow
    visible: true
    width: 720
    height: 460
    Layout.fillHeight: true
    Layout.fillWidth: true
    signal backToMessager()
    signal applyChanges(string newPassword, string oldPassword, string newNickname, date birthday, string about)

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
                Layout.fillWidth: true

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

                Button {
                    id: applyButton
                    text: qsTr("Apply")
                    onClicked: applyChanges(newPasswordTextInput.text, oldPasswordTextInput.text, nicknameTextInput.text, calendar.selectedDate, aboutEdit.text)
                }

            }


            RowLayout {
                id: newPasswordRowLayout
                width: 100
                height: 100
                Layout.fillWidth: true

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
                        width: 200
                        color: "#827fb2"
                        radius: 0
                        border.width: 2
                        border.color: "#585594"
                        z: -1
                        anchors.fill: parent
                    }

                    Text{
                        width: newPasswordTextInput.width
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
                Layout.fillWidth: true

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
                        width: 200
                        color: "#827fb2"
                        radius: 0
                        border.width: 2
                        border.color: "#585594"
                        z: -1
                        anchors.fill: parent
                    }

                    Text{
                        width: oldPasswordTextInput.width
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
                id: nickNameRowLayout
                width: 100
                height: 100

                Text {
                    id: nicknameText
                    text: qsTr("Your nickname")
                    font.pixelSize: 16
                }
                TextInput {
                    id: nicknameTextInput
                    width: 200
                    height: 20
                    text: qsTr("")
                    echoMode: TextInput.Password
                    Layout.minimumHeight: 25
                    Layout.minimumWidth: 100
                    font.pixelSize: 12
                    Rectangle {
                        id: rectangleNicknameInput
                        width: 200
                        color: "#827fb2"
                        radius: 0
                        border.width: 2
                        border.color: "#585594"
                        z: -1
                        anchors.fill: parent
                    }


                }
            }

            RowLayout {
                id: birthdayRowLayout
                width: 100
                height: 100

                Calendar {
                    id: calendar
                    width: (parent.width > parent.height ? parent.width * 0.6 - parent.spacing : parent.width)
                    height: (parent.height > parent.width ? parent.height * 0.6 - parent.spacing : parent.height)
                    frameVisible: true
                    weekNumbersVisible: true
                    selectedDate: new Date(1995, 0, 1)
                    onSelectedDateChanged: console.debug(calendar.selectedDate)
                    focus: true
                    style: CalendarStyle {
                        dayDelegate: Item {
                            readonly property color sameMonthDateTextColor: "#444"
                            readonly property color selectedDateColor: Qt.platform.os === "osx" ? "#3778d0" : systemPalette.highlight
                            readonly property color selectedDateTextColor: "white"
                            readonly property color differentMonthDateTextColor: "#bbb"
                            readonly property color invalidDatecolor: "#dddddd"

                            Rectangle {
                                anchors.fill: parent
                                border.color: "transparent"
                                color: styleData.date !== undefined && styleData.selected ? selectedDateColor : "transparent"
                                anchors.margins: styleData.selected ? -1 : 0
                            }



                            Label {
                                id: dayDelegateText
                                text: styleData.date.getDate()
                                anchors.centerIn: parent
                                color: {
                                    var color = invalidDatecolor;
                                    if (styleData.valid) {
                                        // Date is within the valid range.
                                        color = styleData.visibleMonth ? sameMonthDateTextColor : differentMonthDateTextColor;
                                        if (styleData.selected) {
                                            color = selectedDateTextColor;
                                        }
                                    }
                                    color;
                                }
                            }
                        }
                    }
                }

                Component {
                    id: eventListHeader

                    Row {
                        id: eventDateRow
                        width: parent.width
                        height: eventDayLabel.height
                        spacing: 10

                        Label {
                            id: eventDayLabel
                            text: calendar.selectedDate.getDate()
                            font.pointSize: 35
                        }

                        Column {
                            height: eventDayLabel.height

                            Label {
                                readonly property var options: { weekday: "long" }
                                text: Qt.locale().standaloneDayName(calendar.selectedDate.getDay(), Locale.LongFormat)
                                font.pointSize: 18
                            }
                            Label {
                                text: Qt.locale().standaloneMonthName(calendar.selectedDate.getMonth())
                                      + calendar.selectedDate.toLocaleDateString(Qt.locale(), " yyyy")
                                font.pointSize: 12
                            }
                        }
                    }
                }
            }

            RowLayout {
                id: aboutRowLayout
                width: 100
                height: 100
                Flickable {
                    id: flick
                    maximumFlickVelocity: 500
                    Layout.minimumHeight: 60
                    Layout.maximumHeight: 60
                    Layout.minimumWidth: 200
                    Layout.fillWidth: true
                    contentWidth: aboutEdit.paintedWidth
                    contentHeight: aboutEdit.paintedHeight
                    clip: true
                    //ScrollBar.vertical: ScrollBar { id: vbar; active: hbar.active }
                    function ensureVisible(r)
                    {
                        if (contentX >= r.x)
                            contentX = r.x;
                        else if (contentX+width <= r.x+r.width)
                            contentX = r.x+r.width-width;
                        if (contentY >= r.y)
                            contentY = r.y;
                        else if (contentY+height <= r.y+r.height)
                            contentY = r.y+r.height-height;
                    }
                    TextEdit {
                        id: aboutEdit
                        width: flick.width
                        height: flick.height
                        text: ""
                        horizontalAlignment: Text.AlignLeft
                        selectionColor: "#817fbe"
                        selectByMouse: true
                        textFormat: Text.RichText
                        focus: true
                        wrapMode: TextEdit.Wrap
                        onCursorRectangleChanged: flick.ensureVisible(cursorRectangle)
                    }
                }
            }
        }
    }

}
