import QtQuick 2.7
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.1
import Qt.labs.calendar 1.0

Item {
    id: messagerWindow
    visible: true
    width: 720
    height: 480
    Layout.fillHeight: true
    Layout.fillWidth: true
    signal send(string message)

    Rectangle {
        color: "#ffffff"
        z: -1
        anchors.fill: parent
    }

    RowLayout {
        id: mainRowLayout
        anchors.rightMargin: 5
        anchors.leftMargin: 5
        anchors.topMargin: 5
        anchors.bottomMargin: 5
        anchors.fill: parent

        ColumnLayout {
            id: leftColumnLayout
            width: 100
            height: 100
            Layout.maximumHeight: 1000
            Layout.maximumWidth: 1000
            Layout.fillHeight: true
            Layout.fillWidth: true
            Layout.minimumHeight: 200
            Layout.minimumWidth: 250
            Layout.columnSpan: 1
            Layout.rowSpan: 1

            Text {
                id: currentTitle
                height: 20
                text: qsTr("Current interlocutor")
                Layout.minimumHeight: 20
                Layout.minimumWidth: 230
                fontSizeMode: Text.Fit
                Layout.fillWidth: true
                wrapMode: Text.WrapAnywhere
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
                font.pixelSize: 12
            }

            Text {
                id: historyTextView
                width: 560
                textFormat: Text.RichText
                verticalAlignment: Text.AlignBottom
                Layout.maximumWidth: 5535
                Layout.rowSpan: 2
                renderType: Text.QtRendering
                enabled: true
                z: 0
                antialiasing: true
                wrapMode: Text.WrapAtWordBoundaryOrAnywhere
                font.weight: Font.Normal
                style: Text.Normal
                font.letterSpacing: 0
                font.wordSpacing: 1
                rightPadding: 5
                leftPadding: 5
                bottomPadding: 0
                topPadding: 0
                fontSizeMode: Text.FixedSize
                font.family: "Courier"
                Layout.fillWidth: true
                Layout.fillHeight: true
                Layout.minimumHeight: 200
                Layout.minimumWidth: 230
                font.pixelSize: 12
            }

            RowLayout {
                id: sendingRowLayout
                width: 100
                height: 100
                Layout.alignment: Qt.AlignLeft | Qt.AlignBottom
                Layout.maximumHeight: 70
                Layout.maximumWidth: 2000
                Layout.minimumHeight: 70
                Layout.minimumWidth: 200
                Layout.columnSpan: 1
                Layout.rowSpan: 1
                transformOrigin: Item.Center
                Layout.fillHeight: false
                Layout.fillWidth: true

                Flickable {
                    id: flick
                    maximumFlickVelocity: 500
                    Layout.minimumHeight: 60
                    Layout.maximumHeight: 60
                    Layout.minimumWidth: 200
                    Layout.fillWidth: true
                    contentWidth: messageEdit.paintedWidth
                    contentHeight: messageEdit.paintedHeight
                    clip: true
                    ScrollBar.vertical: ScrollBar { id: vbar; active: hbar.active }
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
                        id: messageEdit
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

                Button {
                    id: sendButton
                    x: 0
                    y: 0
                    width: 75
                    height: 50
                    text: qsTr("Send")
                    anchors.right: parent.right
                    anchors.rightMargin: 0
                    Layout.maximumHeight: 60
                    Layout.maximumWidth: 75
                    Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                    z: 0
                    scale: 1
                    Layout.minimumHeight: 60
                    Layout.minimumWidth: 75
                    Layout.fillHeight: false
                    Layout.fillWidth: false
                    onClicked: {
                        if( messageEdit.getText(0, messageEdit.length) !== "") {
                            messagerWindow.send(messageEdit.getText(0, messageEdit.length))
                            messageEdit.text = ""
                        }
                    }
                }
            }


        }

        ColumnLayout {
            id: rightColumnLayout
            width: 110
            height: 300
            clip: false
            Layout.maximumWidth: 170
            visible: true
            Layout.minimumHeight: 300
            Layout.minimumWidth: 170
            Layout.fillHeight: true
            Layout.fillWidth: true

            RowLayout {
                id: searchingRowLayout
                width: 100
                height: 100
                Layout.maximumHeight: 30
                Layout.minimumHeight: 30

                TextInput {
                    id: searchTextInput
                    width: 80
                    height: 20
                    text: qsTr("Text Input")
                    Layout.maximumHeight: 30
                    Layout.minimumHeight: 30
                    Layout.minimumWidth: 130
                    Layout.fillWidth: true
                    font.pixelSize: 13
                }

                Button {
                    id: button
                    text: qsTr("S")
                    enabled: true
                    autoExclusive: false
                    checked: false
                    checkable: false
                    highlighted: false
                    Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                    Layout.maximumHeight: 30
                    Layout.minimumHeight: 30
                    Layout.maximumWidth: 30
                    Layout.minimumWidth: 30
                }
            }

            ListView {
                id: contactsListView
                width: 110
                height: 160
                Layout.minimumHeight: 300
                Layout.minimumWidth: 150
                highlightRangeMode: ListView.NoHighlightRange
                orientation: ListView.Vertical
                maximumFlickVelocity: 2492
                Layout.fillWidth: true
                Layout.fillHeight: true
                delegate: Item {
                    x: 5
                    width: 80
                    height: 40
                    Row {
                        id: row1
                        Rectangle {
                            width: 40
                            height: 40
                            color: colorCode
                        }

                        Text {
                            text: name
                            anchors.verticalCenter: parent.verticalCenter
                            font.bold: true
                        }
                        spacing: 10
                    }
                }
                model: ListModel {
                    ListElement {
                        name: "Bohdan"
                        colorCode: "grey"
                    }

                    ListElement {
                        name: "Vitaliy"
                        colorCode: "red"
                    }

                    ListElement {
                        name: "Valeriy"
                        colorCode: "blue"
                    }

                    ListElement {
                        name: "Volodymyr"
                        colorCode: "green"
                    }
                }
            }

        }

    }
}
