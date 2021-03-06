package qtwidgets

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h
// #include <qstyleditemdelegate.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
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

// void initStyleOption(class QStyleOptionViewItem *, const class QModelIndex &)
func (this *QStyledItemDelegate) InheritInitStyleOption(f func(option *QStyleOptionViewItem /*777 QStyleOptionViewItem **/, index *qtcore.QModelIndex) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

// bool eventFilter(class QObject *, class QEvent *)
func (this *QStyledItemDelegate) InheritEventFilter(f func(object *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// bool editorEvent(class QEvent *, class QAbstractItemModel *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QStyledItemDelegate) InheritEditorEvent(f func(event *qtcore.QEvent /*777 QEvent **/, model *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/, option *QStyleOptionViewItem, index *qtcore.QModelIndex) bool) {
	qtrt.SetAllInheritCallback(this, "editorEvent", f)
}

type QStyledItemDelegate struct {
	*QAbstractItemDelegate
}
type QStyledItemDelegate_ITF interface {
	QAbstractItemDelegate_ITF
	QStyledItemDelegate_PTR() *QStyledItemDelegate
}

func (ptr *QStyledItemDelegate) QStyledItemDelegate_PTR() *QStyledItemDelegate { return ptr }

func (this *QStyledItemDelegate) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractItemDelegate.GetCthis()
	}
}
func (this *QStyledItemDelegate) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractItemDelegate = NewQAbstractItemDelegateFromPointer(cthis)
}
func NewQStyledItemDelegateFromPointer(cthis unsafe.Pointer) *QStyledItemDelegate {
	bcthis0 := NewQAbstractItemDelegateFromPointer(cthis)
	return &QStyledItemDelegate{bcthis0}
}
func (*QStyledItemDelegate) NewFromPointer(cthis unsafe.Pointer) *QStyledItemDelegate {
	return NewQStyledItemDelegateFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QStyledItemDelegate) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyledItemDelegate(QObject *)
func NewQStyledItemDelegate(parent qtcore.QObject_ITF /*777 QObject **/) *QStyledItemDelegate {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyledItemDelegateC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyledItemDelegateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QStyledItemDelegate")
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyledItemDelegate(QObject *)
func NewQStyledItemDelegate__() *QStyledItemDelegate {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyledItemDelegateC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyledItemDelegateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QStyledItemDelegate")
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QStyledItemDelegate()
func DeleteQStyledItemDelegate(this *QStyledItemDelegate) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyledItemDelegateD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionViewItem &, const QModelIndex &) const
func (this *QStyledItemDelegate) Paint(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg2 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint(const QStyleOptionViewItem &, const QModelIndex &) const
func (this *QStyledItemDelegate) SizeHint(option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) *qtcore.QSize /*123*/ {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg0 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg1 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QWidget * createEditor(QWidget *, const QStyleOptionViewItem &, const QModelIndex &) const
func (this *QStyledItemDelegate) CreateEditor(parent QWidget_ITF /*777 QWidget **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) *QWidget /*777 QWidget **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg2 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setEditorData(QWidget *, const QModelIndex &) const
func (this *QStyledItemDelegate) SetEditorData(editor QWidget_ITF /*777 QWidget **/, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QWidget_PTR() != nil {
		convArg0 = editor.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg1 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setModelData(QWidget *, QAbstractItemModel *, const QModelIndex &) const
func (this *QStyledItemDelegate) SetModelData(editor QWidget_ITF /*777 QWidget **/, model qtcore.QAbstractItemModel_ITF /*777 QAbstractItemModel **/, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QWidget_PTR() != nil {
		convArg0 = editor.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if model != nil && model.QAbstractItemModel_PTR() != nil {
		convArg1 = model.QAbstractItemModel_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg2 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void updateEditorGeometry(QWidget *, const QStyleOptionViewItem &, const QModelIndex &) const
func (this *QStyledItemDelegate) UpdateEditorGeometry(editor QWidget_ITF /*777 QWidget **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QWidget_PTR() != nil {
		convArg0 = editor.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg2 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:85
// index:0
// Public Visibility=Default Availability=Available
// [8] QItemEditorFactory * itemEditorFactory() const
func (this *QStyledItemDelegate) ItemEditorFactory() *QItemEditorFactory /*777 QItemEditorFactory **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate17itemEditorFactoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQItemEditorFactoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemEditorFactory(QItemEditorFactory *)
func (this *QStyledItemDelegate) SetItemEditorFactory(factory QItemEditorFactory_ITF /*777 QItemEditorFactory **/) {
	var convArg0 unsafe.Pointer
	if factory != nil && factory.QItemEditorFactory_PTR() != nil {
		convArg0 = factory.QItemEditorFactory_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyledItemDelegate20setItemEditorFactoryEP18QItemEditorFactory", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString displayText(const QVariant &, const QLocale &) const
func (this *QStyledItemDelegate) DisplayText(value qtcore.QVariant_ITF, locale qtcore.QLocale_ITF) string {
	var convArg0 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg0 = value.QVariant_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if locale != nil && locale.QLocale_PTR() != nil {
		convArg1 = locale.QLocale_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate11displayTextERK8QVariantRK7QLocale", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:91
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionViewItem *, const QModelIndex &) const
func (this *QStyledItemDelegate) InitStyleOption(option QStyleOptionViewItem_ITF /*777 QStyleOptionViewItem **/, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg0 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg1 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate15initStyleOptionEP20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:94
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)
func (this *QStyledItemDelegate) EventFilter(object qtcore.QObject_ITF /*777 QObject **/, event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg1 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyledItemDelegate11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:95
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool editorEvent(QEvent *, QAbstractItemModel *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QStyledItemDelegate) EditorEvent(event qtcore.QEvent_ITF /*777 QEvent **/, model qtcore.QAbstractItemModel_ITF /*777 QAbstractItemModel **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if model != nil && model.QAbstractItemModel_PTR() != nil {
		convArg1 = model.QAbstractItemModel_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg2 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg3 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyledItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
