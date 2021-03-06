package qtwidgets

// /usr/include/qt/QtWidgets/qsizepolicy.h
// #include <qsizepolicy.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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

type QSizePolicy struct {
	*qtrt.CObject
}
type QSizePolicy_ITF interface {
	QSizePolicy_PTR() *QSizePolicy
}

func (ptr *QSizePolicy) QSizePolicy_PTR() *QSizePolicy { return ptr }

func (this *QSizePolicy) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSizePolicy) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSizePolicyFromPointer(cthis unsafe.Pointer) *QSizePolicy {
	return &QSizePolicy{&qtrt.CObject{cthis}}
}
func (*QSizePolicy) NewFromPointer(cthis unsafe.Pointer) *QSizePolicy {
	return NewQSizePolicyFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:113
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QSizePolicy()
func NewQSizePolicy() *QSizePolicy {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicyC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSizePolicyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSizePolicy)
	return gothis
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:116
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QSizePolicy(enum QSizePolicy::Policy, enum QSizePolicy::Policy, enum QSizePolicy::ControlType)
func NewQSizePolicy_1(horizontal int, vertical int, type_ int) *QSizePolicy {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicyC2ENS_6PolicyES0_NS_11ControlTypeE", qtrt.FFI_TYPE_POINTER, horizontal, vertical, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSizePolicyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSizePolicy)
	return gothis
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:116
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QSizePolicy(enum QSizePolicy::Policy, enum QSizePolicy::Policy, enum QSizePolicy::ControlType)
func NewQSizePolicy_1_(horizontal int, vertical int) *QSizePolicy {
	// arg: 2, QSizePolicy::ControlType=Enum, QSizePolicy::ControlType=Enum,
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicyC2ENS_6PolicyES0_NS_11ControlTypeE", qtrt.FFI_TYPE_POINTER, horizontal, vertical, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSizePolicyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSizePolicy)
	return gothis
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:128
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QSizePolicy::Policy horizontalPolicy() const
func (this *QSizePolicy) HorizontalPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicy16horizontalPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:129
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QSizePolicy::Policy verticalPolicy() const
func (this *QSizePolicy) VerticalPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicy14verticalPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:130
// index:0
// Public Visibility=Default Availability=Available
// [4] QSizePolicy::ControlType controlType() const
func (this *QSizePolicy) ControlType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicy11controlTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:132
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setHorizontalPolicy(enum QSizePolicy::Policy)
func (this *QSizePolicy) SetHorizontalPolicy(d int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicy19setHorizontalPolicyENS_6PolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), d)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:133
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setVerticalPolicy(enum QSizePolicy::Policy)
func (this *QSizePolicy) SetVerticalPolicy(d int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicy17setVerticalPolicyENS_6PolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), d)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setControlType(enum QSizePolicy::ControlType)
func (this *QSizePolicy) SetControlType(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicy14setControlTypeENS_11ControlTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:136
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::Orientations expandingDirections() const
func (this *QSizePolicy) ExpandingDirections() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicy19expandingDirectionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:141
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setHeightForWidth(_Bool)
func (this *QSizePolicy) SetHeightForWidth(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicy17setHeightForWidthEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:142
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool hasHeightForWidth() const
func (this *QSizePolicy) HasHeightForWidth() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicy17hasHeightForWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:143
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWidthForHeight(_Bool)
func (this *QSizePolicy) SetWidthForHeight(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicy17setWidthForHeightEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:144
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool hasWidthForHeight() const
func (this *QSizePolicy) HasWidthForHeight() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicy17hasWidthForHeightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:146
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QSizePolicy &) const
func (this *QSizePolicy) Operator_equal_equal(s QSizePolicy_ITF) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QSizePolicy_PTR() != nil {
		convArg0 = s.QSizePolicy_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicyeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:147
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QSizePolicy &) const
func (this *QSizePolicy) Operator_not_equal(s QSizePolicy_ITF) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QSizePolicy_PTR() != nil {
		convArg0 = s.QSizePolicy_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicyneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:153
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int horizontalStretch() const
func (this *QSizePolicy) HorizontalStretch() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicy17horizontalStretchEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:154
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int verticalStretch() const
func (this *QSizePolicy) VerticalStretch() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicy15verticalStretchEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:155
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setHorizontalStretch(int)
func (this *QSizePolicy) SetHorizontalStretch(stretchFactor int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicy20setHorizontalStretchEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), stretchFactor)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:156
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setVerticalStretch(int)
func (this *QSizePolicy) SetVerticalStretch(stretchFactor int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicy18setVerticalStretchEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), stretchFactor)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:158
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool retainSizeWhenHidden() const
func (this *QSizePolicy) RetainSizeWhenHidden() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicy20retainSizeWhenHiddenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:159
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setRetainSizeWhenHidden(_Bool)
func (this *QSizePolicy) SetRetainSizeWhenHidden(retainSize bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicy23setRetainSizeWhenHiddenEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), retainSize)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:161
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void transpose()
func (this *QSizePolicy) Transpose() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicy9transposeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:166
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QSizePolicy transposed() const
func (this *QSizePolicy) Transposed() *QSizePolicy /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicy10transposedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizePolicyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizePolicy)
	return rv2
}

func DeleteQSizePolicy(this *QSizePolicy) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicyD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QSizePolicy__PolicyFlag = int

const QSizePolicy__GrowFlag QSizePolicy__PolicyFlag = 1
const QSizePolicy__ExpandFlag QSizePolicy__PolicyFlag = 2
const QSizePolicy__ShrinkFlag QSizePolicy__PolicyFlag = 4
const QSizePolicy__IgnoreFlag QSizePolicy__PolicyFlag = 8

type QSizePolicy__Policy = int

const QSizePolicy__Fixed QSizePolicy__Policy = 0
const QSizePolicy__Minimum QSizePolicy__Policy = 1
const QSizePolicy__Maximum QSizePolicy__Policy = 4
const QSizePolicy__Preferred QSizePolicy__Policy = 5
const QSizePolicy__MinimumExpanding QSizePolicy__Policy = 3
const QSizePolicy__Expanding QSizePolicy__Policy = 7
const QSizePolicy__Ignored QSizePolicy__Policy = 13

type QSizePolicy__ControlType = int

const QSizePolicy__DefaultType QSizePolicy__ControlType = 1
const QSizePolicy__ButtonBox QSizePolicy__ControlType = 2
const QSizePolicy__CheckBox QSizePolicy__ControlType = 4
const QSizePolicy__ComboBox QSizePolicy__ControlType = 8
const QSizePolicy__Frame QSizePolicy__ControlType = 16
const QSizePolicy__GroupBox QSizePolicy__ControlType = 32
const QSizePolicy__Label QSizePolicy__ControlType = 64
const QSizePolicy__Line QSizePolicy__ControlType = 128
const QSizePolicy__LineEdit QSizePolicy__ControlType = 256
const QSizePolicy__PushButton QSizePolicy__ControlType = 512
const QSizePolicy__RadioButton QSizePolicy__ControlType = 1024
const QSizePolicy__Slider QSizePolicy__ControlType = 2048
const QSizePolicy__SpinBox QSizePolicy__ControlType = 4096
const QSizePolicy__TabWidget QSizePolicy__ControlType = 8192
const QSizePolicy__ToolButton QSizePolicy__ControlType = 16384

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
