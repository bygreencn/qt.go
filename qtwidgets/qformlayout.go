package qtwidgets

// /usr/include/qt/QtWidgets/qformlayout.h
// #include <qformlayout.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

type QFormLayout struct {
	*QLayout
}
type QFormLayout_ITF interface {
	QLayout_ITF
	QFormLayout_PTR() *QFormLayout
}

func (ptr *QFormLayout) QFormLayout_PTR() *QFormLayout { return ptr }

func (this *QFormLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QLayout.GetCthis()
	}
}
func (this *QFormLayout) SetCthis(cthis unsafe.Pointer) {
	this.QLayout = NewQLayoutFromPointer(cthis)
}
func NewQFormLayoutFromPointer(cthis unsafe.Pointer) *QFormLayout {
	bcthis0 := NewQLayoutFromPointer(cthis)
	return &QFormLayout{bcthis0}
}
func (*QFormLayout) NewFromPointer(cthis unsafe.Pointer) *QFormLayout {
	return NewQFormLayoutFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qformlayout.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QFormLayout) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qformlayout.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFormLayout(QWidget *)
func NewQFormLayout(parent QWidget_ITF /*777 QWidget **/) *QFormLayout {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayoutC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFormLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFormLayout")
	return gothis
}

// /usr/include/qt/QtWidgets/qformlayout.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFormLayout(QWidget *)
func NewQFormLayout__() *QFormLayout {
	// arg: 0, QWidget *=Pointer, QWidget=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayoutC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFormLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFormLayout")
	return gothis
}

// /usr/include/qt/QtWidgets/qformlayout.h:92
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFormLayout()
func DeleteQFormLayout(this *QFormLayout) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayoutD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qformlayout.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFieldGrowthPolicy(enum QFormLayout::FieldGrowthPolicy)
func (this *QFormLayout) SetFieldGrowthPolicy(policy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout20setFieldGrowthPolicyENS_17FieldGrowthPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:95
// index:0
// Public Visibility=Default Availability=Available
// [4] QFormLayout::FieldGrowthPolicy fieldGrowthPolicy() const
func (this *QFormLayout) FieldGrowthPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout17fieldGrowthPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowWrapPolicy(enum QFormLayout::RowWrapPolicy)
func (this *QFormLayout) SetRowWrapPolicy(policy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout16setRowWrapPolicyENS_13RowWrapPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:97
// index:0
// Public Visibility=Default Availability=Available
// [4] QFormLayout::RowWrapPolicy rowWrapPolicy() const
func (this *QFormLayout) RowWrapPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout13rowWrapPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLabelAlignment(Qt::Alignment)
func (this *QFormLayout) SetLabelAlignment(alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout17setLabelAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment labelAlignment() const
func (this *QFormLayout) LabelAlignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout14labelAlignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFormAlignment(Qt::Alignment)
func (this *QFormLayout) SetFormAlignment(alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout16setFormAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment formAlignment() const
func (this *QFormLayout) FormAlignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout13formAlignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalSpacing(int)
func (this *QFormLayout) SetHorizontalSpacing(spacing int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout20setHorizontalSpacingEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:104
// index:0
// Public Visibility=Default Availability=Available
// [4] int horizontalSpacing() const
func (this *QFormLayout) HorizontalSpacing() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout17horizontalSpacingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qformlayout.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalSpacing(int)
func (this *QFormLayout) SetVerticalSpacing(spacing int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout18setVerticalSpacingEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:106
// index:0
// Public Visibility=Default Availability=Available
// [4] int verticalSpacing() const
func (this *QFormLayout) VerticalSpacing() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout15verticalSpacingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qformlayout.h:108
// index:0
// Public Visibility=Default Availability=Available
// [4] int spacing() const
func (this *QFormLayout) Spacing() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout7spacingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qformlayout.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpacing(int)
func (this *QFormLayout) SetSpacing(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout10setSpacingEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addRow(QWidget *, QWidget *)
func (this *QFormLayout) AddRow(label QWidget_ITF /*777 QWidget **/, field QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if label != nil && label.QWidget_PTR() != nil {
		convArg0 = label.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if field != nil && field.QWidget_PTR() != nil {
		convArg1 = field.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout6addRowEP7QWidgetS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:112
// index:1
// Public Visibility=Default Availability=Available
// [-2] void addRow(QWidget *, QLayout *)
func (this *QFormLayout) AddRow_1(label QWidget_ITF /*777 QWidget **/, field QLayout_ITF /*777 QLayout **/) {
	var convArg0 unsafe.Pointer
	if label != nil && label.QWidget_PTR() != nil {
		convArg0 = label.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if field != nil && field.QLayout_PTR() != nil {
		convArg1 = field.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout6addRowEP7QWidgetP7QLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:113
// index:2
// Public Visibility=Default Availability=Available
// [-2] void addRow(const QString &, QWidget *)
func (this *QFormLayout) AddRow_2(labelText string, field QWidget_ITF /*777 QWidget **/) {
	var tmpArg0 = qtcore.NewQString_5(labelText)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if field != nil && field.QWidget_PTR() != nil {
		convArg1 = field.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout6addRowERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:114
// index:3
// Public Visibility=Default Availability=Available
// [-2] void addRow(const QString &, QLayout *)
func (this *QFormLayout) AddRow_3(labelText string, field QLayout_ITF /*777 QLayout **/) {
	var tmpArg0 = qtcore.NewQString_5(labelText)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if field != nil && field.QLayout_PTR() != nil {
		convArg1 = field.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout6addRowERK7QStringP7QLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:115
// index:4
// Public Visibility=Default Availability=Available
// [-2] void addRow(QWidget *)
func (this *QFormLayout) AddRow_4(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout6addRowEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:116
// index:5
// Public Visibility=Default Availability=Available
// [-2] void addRow(QLayout *)
func (this *QFormLayout) AddRow_5(layout QLayout_ITF /*777 QLayout **/) {
	var convArg0 unsafe.Pointer
	if layout != nil && layout.QLayout_PTR() != nil {
		convArg0 = layout.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout6addRowEP7QLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertRow(int, QWidget *, QWidget *)
func (this *QFormLayout) InsertRow(row int, label QWidget_ITF /*777 QWidget **/, field QWidget_ITF /*777 QWidget **/) {
	var convArg1 unsafe.Pointer
	if label != nil && label.QWidget_PTR() != nil {
		convArg1 = label.QWidget_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if field != nil && field.QWidget_PTR() != nil {
		convArg2 = field.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiP7QWidgetS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:119
// index:1
// Public Visibility=Default Availability=Available
// [-2] void insertRow(int, QWidget *, QLayout *)
func (this *QFormLayout) InsertRow_1(row int, label QWidget_ITF /*777 QWidget **/, field QLayout_ITF /*777 QLayout **/) {
	var convArg1 unsafe.Pointer
	if label != nil && label.QWidget_PTR() != nil {
		convArg1 = label.QWidget_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if field != nil && field.QLayout_PTR() != nil {
		convArg2 = field.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiP7QWidgetP7QLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:120
// index:2
// Public Visibility=Default Availability=Available
// [-2] void insertRow(int, const QString &, QWidget *)
func (this *QFormLayout) InsertRow_2(row int, labelText string, field QWidget_ITF /*777 QWidget **/) {
	var tmpArg1 = qtcore.NewQString_5(labelText)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if field != nil && field.QWidget_PTR() != nil {
		convArg2 = field.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiRK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:121
// index:3
// Public Visibility=Default Availability=Available
// [-2] void insertRow(int, const QString &, QLayout *)
func (this *QFormLayout) InsertRow_3(row int, labelText string, field QLayout_ITF /*777 QLayout **/) {
	var tmpArg1 = qtcore.NewQString_5(labelText)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if field != nil && field.QLayout_PTR() != nil {
		convArg2 = field.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiRK7QStringP7QLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:122
// index:4
// Public Visibility=Default Availability=Available
// [-2] void insertRow(int, QWidget *)
func (this *QFormLayout) InsertRow_4(row int, widget QWidget_ITF /*777 QWidget **/) {
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:123
// index:5
// Public Visibility=Default Availability=Available
// [-2] void insertRow(int, QLayout *)
func (this *QFormLayout) InsertRow_5(row int, layout QLayout_ITF /*777 QLayout **/) {
	var convArg1 unsafe.Pointer
	if layout != nil && layout.QLayout_PTR() != nil {
		convArg1 = layout.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiP7QLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeRow(int)
func (this *QFormLayout) RemoveRow(row int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout9removeRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:126
// index:1
// Public Visibility=Default Availability=Available
// [-2] void removeRow(QWidget *)
func (this *QFormLayout) RemoveRow_1(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout9removeRowEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:127
// index:2
// Public Visibility=Default Availability=Available
// [-2] void removeRow(QLayout *)
func (this *QFormLayout) RemoveRow_2(layout QLayout_ITF /*777 QLayout **/) {
	var convArg0 unsafe.Pointer
	if layout != nil && layout.QLayout_PTR() != nil {
		convArg0 = layout.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout9removeRowEP7QLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:129
// index:0
// Public Visibility=Default Availability=Available
// [16] QFormLayout::TakeRowResult takeRow(int)
func (this *QFormLayout) TakeRow(row int) unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout7takeRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtWidgets/qformlayout.h:130
// index:1
// Public Visibility=Default Availability=Available
// [16] QFormLayout::TakeRowResult takeRow(QWidget *)
func (this *QFormLayout) TakeRow_1(widget QWidget_ITF /*777 QWidget **/) unsafe.Pointer /*444*/ {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout7takeRowEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtWidgets/qformlayout.h:131
// index:2
// Public Visibility=Default Availability=Available
// [16] QFormLayout::TakeRowResult takeRow(QLayout *)
func (this *QFormLayout) TakeRow_2(layout QLayout_ITF /*777 QLayout **/) unsafe.Pointer /*444*/ {
	var convArg0 unsafe.Pointer
	if layout != nil && layout.QLayout_PTR() != nil {
		convArg0 = layout.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout7takeRowEP7QLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtWidgets/qformlayout.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItem(int, enum QFormLayout::ItemRole, QLayoutItem *)
func (this *QFormLayout) SetItem(row int, role int, item QLayoutItem_ITF /*777 QLayoutItem **/) {
	var convArg2 unsafe.Pointer
	if item != nil && item.QLayoutItem_PTR() != nil {
		convArg2 = item.QLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout7setItemEiNS_8ItemRoleEP11QLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, role, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidget(int, enum QFormLayout::ItemRole, QWidget *)
func (this *QFormLayout) SetWidget(row int, role int, widget QWidget_ITF /*777 QWidget **/) {
	var convArg2 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg2 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout9setWidgetEiNS_8ItemRoleEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, role, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLayout(int, enum QFormLayout::ItemRole, QLayout *)
func (this *QFormLayout) SetLayout(row int, role int, layout QLayout_ITF /*777 QLayout **/) {
	var convArg2 unsafe.Pointer
	if layout != nil && layout.QLayout_PTR() != nil {
		convArg2 = layout.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout9setLayoutEiNS_8ItemRoleEP7QLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, role, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:137
// index:0
// Public Visibility=Default Availability=Available
// [8] QLayoutItem * itemAt(int, enum QFormLayout::ItemRole) const
func (this *QFormLayout) ItemAt(row int, role int) *QLayoutItem /*777 QLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout6itemAtEiNS_8ItemRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, role)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qformlayout.h:146
// index:1
// Public virtual Visibility=Default Availability=Available
// [8] QLayoutItem * itemAt(int) const
func (this *QFormLayout) ItemAt_1(index int) *QLayoutItem /*777 QLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout6itemAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qformlayout.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getItemPosition(int, int *, enum QFormLayout::ItemRole *) const
func (this *QFormLayout) GetItemPosition(index int, rowPtr unsafe.Pointer /*666*/, rolePtr unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout15getItemPositionEiPiPNS_8ItemRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, rowPtr, rolePtr)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getWidgetPosition(QWidget *, int *, enum QFormLayout::ItemRole *) const
func (this *QFormLayout) GetWidgetPosition(widget QWidget_ITF /*777 QWidget **/, rowPtr unsafe.Pointer /*666*/, rolePtr unsafe.Pointer /*666*/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout17getWidgetPositionEP7QWidgetPiPNS_8ItemRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, rowPtr, rolePtr)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getLayoutPosition(QLayout *, int *, enum QFormLayout::ItemRole *) const
func (this *QFormLayout) GetLayoutPosition(layout QLayout_ITF /*777 QLayout **/, rowPtr unsafe.Pointer /*666*/, rolePtr unsafe.Pointer /*666*/) {
	var convArg0 unsafe.Pointer
	if layout != nil && layout.QLayout_PTR() != nil {
		convArg0 = layout.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout17getLayoutPositionEP7QLayoutPiPNS_8ItemRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, rowPtr, rolePtr)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:141
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * labelForField(QWidget *) const
func (this *QFormLayout) LabelForField(field QWidget_ITF /*777 QWidget **/) *QWidget /*777 QWidget **/ {
	var convArg0 unsafe.Pointer
	if field != nil && field.QWidget_PTR() != nil {
		convArg0 = field.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout13labelForFieldEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qformlayout.h:142
// index:1
// Public Visibility=Default Availability=Available
// [8] QWidget * labelForField(QLayout *) const
func (this *QFormLayout) LabelForField_1(field QLayout_ITF /*777 QLayout **/) *QWidget /*777 QWidget **/ {
	var convArg0 unsafe.Pointer
	if field != nil && field.QLayout_PTR() != nil {
		convArg0 = field.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout13labelForFieldEP7QLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qformlayout.h:145
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void addItem(QLayoutItem *)
func (this *QFormLayout) AddItem(item QLayoutItem_ITF /*777 QLayoutItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QLayoutItem_PTR() != nil {
		convArg0 = item.QLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout7addItemEP11QLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:147
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QLayoutItem * takeAt(int)
func (this *QFormLayout) TakeAt(index int) *QLayoutItem /*777 QLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout6takeAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qformlayout.h:149
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setGeometry(const QRect &)
func (this *QFormLayout) SetGeometry(rect qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout11setGeometryERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:150
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSize() const
func (this *QFormLayout) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout11minimumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qformlayout.h:151
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const
func (this *QFormLayout) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qformlayout.h:152
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void invalidate()
func (this *QFormLayout) Invalidate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout10invalidateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:154
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasHeightForWidth() const
func (this *QFormLayout) HasHeightForWidth() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout17hasHeightForWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qformlayout.h:155
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int heightForWidth(int) const
func (this *QFormLayout) HeightForWidth(width int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout14heightForWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qformlayout.h:156
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::Orientations expandingDirections() const
func (this *QFormLayout) ExpandingDirections() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout19expandingDirectionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:157
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int count() const
func (this *QFormLayout) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qformlayout.h:159
// index:0
// Public Visibility=Default Availability=Available
// [4] int rowCount() const
func (this *QFormLayout) RowCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout8rowCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

type QFormLayout__FieldGrowthPolicy = int

const QFormLayout__FieldsStayAtSizeHint QFormLayout__FieldGrowthPolicy = 0
const QFormLayout__ExpandingFieldsGrow QFormLayout__FieldGrowthPolicy = 1
const QFormLayout__AllNonFixedFieldsGrow QFormLayout__FieldGrowthPolicy = 2

type QFormLayout__RowWrapPolicy = int

const QFormLayout__DontWrapRows QFormLayout__RowWrapPolicy = 0
const QFormLayout__WrapLongRows QFormLayout__RowWrapPolicy = 1
const QFormLayout__WrapAllRows QFormLayout__RowWrapPolicy = 2

type QFormLayout__ItemRole = int

const QFormLayout__LabelRole QFormLayout__ItemRole = 0
const QFormLayout__FieldRole QFormLayout__ItemRole = 1
const QFormLayout__SpanningRole QFormLayout__ItemRole = 2

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
		qtgui.KeepMe()
	}
}

//  keep block end
