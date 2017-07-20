

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QCamera>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLayout>
#include <QList>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
#include <QSignalSpy>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QWidget>
#include <QWindow>


class QmlContacts: public QObject
{
Q_OBJECT
public:
	QmlContacts(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QmlContacts_QmlContacts_QRegisterMetaType();QmlContacts_QmlContacts_QRegisterMetaTypes();callbackQmlContacts_Constructor(this);};
	void Signal_SendLastMessage(QString lastMessage) { QByteArray tb39d3c = lastMessage.toUtf8(); Moc_PackedString lastMessagePacked = { const_cast<char*>(tb39d3c.prepend("WHITESPACE").constData()+10), tb39d3c.size()-10 };callbackQmlContacts_SendLastMessage(this, lastMessagePacked); };
	void Signal_SendHistory(QString history) { QByteArray t66f79d = history.toUtf8(); Moc_PackedString historyPacked = { const_cast<char*>(t66f79d.prepend("WHITESPACE").constData()+10), t66f79d.size()-10 };callbackQmlContacts_SendHistory(this, historyPacked); };
	 ~QmlContacts() { callbackQmlContacts_DestroyQmlContacts(this); };
	bool event(QEvent * e) { return callbackQmlContacts_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQmlContacts_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQmlContacts_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQmlContacts_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQmlContacts_CustomEvent(this, event); };
	void deleteLater() { callbackQmlContacts_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQmlContacts_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQmlContacts_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQmlContacts_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQmlContacts_TimerEvent(this, event); };
	
signals:
	void sendLastMessage(QString lastMessage);
	void sendHistory(QString history);
public slots:
	void changeCurrentContact(qint32 newIndex) { callbackQmlContacts_ChangeCurrentContact(this, newIndex); };
	void searchUser(QString searchedUser) { QByteArray t5dc753 = searchedUser.toUtf8(); Moc_PackedString searchedUserPacked = { const_cast<char*>(t5dc753.prepend("WHITESPACE").constData()+10), t5dc753.size()-10 };callbackQmlContacts_SearchUser(this, searchedUserPacked); };
private:
};

Q_DECLARE_METATYPE(QmlContacts*)


void QmlContacts_QmlContacts_QRegisterMetaTypes() {
}

class QmlLogin: public QObject
{
Q_OBJECT
public:
	QmlLogin(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QmlLogin_QmlLogin_QRegisterMetaType();QmlLogin_QmlLogin_QRegisterMetaTypes();callbackQmlLogin_Constructor(this);};
	void Signal_LoginDataIsValid(bool isLoginValid) { callbackQmlLogin_LoginDataIsValid(this, isLoginValid); };
	 ~QmlLogin() { callbackQmlLogin_DestroyQmlLogin(this); };
	bool event(QEvent * e) { return callbackQmlLogin_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQmlLogin_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQmlLogin_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQmlLogin_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQmlLogin_CustomEvent(this, event); };
	void deleteLater() { callbackQmlLogin_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQmlLogin_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQmlLogin_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQmlLogin_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQmlLogin_TimerEvent(this, event); };
	
signals:
	void loginDataIsValid(bool isLoginValid);
public slots:
	void checkLoginData(QString userName, QString password) { QByteArray t2d867a = userName.toUtf8(); Moc_PackedString userNamePacked = { const_cast<char*>(t2d867a.prepend("WHITESPACE").constData()+10), t2d867a.size()-10 };QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackQmlLogin_CheckLoginData(this, userNamePacked, passwordPacked); };
private:
};

Q_DECLARE_METATYPE(QmlLogin*)


void QmlLogin_QmlLogin_QRegisterMetaTypes() {
}

class QmlMessage: public QObject
{
Q_OBJECT
public:
	QmlMessage(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QmlMessage_QmlMessage_QRegisterMetaType();QmlMessage_QmlMessage_QRegisterMetaTypes();callbackQmlMessage_Constructor(this);};
	 ~QmlMessage() { callbackQmlMessage_DestroyQmlMessage(this); };
	bool event(QEvent * e) { return callbackQmlMessage_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQmlMessage_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQmlMessage_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQmlMessage_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQmlMessage_CustomEvent(this, event); };
	void deleteLater() { callbackQmlMessage_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQmlMessage_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQmlMessage_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQmlMessage_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQmlMessage_TimerEvent(this, event); };
	
signals:
public slots:
	void sendMessage(QString message) { QByteArray t6f9b9a = message.toUtf8(); Moc_PackedString messagePacked = { const_cast<char*>(t6f9b9a.prepend("WHITESPACE").constData()+10), t6f9b9a.size()-10 };callbackQmlMessage_SendMessage(this, messagePacked); };
private:
};

Q_DECLARE_METATYPE(QmlMessage*)


void QmlMessage_QmlMessage_QRegisterMetaTypes() {
}

class QmlRegister: public QObject
{
Q_OBJECT
public:
	QmlRegister(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QmlRegister_QmlRegister_QRegisterMetaType();QmlRegister_QmlRegister_QRegisterMetaTypes();callbackQmlRegister_Constructor(this);};
	void Signal_RegisterDataIsValid(bool isRegisterValid) { callbackQmlRegister_RegisterDataIsValid(this, isRegisterValid); };
	 ~QmlRegister() { callbackQmlRegister_DestroyQmlRegister(this); };
	bool event(QEvent * e) { return callbackQmlRegister_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQmlRegister_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQmlRegister_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQmlRegister_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQmlRegister_CustomEvent(this, event); };
	void deleteLater() { callbackQmlRegister_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQmlRegister_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQmlRegister_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQmlRegister_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQmlRegister_TimerEvent(this, event); };
	
signals:
	void registerDataIsValid(bool isRegisterValid);
public slots:
	void checkRegisterData(QString userName, QString nickName, QString password) { QByteArray t2d867a = userName.toUtf8(); Moc_PackedString userNamePacked = { const_cast<char*>(t2d867a.prepend("WHITESPACE").constData()+10), t2d867a.size()-10 };QByteArray t2307d9 = nickName.toUtf8(); Moc_PackedString nickNamePacked = { const_cast<char*>(t2307d9.prepend("WHITESPACE").constData()+10), t2307d9.size()-10 };QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackQmlRegister_CheckRegisterData(this, userNamePacked, nickNamePacked, passwordPacked); };
private:
};

Q_DECLARE_METATYPE(QmlRegister*)


void QmlRegister_QmlRegister_QRegisterMetaTypes() {
}

void QmlRegister_CheckRegisterData(void* ptr, struct Moc_PackedString userName, struct Moc_PackedString nickName, struct Moc_PackedString password)
{
	QMetaObject::invokeMethod(static_cast<QmlRegister*>(ptr), "checkRegisterData", Q_ARG(QString, QString::fromUtf8(userName.data, userName.len)), Q_ARG(QString, QString::fromUtf8(nickName.data, nickName.len)), Q_ARG(QString, QString::fromUtf8(password.data, password.len)));
}

void QmlRegister_ConnectRegisterDataIsValid(void* ptr)
{
	QObject::connect(static_cast<QmlRegister*>(ptr), static_cast<void (QmlRegister::*)(bool)>(&QmlRegister::registerDataIsValid), static_cast<QmlRegister*>(ptr), static_cast<void (QmlRegister::*)(bool)>(&QmlRegister::Signal_RegisterDataIsValid));
}

void QmlRegister_DisconnectRegisterDataIsValid(void* ptr)
{
	QObject::disconnect(static_cast<QmlRegister*>(ptr), static_cast<void (QmlRegister::*)(bool)>(&QmlRegister::registerDataIsValid), static_cast<QmlRegister*>(ptr), static_cast<void (QmlRegister::*)(bool)>(&QmlRegister::Signal_RegisterDataIsValid));
}

void QmlRegister_RegisterDataIsValid(void* ptr, char isRegisterValid)
{
	static_cast<QmlRegister*>(ptr)->registerDataIsValid(isRegisterValid != 0);
}

int QmlRegister_QmlRegister_QRegisterMetaType()
{
	return qRegisterMetaType<QmlRegister*>();
}

int QmlRegister_QmlRegister_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QmlRegister*>(const_cast<const char*>(typeName));
}

int QmlRegister_QmlRegister_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlRegister>();
#else
	return 0;
#endif
}

int QmlRegister_QmlRegister_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlRegister>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QmlRegister___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QmlRegister___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QmlRegister___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QmlRegister___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QmlRegister___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlRegister___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QmlRegister___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QmlRegister___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlRegister___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QmlRegister___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QmlRegister___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlRegister___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QmlRegister___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QmlRegister___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlRegister___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* QmlRegister_NewQmlRegister(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new QmlRegister(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new QmlRegister(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new QmlRegister(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new QmlRegister(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new QmlRegister(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new QmlRegister(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new QmlRegister(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new QmlRegister(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new QmlRegister(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QmlRegister(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QmlRegister(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QmlRegister(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new QmlRegister(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new QmlRegister(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new QmlRegister(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new QmlRegister(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QmlRegister(static_cast<QWindow*>(parent));
	} else {
		return new QmlRegister(static_cast<QObject*>(parent));
	}
}

void QmlRegister_DestroyQmlRegister(void* ptr)
{
	static_cast<QmlRegister*>(ptr)->~QmlRegister();
}

void QmlRegister_DestroyQmlRegisterDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char QmlRegister_EventDefault(void* ptr, void* e)
{
	return static_cast<QmlRegister*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QmlRegister_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QmlRegister*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QmlRegister_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QmlRegister*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QmlRegister_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlRegister*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlRegister_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QmlRegister*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QmlRegister_DeleteLaterDefault(void* ptr)
{
	static_cast<QmlRegister*>(ptr)->QObject::deleteLater();
}

void QmlRegister_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlRegister*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlRegister_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QmlRegister*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



void QmlContacts_ChangeCurrentContact(void* ptr, int newIndex)
{
	QMetaObject::invokeMethod(static_cast<QmlContacts*>(ptr), "changeCurrentContact", Q_ARG(qint32, newIndex));
}

void QmlContacts_SearchUser(void* ptr, struct Moc_PackedString searchedUser)
{
	QMetaObject::invokeMethod(static_cast<QmlContacts*>(ptr), "searchUser", Q_ARG(QString, QString::fromUtf8(searchedUser.data, searchedUser.len)));
}

void QmlContacts_ConnectSendLastMessage(void* ptr)
{
	QObject::connect(static_cast<QmlContacts*>(ptr), static_cast<void (QmlContacts::*)(QString)>(&QmlContacts::sendLastMessage), static_cast<QmlContacts*>(ptr), static_cast<void (QmlContacts::*)(QString)>(&QmlContacts::Signal_SendLastMessage));
}

void QmlContacts_DisconnectSendLastMessage(void* ptr)
{
	QObject::disconnect(static_cast<QmlContacts*>(ptr), static_cast<void (QmlContacts::*)(QString)>(&QmlContacts::sendLastMessage), static_cast<QmlContacts*>(ptr), static_cast<void (QmlContacts::*)(QString)>(&QmlContacts::Signal_SendLastMessage));
}

void QmlContacts_SendLastMessage(void* ptr, struct Moc_PackedString lastMessage)
{
	static_cast<QmlContacts*>(ptr)->sendLastMessage(QString::fromUtf8(lastMessage.data, lastMessage.len));
}

void QmlContacts_ConnectSendHistory(void* ptr)
{
	QObject::connect(static_cast<QmlContacts*>(ptr), static_cast<void (QmlContacts::*)(QString)>(&QmlContacts::sendHistory), static_cast<QmlContacts*>(ptr), static_cast<void (QmlContacts::*)(QString)>(&QmlContacts::Signal_SendHistory));
}

void QmlContacts_DisconnectSendHistory(void* ptr)
{
	QObject::disconnect(static_cast<QmlContacts*>(ptr), static_cast<void (QmlContacts::*)(QString)>(&QmlContacts::sendHistory), static_cast<QmlContacts*>(ptr), static_cast<void (QmlContacts::*)(QString)>(&QmlContacts::Signal_SendHistory));
}

void QmlContacts_SendHistory(void* ptr, struct Moc_PackedString history)
{
	static_cast<QmlContacts*>(ptr)->sendHistory(QString::fromUtf8(history.data, history.len));
}

int QmlContacts_QmlContacts_QRegisterMetaType()
{
	return qRegisterMetaType<QmlContacts*>();
}

int QmlContacts_QmlContacts_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QmlContacts*>(const_cast<const char*>(typeName));
}

int QmlContacts_QmlContacts_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlContacts>();
#else
	return 0;
#endif
}

int QmlContacts_QmlContacts_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlContacts>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QmlContacts___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QmlContacts___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QmlContacts___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QmlContacts___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QmlContacts___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlContacts___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QmlContacts___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QmlContacts___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlContacts___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QmlContacts___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QmlContacts___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlContacts___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QmlContacts___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QmlContacts___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlContacts___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* QmlContacts_NewQmlContacts(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new QmlContacts(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new QmlContacts(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new QmlContacts(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new QmlContacts(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new QmlContacts(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new QmlContacts(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new QmlContacts(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new QmlContacts(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new QmlContacts(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QmlContacts(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QmlContacts(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QmlContacts(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new QmlContacts(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new QmlContacts(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new QmlContacts(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new QmlContacts(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QmlContacts(static_cast<QWindow*>(parent));
	} else {
		return new QmlContacts(static_cast<QObject*>(parent));
	}
}

void QmlContacts_DestroyQmlContacts(void* ptr)
{
	static_cast<QmlContacts*>(ptr)->~QmlContacts();
}

void QmlContacts_DestroyQmlContactsDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char QmlContacts_EventDefault(void* ptr, void* e)
{
	return static_cast<QmlContacts*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QmlContacts_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QmlContacts*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QmlContacts_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QmlContacts*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QmlContacts_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlContacts*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlContacts_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QmlContacts*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QmlContacts_DeleteLaterDefault(void* ptr)
{
	static_cast<QmlContacts*>(ptr)->QObject::deleteLater();
}

void QmlContacts_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlContacts*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlContacts_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QmlContacts*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



void QmlLogin_CheckLoginData(void* ptr, struct Moc_PackedString userName, struct Moc_PackedString password)
{
	QMetaObject::invokeMethod(static_cast<QmlLogin*>(ptr), "checkLoginData", Q_ARG(QString, QString::fromUtf8(userName.data, userName.len)), Q_ARG(QString, QString::fromUtf8(password.data, password.len)));
}

void QmlLogin_ConnectLoginDataIsValid(void* ptr)
{
	QObject::connect(static_cast<QmlLogin*>(ptr), static_cast<void (QmlLogin::*)(bool)>(&QmlLogin::loginDataIsValid), static_cast<QmlLogin*>(ptr), static_cast<void (QmlLogin::*)(bool)>(&QmlLogin::Signal_LoginDataIsValid));
}

void QmlLogin_DisconnectLoginDataIsValid(void* ptr)
{
	QObject::disconnect(static_cast<QmlLogin*>(ptr), static_cast<void (QmlLogin::*)(bool)>(&QmlLogin::loginDataIsValid), static_cast<QmlLogin*>(ptr), static_cast<void (QmlLogin::*)(bool)>(&QmlLogin::Signal_LoginDataIsValid));
}

void QmlLogin_LoginDataIsValid(void* ptr, char isLoginValid)
{
	static_cast<QmlLogin*>(ptr)->loginDataIsValid(isLoginValid != 0);
}

int QmlLogin_QmlLogin_QRegisterMetaType()
{
	return qRegisterMetaType<QmlLogin*>();
}

int QmlLogin_QmlLogin_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QmlLogin*>(const_cast<const char*>(typeName));
}

int QmlLogin_QmlLogin_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlLogin>();
#else
	return 0;
#endif
}

int QmlLogin_QmlLogin_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlLogin>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QmlLogin___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QmlLogin___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QmlLogin___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QmlLogin___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QmlLogin___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlLogin___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QmlLogin___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QmlLogin___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlLogin___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QmlLogin___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QmlLogin___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlLogin___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QmlLogin___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QmlLogin___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlLogin___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* QmlLogin_NewQmlLogin(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new QmlLogin(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new QmlLogin(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new QmlLogin(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new QmlLogin(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new QmlLogin(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new QmlLogin(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new QmlLogin(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new QmlLogin(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new QmlLogin(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QmlLogin(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QmlLogin(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QmlLogin(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new QmlLogin(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new QmlLogin(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new QmlLogin(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new QmlLogin(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QmlLogin(static_cast<QWindow*>(parent));
	} else {
		return new QmlLogin(static_cast<QObject*>(parent));
	}
}

void QmlLogin_DestroyQmlLogin(void* ptr)
{
	static_cast<QmlLogin*>(ptr)->~QmlLogin();
}

void QmlLogin_DestroyQmlLoginDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char QmlLogin_EventDefault(void* ptr, void* e)
{
	return static_cast<QmlLogin*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QmlLogin_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QmlLogin*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QmlLogin_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QmlLogin*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QmlLogin_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlLogin*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlLogin_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QmlLogin*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QmlLogin_DeleteLaterDefault(void* ptr)
{
	static_cast<QmlLogin*>(ptr)->QObject::deleteLater();
}

void QmlLogin_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlLogin*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlLogin_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QmlLogin*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



void QmlMessage_SendMessage(void* ptr, struct Moc_PackedString message)
{
	QMetaObject::invokeMethod(static_cast<QmlMessage*>(ptr), "sendMessage", Q_ARG(QString, QString::fromUtf8(message.data, message.len)));
}

int QmlMessage_QmlMessage_QRegisterMetaType()
{
	return qRegisterMetaType<QmlMessage*>();
}

int QmlMessage_QmlMessage_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QmlMessage*>(const_cast<const char*>(typeName));
}

int QmlMessage_QmlMessage_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlMessage>();
#else
	return 0;
#endif
}

int QmlMessage_QmlMessage_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlMessage>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QmlMessage___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QmlMessage___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QmlMessage___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QmlMessage___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QmlMessage___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlMessage___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QmlMessage___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QmlMessage___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlMessage___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QmlMessage___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QmlMessage___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlMessage___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QmlMessage___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QmlMessage___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlMessage___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* QmlMessage_NewQmlMessage(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new QmlMessage(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new QmlMessage(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new QmlMessage(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new QmlMessage(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new QmlMessage(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new QmlMessage(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new QmlMessage(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new QmlMessage(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new QmlMessage(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QmlMessage(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QmlMessage(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QmlMessage(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new QmlMessage(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new QmlMessage(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new QmlMessage(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new QmlMessage(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QmlMessage(static_cast<QWindow*>(parent));
	} else {
		return new QmlMessage(static_cast<QObject*>(parent));
	}
}

void QmlMessage_DestroyQmlMessage(void* ptr)
{
	static_cast<QmlMessage*>(ptr)->~QmlMessage();
}

void QmlMessage_DestroyQmlMessageDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char QmlMessage_EventDefault(void* ptr, void* e)
{
	return static_cast<QmlMessage*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QmlMessage_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QmlMessage*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QmlMessage_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QmlMessage*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QmlMessage_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlMessage*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlMessage_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QmlMessage*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QmlMessage_DeleteLaterDefault(void* ptr)
{
	static_cast<QmlMessage*>(ptr)->QObject::deleteLater();
}

void QmlMessage_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlMessage*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlMessage_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QmlMessage*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



#include "moc_moc.h"
