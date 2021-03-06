package qtgui

// /usr/include/qt/QtGui/qinputmethod.h
// #include <qinputmethod.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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

type QInputMethod struct {
	*qtcore.QObject
}
type QInputMethod_ITF interface {
	qtcore.QObject_ITF
	QInputMethod_PTR() *QInputMethod
}

func (ptr *QInputMethod) QInputMethod_PTR() *QInputMethod { return ptr }

func (this *QInputMethod) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QInputMethod) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQInputMethodFromPointer(cthis unsafe.Pointer) *QInputMethod {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QInputMethod{bcthis0}
}
func (*QInputMethod) NewFromPointer(cthis unsafe.Pointer) *QInputMethod {
	return NewQInputMethodFromPointer(cthis)
}

// /usr/include/qt/QtGui/qinputmethod.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QInputMethod) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputMethod10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qinputmethod.h:68
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform inputItemTransform() const
func (this *QInputMethod) InputItemTransform() *QTransform /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputMethod18inputItemTransformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qinputmethod.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInputItemTransform(const QTransform &)
func (this *QInputMethod) SetInputItemTransform(transform QTransform_ITF) {
	var convArg0 unsafe.Pointer
	if transform != nil && transform.QTransform_PTR() != nil {
		convArg0 = transform.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputMethod21setInputItemTransformERK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:71
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF inputItemRectangle() const
func (this *QInputMethod) InputItemRectangle() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputMethod18inputItemRectangleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qinputmethod.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInputItemRectangle(const QRectF &)
func (this *QInputMethod) SetInputItemRectangle(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputMethod21setInputItemRectangleERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:75
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF cursorRectangle() const
func (this *QInputMethod) CursorRectangle() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputMethod15cursorRectangleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qinputmethod.h:76
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF anchorRectangle() const
func (this *QInputMethod) AnchorRectangle() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputMethod15anchorRectangleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qinputmethod.h:79
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF keyboardRectangle() const
func (this *QInputMethod) KeyboardRectangle() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputMethod17keyboardRectangleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qinputmethod.h:81
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF inputItemClipRectangle() const
func (this *QInputMethod) InputItemClipRectangle() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputMethod22inputItemClipRectangleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qinputmethod.h:89
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisible() const
func (this *QInputMethod) IsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputMethod9isVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qinputmethod.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVisible(_Bool)
func (this *QInputMethod) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputMethod10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:92
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAnimating() const
func (this *QInputMethod) IsAnimating() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputMethod11isAnimatingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qinputmethod.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] QLocale locale() const
func (this *QInputMethod) Locale() *qtcore.QLocale /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputMethod6localeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQLocale)
	return rv2
}

// /usr/include/qt/QtGui/qinputmethod.h:95
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::LayoutDirection inputDirection() const
func (this *QInputMethod) InputDirection() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputMethod14inputDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:97
// index:0
// Public static Visibility=Default Availability=Available
// [16] QVariant queryFocusObject(Qt::InputMethodQuery, QVariant)
func (this *QInputMethod) QueryFocusObject(query int, argument qtcore.QVariant_ITF /*123*/) *qtcore.QVariant /*123*/ {
	var convArg1 unsafe.Pointer
	if argument != nil && argument.QVariant_PTR() != nil {
		convArg1 = argument.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputMethod16queryFocusObjectEN2Qt16InputMethodQueryE8QVariant", qtrt.FFI_TYPE_POINTER, query, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}
func QInputMethod_QueryFocusObject(query int, argument qtcore.QVariant_ITF /*123*/) *qtcore.QVariant /*123*/ {
	var nilthis *QInputMethod
	rv := nilthis.QueryFocusObject(query, argument)
	return rv
}

// /usr/include/qt/QtGui/qinputmethod.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void show()
func (this *QInputMethod) Show() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputMethod4showEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hide()
func (this *QInputMethod) Hide() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputMethod4hideEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void update(Qt::InputMethodQueries)
func (this *QInputMethod) Update(queries int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputMethod6updateE6QFlagsIN2Qt16InputMethodQueryEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), queries)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reset()
func (this *QInputMethod) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputMethod5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void commit()
func (this *QInputMethod) Commit() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputMethod6commitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void invokeAction(enum QInputMethod::Action, int)
func (this *QInputMethod) InvokeAction(a int, cursorPosition int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputMethod12invokeActionENS_6ActionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a, cursorPosition)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cursorRectangleChanged()
func (this *QInputMethod) CursorRectangleChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputMethod22cursorRectangleChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void anchorRectangleChanged()
func (this *QInputMethod) AnchorRectangleChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputMethod22anchorRectangleChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void keyboardRectangleChanged()
func (this *QInputMethod) KeyboardRectangleChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputMethod24keyboardRectangleChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void inputItemClipRectangleChanged()
func (this *QInputMethod) InputItemClipRectangleChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputMethod29inputItemClipRectangleChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void visibleChanged()
func (this *QInputMethod) VisibleChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputMethod14visibleChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void animatingChanged()
func (this *QInputMethod) AnimatingChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputMethod16animatingChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void localeChanged()
func (this *QInputMethod) LocaleChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputMethod13localeChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void inputDirectionChanged(Qt::LayoutDirection)
func (this *QInputMethod) InputDirectionChanged(newDirection int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputMethod21inputDirectionChangedEN2Qt15LayoutDirectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newDirection)
	qtrt.ErrPrint(err, rv)
}

func DeleteQInputMethod(this *QInputMethod) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputMethodD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QInputMethod__Action = int

const QInputMethod__Click QInputMethod__Action = 0
const QInputMethod__ContextMenu QInputMethod__Action = 1

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
