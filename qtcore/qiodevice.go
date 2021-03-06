package qtcore

// /usr/include/qt/QtCore/qiodevice.h
// #include <qiodevice.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// long long readData(char *, qint64)
func (this *QIODevice) InheritReadData(f func(data string, maxlen int64) int64) {
	qtrt.SetAllInheritCallback(this, "readData", f)
}

// long long readLineData(char *, qint64)
func (this *QIODevice) InheritReadLineData(f func(data string, maxlen int64) int64) {
	qtrt.SetAllInheritCallback(this, "readLineData", f)
}

// long long writeData(const char *, qint64)
func (this *QIODevice) InheritWriteData(f func(data string, len_ int64) int64) {
	qtrt.SetAllInheritCallback(this, "writeData", f)
}

// void setOpenMode(QIODevice::OpenMode)
func (this *QIODevice) InheritSetOpenMode(f func(openMode int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setOpenMode", f)
}

// void setErrorString(const class QString &)
func (this *QIODevice) InheritSetErrorString(f func(errorString string) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setErrorString", f)
}

type QIODevice struct {
	*QObject
}
type QIODevice_ITF interface {
	QObject_ITF
	QIODevice_PTR() *QIODevice
}

func (ptr *QIODevice) QIODevice_PTR() *QIODevice { return ptr }

func (this *QIODevice) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QIODevice) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQIODeviceFromPointer(cthis unsafe.Pointer) *QIODevice {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QIODevice{bcthis0}
}
func (*QIODevice) NewFromPointer(cthis unsafe.Pointer) *QIODevice {
	return NewQIODeviceFromPointer(cthis)
}

// /usr/include/qt/QtCore/qiodevice.h:68
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QIODevice) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QIODevice10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qiodevice.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QIODevice()
func NewQIODevice() *QIODevice {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODeviceC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QIODevice")
	return gothis
}

// /usr/include/qt/QtCore/qiodevice.h:85
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QIODevice(QObject *)
func NewQIODevice_1(parent QObject_ITF /*777 QObject **/) *QIODevice {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODeviceC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QIODevice")
	return gothis
}

// /usr/include/qt/QtCore/qiodevice.h:87
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QIODevice()
func DeleteQIODevice(this *QIODevice) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODeviceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qiodevice.h:89
// index:0
// Public Visibility=Default Availability=Available
// [4] QIODevice::OpenMode openMode() const
func (this *QIODevice) OpenMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QIODevice8openModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qiodevice.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextModeEnabled(_Bool)
func (this *QIODevice) SetTextModeEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice18setTextModeEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:92
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTextModeEnabled() const
func (this *QIODevice) IsTextModeEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QIODevice17isTextModeEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qiodevice.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isOpen() const
func (this *QIODevice) IsOpen() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QIODevice6isOpenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qiodevice.h:95
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadable() const
func (this *QIODevice) IsReadable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QIODevice10isReadableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qiodevice.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isWritable() const
func (this *QIODevice) IsWritable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QIODevice10isWritableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qiodevice.h:97
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isSequential() const
func (this *QIODevice) IsSequential() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QIODevice12isSequentialEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qiodevice.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] int readChannelCount() const
func (this *QIODevice) ReadChannelCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QIODevice16readChannelCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qiodevice.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] int writeChannelCount() const
func (this *QIODevice) WriteChannelCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QIODevice17writeChannelCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qiodevice.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentReadChannel() const
func (this *QIODevice) CurrentReadChannel() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QIODevice18currentReadChannelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qiodevice.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentReadChannel(int)
func (this *QIODevice) SetCurrentReadChannel(channel int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice21setCurrentReadChannelEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), channel)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:103
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentWriteChannel() const
func (this *QIODevice) CurrentWriteChannel() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QIODevice19currentWriteChannelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qiodevice.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentWriteChannel(int)
func (this *QIODevice) SetCurrentWriteChannel(channel int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice22setCurrentWriteChannelEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), channel)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:106
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool open(QIODevice::OpenMode)
func (this *QIODevice) Open(mode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice4openE6QFlagsINS_12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qiodevice.h:107
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void close()
func (this *QIODevice) Close() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:111
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 pos() const
func (this *QIODevice) Pos() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QIODevice3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qiodevice.h:112
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 size() const
func (this *QIODevice) Size() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QIODevice4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qiodevice.h:113
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool seek(qint64)
func (this *QIODevice) Seek(pos int64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice4seekEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qiodevice.h:114
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool atEnd() const
func (this *QIODevice) AtEnd() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QIODevice5atEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qiodevice.h:115
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool reset()
func (this *QIODevice) Reset() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qiodevice.h:117
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 bytesAvailable() const
func (this *QIODevice) BytesAvailable() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QIODevice14bytesAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qiodevice.h:118
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 bytesToWrite() const
func (this *QIODevice) BytesToWrite() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QIODevice12bytesToWriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qiodevice.h:120
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 read(char *, qint64)
func (this *QIODevice) Read(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice4readEPcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qiodevice.h:121
// index:1
// Public Visibility=Default Availability=Available
// [8] QByteArray read(qint64)
func (this *QIODevice) Read_1(maxlen int64) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice4readEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxlen)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qiodevice.h:122
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray readAll()
func (this *QIODevice) ReadAll() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice7readAllEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qiodevice.h:123
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 readLine(char *, qint64)
func (this *QIODevice) ReadLine(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice8readLineEPcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qiodevice.h:124
// index:1
// Public Visibility=Default Availability=Available
// [8] QByteArray readLine(qint64)
func (this *QIODevice) ReadLine_1(maxlen int64) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice8readLineEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxlen)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qiodevice.h:124
// index:1
// Public Visibility=Default Availability=Available
// [8] QByteArray readLine(qint64)
func (this *QIODevice) ReadLine_1_() *QByteArray /*123*/ {
	// arg: 0, qint64=Typedef, qint64=Typedef, long long
	maxlen := int64(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice8readLineEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxlen)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qiodevice.h:125
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool canReadLine() const
func (this *QIODevice) CanReadLine() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QIODevice11canReadLineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qiodevice.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void startTransaction()
func (this *QIODevice) StartTransaction() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice16startTransactionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void commitTransaction()
func (this *QIODevice) CommitTransaction() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice17commitTransactionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rollbackTransaction()
func (this *QIODevice) RollbackTransaction() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice19rollbackTransactionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:130
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTransactionStarted() const
func (this *QIODevice) IsTransactionStarted() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QIODevice20isTransactionStartedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qiodevice.h:132
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 write(const char *, qint64)
func (this *QIODevice) Write(data string, len_ int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice5writeEPKcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qiodevice.h:133
// index:1
// Public Visibility=Default Availability=Available
// [8] qint64 write(const char *)
func (this *QIODevice) Write_1(data string) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice5writeEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qiodevice.h:134
// index:2
// Public inline Visibility=Default Availability=Available
// [8] qint64 write(const QByteArray &)
func (this *QIODevice) Write_2(data QByteArray_ITF) int64 {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice5writeERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qiodevice.h:137
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 peek(char *, qint64)
func (this *QIODevice) Peek(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice4peekEPcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qiodevice.h:138
// index:1
// Public Visibility=Default Availability=Available
// [8] QByteArray peek(qint64)
func (this *QIODevice) Peek_1(maxlen int64) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice4peekEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxlen)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qiodevice.h:139
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 skip(qint64)
func (this *QIODevice) Skip(maxSize int64) int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice4skipEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxSize)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qiodevice.h:141
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForReadyRead(int)
func (this *QIODevice) WaitForReadyRead(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice16waitForReadyReadEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qiodevice.h:142
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForBytesWritten(int)
func (this *QIODevice) WaitForBytesWritten(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice19waitForBytesWrittenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qiodevice.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ungetChar(char)
func (this *QIODevice) UngetChar(c byte) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice9ungetCharEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:145
// index:0
// Public Visibility=Default Availability=Available
// [1] bool putChar(char)
func (this *QIODevice) PutChar(c byte) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice7putCharEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qiodevice.h:146
// index:0
// Public Visibility=Default Availability=Available
// [1] bool getChar(char *)
func (this *QIODevice) GetChar(c string) bool {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice7getCharEPc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qiodevice.h:148
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const
func (this *QIODevice) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QIODevice11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qiodevice.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void readyRead()
func (this *QIODevice) ReadyRead() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice9readyReadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void channelReadyRead(int)
func (this *QIODevice) ChannelReadyRead(channel int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice16channelReadyReadEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), channel)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:154
// index:0
// Public Visibility=Default Availability=Available
// [-2] void bytesWritten(qint64)
func (this *QIODevice) BytesWritten(bytes int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice12bytesWrittenEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bytes)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void channelBytesWritten(int, qint64)
func (this *QIODevice) ChannelBytesWritten(channel int, bytes int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice19channelBytesWrittenEix", qtrt.FFI_TYPE_POINTER, this.GetCthis(), channel, bytes)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void aboutToClose()
func (this *QIODevice) AboutToClose() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice12aboutToCloseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void readChannelFinished()
func (this *QIODevice) ReadChannelFinished() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice19readChannelFinishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:166
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [8] qint64 readData(char *, qint64)
func (this *QIODevice) ReadData(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice8readDataEPcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qiodevice.h:167
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 readLineData(char *, qint64)
func (this *QIODevice) ReadLineData(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice12readLineDataEPcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qiodevice.h:168
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [8] qint64 writeData(const char *, qint64)
func (this *QIODevice) WriteData(data string, len_ int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice9writeDataEPKcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qiodevice.h:170
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setOpenMode(QIODevice::OpenMode)
func (this *QIODevice) SetOpenMode(openMode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice11setOpenModeE6QFlagsINS_12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), openMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:172
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setErrorString(const QString &)
func (this *QIODevice) SetErrorString(errorString string) {
	var tmpArg0 = NewQString_5(errorString)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QIODevice14setErrorStringERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

type QIODevice__OpenModeFlag = int

const QIODevice__NotOpen QIODevice__OpenModeFlag = 0
const QIODevice__ReadOnly QIODevice__OpenModeFlag = 1
const QIODevice__WriteOnly QIODevice__OpenModeFlag = 2
const QIODevice__ReadWrite QIODevice__OpenModeFlag = 3
const QIODevice__Append QIODevice__OpenModeFlag = 4
const QIODevice__Truncate QIODevice__OpenModeFlag = 8
const QIODevice__Text QIODevice__OpenModeFlag = 16
const QIODevice__Unbuffered QIODevice__OpenModeFlag = 32

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
