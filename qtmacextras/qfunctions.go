package qtmacextras

import "unsafe"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

func init() {
	if false {
		_ = unsafe.Pointer(uintptr(0))
	}
	if false {
		qtrt.KeepMe()
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

//  header block end

//  body block begin
// /usr/include/qt/QtMacExtras/../../src/macextras/qmacpasteboardmime.h:56
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qRegisterDraggedTypes(const QStringList &)
func QRegisterDraggedTypes(types qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if types != nil && types.QStringList_PTR() != nil {
		convArg0 = types.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z21qRegisterDraggedTypesRK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end
