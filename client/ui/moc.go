package ui

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type QmlContacts_ITF interface {
	std_core.QObject_ITF
	QmlContacts_PTR() *QmlContacts
}

func (ptr *QmlContacts) QmlContacts_PTR() *QmlContacts {
	return ptr
}

func (ptr *QmlContacts) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlContacts) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQmlContacts(ptr QmlContacts_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlContacts_PTR().Pointer()
	}
	return nil
}

func NewQmlContactsFromPointer(ptr unsafe.Pointer) *QmlContacts {
	var n *QmlContacts
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(QmlContacts)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *QmlContacts:
			n = deduced

		case *std_core.QObject:
			n = &QmlContacts{QObject: *deduced}

		default:
			n = new(QmlContacts)
			n.SetPointer(ptr)
		}
	}
	return n
}

//export callbackQmlContacts_Constructor
func callbackQmlContacts_Constructor(ptr unsafe.Pointer) {
	gPtr := NewQmlContactsFromPointer(ptr)
	qt.Register(ptr, gPtr)
}

//export callbackQmlContacts_ChangeCurrentContact
func callbackQmlContacts_ChangeCurrentContact(ptr unsafe.Pointer, newIndex C.int) {
	if signal := qt.GetSignal(ptr, "changeCurrentContact"); signal != nil {
		signal.(func(int))(int(int32(newIndex)))
	}

}

func (ptr *QmlContacts) ConnectChangeCurrentContact(f func(newIndex int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "changeCurrentContact"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "changeCurrentContact", func(newIndex int) {
				signal.(func(int))(newIndex)
				f(newIndex)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "changeCurrentContact", f)
		}
	}
}

func (ptr *QmlContacts) DisconnectChangeCurrentContact() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "changeCurrentContact")
	}
}

func (ptr *QmlContacts) ChangeCurrentContact(newIndex int) {
	if ptr.Pointer() != nil {
		C.QmlContacts_ChangeCurrentContact(ptr.Pointer(), C.int(int32(newIndex)))
	}
}

//export callbackQmlContacts_SearchUser
func callbackQmlContacts_SearchUser(ptr unsafe.Pointer, searchedUser C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "searchUser"); signal != nil {
		signal.(func(string))(cGoUnpackString(searchedUser))
	}

}

func (ptr *QmlContacts) ConnectSearchUser(f func(searchedUser string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "searchUser"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "searchUser", func(searchedUser string) {
				signal.(func(string))(searchedUser)
				f(searchedUser)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "searchUser", f)
		}
	}
}

func (ptr *QmlContacts) DisconnectSearchUser() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "searchUser")
	}
}

func (ptr *QmlContacts) SearchUser(searchedUser string) {
	if ptr.Pointer() != nil {
		var searchedUserC *C.char
		if searchedUser != "" {
			searchedUserC = C.CString(searchedUser)
			defer C.free(unsafe.Pointer(searchedUserC))
		}
		C.QmlContacts_SearchUser(ptr.Pointer(), C.struct_Moc_PackedString{data: searchedUserC, len: C.longlong(len(searchedUser))})
	}
}

//export callbackQmlContacts_SendLastMessage
func callbackQmlContacts_SendLastMessage(ptr unsafe.Pointer, lastMessage C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "sendLastMessage"); signal != nil {
		signal.(func(string))(cGoUnpackString(lastMessage))
	}

}

func (ptr *QmlContacts) ConnectSendLastMessage(f func(lastMessage string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sendLastMessage") {
			C.QmlContacts_ConnectSendLastMessage(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sendLastMessage"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "sendLastMessage", func(lastMessage string) {
				signal.(func(string))(lastMessage)
				f(lastMessage)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sendLastMessage", f)
		}
	}
}

func (ptr *QmlContacts) DisconnectSendLastMessage() {
	if ptr.Pointer() != nil {
		C.QmlContacts_DisconnectSendLastMessage(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sendLastMessage")
	}
}

func (ptr *QmlContacts) SendLastMessage(lastMessage string) {
	if ptr.Pointer() != nil {
		var lastMessageC *C.char
		if lastMessage != "" {
			lastMessageC = C.CString(lastMessage)
			defer C.free(unsafe.Pointer(lastMessageC))
		}
		C.QmlContacts_SendLastMessage(ptr.Pointer(), C.struct_Moc_PackedString{data: lastMessageC, len: C.longlong(len(lastMessage))})
	}
}

//export callbackQmlContacts_SendHistory
func callbackQmlContacts_SendHistory(ptr unsafe.Pointer, history C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "sendHistory"); signal != nil {
		signal.(func(string))(cGoUnpackString(history))
	}

}

func (ptr *QmlContacts) ConnectSendHistory(f func(history string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sendHistory") {
			C.QmlContacts_ConnectSendHistory(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sendHistory"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "sendHistory", func(history string) {
				signal.(func(string))(history)
				f(history)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sendHistory", f)
		}
	}
}

func (ptr *QmlContacts) DisconnectSendHistory() {
	if ptr.Pointer() != nil {
		C.QmlContacts_DisconnectSendHistory(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sendHistory")
	}
}

func (ptr *QmlContacts) SendHistory(history string) {
	if ptr.Pointer() != nil {
		var historyC *C.char
		if history != "" {
			historyC = C.CString(history)
			defer C.free(unsafe.Pointer(historyC))
		}
		C.QmlContacts_SendHistory(ptr.Pointer(), C.struct_Moc_PackedString{data: historyC, len: C.longlong(len(history))})
	}
}

func QmlContacts_QRegisterMetaType() int {
	return int(int32(C.QmlContacts_QmlContacts_QRegisterMetaType()))
}

func (ptr *QmlContacts) QRegisterMetaType() int {
	return int(int32(C.QmlContacts_QmlContacts_QRegisterMetaType()))
}

func QmlContacts_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QmlContacts_QmlContacts_QRegisterMetaType2(typeNameC)))
}

func (ptr *QmlContacts) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QmlContacts_QmlContacts_QRegisterMetaType2(typeNameC)))
}

func QmlContacts_QmlRegisterType() int {
	return int(int32(C.QmlContacts_QmlContacts_QmlRegisterType()))
}

func (ptr *QmlContacts) QmlRegisterType() int {
	return int(int32(C.QmlContacts_QmlContacts_QmlRegisterType()))
}

func QmlContacts_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QmlContacts_QmlContacts_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QmlContacts) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QmlContacts_QmlContacts_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QmlContacts) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.QmlContacts___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QmlContacts) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QmlContacts___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *QmlContacts) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QmlContacts___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QmlContacts) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlContacts___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlContacts) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlContacts___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlContacts) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QmlContacts___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QmlContacts) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlContacts___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlContacts) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlContacts___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlContacts) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QmlContacts___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QmlContacts) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlContacts___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlContacts) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlContacts___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlContacts) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QmlContacts___findChildren_newList(ptr.Pointer()))
}

func (ptr *QmlContacts) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlContacts___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlContacts) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlContacts___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlContacts) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QmlContacts___children_newList(ptr.Pointer()))
}

func NewQmlContacts(parent std_core.QObject_ITF) *QmlContacts {
	var tmpValue = NewQmlContactsFromPointer(C.QmlContacts_NewQmlContacts(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQmlContacts_DestroyQmlContacts
func callbackQmlContacts_DestroyQmlContacts(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QmlContacts"); signal != nil {
		signal.(func())()
	} else {
		NewQmlContactsFromPointer(ptr).DestroyQmlContactsDefault()
	}
}

func (ptr *QmlContacts) ConnectDestroyQmlContacts(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QmlContacts"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QmlContacts", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QmlContacts", f)
		}
	}
}

func (ptr *QmlContacts) DisconnectDestroyQmlContacts() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QmlContacts")
	}
}

func (ptr *QmlContacts) DestroyQmlContacts() {
	if ptr.Pointer() != nil {
		C.QmlContacts_DestroyQmlContacts(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QmlContacts) DestroyQmlContactsDefault() {
	if ptr.Pointer() != nil {
		C.QmlContacts_DestroyQmlContactsDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQmlContacts_Event
func callbackQmlContacts_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlContactsFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *QmlContacts) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QmlContacts_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQmlContacts_EventFilter
func callbackQmlContacts_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlContactsFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *QmlContacts) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QmlContacts_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQmlContacts_ChildEvent
func callbackQmlContacts_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewQmlContactsFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QmlContacts) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlContacts_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackQmlContacts_ConnectNotify
func callbackQmlContacts_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlContactsFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlContacts) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QmlContacts_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlContacts_CustomEvent
func callbackQmlContacts_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewQmlContactsFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *QmlContacts) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlContacts_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackQmlContacts_DeleteLater
func callbackQmlContacts_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQmlContactsFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QmlContacts) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QmlContacts_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQmlContacts_Destroyed
func callbackQmlContacts_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackQmlContacts_DisconnectNotify
func callbackQmlContacts_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlContactsFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlContacts) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QmlContacts_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlContacts_ObjectNameChanged
func callbackQmlContacts_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQmlContacts_TimerEvent
func callbackQmlContacts_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewQmlContactsFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QmlContacts) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlContacts_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type QmlLogin_ITF interface {
	std_core.QObject_ITF
	QmlLogin_PTR() *QmlLogin
}

func (ptr *QmlLogin) QmlLogin_PTR() *QmlLogin {
	return ptr
}

func (ptr *QmlLogin) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlLogin) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQmlLogin(ptr QmlLogin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlLogin_PTR().Pointer()
	}
	return nil
}

func NewQmlLoginFromPointer(ptr unsafe.Pointer) *QmlLogin {
	var n *QmlLogin
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(QmlLogin)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *QmlLogin:
			n = deduced

		case *std_core.QObject:
			n = &QmlLogin{QObject: *deduced}

		default:
			n = new(QmlLogin)
			n.SetPointer(ptr)
		}
	}
	return n
}

//export callbackQmlLogin_Constructor
func callbackQmlLogin_Constructor(ptr unsafe.Pointer) {
	gPtr := NewQmlLoginFromPointer(ptr)
	qt.Register(ptr, gPtr)
}

//export callbackQmlLogin_CheckLoginData
func callbackQmlLogin_CheckLoginData(ptr unsafe.Pointer, userName C.struct_Moc_PackedString, password C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "checkLoginData"); signal != nil {
		signal.(func(string, string))(cGoUnpackString(userName), cGoUnpackString(password))
	}

}

func (ptr *QmlLogin) ConnectCheckLoginData(f func(userName string, password string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "checkLoginData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "checkLoginData", func(userName string, password string) {
				signal.(func(string, string))(userName, password)
				f(userName, password)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "checkLoginData", f)
		}
	}
}

func (ptr *QmlLogin) DisconnectCheckLoginData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "checkLoginData")
	}
}

func (ptr *QmlLogin) CheckLoginData(userName string, password string) {
	if ptr.Pointer() != nil {
		var userNameC *C.char
		if userName != "" {
			userNameC = C.CString(userName)
			defer C.free(unsafe.Pointer(userNameC))
		}
		var passwordC *C.char
		if password != "" {
			passwordC = C.CString(password)
			defer C.free(unsafe.Pointer(passwordC))
		}
		C.QmlLogin_CheckLoginData(ptr.Pointer(), C.struct_Moc_PackedString{data: userNameC, len: C.longlong(len(userName))}, C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

//export callbackQmlLogin_LoginDataIsValid
func callbackQmlLogin_LoginDataIsValid(ptr unsafe.Pointer, isLoginValid C.char) {
	if signal := qt.GetSignal(ptr, "loginDataIsValid"); signal != nil {
		signal.(func(bool))(int8(isLoginValid) != 0)
	}

}

func (ptr *QmlLogin) ConnectLoginDataIsValid(f func(isLoginValid bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "loginDataIsValid") {
			C.QmlLogin_ConnectLoginDataIsValid(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "loginDataIsValid"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "loginDataIsValid", func(isLoginValid bool) {
				signal.(func(bool))(isLoginValid)
				f(isLoginValid)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loginDataIsValid", f)
		}
	}
}

func (ptr *QmlLogin) DisconnectLoginDataIsValid() {
	if ptr.Pointer() != nil {
		C.QmlLogin_DisconnectLoginDataIsValid(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "loginDataIsValid")
	}
}

func (ptr *QmlLogin) LoginDataIsValid(isLoginValid bool) {
	if ptr.Pointer() != nil {
		C.QmlLogin_LoginDataIsValid(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isLoginValid))))
	}
}

func QmlLogin_QRegisterMetaType() int {
	return int(int32(C.QmlLogin_QmlLogin_QRegisterMetaType()))
}

func (ptr *QmlLogin) QRegisterMetaType() int {
	return int(int32(C.QmlLogin_QmlLogin_QRegisterMetaType()))
}

func QmlLogin_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QmlLogin_QmlLogin_QRegisterMetaType2(typeNameC)))
}

func (ptr *QmlLogin) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QmlLogin_QmlLogin_QRegisterMetaType2(typeNameC)))
}

func QmlLogin_QmlRegisterType() int {
	return int(int32(C.QmlLogin_QmlLogin_QmlRegisterType()))
}

func (ptr *QmlLogin) QmlRegisterType() int {
	return int(int32(C.QmlLogin_QmlLogin_QmlRegisterType()))
}

func QmlLogin_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QmlLogin_QmlLogin_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QmlLogin) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QmlLogin_QmlLogin_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QmlLogin) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.QmlLogin___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QmlLogin) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QmlLogin___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *QmlLogin) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QmlLogin___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QmlLogin) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlLogin___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlLogin) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlLogin___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlLogin) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QmlLogin___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QmlLogin) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlLogin___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlLogin) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlLogin___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlLogin) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QmlLogin___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QmlLogin) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlLogin___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlLogin) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlLogin___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlLogin) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QmlLogin___findChildren_newList(ptr.Pointer()))
}

func (ptr *QmlLogin) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlLogin___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlLogin) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlLogin___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlLogin) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QmlLogin___children_newList(ptr.Pointer()))
}

func NewQmlLogin(parent std_core.QObject_ITF) *QmlLogin {
	var tmpValue = NewQmlLoginFromPointer(C.QmlLogin_NewQmlLogin(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQmlLogin_DestroyQmlLogin
func callbackQmlLogin_DestroyQmlLogin(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QmlLogin"); signal != nil {
		signal.(func())()
	} else {
		NewQmlLoginFromPointer(ptr).DestroyQmlLoginDefault()
	}
}

func (ptr *QmlLogin) ConnectDestroyQmlLogin(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QmlLogin"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QmlLogin", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QmlLogin", f)
		}
	}
}

func (ptr *QmlLogin) DisconnectDestroyQmlLogin() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QmlLogin")
	}
}

func (ptr *QmlLogin) DestroyQmlLogin() {
	if ptr.Pointer() != nil {
		C.QmlLogin_DestroyQmlLogin(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QmlLogin) DestroyQmlLoginDefault() {
	if ptr.Pointer() != nil {
		C.QmlLogin_DestroyQmlLoginDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQmlLogin_Event
func callbackQmlLogin_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlLoginFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *QmlLogin) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QmlLogin_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQmlLogin_EventFilter
func callbackQmlLogin_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlLoginFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *QmlLogin) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QmlLogin_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQmlLogin_ChildEvent
func callbackQmlLogin_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewQmlLoginFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QmlLogin) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlLogin_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackQmlLogin_ConnectNotify
func callbackQmlLogin_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlLoginFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlLogin) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QmlLogin_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlLogin_CustomEvent
func callbackQmlLogin_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewQmlLoginFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *QmlLogin) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlLogin_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackQmlLogin_DeleteLater
func callbackQmlLogin_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQmlLoginFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QmlLogin) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QmlLogin_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQmlLogin_Destroyed
func callbackQmlLogin_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackQmlLogin_DisconnectNotify
func callbackQmlLogin_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlLoginFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlLogin) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QmlLogin_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlLogin_ObjectNameChanged
func callbackQmlLogin_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQmlLogin_TimerEvent
func callbackQmlLogin_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewQmlLoginFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QmlLogin) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlLogin_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type QmlMessage_ITF interface {
	std_core.QObject_ITF
	QmlMessage_PTR() *QmlMessage
}

func (ptr *QmlMessage) QmlMessage_PTR() *QmlMessage {
	return ptr
}

func (ptr *QmlMessage) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlMessage) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQmlMessage(ptr QmlMessage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlMessage_PTR().Pointer()
	}
	return nil
}

func NewQmlMessageFromPointer(ptr unsafe.Pointer) *QmlMessage {
	var n *QmlMessage
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(QmlMessage)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *QmlMessage:
			n = deduced

		case *std_core.QObject:
			n = &QmlMessage{QObject: *deduced}

		default:
			n = new(QmlMessage)
			n.SetPointer(ptr)
		}
	}
	return n
}

//export callbackQmlMessage_Constructor
func callbackQmlMessage_Constructor(ptr unsafe.Pointer) {
	gPtr := NewQmlMessageFromPointer(ptr)
	qt.Register(ptr, gPtr)
}

//export callbackQmlMessage_SendMessage
func callbackQmlMessage_SendMessage(ptr unsafe.Pointer, message C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "sendMessage"); signal != nil {
		signal.(func(string))(cGoUnpackString(message))
	}

}

func (ptr *QmlMessage) ConnectSendMessage(f func(message string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "sendMessage"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "sendMessage", func(message string) {
				signal.(func(string))(message)
				f(message)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sendMessage", f)
		}
	}
}

func (ptr *QmlMessage) DisconnectSendMessage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "sendMessage")
	}
}

func (ptr *QmlMessage) SendMessage(message string) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		C.QmlMessage_SendMessage(ptr.Pointer(), C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})
	}
}

func QmlMessage_QRegisterMetaType() int {
	return int(int32(C.QmlMessage_QmlMessage_QRegisterMetaType()))
}

func (ptr *QmlMessage) QRegisterMetaType() int {
	return int(int32(C.QmlMessage_QmlMessage_QRegisterMetaType()))
}

func QmlMessage_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QmlMessage_QmlMessage_QRegisterMetaType2(typeNameC)))
}

func (ptr *QmlMessage) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QmlMessage_QmlMessage_QRegisterMetaType2(typeNameC)))
}

func QmlMessage_QmlRegisterType() int {
	return int(int32(C.QmlMessage_QmlMessage_QmlRegisterType()))
}

func (ptr *QmlMessage) QmlRegisterType() int {
	return int(int32(C.QmlMessage_QmlMessage_QmlRegisterType()))
}

func QmlMessage_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QmlMessage_QmlMessage_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QmlMessage) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QmlMessage_QmlMessage_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QmlMessage) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.QmlMessage___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QmlMessage) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QmlMessage___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *QmlMessage) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QmlMessage___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QmlMessage) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlMessage___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlMessage) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlMessage___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlMessage) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QmlMessage___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QmlMessage) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlMessage___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlMessage) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlMessage___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlMessage) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QmlMessage___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QmlMessage) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlMessage___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlMessage) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlMessage___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlMessage) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QmlMessage___findChildren_newList(ptr.Pointer()))
}

func (ptr *QmlMessage) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlMessage___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlMessage) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlMessage___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlMessage) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QmlMessage___children_newList(ptr.Pointer()))
}

func NewQmlMessage(parent std_core.QObject_ITF) *QmlMessage {
	var tmpValue = NewQmlMessageFromPointer(C.QmlMessage_NewQmlMessage(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQmlMessage_DestroyQmlMessage
func callbackQmlMessage_DestroyQmlMessage(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QmlMessage"); signal != nil {
		signal.(func())()
	} else {
		NewQmlMessageFromPointer(ptr).DestroyQmlMessageDefault()
	}
}

func (ptr *QmlMessage) ConnectDestroyQmlMessage(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QmlMessage"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QmlMessage", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QmlMessage", f)
		}
	}
}

func (ptr *QmlMessage) DisconnectDestroyQmlMessage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QmlMessage")
	}
}

func (ptr *QmlMessage) DestroyQmlMessage() {
	if ptr.Pointer() != nil {
		C.QmlMessage_DestroyQmlMessage(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QmlMessage) DestroyQmlMessageDefault() {
	if ptr.Pointer() != nil {
		C.QmlMessage_DestroyQmlMessageDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQmlMessage_Event
func callbackQmlMessage_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlMessageFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *QmlMessage) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QmlMessage_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQmlMessage_EventFilter
func callbackQmlMessage_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlMessageFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *QmlMessage) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QmlMessage_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQmlMessage_ChildEvent
func callbackQmlMessage_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewQmlMessageFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QmlMessage) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlMessage_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackQmlMessage_ConnectNotify
func callbackQmlMessage_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlMessageFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlMessage) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QmlMessage_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlMessage_CustomEvent
func callbackQmlMessage_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewQmlMessageFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *QmlMessage) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlMessage_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackQmlMessage_DeleteLater
func callbackQmlMessage_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQmlMessageFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QmlMessage) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QmlMessage_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQmlMessage_Destroyed
func callbackQmlMessage_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackQmlMessage_DisconnectNotify
func callbackQmlMessage_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlMessageFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlMessage) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QmlMessage_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlMessage_ObjectNameChanged
func callbackQmlMessage_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQmlMessage_TimerEvent
func callbackQmlMessage_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewQmlMessageFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QmlMessage) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlMessage_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type QmlRegister_ITF interface {
	std_core.QObject_ITF
	QmlRegister_PTR() *QmlRegister
}

func (ptr *QmlRegister) QmlRegister_PTR() *QmlRegister {
	return ptr
}

func (ptr *QmlRegister) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlRegister) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQmlRegister(ptr QmlRegister_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlRegister_PTR().Pointer()
	}
	return nil
}

func NewQmlRegisterFromPointer(ptr unsafe.Pointer) *QmlRegister {
	var n *QmlRegister
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(QmlRegister)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *QmlRegister:
			n = deduced

		case *std_core.QObject:
			n = &QmlRegister{QObject: *deduced}

		default:
			n = new(QmlRegister)
			n.SetPointer(ptr)
		}
	}
	return n
}

//export callbackQmlRegister_Constructor
func callbackQmlRegister_Constructor(ptr unsafe.Pointer) {
	gPtr := NewQmlRegisterFromPointer(ptr)
	qt.Register(ptr, gPtr)
}

//export callbackQmlRegister_CheckRegisterData
func callbackQmlRegister_CheckRegisterData(ptr unsafe.Pointer, userName C.struct_Moc_PackedString, nickName C.struct_Moc_PackedString, password C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "checkRegisterData"); signal != nil {
		signal.(func(string, string, string))(cGoUnpackString(userName), cGoUnpackString(nickName), cGoUnpackString(password))
	}

}

func (ptr *QmlRegister) ConnectCheckRegisterData(f func(userName string, nickName string, password string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "checkRegisterData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "checkRegisterData", func(userName string, nickName string, password string) {
				signal.(func(string, string, string))(userName, nickName, password)
				f(userName, nickName, password)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "checkRegisterData", f)
		}
	}
}

func (ptr *QmlRegister) DisconnectCheckRegisterData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "checkRegisterData")
	}
}

func (ptr *QmlRegister) CheckRegisterData(userName string, nickName string, password string) {
	if ptr.Pointer() != nil {
		var userNameC *C.char
		if userName != "" {
			userNameC = C.CString(userName)
			defer C.free(unsafe.Pointer(userNameC))
		}
		var nickNameC *C.char
		if nickName != "" {
			nickNameC = C.CString(nickName)
			defer C.free(unsafe.Pointer(nickNameC))
		}
		var passwordC *C.char
		if password != "" {
			passwordC = C.CString(password)
			defer C.free(unsafe.Pointer(passwordC))
		}
		C.QmlRegister_CheckRegisterData(ptr.Pointer(), C.struct_Moc_PackedString{data: userNameC, len: C.longlong(len(userName))}, C.struct_Moc_PackedString{data: nickNameC, len: C.longlong(len(nickName))}, C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

//export callbackQmlRegister_RegisterDataIsValid
func callbackQmlRegister_RegisterDataIsValid(ptr unsafe.Pointer, isRegisterValid C.char) {
	if signal := qt.GetSignal(ptr, "registerDataIsValid"); signal != nil {
		signal.(func(bool))(int8(isRegisterValid) != 0)
	}

}

func (ptr *QmlRegister) ConnectRegisterDataIsValid(f func(isRegisterValid bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "registerDataIsValid") {
			C.QmlRegister_ConnectRegisterDataIsValid(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "registerDataIsValid"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "registerDataIsValid", func(isRegisterValid bool) {
				signal.(func(bool))(isRegisterValid)
				f(isRegisterValid)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "registerDataIsValid", f)
		}
	}
}

func (ptr *QmlRegister) DisconnectRegisterDataIsValid() {
	if ptr.Pointer() != nil {
		C.QmlRegister_DisconnectRegisterDataIsValid(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "registerDataIsValid")
	}
}

func (ptr *QmlRegister) RegisterDataIsValid(isRegisterValid bool) {
	if ptr.Pointer() != nil {
		C.QmlRegister_RegisterDataIsValid(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isRegisterValid))))
	}
}

func QmlRegister_QRegisterMetaType() int {
	return int(int32(C.QmlRegister_QmlRegister_QRegisterMetaType()))
}

func (ptr *QmlRegister) QRegisterMetaType() int {
	return int(int32(C.QmlRegister_QmlRegister_QRegisterMetaType()))
}

func QmlRegister_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QmlRegister_QmlRegister_QRegisterMetaType2(typeNameC)))
}

func (ptr *QmlRegister) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QmlRegister_QmlRegister_QRegisterMetaType2(typeNameC)))
}

func QmlRegister_QmlRegisterType() int {
	return int(int32(C.QmlRegister_QmlRegister_QmlRegisterType()))
}

func (ptr *QmlRegister) QmlRegisterType() int {
	return int(int32(C.QmlRegister_QmlRegister_QmlRegisterType()))
}

func QmlRegister_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QmlRegister_QmlRegister_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QmlRegister) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QmlRegister_QmlRegister_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QmlRegister) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.QmlRegister___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QmlRegister) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QmlRegister___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *QmlRegister) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QmlRegister___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QmlRegister) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlRegister___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlRegister) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlRegister___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlRegister) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QmlRegister___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QmlRegister) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlRegister___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlRegister) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlRegister___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlRegister) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QmlRegister___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QmlRegister) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlRegister___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlRegister) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlRegister___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlRegister) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QmlRegister___findChildren_newList(ptr.Pointer()))
}

func (ptr *QmlRegister) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlRegister___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlRegister) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlRegister___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlRegister) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QmlRegister___children_newList(ptr.Pointer()))
}

func NewQmlRegister(parent std_core.QObject_ITF) *QmlRegister {
	var tmpValue = NewQmlRegisterFromPointer(C.QmlRegister_NewQmlRegister(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQmlRegister_DestroyQmlRegister
func callbackQmlRegister_DestroyQmlRegister(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QmlRegister"); signal != nil {
		signal.(func())()
	} else {
		NewQmlRegisterFromPointer(ptr).DestroyQmlRegisterDefault()
	}
}

func (ptr *QmlRegister) ConnectDestroyQmlRegister(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QmlRegister"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QmlRegister", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QmlRegister", f)
		}
	}
}

func (ptr *QmlRegister) DisconnectDestroyQmlRegister() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QmlRegister")
	}
}

func (ptr *QmlRegister) DestroyQmlRegister() {
	if ptr.Pointer() != nil {
		C.QmlRegister_DestroyQmlRegister(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QmlRegister) DestroyQmlRegisterDefault() {
	if ptr.Pointer() != nil {
		C.QmlRegister_DestroyQmlRegisterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQmlRegister_Event
func callbackQmlRegister_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlRegisterFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *QmlRegister) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QmlRegister_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQmlRegister_EventFilter
func callbackQmlRegister_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlRegisterFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *QmlRegister) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QmlRegister_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQmlRegister_ChildEvent
func callbackQmlRegister_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewQmlRegisterFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QmlRegister) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlRegister_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackQmlRegister_ConnectNotify
func callbackQmlRegister_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlRegisterFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlRegister) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QmlRegister_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlRegister_CustomEvent
func callbackQmlRegister_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewQmlRegisterFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *QmlRegister) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlRegister_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackQmlRegister_DeleteLater
func callbackQmlRegister_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQmlRegisterFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QmlRegister) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QmlRegister_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQmlRegister_Destroyed
func callbackQmlRegister_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackQmlRegister_DisconnectNotify
func callbackQmlRegister_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlRegisterFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlRegister) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QmlRegister_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlRegister_ObjectNameChanged
func callbackQmlRegister_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQmlRegister_TimerEvent
func callbackQmlRegister_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewQmlRegisterFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QmlRegister) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlRegister_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
