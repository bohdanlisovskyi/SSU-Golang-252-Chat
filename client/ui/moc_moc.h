/****************************************************************************
** Meta object code from reading C++ file 'moc.cpp'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.9.1)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'moc.cpp' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.9.1. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
struct qt_meta_stringdata_QmlContacts_t {
    QByteArrayData data[10];
    char stringdata0[115];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_QmlContacts_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_QmlContacts_t qt_meta_stringdata_QmlContacts = {
    {
QT_MOC_LITERAL(0, 0, 11), // "QmlContacts"
QT_MOC_LITERAL(1, 12, 15), // "sendLastMessage"
QT_MOC_LITERAL(2, 28, 0), // ""
QT_MOC_LITERAL(3, 29, 11), // "lastMessage"
QT_MOC_LITERAL(4, 41, 11), // "sendHistory"
QT_MOC_LITERAL(5, 53, 7), // "history"
QT_MOC_LITERAL(6, 61, 20), // "changeCurrentContact"
QT_MOC_LITERAL(7, 82, 8), // "newIndex"
QT_MOC_LITERAL(8, 91, 10), // "searchUser"
QT_MOC_LITERAL(9, 102, 12) // "searchedUser"

    },
    "QmlContacts\0sendLastMessage\0\0lastMessage\0"
    "sendHistory\0history\0changeCurrentContact\0"
    "newIndex\0searchUser\0searchedUser"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_QmlContacts[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       4,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       2,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   34,    2, 0x06 /* Public */,
       4,    1,   37,    2, 0x06 /* Public */,

 // slots: name, argc, parameters, tag, flags
       6,    1,   40,    2, 0x0a /* Public */,
       8,    1,   43,    2, 0x0a /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::QString,    5,

 // slots: parameters
    QMetaType::Void, QMetaType::Int,    7,
    QMetaType::Void, QMetaType::QString,    9,

       0        // eod
};

void QmlContacts::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        QmlContacts *_t = static_cast<QmlContacts *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->sendLastMessage((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 1: _t->sendHistory((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 2: _t->changeCurrentContact((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 3: _t->searchUser((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        void **func = reinterpret_cast<void **>(_a[1]);
        {
            typedef void (QmlContacts::*_t)(QString );
            if (*reinterpret_cast<_t *>(func) == static_cast<_t>(&QmlContacts::sendLastMessage)) {
                *result = 0;
                return;
            }
        }
        {
            typedef void (QmlContacts::*_t)(QString );
            if (*reinterpret_cast<_t *>(func) == static_cast<_t>(&QmlContacts::sendHistory)) {
                *result = 1;
                return;
            }
        }
    }
}

const QMetaObject QmlContacts::staticMetaObject = {
    { &QObject::staticMetaObject, qt_meta_stringdata_QmlContacts.data,
      qt_meta_data_QmlContacts,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *QmlContacts::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *QmlContacts::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_QmlContacts.stringdata0))
        return static_cast<void*>(const_cast< QmlContacts*>(this));
    return QObject::qt_metacast(_clname);
}

int QmlContacts::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 4)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 4;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 4)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 4;
    }
    return _id;
}

// SIGNAL 0
void QmlContacts::sendLastMessage(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void QmlContacts::sendHistory(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}
struct qt_meta_stringdata_QmlLogin_t {
    QByteArrayData data[7];
    char stringdata0[73];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_QmlLogin_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_QmlLogin_t qt_meta_stringdata_QmlLogin = {
    {
QT_MOC_LITERAL(0, 0, 8), // "QmlLogin"
QT_MOC_LITERAL(1, 9, 16), // "loginDataIsValid"
QT_MOC_LITERAL(2, 26, 0), // ""
QT_MOC_LITERAL(3, 27, 12), // "isLoginValid"
QT_MOC_LITERAL(4, 40, 14), // "checkLoginData"
QT_MOC_LITERAL(5, 55, 8), // "userName"
QT_MOC_LITERAL(6, 64, 8) // "password"

    },
    "QmlLogin\0loginDataIsValid\0\0isLoginValid\0"
    "checkLoginData\0userName\0password"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_QmlLogin[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       2,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       1,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   24,    2, 0x06 /* Public */,

 // slots: name, argc, parameters, tag, flags
       4,    2,   27,    2, 0x0a /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::Bool,    3,

 // slots: parameters
    QMetaType::Void, QMetaType::QString, QMetaType::QString,    5,    6,

       0        // eod
};

void QmlLogin::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        QmlLogin *_t = static_cast<QmlLogin *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->loginDataIsValid((*reinterpret_cast< bool(*)>(_a[1]))); break;
        case 1: _t->checkLoginData((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        void **func = reinterpret_cast<void **>(_a[1]);
        {
            typedef void (QmlLogin::*_t)(bool );
            if (*reinterpret_cast<_t *>(func) == static_cast<_t>(&QmlLogin::loginDataIsValid)) {
                *result = 0;
                return;
            }
        }
    }
}

const QMetaObject QmlLogin::staticMetaObject = {
    { &QObject::staticMetaObject, qt_meta_stringdata_QmlLogin.data,
      qt_meta_data_QmlLogin,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *QmlLogin::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *QmlLogin::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_QmlLogin.stringdata0))
        return static_cast<void*>(const_cast< QmlLogin*>(this));
    return QObject::qt_metacast(_clname);
}

int QmlLogin::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 2)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 2;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 2)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 2;
    }
    return _id;
}

// SIGNAL 0
void QmlLogin::loginDataIsValid(bool _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}
struct qt_meta_stringdata_QmlMessage_t {
    QByteArrayData data[4];
    char stringdata0[32];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_QmlMessage_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_QmlMessage_t qt_meta_stringdata_QmlMessage = {
    {
QT_MOC_LITERAL(0, 0, 10), // "QmlMessage"
QT_MOC_LITERAL(1, 11, 11), // "sendMessage"
QT_MOC_LITERAL(2, 23, 0), // ""
QT_MOC_LITERAL(3, 24, 7) // "message"

    },
    "QmlMessage\0sendMessage\0\0message"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_QmlMessage[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       1,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       0,       // signalCount

 // slots: name, argc, parameters, tag, flags
       1,    1,   19,    2, 0x0a /* Public */,

 // slots: parameters
    QMetaType::Void, QMetaType::QString,    3,

       0        // eod
};

void QmlMessage::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        QmlMessage *_t = static_cast<QmlMessage *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->sendMessage((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    }
}

const QMetaObject QmlMessage::staticMetaObject = {
    { &QObject::staticMetaObject, qt_meta_stringdata_QmlMessage.data,
      qt_meta_data_QmlMessage,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *QmlMessage::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *QmlMessage::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_QmlMessage.stringdata0))
        return static_cast<void*>(const_cast< QmlMessage*>(this));
    return QObject::qt_metacast(_clname);
}

int QmlMessage::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 1)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 1;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 1)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 1;
    }
    return _id;
}
struct qt_meta_stringdata_QmlRegister_t {
    QByteArrayData data[8];
    char stringdata0[94];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_QmlRegister_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_QmlRegister_t qt_meta_stringdata_QmlRegister = {
    {
QT_MOC_LITERAL(0, 0, 11), // "QmlRegister"
QT_MOC_LITERAL(1, 12, 19), // "registerDataIsValid"
QT_MOC_LITERAL(2, 32, 0), // ""
QT_MOC_LITERAL(3, 33, 15), // "isRegisterValid"
QT_MOC_LITERAL(4, 49, 17), // "checkRegisterData"
QT_MOC_LITERAL(5, 67, 8), // "userName"
QT_MOC_LITERAL(6, 76, 8), // "nickName"
QT_MOC_LITERAL(7, 85, 8) // "password"

    },
    "QmlRegister\0registerDataIsValid\0\0"
    "isRegisterValid\0checkRegisterData\0"
    "userName\0nickName\0password"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_QmlRegister[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       2,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       1,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   24,    2, 0x06 /* Public */,

 // slots: name, argc, parameters, tag, flags
       4,    3,   27,    2, 0x0a /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::Bool,    3,

 // slots: parameters
    QMetaType::Void, QMetaType::QString, QMetaType::QString, QMetaType::QString,    5,    6,    7,

       0        // eod
};

void QmlRegister::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        QmlRegister *_t = static_cast<QmlRegister *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->registerDataIsValid((*reinterpret_cast< bool(*)>(_a[1]))); break;
        case 1: _t->checkRegisterData((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< QString(*)>(_a[3]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        void **func = reinterpret_cast<void **>(_a[1]);
        {
            typedef void (QmlRegister::*_t)(bool );
            if (*reinterpret_cast<_t *>(func) == static_cast<_t>(&QmlRegister::registerDataIsValid)) {
                *result = 0;
                return;
            }
        }
    }
}

const QMetaObject QmlRegister::staticMetaObject = {
    { &QObject::staticMetaObject, qt_meta_stringdata_QmlRegister.data,
      qt_meta_data_QmlRegister,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *QmlRegister::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *QmlRegister::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_QmlRegister.stringdata0))
        return static_cast<void*>(const_cast< QmlRegister*>(this));
    return QObject::qt_metacast(_clname);
}

int QmlRegister::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 2)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 2;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 2)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 2;
    }
    return _id;
}

// SIGNAL 0
void QmlRegister::registerDataIsValid(bool _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
