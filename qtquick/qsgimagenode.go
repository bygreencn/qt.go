package qtquick

// /usr/include/qt/QtQuick/qsgimagenode.h
// #include <qsgimagenode.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

type QSGImageNode struct {
	*QSGGeometryNode
}
type QSGImageNode_ITF interface {
	QSGGeometryNode_ITF
	QSGImageNode_PTR() *QSGImageNode
}

func (ptr *QSGImageNode) QSGImageNode_PTR() *QSGImageNode { return ptr }

func (this *QSGImageNode) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QSGGeometryNode.GetCthis()
	}
}
func (this *QSGImageNode) SetCthis(cthis unsafe.Pointer) {
	this.QSGGeometryNode = NewQSGGeometryNodeFromPointer(cthis)
}
func NewQSGImageNodeFromPointer(cthis unsafe.Pointer) *QSGImageNode {
	bcthis0 := NewQSGGeometryNodeFromPointer(cthis)
	return &QSGImageNode{bcthis0}
}
func (*QSGImageNode) NewFromPointer(cthis unsafe.Pointer) *QSGImageNode {
	return NewQSGImageNodeFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:51
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [-2] void ~QSGImageNode()
func DeleteQSGImageNode(this *QSGImageNode) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSGImageNodeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 144)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:53
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setRect(const QRectF &)
func (this *QSGImageNode) SetRect(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSGImageNode7setRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:54
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setRect(qreal, qreal, qreal, qreal)
func (this *QSGImageNode) SetRect_1(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSGImageNode7setRectEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:55
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [32] QRectF rect() const
func (this *QSGImageNode) Rect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSGImageNode4rectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtQuick/qsgimagenode.h:57
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setSourceRect(const QRectF &)
func (this *QSGImageNode) SetSourceRect(r qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSGImageNode13setSourceRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:58
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setSourceRect(qreal, qreal, qreal, qreal)
func (this *QSGImageNode) SetSourceRect_1(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSGImageNode13setSourceRectEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:59
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [32] QRectF sourceRect() const
func (this *QSGImageNode) SourceRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSGImageNode10sourceRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtQuick/qsgimagenode.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setTexture(QSGTexture *)
func (this *QSGImageNode) SetTexture(texture QSGTexture_ITF /*777 QSGTexture **/) {
	var convArg0 unsafe.Pointer
	if texture != nil && texture.QSGTexture_PTR() != nil {
		convArg0 = texture.QSGTexture_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSGImageNode10setTextureEP10QSGTexture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:62
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSGTexture * texture() const
func (this *QSGImageNode) Texture() *QSGTexture /*777 QSGTexture **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSGImageNode7textureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgimagenode.h:64
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setFiltering(QSGTexture::Filtering)
func (this *QSGImageNode) SetFiltering(filtering int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSGImageNode12setFilteringEN10QSGTexture9FilteringE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filtering)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:65
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QSGTexture::Filtering filtering() const
func (this *QSGImageNode) Filtering() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSGImageNode9filteringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:67
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setMipmapFiltering(QSGTexture::Filtering)
func (this *QSGImageNode) SetMipmapFiltering(filtering int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSGImageNode18setMipmapFilteringEN10QSGTexture9FilteringE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filtering)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:68
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QSGTexture::Filtering mipmapFiltering() const
func (this *QSGImageNode) MipmapFiltering() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSGImageNode15mipmapFilteringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:79
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setTextureCoordinatesTransform(QSGImageNode::TextureCoordinatesTransformMode)
func (this *QSGImageNode) SetTextureCoordinatesTransform(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSGImageNode30setTextureCoordinatesTransformE6QFlagsINS_31TextureCoordinatesTransformFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:80
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QSGImageNode::TextureCoordinatesTransformMode textureCoordinatesTransform() const
func (this *QSGImageNode) TextureCoordinatesTransform() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSGImageNode27textureCoordinatesTransformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:82
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setOwnsTexture(_Bool)
func (this *QSGImageNode) SetOwnsTexture(owns bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSGImageNode14setOwnsTextureEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), owns)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:83
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool ownsTexture() const
func (this *QSGImageNode) OwnsTexture() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSGImageNode11ownsTextureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qsgimagenode.h:85
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void rebuildGeometry(QSGGeometry *, QSGTexture *, const QRectF &, QRectF, QSGImageNode::TextureCoordinatesTransformMode)
func (this *QSGImageNode) RebuildGeometry(g QSGGeometry_ITF /*777 QSGGeometry **/, texture QSGTexture_ITF /*777 QSGTexture **/, rect qtcore.QRectF_ITF, sourceRect qtcore.QRectF_ITF /*123*/, texCoordMode int) {
	var convArg0 unsafe.Pointer
	if g != nil && g.QSGGeometry_PTR() != nil {
		convArg0 = g.QSGGeometry_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if texture != nil && texture.QSGTexture_PTR() != nil {
		convArg1 = texture.QSGTexture_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg2 = rect.QRectF_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if sourceRect != nil && sourceRect.QRectF_PTR() != nil {
		convArg3 = sourceRect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSGImageNode15rebuildGeometryEP11QSGGeometryP10QSGTextureRK6QRectFS4_6QFlagsINS_31TextureCoordinatesTransformFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, texCoordMode)
	qtrt.ErrPrint(err, rv)
}
func QSGImageNode_RebuildGeometry(g QSGGeometry_ITF /*777 QSGGeometry **/, texture QSGTexture_ITF /*777 QSGTexture **/, rect qtcore.QRectF_ITF, sourceRect qtcore.QRectF_ITF /*123*/, texCoordMode int) {
	var nilthis *QSGImageNode
	nilthis.RebuildGeometry(g, texture, rect, sourceRect, texCoordMode)
}

type QSGImageNode__TextureCoordinatesTransformFlag = int

const QSGImageNode__NoTransform QSGImageNode__TextureCoordinatesTransformFlag = 0
const QSGImageNode__MirrorHorizontally QSGImageNode__TextureCoordinatesTransformFlag = 1
const QSGImageNode__MirrorVertically QSGImageNode__TextureCoordinatesTransformFlag = 2

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
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  keep block end
