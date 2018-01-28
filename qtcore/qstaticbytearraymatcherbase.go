package qtcore

// /usr/include/qt/QtCore/qbytearraymatcher.h
// #include <qbytearraymatcher.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

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
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QStaticByteArrayMatcherBase struct {
	*qtrt.CObject
}

func (this *QStaticByteArrayMatcherBase) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStaticByteArrayMatcherBase) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQStaticByteArrayMatcherBaseFromPointer(cthis unsafe.Pointer) *QStaticByteArrayMatcherBase {
	return &QStaticByteArrayMatcherBase{&qtrt.CObject{cthis}}
}
func (*QStaticByteArrayMatcherBase) NewFromPointer(cthis unsafe.Pointer) *QStaticByteArrayMatcherBase {
	return NewQStaticByteArrayMatcherBaseFromPointer(cthis)
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:93
// index:0
// Protected inline
// void QStaticByteArrayMatcherBase(const char *, uint)
func NewQStaticByteArrayMatcherBase(pattern string, n uint) *QStaticByteArrayMatcherBase {
	cthis := qtrt.Calloc(1, 256) // 256
	var convArg0 = qtrt.CString(pattern)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QStaticByteArrayMatcherBaseC2EPKcj", ffiqt.FFI_TYPE_VOID, cthis, convArg0, n)
	gopp.ErrPrint(err, rv)
	gothis := NewQStaticByteArrayMatcherBaseFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:98
// index:0
// Protected
// int indexOfIn(const char *, uint, const char *, int, int)
func (this *QStaticByteArrayMatcherBase) IndexOfIn(needle string, nlen uint, haystack string, hlen int, from int) int {
	var convArg0 = qtrt.CString(needle)
	defer qtrt.FreeMem(convArg0)
	var convArg2 = qtrt.CString(haystack)
	defer qtrt.FreeMem(convArg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QStaticByteArrayMatcherBase9indexOfInEPKcjS1_ii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, nlen, convArg2, hlen, from)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end