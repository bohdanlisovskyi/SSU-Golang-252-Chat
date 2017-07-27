

#pragma once

#ifndef GO_MOC_H
#define GO_MOC_H

#include <stdint.h>

#ifdef __cplusplus
class QmlContacts;
void QmlContacts_QmlContacts_QRegisterMetaTypes();
class QmlLogin;
void QmlLogin_QmlLogin_QRegisterMetaTypes();
class QmlMessage;
void QmlMessage_QmlMessage_QRegisterMetaTypes();
class QmlRegister;
void QmlRegister_QmlRegister_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; };
struct Moc_PackedList { void* data; long long len; };
void QmlContacts_ChangeCurrentContact(void* ptr, int newIndex);
void QmlContacts_SearchUser(void* ptr, struct Moc_PackedString searchedUser);
void QmlContacts_ConnectSendLastMessage(void* ptr);
void QmlContacts_DisconnectSendLastMessage(void* ptr);
void QmlContacts_SendLastMessage(void* ptr, struct Moc_PackedString lastMessage);
void QmlContacts_ConnectSendHistory(void* ptr);
void QmlContacts_DisconnectSendHistory(void* ptr);
void QmlContacts_SendHistory(void* ptr, struct Moc_PackedString history);
int QmlContacts_QmlContacts_QRegisterMetaType();
int QmlContacts_QmlContacts_QRegisterMetaType2(char* typeName);
int QmlContacts_QmlContacts_QmlRegisterType();
int QmlContacts_QmlContacts_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* QmlContacts___dynamicPropertyNames_atList(void* ptr, int i);
void QmlContacts___dynamicPropertyNames_setList(void* ptr, void* i);
void* QmlContacts___dynamicPropertyNames_newList(void* ptr);
void* QmlContacts___findChildren_atList2(void* ptr, int i);
void QmlContacts___findChildren_setList2(void* ptr, void* i);
void* QmlContacts___findChildren_newList2(void* ptr);
void* QmlContacts___findChildren_atList3(void* ptr, int i);
void QmlContacts___findChildren_setList3(void* ptr, void* i);
void* QmlContacts___findChildren_newList3(void* ptr);
void* QmlContacts___findChildren_atList(void* ptr, int i);
void QmlContacts___findChildren_setList(void* ptr, void* i);
void* QmlContacts___findChildren_newList(void* ptr);
void* QmlContacts___children_atList(void* ptr, int i);
void QmlContacts___children_setList(void* ptr, void* i);
void* QmlContacts___children_newList(void* ptr);
void* QmlContacts_NewQmlContacts(void* parent);
void QmlContacts_DestroyQmlContacts(void* ptr);
void QmlContacts_DestroyQmlContactsDefault(void* ptr);
char QmlContacts_EventDefault(void* ptr, void* e);
char QmlContacts_EventFilterDefault(void* ptr, void* watched, void* event);
void QmlContacts_ChildEventDefault(void* ptr, void* event);
void QmlContacts_ConnectNotifyDefault(void* ptr, void* sign);
void QmlContacts_CustomEventDefault(void* ptr, void* event);
void QmlContacts_DeleteLaterDefault(void* ptr);
void QmlContacts_DisconnectNotifyDefault(void* ptr, void* sign);
void QmlContacts_TimerEventDefault(void* ptr, void* event);
;
void QmlLogin_CheckLoginData(void* ptr, struct Moc_PackedString userName, struct Moc_PackedString password);
void QmlLogin_ConnectLoginDataIsValid(void* ptr);
void QmlLogin_DisconnectLoginDataIsValid(void* ptr);
void QmlLogin_LoginDataIsValid(void* ptr, char isLoginValid);
int QmlLogin_QmlLogin_QRegisterMetaType();
int QmlLogin_QmlLogin_QRegisterMetaType2(char* typeName);
int QmlLogin_QmlLogin_QmlRegisterType();
int QmlLogin_QmlLogin_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* QmlLogin___dynamicPropertyNames_atList(void* ptr, int i);
void QmlLogin___dynamicPropertyNames_setList(void* ptr, void* i);
void* QmlLogin___dynamicPropertyNames_newList(void* ptr);
void* QmlLogin___findChildren_atList2(void* ptr, int i);
void QmlLogin___findChildren_setList2(void* ptr, void* i);
void* QmlLogin___findChildren_newList2(void* ptr);
void* QmlLogin___findChildren_atList3(void* ptr, int i);
void QmlLogin___findChildren_setList3(void* ptr, void* i);
void* QmlLogin___findChildren_newList3(void* ptr);
void* QmlLogin___findChildren_atList(void* ptr, int i);
void QmlLogin___findChildren_setList(void* ptr, void* i);
void* QmlLogin___findChildren_newList(void* ptr);
void* QmlLogin___children_atList(void* ptr, int i);
void QmlLogin___children_setList(void* ptr, void* i);
void* QmlLogin___children_newList(void* ptr);
void* QmlLogin_NewQmlLogin(void* parent);
void QmlLogin_DestroyQmlLogin(void* ptr);
void QmlLogin_DestroyQmlLoginDefault(void* ptr);
char QmlLogin_EventDefault(void* ptr, void* e);
char QmlLogin_EventFilterDefault(void* ptr, void* watched, void* event);
void QmlLogin_ChildEventDefault(void* ptr, void* event);
void QmlLogin_ConnectNotifyDefault(void* ptr, void* sign);
void QmlLogin_CustomEventDefault(void* ptr, void* event);
void QmlLogin_DeleteLaterDefault(void* ptr);
void QmlLogin_DisconnectNotifyDefault(void* ptr, void* sign);
void QmlLogin_TimerEventDefault(void* ptr, void* event);
;
void QmlMessage_SendMessage(void* ptr, struct Moc_PackedString message);
int QmlMessage_QmlMessage_QRegisterMetaType();
int QmlMessage_QmlMessage_QRegisterMetaType2(char* typeName);
int QmlMessage_QmlMessage_QmlRegisterType();
int QmlMessage_QmlMessage_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* QmlMessage___dynamicPropertyNames_atList(void* ptr, int i);
void QmlMessage___dynamicPropertyNames_setList(void* ptr, void* i);
void* QmlMessage___dynamicPropertyNames_newList(void* ptr);
void* QmlMessage___findChildren_atList2(void* ptr, int i);
void QmlMessage___findChildren_setList2(void* ptr, void* i);
void* QmlMessage___findChildren_newList2(void* ptr);
void* QmlMessage___findChildren_atList3(void* ptr, int i);
void QmlMessage___findChildren_setList3(void* ptr, void* i);
void* QmlMessage___findChildren_newList3(void* ptr);
void* QmlMessage___findChildren_atList(void* ptr, int i);
void QmlMessage___findChildren_setList(void* ptr, void* i);
void* QmlMessage___findChildren_newList(void* ptr);
void* QmlMessage___children_atList(void* ptr, int i);
void QmlMessage___children_setList(void* ptr, void* i);
void* QmlMessage___children_newList(void* ptr);
void* QmlMessage_NewQmlMessage(void* parent);
void QmlMessage_DestroyQmlMessage(void* ptr);
void QmlMessage_DestroyQmlMessageDefault(void* ptr);
char QmlMessage_EventDefault(void* ptr, void* e);
char QmlMessage_EventFilterDefault(void* ptr, void* watched, void* event);
void QmlMessage_ChildEventDefault(void* ptr, void* event);
void QmlMessage_ConnectNotifyDefault(void* ptr, void* sign);
void QmlMessage_CustomEventDefault(void* ptr, void* event);
void QmlMessage_DeleteLaterDefault(void* ptr);
void QmlMessage_DisconnectNotifyDefault(void* ptr, void* sign);
void QmlMessage_TimerEventDefault(void* ptr, void* event);
;
void QmlRegister_CheckRegisterData(void* ptr, struct Moc_PackedString userName, struct Moc_PackedString nickName, struct Moc_PackedString password);
void QmlRegister_ConnectRegisterDataIsValid(void* ptr);
void QmlRegister_DisconnectRegisterDataIsValid(void* ptr);
void QmlRegister_RegisterDataIsValid(void* ptr, char isRegisterValid);
int QmlRegister_QmlRegister_QRegisterMetaType();
int QmlRegister_QmlRegister_QRegisterMetaType2(char* typeName);
int QmlRegister_QmlRegister_QmlRegisterType();
int QmlRegister_QmlRegister_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* QmlRegister___dynamicPropertyNames_atList(void* ptr, int i);
void QmlRegister___dynamicPropertyNames_setList(void* ptr, void* i);
void* QmlRegister___dynamicPropertyNames_newList(void* ptr);
void* QmlRegister___findChildren_atList2(void* ptr, int i);
void QmlRegister___findChildren_setList2(void* ptr, void* i);
void* QmlRegister___findChildren_newList2(void* ptr);
void* QmlRegister___findChildren_atList3(void* ptr, int i);
void QmlRegister___findChildren_setList3(void* ptr, void* i);
void* QmlRegister___findChildren_newList3(void* ptr);
void* QmlRegister___findChildren_atList(void* ptr, int i);
void QmlRegister___findChildren_setList(void* ptr, void* i);
void* QmlRegister___findChildren_newList(void* ptr);
void* QmlRegister___children_atList(void* ptr, int i);
void QmlRegister___children_setList(void* ptr, void* i);
void* QmlRegister___children_newList(void* ptr);
void* QmlRegister_NewQmlRegister(void* parent);
void QmlRegister_DestroyQmlRegister(void* ptr);
void QmlRegister_DestroyQmlRegisterDefault(void* ptr);
char QmlRegister_EventDefault(void* ptr, void* e);
char QmlRegister_EventFilterDefault(void* ptr, void* watched, void* event);
void QmlRegister_ChildEventDefault(void* ptr, void* event);
void QmlRegister_ConnectNotifyDefault(void* ptr, void* sign);
void QmlRegister_CustomEventDefault(void* ptr, void* event);
void QmlRegister_DeleteLaterDefault(void* ptr);
void QmlRegister_DisconnectNotifyDefault(void* ptr, void* sign);
void QmlRegister_TimerEventDefault(void* ptr, void* event);
;

#ifdef __cplusplus
}
#endif

#endif