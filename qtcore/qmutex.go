package qtcore

// /usr/include/qt/QtCore/qmutex.h
// #include <qmutex.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QMutex struct {
	*QBasicMutex
}
type QMutex_ITF interface {
	QBasicMutex_ITF
	QMutex_PTR() *QMutex
}

func (ptr *QMutex) QMutex_PTR() *QMutex { return ptr }

func (this *QMutex) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QBasicMutex.GetCthis()
	}
}
func (this *QMutex) SetCthis(cthis unsafe.Pointer) {
	this.QBasicMutex = NewQBasicMutexFromPointer(cthis)
}
func NewQMutexFromPointer(cthis unsafe.Pointer) *QMutex {
	bcthis0 := NewQBasicMutexFromPointer(cthis)
	return &QMutex{bcthis0}
}
func (*QMutex) NewFromPointer(cthis unsafe.Pointer) *QMutex {
	return NewQMutexFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmutex.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMutex(enum QMutex::RecursionMode)
func NewQMutex(mode int) *QMutex {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMutexC2ENS_13RecursionModeE", qtrt.FFI_TYPE_POINTER, mode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMutexFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMutex)
	return gothis
}

// /usr/include/qt/QtCore/qmutex.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMutex(enum QMutex::RecursionMode)
func NewQMutex__() *QMutex {
	// arg: 0, QMutex::RecursionMode=Enum, QMutex::RecursionMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMutexC2ENS_13RecursionModeE", qtrt.FFI_TYPE_POINTER, mode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMutexFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMutex)
	return gothis
}

// /usr/include/qt/QtCore/qmutex.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QMutex()
func DeleteQMutex(this *QMutex) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMutexD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qmutex.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void lock()
func (this *QMutex) Lock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMutex4lockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:135
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tryLock(int)
func (this *QMutex) TryLock(timeout int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMutex7tryLockEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmutex.h:135
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tryLock(int)
func (this *QMutex) TryLock__() bool {
	// arg: 0, int=Int, =Invalid,
	timeout := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMutex7tryLockEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmutex.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unlock()
func (this *QMutex) Unlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMutex6unlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:140
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool try_lock()
func (this *QMutex) Try_lock() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMutex8try_lockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmutex.h:161
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isRecursive() const
func (this *QMutex) IsRecursive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMutex11isRecursiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

type QMutex__RecursionMode = int

const QMutex__NonRecursive QMutex__RecursionMode = 0
const QMutex__Recursive QMutex__RecursionMode = 1

//  body block end

//  keep block begin

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
