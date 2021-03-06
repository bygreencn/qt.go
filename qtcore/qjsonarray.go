package qtcore

// /usr/include/qt/QtCore/qjsonarray.h
// #include <qjsonarray.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QJsonArray struct {
	*qtrt.CObject
}
type QJsonArray_ITF interface {
	QJsonArray_PTR() *QJsonArray
}

func (ptr *QJsonArray) QJsonArray_PTR() *QJsonArray { return ptr }

func (this *QJsonArray) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QJsonArray) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQJsonArrayFromPointer(cthis unsafe.Pointer) *QJsonArray {
	return &QJsonArray{&qtrt.CObject{cthis}}
}
func (*QJsonArray) NewFromPointer(cthis unsafe.Pointer) *QJsonArray {
	return NewQJsonArrayFromPointer(cthis)
}

// /usr/include/qt/QtCore/qjsonarray.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QJsonArray()
func NewQJsonArray() *QJsonArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArrayC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonArray)
	return gothis
}

// /usr/include/qt/QtCore/qjsonarray.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QJsonArray()
func DeleteQJsonArray(this *QJsonArray) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArrayD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qjsonarray.h:73
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonArray & operator=(const QJsonArray &)
func (this *QJsonArray) Operator_equal(other QJsonArray_ITF) *QJsonArray {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonArray_PTR() != nil {
		convArg0 = other.QJsonArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArrayaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:83
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray & operator=(QJsonArray &&)
func (this *QJsonArray) Operator_equal_1(other unsafe.Pointer /*333*/) *QJsonArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArrayaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:89
// index:0
// Public static Visibility=Default Availability=Available
// [16] QJsonArray fromStringList(const QStringList &)
func (this *QJsonArray) FromStringList(list QStringList_ITF) *QJsonArray /*123*/ {
	var convArg0 unsafe.Pointer
	if list != nil && list.QStringList_PTR() != nil {
		convArg0 = list.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray14fromStringListERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}
func QJsonArray_FromStringList(list QStringList_ITF) *QJsonArray /*123*/ {
	var nilthis *QJsonArray
	rv := nilthis.FromStringList(list)
	return rv
}

// /usr/include/qt/QtCore/qjsonarray.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QVariantList toVariantList() const
func (this *QJsonArray) ToVariantList() *QVariantList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray13toVariantListEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] int size() const
func (this *QJsonArray) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qjsonarray.h:94
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int count() const
func (this *QJsonArray) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qjsonarray.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const
func (this *QJsonArray) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonarray.h:97
// index:0
// Public Visibility=Default Availability=Available
// [24] QJsonValue at(int) const
func (this *QJsonArray) At(i int) *QJsonValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray2atEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:98
// index:0
// Public Visibility=Default Availability=Available
// [24] QJsonValue first() const
func (this *QJsonArray) First() *QJsonValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray5firstEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:99
// index:0
// Public Visibility=Default Availability=Available
// [24] QJsonValue last() const
func (this *QJsonArray) Last() *QJsonValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray4lastEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeAt(int)
func (this *QJsonArray) RemoveAt(i int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray8removeAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:104
// index:0
// Public Visibility=Default Availability=Available
// [24] QJsonValue takeAt(int)
func (this *QJsonArray) TakeAt(i int) *QJsonValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray6takeAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:105
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void removeFirst()
func (this *QJsonArray) RemoveFirst() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray11removeFirstEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:106
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void removeLast()
func (this *QJsonArray) RemoveLast() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray10removeLastEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void replace(int, const QJsonValue &)
func (this *QJsonArray) Replace(i int, value QJsonValue_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QJsonValue_PTR() != nil {
		convArg1 = value.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray7replaceEiRK10QJsonValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:111
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QJsonValue &) const
func (this *QJsonArray) Contains(element QJsonValue_ITF) bool {
	var convArg0 unsafe.Pointer
	if element != nil && element.QJsonValue_PTR() != nil {
		convArg0 = element.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray8containsERK10QJsonValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonarray.h:112
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonValueRef operator[](int)
func (this *QJsonArray) Operator_get_index(i int) *QJsonValueRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArrayixEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValueRef)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:113
// index:1
// Public Visibility=Default Availability=Available
// [24] QJsonValue operator[](int) const
func (this *QJsonArray) Operator_get_index_1(i int) *QJsonValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArrayixEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:115
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QJsonArray &) const
func (this *QJsonArray) Operator_equal_equal(other QJsonArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonArray_PTR() != nil {
		convArg0 = other.QJsonArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArrayeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonarray.h:116
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QJsonArray &) const
func (this *QJsonArray) Operator_not_equal(other QJsonArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonArray_PTR() != nil {
		convArg0 = other.QJsonArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArrayneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonarray.h:118
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QJsonArray &)
func (this *QJsonArray) Swap(other QJsonArray_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonArray_PTR() != nil {
		convArg0 = other.QJsonArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:214
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray::iterator begin()
func (this *QJsonArray) Begin() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonarray.h:215
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray::const_iterator begin() const
func (this *QJsonArray) Begin_1() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonarray.h:216
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray::const_iterator constBegin() const
func (this *QJsonArray) ConstBegin() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray10constBeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonarray.h:217
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray::iterator end()
func (this *QJsonArray) End() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonarray.h:218
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray::const_iterator end() const
func (this *QJsonArray) End_1() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonarray.h:219
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray::const_iterator constEnd() const
func (this *QJsonArray) ConstEnd() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray8constEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonarray.h:228
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray operator+(const QJsonValue &) const
func (this *QJsonArray) Operator_add(v QJsonValue_ITF) *QJsonArray /*123*/ {
	var convArg0 unsafe.Pointer
	if v != nil && v.QJsonValue_PTR() != nil {
		convArg0 = v.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArrayplERK10QJsonValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:230
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray & operator+=(const QJsonValue &)
func (this *QJsonArray) Operator_add_equal(v QJsonValue_ITF) *QJsonArray {
	var convArg0 unsafe.Pointer
	if v != nil && v.QJsonValue_PTR() != nil {
		convArg0 = v.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArraypLERK10QJsonValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:232
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray & operator<<(const QJsonValue &)
func (this *QJsonArray) Operator_left_shift(v QJsonValue_ITF) *QJsonArray {
	var convArg0 unsafe.Pointer
	if v != nil && v.QJsonValue_PTR() != nil {
		convArg0 = v.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArraylsERK10QJsonValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:236
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void push_back(const QJsonValue &)
func (this *QJsonArray) Push_back(t QJsonValue_ITF) {
	var convArg0 unsafe.Pointer
	if t != nil && t.QJsonValue_PTR() != nil {
		convArg0 = t.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray9push_backERK10QJsonValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:237
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void push_front(const QJsonValue &)
func (this *QJsonArray) Push_front(t QJsonValue_ITF) {
	var convArg0 unsafe.Pointer
	if t != nil && t.QJsonValue_PTR() != nil {
		convArg0 = t.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray10push_frontERK10QJsonValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:238
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void pop_front()
func (this *QJsonArray) Pop_front() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray9pop_frontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:239
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void pop_back()
func (this *QJsonArray) Pop_back() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray8pop_backEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:240
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool empty() const
func (this *QJsonArray) Empty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray5emptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
}

//  keep block end
