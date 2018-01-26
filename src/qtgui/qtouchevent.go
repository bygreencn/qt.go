package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QTouchEvent struct {
	*QInputEvent
}

func (this *QTouchEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QInputEvent.GetCthis()
	}
}
func (this *QTouchEvent) SetCthis(cthis unsafe.Pointer) {
	this.QInputEvent = NewQInputEventFromPointer(cthis)
}
func NewQTouchEventFromPointer(cthis unsafe.Pointer) *QTouchEvent {
	bcthis0 := NewQInputEventFromPointer(cthis)
	return &QTouchEvent{bcthis0}
}
func (*QTouchEvent) NewFromPointer(cthis unsafe.Pointer) *QTouchEvent {
	return NewQTouchEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:947
// index:0
// Public virtual
// void ~QTouchEvent()
func DeleteQTouchEvent(*QTouchEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTouchEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:949
// index:0
// Public inline
// QWindow * window()
func (this *QTouchEvent) Window() *QWindow /*777 QWindow **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTouchEvent6windowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:950
// index:0
// Public inline
// QObject * target()
func (this *QTouchEvent) Target() *qtcore.QObject /*777 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTouchEvent6targetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:954
// index:0
// Public inline
// Qt::TouchPointStates touchPointStates()
func (this *QTouchEvent) TouchPointStates() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTouchEvent16touchPointStatesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:955
// index:0
// Public inline
// const QList<QTouchEvent::TouchPoint> & touchPoints()
func (this *QTouchEvent) TouchPoints() unsafe.Pointer /*555*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTouchEvent11touchPointsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qevent.h:956
// index:0
// Public inline
// QTouchDevice * device()
func (this *QTouchEvent) Device() *QTouchDevice /*777 QTouchDevice **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTouchEvent6deviceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTouchDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:959
// index:0
// Public inline
// void setWindow(class QWindow *)
func (this *QTouchEvent) SetWindow(awindow *QWindow /*777 QWindow **/) {
	var convArg0 = awindow.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTouchEvent9setWindowEP7QWindow", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:960
// index:0
// Public inline
// void setTarget(class QObject *)
func (this *QTouchEvent) SetTarget(atarget *qtcore.QObject /*777 QObject **/) {
	var convArg0 = atarget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTouchEvent9setTargetEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:963
// index:0
// Public inline
// void setDevice(class QTouchDevice *)
func (this *QTouchEvent) SetDevice(adevice *QTouchDevice /*777 QTouchDevice **/) {
	var convArg0 = adevice.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTouchEvent9setDeviceEP12QTouchDevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end