package qtcore

// /usr/include/qt/QtCore/qabstractanimation.h
// #include <qabstractanimation.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 28
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// void advanceAnimation(qint64)
func (this *QAnimationDriver) InheritAdvanceAnimation(f func(timeStep int64) /*void*/) {
	qtrt.SetAllInheritCallback(this, "advanceAnimation", f)
}

// void start()
func (this *QAnimationDriver) InheritStart(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "start", f)
}

// void stop()
func (this *QAnimationDriver) InheritStop(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "stop", f)
}

type QAnimationDriver struct {
	*QObject
}
type QAnimationDriver_ITF interface {
	QObject_ITF
	QAnimationDriver_PTR() *QAnimationDriver
}

func (ptr *QAnimationDriver) QAnimationDriver_PTR() *QAnimationDriver { return ptr }

func (this *QAnimationDriver) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAnimationDriver) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQAnimationDriverFromPointer(cthis unsafe.Pointer) *QAnimationDriver {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QAnimationDriver{bcthis0}
}
func (*QAnimationDriver) NewFromPointer(cthis unsafe.Pointer) *QAnimationDriver {
	return NewQAnimationDriverFromPointer(cthis)
}

// /usr/include/qt/QtCore/qabstractanimation.h:135
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QAnimationDriver) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAnimationDriver10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstractanimation.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAnimationDriver(QObject *)
func NewQAnimationDriver(parent QObject_ITF /*777 QObject **/) *QAnimationDriver {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAnimationDriverC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAnimationDriverFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAnimationDriver")
	return gothis
}

// /usr/include/qt/QtCore/qabstractanimation.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAnimationDriver(QObject *)
func NewQAnimationDriver__() *QAnimationDriver {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAnimationDriverC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAnimationDriverFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAnimationDriver")
	return gothis
}

// /usr/include/qt/QtCore/qabstractanimation.h:140
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAnimationDriver()
func DeleteQAnimationDriver(this *QAnimationDriver) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAnimationDriverD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qabstractanimation.h:142
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void advance()
func (this *QAnimationDriver) Advance() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAnimationDriver7advanceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void install()
func (this *QAnimationDriver) Install() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAnimationDriver7installEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void uninstall()
func (this *QAnimationDriver) Uninstall() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAnimationDriver9uninstallEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:147
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRunning() const
func (this *QAnimationDriver) IsRunning() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAnimationDriver9isRunningEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractanimation.h:149
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 elapsed() const
func (this *QAnimationDriver) Elapsed() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAnimationDriver7elapsedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qabstractanimation.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStartTime(qint64)
func (this *QAnimationDriver) SetStartTime(startTime int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAnimationDriver12setStartTimeEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), startTime)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:153
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 startTime() const
func (this *QAnimationDriver) StartTime() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAnimationDriver9startTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qabstractanimation.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void started()
func (this *QAnimationDriver) Started() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAnimationDriver7startedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stopped()
func (this *QAnimationDriver) Stopped() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAnimationDriver7stoppedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:161
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void advanceAnimation(qint64)
func (this *QAnimationDriver) AdvanceAnimation(timeStep int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAnimationDriver16advanceAnimationEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeStep)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:161
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void advanceAnimation(qint64)
func (this *QAnimationDriver) AdvanceAnimation__() {
	// arg: 0, qint64=Typedef, qint64=Typedef, long long
	timeStep := int64(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAnimationDriver16advanceAnimationEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeStep)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:162
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void start()
func (this *QAnimationDriver) Start() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAnimationDriver5startEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:163
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void stop()
func (this *QAnimationDriver) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAnimationDriver4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
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
