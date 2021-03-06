package qtcore

// /usr/include/qt/QtCore/qcoreevent.h
// #include <qcoreevent.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QDynamicPropertyChangeEvent struct {
	*QEvent
}
type QDynamicPropertyChangeEvent_ITF interface {
	QEvent_ITF
	QDynamicPropertyChangeEvent_PTR() *QDynamicPropertyChangeEvent
}

func (ptr *QDynamicPropertyChangeEvent) QDynamicPropertyChangeEvent_PTR() *QDynamicPropertyChangeEvent {
	return ptr
}

func (this *QDynamicPropertyChangeEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QDynamicPropertyChangeEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = NewQEventFromPointer(cthis)
}
func NewQDynamicPropertyChangeEventFromPointer(cthis unsafe.Pointer) *QDynamicPropertyChangeEvent {
	bcthis0 := NewQEventFromPointer(cthis)
	return &QDynamicPropertyChangeEvent{bcthis0}
}
func (*QDynamicPropertyChangeEvent) NewFromPointer(cthis unsafe.Pointer) *QDynamicPropertyChangeEvent {
	return NewQDynamicPropertyChangeEventFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcoreevent.h:365
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDynamicPropertyChangeEvent(const QByteArray &)
func NewQDynamicPropertyChangeEvent(name QByteArray_ITF) *QDynamicPropertyChangeEvent {
	var convArg0 unsafe.Pointer
	if name != nil && name.QByteArray_PTR() != nil {
		convArg0 = name.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QDynamicPropertyChangeEventC2ERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDynamicPropertyChangeEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDynamicPropertyChangeEvent)
	return gothis
}

// /usr/include/qt/QtCore/qcoreevent.h:366
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDynamicPropertyChangeEvent()
func DeleteQDynamicPropertyChangeEvent(this *QDynamicPropertyChangeEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QDynamicPropertyChangeEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qcoreevent.h:368
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray propertyName() const
func (this *QDynamicPropertyChangeEvent) PropertyName() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QDynamicPropertyChangeEvent12propertyNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

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
