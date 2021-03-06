package qtwidgets

// /usr/include/qt/QtWidgets/qundogroup.h
// #include <qundogroup.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 66
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

type QUndoGroup struct {
	*qtcore.QObject
}
type QUndoGroup_ITF interface {
	qtcore.QObject_ITF
	QUndoGroup_PTR() *QUndoGroup
}

func (ptr *QUndoGroup) QUndoGroup_PTR() *QUndoGroup { return ptr }

func (this *QUndoGroup) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QUndoGroup) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQUndoGroupFromPointer(cthis unsafe.Pointer) *QUndoGroup {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QUndoGroup{bcthis0}
}
func (*QUndoGroup) NewFromPointer(cthis unsafe.Pointer) *QUndoGroup {
	return NewQUndoGroupFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qundogroup.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QUndoGroup) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoGroup10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qundogroup.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QUndoGroup(QObject *)
func NewQUndoGroup(parent qtcore.QObject_ITF /*777 QObject **/) *QUndoGroup {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroupC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUndoGroupFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QUndoGroup")
	return gothis
}

// /usr/include/qt/QtWidgets/qundogroup.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QUndoGroup(QObject *)
func NewQUndoGroup__() *QUndoGroup {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroupC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUndoGroupFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QUndoGroup")
	return gothis
}

// /usr/include/qt/QtWidgets/qundogroup.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QUndoGroup()
func DeleteQUndoGroup(this *QUndoGroup) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroupD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qundogroup.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addStack(QUndoStack *)
func (this *QUndoGroup) AddStack(stack QUndoStack_ITF /*777 QUndoStack **/) {
	var convArg0 unsafe.Pointer
	if stack != nil && stack.QUndoStack_PTR() != nil {
		convArg0 = stack.QUndoStack_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup8addStackEP10QUndoStack", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeStack(QUndoStack *)
func (this *QUndoGroup) RemoveStack(stack QUndoStack_ITF /*777 QUndoStack **/) {
	var convArg0 unsafe.Pointer
	if stack != nil && stack.QUndoStack_PTR() != nil {
		convArg0 = stack.QUndoStack_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup11removeStackEP10QUndoStack", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QUndoStack * activeStack() const
func (this *QUndoGroup) ActiveStack() *QUndoStack /*777 QUndoStack **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoGroup11activeStackEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQUndoStackFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qundogroup.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * createUndoAction(QObject *, const QString &) const
func (this *QUndoGroup) CreateUndoAction(parent qtcore.QObject_ITF /*777 QObject **/, prefix string) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(prefix)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoGroup16createUndoActionEP7QObjectRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qundogroup.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * createUndoAction(QObject *, const QString &) const
func (this *QUndoGroup) CreateUndoAction__(parent qtcore.QObject_ITF /*777 QObject **/) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	// arg: 1, const QString &=LValueReference, QString=Record,
	var convArg1 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoGroup16createUndoActionEP7QObjectRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qundogroup.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * createRedoAction(QObject *, const QString &) const
func (this *QUndoGroup) CreateRedoAction(parent qtcore.QObject_ITF /*777 QObject **/, prefix string) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(prefix)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoGroup16createRedoActionEP7QObjectRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qundogroup.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * createRedoAction(QObject *, const QString &) const
func (this *QUndoGroup) CreateRedoAction__(parent qtcore.QObject_ITF /*777 QObject **/) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	// arg: 1, const QString &=LValueReference, QString=Record,
	var convArg1 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoGroup16createRedoActionEP7QObjectRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qundogroup.h:75
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canUndo() const
func (this *QUndoGroup) CanUndo() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoGroup7canUndoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qundogroup.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canRedo() const
func (this *QUndoGroup) CanRedo() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoGroup7canRedoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qundogroup.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QString undoText() const
func (this *QUndoGroup) UndoText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoGroup8undoTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qundogroup.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QString redoText() const
func (this *QUndoGroup) RedoText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoGroup8redoTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qundogroup.h:79
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isClean() const
func (this *QUndoGroup) IsClean() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoGroup7isCleanEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qundogroup.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void undo()
func (this *QUndoGroup) Undo() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup4undoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redo()
func (this *QUndoGroup) Redo() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup4redoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setActiveStack(QUndoStack *)
func (this *QUndoGroup) SetActiveStack(stack QUndoStack_ITF /*777 QUndoStack **/) {
	var convArg0 unsafe.Pointer
	if stack != nil && stack.QUndoStack_PTR() != nil {
		convArg0 = stack.QUndoStack_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup14setActiveStackEP10QUndoStack", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activeStackChanged(QUndoStack *)
func (this *QUndoGroup) ActiveStackChanged(stack QUndoStack_ITF /*777 QUndoStack **/) {
	var convArg0 unsafe.Pointer
	if stack != nil && stack.QUndoStack_PTR() != nil {
		convArg0 = stack.QUndoStack_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup18activeStackChangedEP10QUndoStack", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void indexChanged(int)
func (this *QUndoGroup) IndexChanged(idx int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup12indexChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), idx)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cleanChanged(_Bool)
func (this *QUndoGroup) CleanChanged(clean bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup12cleanChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), clean)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void canUndoChanged(_Bool)
func (this *QUndoGroup) CanUndoChanged(canUndo bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup14canUndoChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), canUndo)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void canRedoChanged(_Bool)
func (this *QUndoGroup) CanRedoChanged(canRedo bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup14canRedoChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), canRedo)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void undoTextChanged(const QString &)
func (this *QUndoGroup) UndoTextChanged(undoText string) {
	var tmpArg0 = qtcore.NewQString_5(undoText)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup15undoTextChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redoTextChanged(const QString &)
func (this *QUndoGroup) RedoTextChanged(redoText string) {
	var tmpArg0 = qtcore.NewQString_5(redoText)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup15redoTextChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
