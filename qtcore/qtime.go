package qtcore

// /usr/include/qt/QtCore/qdatetime.h
// #include <qdatetime.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 39
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QTime struct {
	*qtrt.CObject
}
type QTime_ITF interface {
	QTime_PTR() *QTime
}

func (ptr *QTime) QTime_PTR() *QTime { return ptr }

func (this *QTime) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTime) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTimeFromPointer(cthis unsafe.Pointer) *QTime {
	return &QTime{&qtrt.CObject{cthis}}
}
func (*QTime) NewFromPointer(cthis unsafe.Pointer) *QTime {
	return NewQTimeFromPointer(cthis)
}

// /usr/include/qt/QtCore/qdatetime.h:159
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QTime()
func NewQTime() *QTime {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QTimeC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTimeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTime)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:161
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTime(int, int, int, int)
func NewQTime_1(h int, m int, s int, ms int) *QTime {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QTimeC2Eiiii", qtrt.FFI_TYPE_POINTER, h, m, s, ms)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTimeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTime)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:161
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTime(int, int, int, int)
func NewQTime_1_(h int, m int) *QTime {
	// arg: 2, int=Int, =Invalid,
	s := int(0)
	// arg: 3, int=Int, =Invalid,
	ms := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN5QTimeC2Eiiii", qtrt.FFI_TYPE_POINTER, h, m, s, ms)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTimeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTime)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:161
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTime(int, int, int, int)
func NewQTime_1_1(h int, m int, s int) *QTime {
	// arg: 3, int=Int, =Invalid,
	ms := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN5QTimeC2Eiiii", qtrt.FFI_TYPE_POINTER, h, m, s, ms)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTimeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTime)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:163
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull() const
func (this *QTime) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTime6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const
func (this *QTime) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTime7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:199
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool isValid(int, int, int, int)
func (this *QTime) IsValid_1(h int, m int, s int, ms int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QTime7isValidEiiii", qtrt.FFI_TYPE_POINTER, h, m, s, ms)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QTime_IsValid_1(h int, m int, s int, ms int) bool {
	var nilthis *QTime
	rv := nilthis.IsValid_1(h, m, s, ms)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:199
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool isValid(int, int, int, int)
func (this *QTime) IsValid_1_(h int, m int, s int) bool {
	// arg: 3, int=Int, =Invalid,
	ms := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN5QTime7isValidEiiii", qtrt.FFI_TYPE_POINTER, h, m, s, ms)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:166
// index:0
// Public Visibility=Default Availability=Available
// [4] int hour() const
func (this *QTime) Hour() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTime4hourEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:167
// index:0
// Public Visibility=Default Availability=Available
// [4] int minute() const
func (this *QTime) Minute() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTime6minuteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:168
// index:0
// Public Visibility=Default Availability=Available
// [4] int second() const
func (this *QTime) Second() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTime6secondEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:169
// index:0
// Public Visibility=Default Availability=Available
// [4] int msec() const
func (this *QTime) Msec() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTime4msecEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:171
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString(Qt::DateFormat) const
func (this *QTime) ToString(f int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTime8toStringEN2Qt10DateFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:171
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString(Qt::DateFormat) const
func (this *QTime) ToString__() string {
	// arg: 0, Qt::DateFormat=Elaborated, Qt::DateFormat=Enum,
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTime8toStringEN2Qt10DateFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:173
// index:1
// Public Visibility=Default Availability=Available
// [8] QString toString(const QString &) const
func (this *QTime) ToString_1(format string) string {
	var tmpArg0 = NewQString_5(format)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTime8toStringERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:175
// index:2
// Public Visibility=Default Availability=Available
// [8] QString toString(QStringView) const
func (this *QTime) ToString_2(format QStringView_ITF /*123*/) string {
	var convArg0 unsafe.Pointer
	if format != nil && format.QStringView_PTR() != nil {
		convArg0 = format.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTime8toStringE11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:177
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setHMS(int, int, int, int)
func (this *QTime) SetHMS(h int, m int, s int, ms int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QTime6setHMSEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), h, m, s, ms)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:177
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setHMS(int, int, int, int)
func (this *QTime) SetHMS__(h int, m int, s int) bool {
	// arg: 3, int=Int, =Invalid,
	ms := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN5QTime6setHMSEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), h, m, s, ms)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:179
// index:0
// Public Visibility=Default Availability=Available
// [4] QTime addSecs(int) const
func (this *QTime) AddSecs(secs int) *QTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTime7addSecsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), secs)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:180
// index:0
// Public Visibility=Default Availability=Available
// [4] int secsTo(const QTime &) const
func (this *QTime) SecsTo(arg0 QTime_ITF) int {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTime_PTR() != nil {
		convArg0 = arg0.QTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTime6secsToERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:181
// index:0
// Public Visibility=Default Availability=Available
// [4] QTime addMSecs(int) const
func (this *QTime) AddMSecs(ms int) *QTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTime8addMSecsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ms)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:182
// index:0
// Public Visibility=Default Availability=Available
// [4] int msecsTo(const QTime &) const
func (this *QTime) MsecsTo(arg0 QTime_ITF) int {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTime_PTR() != nil {
		convArg0 = arg0.QTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTime7msecsToERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:184
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QTime &) const
func (this *QTime) Operator_equal_equal(other QTime_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTime_PTR() != nil {
		convArg0 = other.QTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTimeeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:185
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QTime &) const
func (this *QTime) Operator_not_equal(other QTime_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTime_PTR() != nil {
		convArg0 = other.QTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTimeneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:186
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<(const QTime &) const
func (this *QTime) Operator_less_than(other QTime_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTime_PTR() != nil {
		convArg0 = other.QTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTimeltERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:187
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<=(const QTime &) const
func (this *QTime) Operator_less_than_equal(other QTime_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTime_PTR() != nil {
		convArg0 = other.QTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTimeleERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:188
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>(const QTime &) const
func (this *QTime) Operator_greater_than(other QTime_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTime_PTR() != nil {
		convArg0 = other.QTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTimegtERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:189
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>=(const QTime &) const
func (this *QTime) Operator_greater_than_equal(other QTime_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTime_PTR() != nil {
		convArg0 = other.QTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTimegeERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:191
// index:0
// Public static inline Visibility=Default Availability=Available
// [4] QTime fromMSecsSinceStartOfDay(int)
func (this *QTime) FromMSecsSinceStartOfDay(msecs int) *QTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QTime24fromMSecsSinceStartOfDayEi", qtrt.FFI_TYPE_POINTER, msecs)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTime)
	return rv2
}
func QTime_FromMSecsSinceStartOfDay(msecs int) *QTime /*123*/ {
	var nilthis *QTime
	rv := nilthis.FromMSecsSinceStartOfDay(msecs)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:192
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int msecsSinceStartOfDay() const
func (this *QTime) MsecsSinceStartOfDay() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTime20msecsSinceStartOfDayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:194
// index:0
// Public static Visibility=Default Availability=Available
// [4] QTime currentTime()
func (this *QTime) CurrentTime() *QTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QTime11currentTimeEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTime)
	return rv2
}
func QTime_CurrentTime() *QTime /*123*/ {
	var nilthis *QTime
	rv := nilthis.CurrentTime()
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:196
// index:0
// Public static Visibility=Default Availability=Available
// [4] QTime fromString(const QString &, Qt::DateFormat)
func (this *QTime) FromString(s string, f int) *QTime /*123*/ {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QTime10fromStringERK7QStringN2Qt10DateFormatE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTime)
	return rv2
}
func QTime_FromString(s string, f int) *QTime /*123*/ {
	var nilthis *QTime
	rv := nilthis.FromString(s, f)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:196
// index:0
// Public static Visibility=Default Availability=Available
// [4] QTime fromString(const QString &, Qt::DateFormat)
func (this *QTime) FromString__(s string) *QTime /*123*/ {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, Qt::DateFormat=Elaborated, Qt::DateFormat=Enum,
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN5QTime10fromStringERK7QStringN2Qt10DateFormatE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:197
// index:1
// Public static Visibility=Default Availability=Available
// [4] QTime fromString(const QString &, const QString &)
func (this *QTime) FromString_1(s string, format string) *QTime /*123*/ {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(format)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QTime10fromStringERK7QStringS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTime)
	return rv2
}
func QTime_FromString_1(s string, format string) *QTime /*123*/ {
	var nilthis *QTime
	rv := nilthis.FromString_1(s, format)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:201
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start()
func (this *QTime) Start() {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QTime5startEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:202
// index:0
// Public Visibility=Default Availability=Available
// [4] int restart()
func (this *QTime) Restart() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QTime7restartEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:203
// index:0
// Public Visibility=Default Availability=Available
// [4] int elapsed() const
func (this *QTime) Elapsed() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTime7elapsedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

func DeleteQTime(this *QTime) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QTimeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QTime__TimeFlag = int

const QTime__NullTime QTime__TimeFlag = -1

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
