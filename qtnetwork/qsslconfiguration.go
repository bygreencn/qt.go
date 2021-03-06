package qtnetwork

// /usr/include/qt/QtNetwork/qsslconfiguration.h
// #include <qsslconfiguration.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 67
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QSslConfiguration struct {
	*qtrt.CObject
}
type QSslConfiguration_ITF interface {
	QSslConfiguration_PTR() *QSslConfiguration
}

func (ptr *QSslConfiguration) QSslConfiguration_PTR() *QSslConfiguration { return ptr }

func (this *QSslConfiguration) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSslConfiguration) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSslConfigurationFromPointer(cthis unsafe.Pointer) *QSslConfiguration {
	return &QSslConfiguration{&qtrt.CObject{cthis}}
}
func (*QSslConfiguration) NewFromPointer(cthis unsafe.Pointer) *QSslConfiguration {
	return NewQSslConfigurationFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSslConfiguration()
func NewQSslConfiguration() *QSslConfiguration {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfigurationC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslConfigurationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslConfiguration)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QSslConfiguration()
func DeleteQSslConfiguration(this *QSslConfiguration) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfigurationD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:83
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSslConfiguration & operator=(QSslConfiguration &&)
func (this *QSslConfiguration) Operator_equal(other unsafe.Pointer /*333*/) *QSslConfiguration {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfigurationaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslConfiguration)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:85
// index:1
// Public Visibility=Default Availability=Available
// [8] QSslConfiguration & operator=(const QSslConfiguration &)
func (this *QSslConfiguration) Operator_equal_1(other QSslConfiguration_ITF) *QSslConfiguration {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslConfiguration_PTR() != nil {
		convArg0 = other.QSslConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfigurationaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslConfiguration)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:87
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QSslConfiguration &)
func (this *QSslConfiguration) Swap(other QSslConfiguration_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslConfiguration_PTR() != nil {
		convArg0 = other.QSslConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:90
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QSslConfiguration &) const
func (this *QSslConfiguration) Operator_equal_equal(other QSslConfiguration_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslConfiguration_PTR() != nil {
		convArg0 = other.QSslConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfigurationeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:91
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QSslConfiguration &) const
func (this *QSslConfiguration) Operator_not_equal(other QSslConfiguration_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslConfiguration_PTR() != nil {
		convArg0 = other.QSslConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfigurationneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const
func (this *QSslConfiguration) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:96
// index:0
// Public Visibility=Default Availability=Available
// [4] QSsl::SslProtocol protocol() const
func (this *QSslConfiguration) Protocol() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration8protocolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProtocol(QSsl::SslProtocol)
func (this *QSslConfiguration) SetProtocol(protocol int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration11setProtocolEN4QSsl11SslProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), protocol)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] QSslSocket::PeerVerifyMode peerVerifyMode() const
func (this *QSslConfiguration) PeerVerifyMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration14peerVerifyModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPeerVerifyMode(QSslSocket::PeerVerifyMode)
func (this *QSslConfiguration) SetPeerVerifyMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration17setPeerVerifyModeEN10QSslSocket14PeerVerifyModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:103
// index:0
// Public Visibility=Default Availability=Available
// [4] int peerVerifyDepth() const
func (this *QSslConfiguration) PeerVerifyDepth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration15peerVerifyDepthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPeerVerifyDepth(int)
func (this *QSslConfiguration) SetPeerVerifyDepth(depth int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration18setPeerVerifyDepthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), depth)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:110
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslCertificate localCertificate() const
func (this *QSslConfiguration) LocalCertificate() *QSslCertificate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration16localCertificateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslCertificateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslCertificate)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLocalCertificate(const QSslCertificate &)
func (this *QSslConfiguration) SetLocalCertificate(certificate QSslCertificate_ITF) {
	var convArg0 unsafe.Pointer
	if certificate != nil && certificate.QSslCertificate_PTR() != nil {
		convArg0 = certificate.QSslCertificate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration19setLocalCertificateERK15QSslCertificate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:113
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslCertificate peerCertificate() const
func (this *QSslConfiguration) PeerCertificate() *QSslCertificate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration15peerCertificateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslCertificateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslCertificate)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslCipher sessionCipher() const
func (this *QSslConfiguration) SessionCipher() *QSslCipher /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration13sessionCipherEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslCipherFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslCipher)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:116
// index:0
// Public Visibility=Default Availability=Available
// [4] QSsl::SslProtocol sessionProtocol() const
func (this *QSslConfiguration) SessionProtocol() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration15sessionProtocolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslKey privateKey() const
func (this *QSslConfiguration) PrivateKey() *QSslKey /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration10privateKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslKey)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPrivateKey(const QSslKey &)
func (this *QSslConfiguration) SetPrivateKey(key QSslKey_ITF) {
	var convArg0 unsafe.Pointer
	if key != nil && key.QSslKey_PTR() != nil {
		convArg0 = key.QSslKey_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration13setPrivateKeyERK7QSslKey", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSslOption(QSsl::SslOption, _Bool)
func (this *QSslConfiguration) SetSslOption(option int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration12setSslOptionEN4QSsl9SslOptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:133
// index:0
// Public Visibility=Default Availability=Available
// [1] bool testSslOption(QSsl::SslOption) const
func (this *QSslConfiguration) TestSslOption(option int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration13testSslOptionEN4QSsl9SslOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:135
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray sessionTicket() const
func (this *QSslConfiguration) SessionTicket() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration13sessionTicketEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSessionTicket(const QByteArray &)
func (this *QSslConfiguration) SetSessionTicket(sessionTicket qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if sessionTicket != nil && sessionTicket.QByteArray_PTR() != nil {
		convArg0 = sessionTicket.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration16setSessionTicketERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] int sessionTicketLifeTimeHint() const
func (this *QSslConfiguration) SessionTicketLifeTimeHint() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration25sessionTicketLifeTimeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:139
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslKey ephemeralServerKey() const
func (this *QSslConfiguration) EphemeralServerKey() *QSslKey /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration18ephemeralServerKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslKey)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:146
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray preSharedKeyIdentityHint() const
func (this *QSslConfiguration) PreSharedKeyIdentityHint() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration24preSharedKeyIdentityHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreSharedKeyIdentityHint(const QByteArray &)
func (this *QSslConfiguration) SetPreSharedKeyIdentityHint(hint qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if hint != nil && hint.QByteArray_PTR() != nil {
		convArg0 = hint.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration27setPreSharedKeyIdentityHintERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:149
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslDiffieHellmanParameters diffieHellmanParameters() const
func (this *QSslConfiguration) DiffieHellmanParameters() *QSslDiffieHellmanParameters /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration23diffieHellmanParametersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslDiffieHellmanParametersFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslDiffieHellmanParameters)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDiffieHellmanParameters(const QSslDiffieHellmanParameters &)
func (this *QSslConfiguration) SetDiffieHellmanParameters(dhparams QSslDiffieHellmanParameters_ITF) {
	var convArg0 unsafe.Pointer
	if dhparams != nil && dhparams.QSslDiffieHellmanParameters_PTR() != nil {
		convArg0 = dhparams.QSslDiffieHellmanParameters_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration26setDiffieHellmanParametersERK27QSslDiffieHellmanParameters", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:152
// index:0
// Public static Visibility=Default Availability=Available
// [8] QSslConfiguration defaultConfiguration()
func (this *QSslConfiguration) DefaultConfiguration() *QSslConfiguration /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration20defaultConfigurationEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslConfiguration)
	return rv2
}
func QSslConfiguration_DefaultConfiguration() *QSslConfiguration /*123*/ {
	var nilthis *QSslConfiguration
	rv := nilthis.DefaultConfiguration()
	return rv
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:153
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setDefaultConfiguration(const QSslConfiguration &)
func (this *QSslConfiguration) SetDefaultConfiguration(configuration QSslConfiguration_ITF) {
	var convArg0 unsafe.Pointer
	if configuration != nil && configuration.QSslConfiguration_PTR() != nil {
		convArg0 = configuration.QSslConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration23setDefaultConfigurationERKS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QSslConfiguration_SetDefaultConfiguration(configuration QSslConfiguration_ITF) {
	var nilthis *QSslConfiguration
	nilthis.SetDefaultConfiguration(configuration)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:168
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray nextNegotiatedProtocol() const
func (this *QSslConfiguration) NextNegotiatedProtocol() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration22nextNegotiatedProtocolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:169
// index:0
// Public Visibility=Default Availability=Available
// [4] QSslConfiguration::NextProtocolNegotiationStatus nextProtocolNegotiationStatus() const
func (this *QSslConfiguration) NextProtocolNegotiationStatus() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration29nextProtocolNegotiationStatusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

type QSslConfiguration__NextProtocolNegotiationStatus = int

const QSslConfiguration__NextProtocolNegotiationNone QSslConfiguration__NextProtocolNegotiationStatus = 0
const QSslConfiguration__NextProtocolNegotiationNegotiated QSslConfiguration__NextProtocolNegotiationStatus = 1
const QSslConfiguration__NextProtocolNegotiationUnsupported QSslConfiguration__NextProtocolNegotiationStatus = 2

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
}

//  keep block end
