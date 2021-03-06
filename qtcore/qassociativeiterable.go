package qtcore

// /usr/include/qt/QtCore/qvariant.h
// #include <qvariant.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QAssociativeIterable struct {
	*qtrt.CObject
}
type QAssociativeIterable_ITF interface {
	QAssociativeIterable_PTR() *QAssociativeIterable
}

func (ptr *QAssociativeIterable) QAssociativeIterable_PTR() *QAssociativeIterable { return ptr }

func (this *QAssociativeIterable) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAssociativeIterable) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAssociativeIterableFromPointer(cthis unsafe.Pointer) *QAssociativeIterable {
	return &QAssociativeIterable{&qtrt.CObject{cthis}}
}
func (*QAssociativeIterable) NewFromPointer(cthis unsafe.Pointer) *QAssociativeIterable {
	return NewQAssociativeIterableFromPointer(cthis)
}

// /usr/include/qt/QtCore/qvariant.h:680
// index:0
// Public Visibility=Default Availability=Available
// [120] QAssociativeIterable::const_iterator begin() const
func (this *QAssociativeIterable) Begin() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAssociativeIterable5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qvariant.h:681
// index:0
// Public Visibility=Default Availability=Available
// [120] QAssociativeIterable::const_iterator end() const
func (this *QAssociativeIterable) End() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAssociativeIterable3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qvariant.h:682
// index:0
// Public Visibility=Default Availability=Available
// [120] QAssociativeIterable::const_iterator find(const QVariant &) const
func (this *QAssociativeIterable) Find(key QVariant_ITF) unsafe.Pointer /*444*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QVariant_PTR() != nil {
		convArg0 = key.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAssociativeIterable4findERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qvariant.h:684
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant value(const QVariant &) const
func (this *QAssociativeIterable) Value(key QVariant_ITF) *QVariant /*123*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QVariant_PTR() != nil {
		convArg0 = key.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAssociativeIterable5valueERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:686
// index:0
// Public Visibility=Default Availability=Available
// [4] int size() const
func (this *QAssociativeIterable) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAssociativeIterable4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

func DeleteQAssociativeIterable(this *QAssociativeIterable) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAssociativeIterableD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
