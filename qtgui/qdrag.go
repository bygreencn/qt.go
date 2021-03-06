package qtgui

// /usr/include/qt/QtGui/qdrag.h
// #include <qdrag.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QDrag struct {
	*qtcore.QObject
}
type QDrag_ITF interface {
	qtcore.QObject_ITF
	QDrag_PTR() *QDrag
}

func (ptr *QDrag) QDrag_PTR() *QDrag { return ptr }

func (this *QDrag) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QDrag) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQDragFromPointer(cthis unsafe.Pointer) *QDrag {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QDrag{bcthis0}
}
func (*QDrag) NewFromPointer(cthis unsafe.Pointer) *QDrag {
	return NewQDragFromPointer(cthis)
}

// /usr/include/qt/QtGui/qdrag.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QDrag) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDrag10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qdrag.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDrag(QObject *)
func NewQDrag(dragSource qtcore.QObject_ITF /*777 QObject **/) *QDrag {
	var convArg0 unsafe.Pointer
	if dragSource != nil && dragSource.QObject_PTR() != nil {
		convArg0 = dragSource.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDragC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDragFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDrag")
	return gothis
}

// /usr/include/qt/QtGui/qdrag.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDrag()
func DeleteQDrag(this *QDrag) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDragD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qdrag.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMimeData(QMimeData *)
func (this *QDrag) SetMimeData(data qtcore.QMimeData_ITF /*777 QMimeData **/) {
	var convArg0 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg0 = data.QMimeData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDrag11setMimeDataEP9QMimeData", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QMimeData * mimeData() const
func (this *QDrag) MimeData() *qtcore.QMimeData /*777 QMimeData **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDrag8mimeDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMimeDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qdrag.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPixmap(const QPixmap &)
func (this *QDrag) SetPixmap(arg0 QPixmap_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPixmap_PTR() != nil {
		convArg0 = arg0.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDrag9setPixmapERK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:69
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap pixmap() const
func (this *QDrag) Pixmap() *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDrag6pixmapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qdrag.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHotSpot(const QPoint &)
func (this *QDrag) SetHotSpot(hotspot qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if hotspot != nil && hotspot.QPoint_PTR() != nil {
		convArg0 = hotspot.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDrag10setHotSpotERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint hotSpot() const
func (this *QDrag) HotSpot() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDrag7hotSpotEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qdrag.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * source() const
func (this *QDrag) Source() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDrag6sourceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qdrag.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * target() const
func (this *QDrag) Target() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDrag6targetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qdrag.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::DropAction start(Qt::DropActions)
func (this *QDrag) Start(supportedActions int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDrag5startE6QFlagsIN2Qt10DropActionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), supportedActions)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qdrag.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::DropAction start(Qt::DropActions)
func (this *QDrag) Start__() int {
	// arg: 0, Qt::DropActions=Elaborated, Qt::DropActions=Typedef, QFlags<Qt::DropAction>
	supportedActions := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDrag5startE6QFlagsIN2Qt10DropActionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), supportedActions)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qdrag.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::DropAction exec(Qt::DropActions)
func (this *QDrag) Exec(supportedActions int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDrag4execE6QFlagsIN2Qt10DropActionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), supportedActions)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qdrag.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::DropAction exec(Qt::DropActions)
func (this *QDrag) Exec__() int {
	// arg: 0, Qt::DropActions=Elaborated, Qt::DropActions=Typedef, QFlags<Qt::DropAction>
	supportedActions := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDrag4execE6QFlagsIN2Qt10DropActionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), supportedActions)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qdrag.h:79
// index:1
// Public Visibility=Default Availability=Available
// [4] Qt::DropAction exec(Qt::DropActions, Qt::DropAction)
func (this *QDrag) Exec_1(supportedActions int, defaultAction int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDrag4execE6QFlagsIN2Qt10DropActionEES2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), supportedActions, defaultAction)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qdrag.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDragCursor(const QPixmap &, Qt::DropAction)
func (this *QDrag) SetDragCursor(cursor QPixmap_ITF, action int) {
	var convArg0 unsafe.Pointer
	if cursor != nil && cursor.QPixmap_PTR() != nil {
		convArg0 = cursor.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDrag13setDragCursorERK7QPixmapN2Qt10DropActionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, action)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:82
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap dragCursor(Qt::DropAction) const
func (this *QDrag) DragCursor(action int) *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDrag10dragCursorEN2Qt10DropActionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qdrag.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::DropActions supportedActions() const
func (this *QDrag) SupportedActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDrag16supportedActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qdrag.h:85
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::DropAction defaultAction() const
func (this *QDrag) DefaultAction() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDrag13defaultActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qdrag.h:87
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void cancel()
func (this *QDrag) Cancel() {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDrag6cancelEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}
func QDrag_Cancel() {
	var nilthis *QDrag
	nilthis.Cancel()
}

// /usr/include/qt/QtGui/qdrag.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void actionChanged(Qt::DropAction)
func (this *QDrag) ActionChanged(action int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDrag13actionChangedEN2Qt10DropActionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void targetChanged(QObject *)
func (this *QDrag) TargetChanged(newTarget qtcore.QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if newTarget != nil && newTarget.QObject_PTR() != nil {
		convArg0 = newTarget.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDrag13targetChangedEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
}

//  keep block end
