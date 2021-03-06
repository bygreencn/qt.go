package qtcore

// /usr/include/qt/QtCore/qtemporarydir.h
// #include <qtemporarydir.h>
// #include <QtCore>

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

//  ext block end

//  body block begin

type QTemporaryDir struct {
	*qtrt.CObject
}
type QTemporaryDir_ITF interface {
	QTemporaryDir_PTR() *QTemporaryDir
}

func (ptr *QTemporaryDir) QTemporaryDir_PTR() *QTemporaryDir { return ptr }

func (this *QTemporaryDir) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTemporaryDir) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTemporaryDirFromPointer(cthis unsafe.Pointer) *QTemporaryDir {
	return &QTemporaryDir{&qtrt.CObject{cthis}}
}
func (*QTemporaryDir) NewFromPointer(cthis unsafe.Pointer) *QTemporaryDir {
	return NewQTemporaryDirFromPointer(cthis)
}

// /usr/include/qt/QtCore/qtemporarydir.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTemporaryDir()
func NewQTemporaryDir() *QTemporaryDir {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTemporaryDirC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTemporaryDirFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTemporaryDir)
	return gothis
}

// /usr/include/qt/QtCore/qtemporarydir.h:57
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTemporaryDir(const QString &)
func NewQTemporaryDir_1(templateName string) *QTemporaryDir {
	var tmpArg0 = NewQString_5(templateName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTemporaryDirC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTemporaryDirFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTemporaryDir)
	return gothis
}

// /usr/include/qt/QtCore/qtemporarydir.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTemporaryDir()
func DeleteQTemporaryDir(this *QTemporaryDir) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTemporaryDirD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qtemporarydir.h:60
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const
func (this *QTemporaryDir) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTemporaryDir7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtemporarydir.h:61
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const
func (this *QTemporaryDir) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTemporaryDir11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtemporarydir.h:63
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoRemove() const
func (this *QTemporaryDir) AutoRemove() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTemporaryDir10autoRemoveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtemporarydir.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoRemove(_Bool)
func (this *QTemporaryDir) SetAutoRemove(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTemporaryDir13setAutoRemoveEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporarydir.h:65
// index:0
// Public Visibility=Default Availability=Available
// [1] bool remove()
func (this *QTemporaryDir) Remove() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTemporaryDir6removeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtemporarydir.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QString path() const
func (this *QTemporaryDir) Path() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTemporaryDir4pathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtemporarydir.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QString filePath(const QString &) const
func (this *QTemporaryDir) FilePath(fileName string) string {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTemporaryDir8filePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
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
