package qtwidgets

// /usr/include/qt/QtWidgets/qmenubar.h
// #include <qmenubar.h>
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

// void changeEvent(class QEvent *)
func (this *QMenuBar) InheritChangeEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QMenuBar) InheritKeyPressEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QMenuBar) InheritMouseReleaseEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QMenuBar) InheritMousePressEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QMenuBar) InheritMouseMoveEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void leaveEvent(class QEvent *)
func (this *QMenuBar) InheritLeaveEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "leaveEvent", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QMenuBar) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QMenuBar) InheritResizeEvent(f func(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void actionEvent(class QActionEvent *)
func (this *QMenuBar) InheritActionEvent(f func(arg0 *qtgui.QActionEvent /*777 QActionEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "actionEvent", f)
}

// void focusOutEvent(class QFocusEvent *)
func (this *QMenuBar) InheritFocusOutEvent(f func(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void focusInEvent(class QFocusEvent *)
func (this *QMenuBar) InheritFocusInEvent(f func(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void timerEvent(class QTimerEvent *)
func (this *QMenuBar) InheritTimerEvent(f func(arg0 *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// bool eventFilter(class QObject *, class QEvent *)
func (this *QMenuBar) InheritEventFilter(f func(arg0 *qtcore.QObject /*777 QObject **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// bool event(class QEvent *)
func (this *QMenuBar) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void initStyleOption(class QStyleOptionMenuItem *, const class QAction *)
func (this *QMenuBar) InheritInitStyleOption(f func(option *QStyleOptionMenuItem /*777 QStyleOptionMenuItem **/, action *QAction /*777 const QAction **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

type QMenuBar struct {
	*QWidget
}
type QMenuBar_ITF interface {
	QWidget_ITF
	QMenuBar_PTR() *QMenuBar
}

func (ptr *QMenuBar) QMenuBar_PTR() *QMenuBar { return ptr }

func (this *QMenuBar) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QMenuBar) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQMenuBarFromPointer(cthis unsafe.Pointer) *QMenuBar {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QMenuBar{bcthis0}
}
func (*QMenuBar) NewFromPointer(cthis unsafe.Pointer) *QMenuBar {
	return NewQMenuBarFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qmenubar.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QMenuBar) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMenuBar10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenubar.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMenuBar(QWidget *)
func NewQMenuBar(parent QWidget_ITF /*777 QWidget **/) *QMenuBar {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBarC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMenuBarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMenuBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qmenubar.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMenuBar(QWidget *)
func NewQMenuBar__() *QMenuBar {
	// arg: 0, QWidget *=Pointer, QWidget=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBarC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMenuBarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMenuBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qmenubar.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMenuBar()
func DeleteQMenuBar(this *QMenuBar) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBarD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qmenubar.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QString &)
func (this *QMenuBar) AddAction(text string) *QAction /*777 QAction **/ {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar9addActionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenubar.h:68
// index:1
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QString &, const QObject *, const char *)
func (this *QMenuBar) AddAction_1(text string, receiver qtcore.QObject_ITF /*777 const QObject **/, member string) *QAction /*777 QAction **/ {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg1 = receiver.QObject_PTR().GetCthis()
	}
	var convArg2 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar9addActionERK7QStringPK7QObjectPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenubar.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * addMenu(QMenu *)
func (this *QMenuBar) AddMenu(menu QMenu_ITF /*777 QMenu **/) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if menu != nil && menu.QMenu_PTR() != nil {
		convArg0 = menu.QMenu_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar7addMenuEP5QMenu", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenubar.h:71
// index:1
// Public Visibility=Default Availability=Available
// [8] QMenu * addMenu(const QString &)
func (this *QMenuBar) AddMenu_1(title string) *QMenu /*777 QMenu **/ {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar7addMenuERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenubar.h:72
// index:2
// Public Visibility=Default Availability=Available
// [8] QMenu * addMenu(const QIcon &, const QString &)
func (this *QMenuBar) AddMenu_2(icon qtgui.QIcon_ITF, title string) *QMenu /*777 QMenu **/ {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar7addMenuERK5QIconRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenubar.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * addSeparator()
func (this *QMenuBar) AddSeparator() *QAction /*777 QAction **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar12addSeparatorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenubar.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * insertSeparator(QAction *)
func (this *QMenuBar) InsertSeparator(before QAction_ITF /*777 QAction **/) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if before != nil && before.QAction_PTR() != nil {
		convArg0 = before.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar15insertSeparatorEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenubar.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * insertMenu(QAction *, QMenu *)
func (this *QMenuBar) InsertMenu(before QAction_ITF /*777 QAction **/, menu QMenu_ITF /*777 QMenu **/) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if before != nil && before.QAction_PTR() != nil {
		convArg0 = before.QAction_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if menu != nil && menu.QMenu_PTR() != nil {
		convArg1 = menu.QMenu_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar10insertMenuEP7QActionP5QMenu", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenubar.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QMenuBar) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * activeAction() const
func (this *QMenuBar) ActiveAction() *QAction /*777 QAction **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMenuBar12activeActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenubar.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setActiveAction(QAction *)
func (this *QMenuBar) SetActiveAction(action QAction_ITF /*777 QAction **/) {
	var convArg0 unsafe.Pointer
	if action != nil && action.QAction_PTR() != nil {
		convArg0 = action.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar15setActiveActionEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultUp(_Bool)
func (this *QMenuBar) SetDefaultUp(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar12setDefaultUpEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:86
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDefaultUp() const
func (this *QMenuBar) IsDefaultUp() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMenuBar11isDefaultUpEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmenubar.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const
func (this *QMenuBar) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMenuBar8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qmenubar.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint() const
func (this *QMenuBar) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMenuBar15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qmenubar.h:90
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int heightForWidth(int) const
func (this *QMenuBar) HeightForWidth(arg0 int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMenuBar14heightForWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qmenubar.h:92
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect actionGeometry(QAction *) const
func (this *QMenuBar) ActionGeometry(arg0 QAction_ITF /*777 QAction **/) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QAction_PTR() != nil {
		convArg0 = arg0.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMenuBar14actionGeometryEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qmenubar.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * actionAt(const QPoint &) const
func (this *QMenuBar) ActionAt(arg0 qtcore.QPoint_ITF) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPoint_PTR() != nil {
		convArg0 = arg0.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMenuBar8actionAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenubar.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCornerWidget(QWidget *, Qt::Corner)
func (this *QMenuBar) SetCornerWidget(w QWidget_ITF /*777 QWidget **/, corner int) {
	var convArg0 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar15setCornerWidgetEP7QWidgetN2Qt6CornerE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, corner)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCornerWidget(QWidget *, Qt::Corner)
func (this *QMenuBar) SetCornerWidget__(w QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	// arg: 1, Qt::Corner=Elaborated, Qt::Corner=Enum,
	corner := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar15setCornerWidgetEP7QWidgetN2Qt6CornerE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, corner)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:96
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * cornerWidget(Qt::Corner) const
func (this *QMenuBar) CornerWidget(corner int) *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMenuBar12cornerWidgetEN2Qt6CornerE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), corner)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenubar.h:96
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * cornerWidget(Qt::Corner) const
func (this *QMenuBar) CornerWidget__() *QWidget /*777 QWidget **/ {
	// arg: 0, Qt::Corner=Elaborated, Qt::Corner=Enum,
	corner := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMenuBar12cornerWidgetEN2Qt6CornerE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), corner)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenubar.h:102
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNativeMenuBar() const
func (this *QMenuBar) IsNativeMenuBar() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMenuBar15isNativeMenuBarEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmenubar.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNativeMenuBar(_Bool)
func (this *QMenuBar) SetNativeMenuBar(nativeMenuBar bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar16setNativeMenuBarEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nativeMenuBar)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:106
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setVisible(_Bool)
func (this *QMenuBar) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void triggered(QAction *)
func (this *QMenuBar) Triggered(action QAction_ITF /*777 QAction **/) {
	var convArg0 unsafe.Pointer
	if action != nil && action.QAction_PTR() != nil {
		convArg0 = action.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar9triggeredEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hovered(QAction *)
func (this *QMenuBar) Hovered(action QAction_ITF /*777 QAction **/) {
	var convArg0 unsafe.Pointer
	if action != nil && action.QAction_PTR() != nil {
		convArg0 = action.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar7hoveredEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:113
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QMenuBar) ChangeEvent(arg0 qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:114
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QMenuBar) KeyPressEvent(arg0 qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QKeyEvent_PTR() != nil {
		convArg0 = arg0.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:115
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QMenuBar) MouseReleaseEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:116
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QMenuBar) MousePressEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:117
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QMenuBar) MouseMoveEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:118
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void leaveEvent(QEvent *)
func (this *QMenuBar) LeaveEvent(arg0 qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar10leaveEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:119
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QMenuBar) PaintEvent(arg0 qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPaintEvent_PTR() != nil {
		convArg0 = arg0.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:120
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QMenuBar) ResizeEvent(arg0 qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QResizeEvent_PTR() != nil {
		convArg0 = arg0.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:121
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void actionEvent(QActionEvent *)
func (this *QMenuBar) ActionEvent(arg0 qtgui.QActionEvent_ITF /*777 QActionEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QActionEvent_PTR() != nil {
		convArg0 = arg0.QActionEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar11actionEventEP12QActionEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:122
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)
func (this *QMenuBar) FocusOutEvent(arg0 qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFocusEvent_PTR() != nil {
		convArg0 = arg0.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:123
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)
func (this *QMenuBar) FocusInEvent(arg0 qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFocusEvent_PTR() != nil {
		convArg0 = arg0.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:124
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)
func (this *QMenuBar) TimerEvent(arg0 qtcore.QTimerEvent_ITF /*777 QTimerEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTimerEvent_PTR() != nil {
		convArg0 = arg0.QTimerEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:125
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)
func (this *QMenuBar) EventFilter(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QEvent_PTR() != nil {
		convArg1 = arg1.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmenubar.h:126
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QMenuBar) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMenuBar5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmenubar.h:127
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionMenuItem *, const QAction *) const
func (this *QMenuBar) InitStyleOption(option QStyleOptionMenuItem_ITF /*777 QStyleOptionMenuItem **/, action QAction_ITF /*777 const QAction **/) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionMenuItem_PTR() != nil {
		convArg0 = option.QStyleOptionMenuItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if action != nil && action.QAction_PTR() != nil {
		convArg1 = action.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMenuBar15initStyleOptionEP20QStyleOptionMenuItemPK7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
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
