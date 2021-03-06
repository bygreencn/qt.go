package qtcore

// /usr/include/qt/QtCore/qfuturewatcher.h
// #include <qfuturewatcher.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 45
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// void connectNotify(const class QMetaMethod &)
func (this *QFutureWatcherBase) InheritConnectNotify(f func(signal *QMetaMethod) /*void*/) {
	qtrt.SetAllInheritCallback(this, "connectNotify", f)
}

// void disconnectNotify(const class QMetaMethod &)
func (this *QFutureWatcherBase) InheritDisconnectNotify(f func(signal *QMetaMethod) /*void*/) {
	qtrt.SetAllInheritCallback(this, "disconnectNotify", f)
}

// void connectOutputInterface()
func (this *QFutureWatcherBase) InheritConnectOutputInterface(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "connectOutputInterface", f)
}

// void disconnectOutputInterface(_Bool)
func (this *QFutureWatcherBase) InheritDisconnectOutputInterface(f func(pendingAssignment bool) /*void*/) {
	qtrt.SetAllInheritCallback(this, "disconnectOutputInterface", f)
}

type QFutureWatcherBase struct {
	*QObject
}
type QFutureWatcherBase_ITF interface {
	QObject_ITF
	QFutureWatcherBase_PTR() *QFutureWatcherBase
}

func (ptr *QFutureWatcherBase) QFutureWatcherBase_PTR() *QFutureWatcherBase { return ptr }

func (this *QFutureWatcherBase) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QFutureWatcherBase) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQFutureWatcherBaseFromPointer(cthis unsafe.Pointer) *QFutureWatcherBase {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QFutureWatcherBase{bcthis0}
}
func (*QFutureWatcherBase) NewFromPointer(cthis unsafe.Pointer) *QFutureWatcherBase {
	return NewQFutureWatcherBaseFromPointer(cthis)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QFutureWatcherBase) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QFutureWatcherBase10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qfuturewatcher.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFutureWatcherBase(QObject *)
func NewQFutureWatcherBase(parent QObject_ITF /*777 QObject **/) *QFutureWatcherBase {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBaseC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFutureWatcherBaseFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFutureWatcherBase")
	return gothis
}

// /usr/include/qt/QtCore/qfuturewatcher.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFutureWatcherBase(QObject *)
func NewQFutureWatcherBase__() *QFutureWatcherBase {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBaseC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFutureWatcherBaseFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFutureWatcherBase")
	return gothis
}

// /usr/include/qt/QtCore/qfuturewatcher.h:64
// index:0
// Public Visibility=Default Availability=Available
// [4] int progressValue() const
func (this *QFutureWatcherBase) ProgressValue() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QFutureWatcherBase13progressValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qfuturewatcher.h:65
// index:0
// Public Visibility=Default Availability=Available
// [4] int progressMinimum() const
func (this *QFutureWatcherBase) ProgressMinimum() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QFutureWatcherBase15progressMinimumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qfuturewatcher.h:66
// index:0
// Public Visibility=Default Availability=Available
// [4] int progressMaximum() const
func (this *QFutureWatcherBase) ProgressMaximum() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QFutureWatcherBase15progressMaximumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qfuturewatcher.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QString progressText() const
func (this *QFutureWatcherBase) ProgressText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QFutureWatcherBase12progressTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfuturewatcher.h:69
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isStarted() const
func (this *QFutureWatcherBase) IsStarted() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QFutureWatcherBase9isStartedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfuturewatcher.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFinished() const
func (this *QFutureWatcherBase) IsFinished() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QFutureWatcherBase10isFinishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfuturewatcher.h:71
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRunning() const
func (this *QFutureWatcherBase) IsRunning() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QFutureWatcherBase9isRunningEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfuturewatcher.h:72
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCanceled() const
func (this *QFutureWatcherBase) IsCanceled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QFutureWatcherBase10isCanceledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfuturewatcher.h:73
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPaused() const
func (this *QFutureWatcherBase) IsPaused() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QFutureWatcherBase8isPausedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfuturewatcher.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void waitForFinished()
func (this *QFutureWatcherBase) WaitForFinished() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBase15waitForFinishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPendingResultsLimit(int)
func (this *QFutureWatcherBase) SetPendingResultsLimit(limit int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBase22setPendingResultsLimitEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), limit)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QFutureWatcherBase) Event(event QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBase5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfuturewatcher.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void started()
func (this *QFutureWatcherBase) Started() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBase7startedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void finished()
func (this *QFutureWatcherBase) Finished() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBase8finishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void canceled()
func (this *QFutureWatcherBase) Canceled() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBase8canceledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void paused()
func (this *QFutureWatcherBase) Paused() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBase6pausedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resumed()
func (this *QFutureWatcherBase) Resumed() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBase7resumedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resultReadyAt(int)
func (this *QFutureWatcherBase) ResultReadyAt(resultIndex int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBase13resultReadyAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), resultIndex)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resultsReadyAt(int, int)
func (this *QFutureWatcherBase) ResultsReadyAt(beginIndex int, endIndex int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBase14resultsReadyAtEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), beginIndex, endIndex)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void progressRangeChanged(int, int)
func (this *QFutureWatcherBase) ProgressRangeChanged(minimum int, maximum int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBase20progressRangeChangedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), minimum, maximum)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void progressValueChanged(int)
func (this *QFutureWatcherBase) ProgressValueChanged(progressValue int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBase20progressValueChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), progressValue)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void progressTextChanged(const QString &)
func (this *QFutureWatcherBase) ProgressTextChanged(progressText string) {
	var tmpArg0 = NewQString_5(progressText)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBase19progressTextChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cancel()
func (this *QFutureWatcherBase) Cancel() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBase6cancelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPaused(_Bool)
func (this *QFutureWatcherBase) SetPaused(paused bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBase9setPausedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), paused)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void pause()
func (this *QFutureWatcherBase) Pause() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBase5pauseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resume()
func (this *QFutureWatcherBase) Resume() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBase6resumeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void togglePaused()
func (this *QFutureWatcherBase) TogglePaused() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBase12togglePausedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:101
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void connectNotify(const QMetaMethod &)
func (this *QFutureWatcherBase) ConnectNotify(signal QMetaMethod_ITF) {
	var convArg0 unsafe.Pointer
	if signal != nil && signal.QMetaMethod_PTR() != nil {
		convArg0 = signal.QMetaMethod_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBase13connectNotifyERK11QMetaMethod", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:102
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void disconnectNotify(const QMetaMethod &)
func (this *QFutureWatcherBase) DisconnectNotify(signal QMetaMethod_ITF) {
	var convArg0 unsafe.Pointer
	if signal != nil && signal.QMetaMethod_PTR() != nil {
		convArg0 = signal.QMetaMethod_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBase16disconnectNotifyERK11QMetaMethod", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:105
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void connectOutputInterface()
func (this *QFutureWatcherBase) ConnectOutputInterface() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBase22connectOutputInterfaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:106
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void disconnectOutputInterface(_Bool)
func (this *QFutureWatcherBase) DisconnectOutputInterface(pendingAssignment bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBase25disconnectOutputInterfaceEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pendingAssignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:106
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void disconnectOutputInterface(_Bool)
func (this *QFutureWatcherBase) DisconnectOutputInterface__() {
	// arg: 0, bool=Bool, =Invalid,
	pendingAssignment := false
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBase25disconnectOutputInterfaceEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pendingAssignment)
	qtrt.ErrPrint(err, rv)
}

func DeleteQFutureWatcherBase(this *QFutureWatcherBase) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFutureWatcherBaseD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
