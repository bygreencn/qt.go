package qtcore

// /usr/include/qt/QtCore/qjsondocument.h
// #include <qjsondocument.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QJsonDocument struct {
	*qtrt.CObject
}
type QJsonDocument_ITF interface {
	QJsonDocument_PTR() *QJsonDocument
}

func (ptr *QJsonDocument) QJsonDocument_PTR() *QJsonDocument { return ptr }

func (this *QJsonDocument) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QJsonDocument) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQJsonDocumentFromPointer(cthis unsafe.Pointer) *QJsonDocument {
	return &QJsonDocument{&qtrt.CObject{cthis}}
}
func (*QJsonDocument) NewFromPointer(cthis unsafe.Pointer) *QJsonDocument {
	return NewQJsonDocumentFromPointer(cthis)
}

// /usr/include/qt/QtCore/qjsondocument.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QJsonDocument()
func NewQJsonDocument() *QJsonDocument {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocumentC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonDocument)
	return gothis
}

// /usr/include/qt/QtCore/qjsondocument.h:89
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QJsonDocument(const QJsonObject &)
func NewQJsonDocument_1(object QJsonObject_ITF) *QJsonDocument {
	var convArg0 unsafe.Pointer
	if object != nil && object.QJsonObject_PTR() != nil {
		convArg0 = object.QJsonObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocumentC2ERK11QJsonObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonDocument)
	return gothis
}

// /usr/include/qt/QtCore/qjsondocument.h:90
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QJsonDocument(const QJsonArray &)
func NewQJsonDocument_2(array QJsonArray_ITF) *QJsonDocument {
	var convArg0 unsafe.Pointer
	if array != nil && array.QJsonArray_PTR() != nil {
		convArg0 = array.QJsonArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocumentC2ERK10QJsonArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonDocument)
	return gothis
}

// /usr/include/qt/QtCore/qjsondocument.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QJsonDocument()
func DeleteQJsonDocument(this *QJsonDocument) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocumentD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qjsondocument.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] QJsonDocument & operator=(const QJsonDocument &)
func (this *QJsonDocument) Operator_equal(other QJsonDocument_ITF) *QJsonDocument {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonDocument_PTR() != nil {
		convArg0 = other.QJsonDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocumentaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonDocument)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:102
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QJsonDocument & operator=(QJsonDocument &&)
func (this *QJsonDocument) Operator_equal_1(other unsafe.Pointer /*333*/) *QJsonDocument {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocumentaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonDocument)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:108
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QJsonDocument &)
func (this *QJsonDocument) Swap(other QJsonDocument_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonDocument_PTR() != nil {
		convArg0 = other.QJsonDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocument4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:118
// index:0
// Public static Visibility=Default Availability=Available
// [8] QJsonDocument fromRawData(const char *, int, enum QJsonDocument::DataValidation)
func (this *QJsonDocument) FromRawData(data string, size int, validation int) *QJsonDocument /*123*/ {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocument11fromRawDataEPKciNS_14DataValidationE", qtrt.FFI_TYPE_POINTER, convArg0, size, validation)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonDocument)
	return rv2
}
func QJsonDocument_FromRawData(data string, size int, validation int) *QJsonDocument /*123*/ {
	var nilthis *QJsonDocument
	rv := nilthis.FromRawData(data, size, validation)
	return rv
}

// /usr/include/qt/QtCore/qjsondocument.h:118
// index:0
// Public static Visibility=Default Availability=Available
// [8] QJsonDocument fromRawData(const char *, int, enum QJsonDocument::DataValidation)
func (this *QJsonDocument) FromRawData__(data string, size int) *QJsonDocument /*123*/ {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	// arg: 2, QJsonDocument::DataValidation=Enum, QJsonDocument::DataValidation=Enum,
	validation := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocument11fromRawDataEPKciNS_14DataValidationE", qtrt.FFI_TYPE_POINTER, convArg0, size, validation)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonDocument)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * rawData(int *) const
func (this *QJsonDocument) RawData(size unsafe.Pointer /*666*/) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocument7rawDataEPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:121
// index:0
// Public static Visibility=Default Availability=Available
// [8] QJsonDocument fromBinaryData(const QByteArray &, enum QJsonDocument::DataValidation)
func (this *QJsonDocument) FromBinaryData(data QByteArray_ITF, validation int) *QJsonDocument /*123*/ {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocument14fromBinaryDataERK10QByteArrayNS_14DataValidationE", qtrt.FFI_TYPE_POINTER, convArg0, validation)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonDocument)
	return rv2
}
func QJsonDocument_FromBinaryData(data QByteArray_ITF, validation int) *QJsonDocument /*123*/ {
	var nilthis *QJsonDocument
	rv := nilthis.FromBinaryData(data, validation)
	return rv
}

// /usr/include/qt/QtCore/qjsondocument.h:121
// index:0
// Public static Visibility=Default Availability=Available
// [8] QJsonDocument fromBinaryData(const QByteArray &, enum QJsonDocument::DataValidation)
func (this *QJsonDocument) FromBinaryData__(data QByteArray_ITF) *QJsonDocument /*123*/ {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	// arg: 1, QJsonDocument::DataValidation=Enum, QJsonDocument::DataValidation=Enum,
	validation := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocument14fromBinaryDataERK10QByteArrayNS_14DataValidationE", qtrt.FFI_TYPE_POINTER, convArg0, validation)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonDocument)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:122
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toBinaryData() const
func (this *QJsonDocument) ToBinaryData() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocument12toBinaryDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:124
// index:0
// Public static Visibility=Default Availability=Available
// [8] QJsonDocument fromVariant(const QVariant &)
func (this *QJsonDocument) FromVariant(variant QVariant_ITF) *QJsonDocument /*123*/ {
	var convArg0 unsafe.Pointer
	if variant != nil && variant.QVariant_PTR() != nil {
		convArg0 = variant.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocument11fromVariantERK8QVariant", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonDocument)
	return rv2
}
func QJsonDocument_FromVariant(variant QVariant_ITF) *QJsonDocument /*123*/ {
	var nilthis *QJsonDocument
	rv := nilthis.FromVariant(variant)
	return rv
}

// /usr/include/qt/QtCore/qjsondocument.h:125
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant toVariant() const
func (this *QJsonDocument) ToVariant() *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocument9toVariantEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:137
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toJson() const
func (this *QJsonDocument) ToJson() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocument6toJsonEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:138
// index:1
// Public Visibility=Default Availability=Available
// [8] QByteArray toJson(enum QJsonDocument::JsonFormat) const
func (this *QJsonDocument) ToJson_1(format int) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocument6toJsonENS_10JsonFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const
func (this *QJsonDocument) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocument7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsondocument.h:142
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isArray() const
func (this *QJsonDocument) IsArray() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocument7isArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsondocument.h:143
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isObject() const
func (this *QJsonDocument) IsObject() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocument8isObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsondocument.h:145
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonObject object() const
func (this *QJsonDocument) Object() *QJsonObject /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocument6objectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonObject)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:146
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonArray array() const
func (this *QJsonDocument) Array() *QJsonArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocument5arrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setObject(const QJsonObject &)
func (this *QJsonDocument) SetObject(object QJsonObject_ITF) {
	var convArg0 unsafe.Pointer
	if object != nil && object.QJsonObject_PTR() != nil {
		convArg0 = object.QJsonObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocument9setObjectERK11QJsonObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setArray(const QJsonArray &)
func (this *QJsonDocument) SetArray(array QJsonArray_ITF) {
	var convArg0 unsafe.Pointer
	if array != nil && array.QJsonArray_PTR() != nil {
		convArg0 = array.QJsonArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocument8setArrayERK10QJsonArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:151
// index:0
// Public Visibility=Default Availability=Available
// [24] const QJsonValue operator[](const QString &) const
func (this *QJsonDocument) Operator_get_index(key string) *QJsonValue /*123*/ {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocumentixERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:152
// index:1
// Public Visibility=Default Availability=Available
// [24] const QJsonValue operator[](QLatin1String) const
func (this *QJsonDocument) Operator_get_index_1(key QLatin1String_ITF /*123*/) *QJsonValue /*123*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QLatin1String_PTR() != nil {
		convArg0 = key.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocumentixE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:153
// index:2
// Public Visibility=Default Availability=Available
// [24] const QJsonValue operator[](int) const
func (this *QJsonDocument) Operator_get_index_2(i int) *QJsonValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocumentixEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:155
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QJsonDocument &) const
func (this *QJsonDocument) Operator_equal_equal(other QJsonDocument_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonDocument_PTR() != nil {
		convArg0 = other.QJsonDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocumenteqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsondocument.h:156
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QJsonDocument &) const
func (this *QJsonDocument) Operator_not_equal(other QJsonDocument_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonDocument_PTR() != nil {
		convArg0 = other.QJsonDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocumentneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsondocument.h:158
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const
func (this *QJsonDocument) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocument6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

type QJsonDocument__DataValidation = int

const QJsonDocument__Validate QJsonDocument__DataValidation = 0
const QJsonDocument__BypassValidation QJsonDocument__DataValidation = 1

type QJsonDocument__JsonFormat = int

const QJsonDocument__Indented QJsonDocument__JsonFormat = 0
const QJsonDocument__Compact QJsonDocument__JsonFormat = 1

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
