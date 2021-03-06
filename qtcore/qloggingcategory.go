package qtcore

// /usr/include/qt/QtCore/qloggingcategory.h
// #include <qloggingcategory.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QLoggingCategory struct {
	*qtrt.CObject
}
type QLoggingCategory_ITF interface {
	QLoggingCategory_PTR() *QLoggingCategory
}

func (ptr *QLoggingCategory) QLoggingCategory_PTR() *QLoggingCategory { return ptr }

func (this *QLoggingCategory) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QLoggingCategory) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQLoggingCategoryFromPointer(cthis unsafe.Pointer) *QLoggingCategory {
	return &QLoggingCategory{&qtrt.CObject{cthis}}
}
func (*QLoggingCategory) NewFromPointer(cthis unsafe.Pointer) *QLoggingCategory {
	return NewQLoggingCategoryFromPointer(cthis)
}

// /usr/include/qt/QtCore/qloggingcategory.h:53
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLoggingCategory(const char *)
func NewQLoggingCategory(category string) *QLoggingCategory {
	var convArg0 = qtrt.CString(category)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QLoggingCategoryC2EPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLoggingCategoryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLoggingCategory)
	return gothis
}

// /usr/include/qt/QtCore/qloggingcategory.h:54
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QLoggingCategory(const char *, enum QtMsgType)
func NewQLoggingCategory_1(category string, severityLevel int) *QLoggingCategory {
	var convArg0 = qtrt.CString(category)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QLoggingCategoryC2EPKc9QtMsgType", qtrt.FFI_TYPE_POINTER, convArg0, severityLevel)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLoggingCategoryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLoggingCategory)
	return gothis
}

// /usr/include/qt/QtCore/qloggingcategory.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QLoggingCategory()
func DeleteQLoggingCategory(this *QLoggingCategory) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QLoggingCategoryD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qloggingcategory.h:57
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabled(enum QtMsgType) const
func (this *QLoggingCategory) IsEnabled(type_ int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QLoggingCategory9isEnabledE9QtMsgType", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qloggingcategory.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnabled(enum QtMsgType, _Bool)
func (this *QLoggingCategory) SetEnabled(type_ int, enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QLoggingCategory10setEnabledE9QtMsgTypeb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qloggingcategory.h:61
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDebugEnabled() const
func (this *QLoggingCategory) IsDebugEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QLoggingCategory14isDebugEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qloggingcategory.h:62
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isInfoEnabled() const
func (this *QLoggingCategory) IsInfoEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QLoggingCategory13isInfoEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qloggingcategory.h:63
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isWarningEnabled() const
func (this *QLoggingCategory) IsWarningEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QLoggingCategory16isWarningEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qloggingcategory.h:64
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isCriticalEnabled() const
func (this *QLoggingCategory) IsCriticalEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QLoggingCategory17isCriticalEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qloggingcategory.h:71
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const char * categoryName() const
func (this *QLoggingCategory) CategoryName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QLoggingCategory12categoryNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qloggingcategory.h:77
// index:0
// Public static Visibility=Default Availability=Available
// [8] QLoggingCategory * defaultCategory()
func (this *QLoggingCategory) DefaultCategory() *QLoggingCategory /*777 QLoggingCategory **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QLoggingCategory15defaultCategoryEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLoggingCategoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QLoggingCategory_DefaultCategory() *QLoggingCategory /*777 QLoggingCategory **/ {
	var nilthis *QLoggingCategory
	rv := nilthis.DefaultCategory()
	return rv
}

// /usr/include/qt/QtCore/qloggingcategory.h:82
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setFilterRules(const QString &)
func (this *QLoggingCategory) SetFilterRules(rules string) {
	var tmpArg0 = NewQString_5(rules)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QLoggingCategory14setFilterRulesERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QLoggingCategory_SetFilterRules(rules string) {
	var nilthis *QLoggingCategory
	nilthis.SetFilterRules(rules)
}

type QLoggingCategory__ = int

const QLoggingCategory__DebugShift QLoggingCategory__ = 0
const QLoggingCategory__WarningShift QLoggingCategory__ = 8
const QLoggingCategory__CriticalShift QLoggingCategory__ = 16
const QLoggingCategory__InfoShift QLoggingCategory__ = 24

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
