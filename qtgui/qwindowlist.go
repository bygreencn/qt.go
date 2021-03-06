package qtgui

// /usr/include/qt/QtGui/qwindow.h
// #include <qwindow.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

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

//  body block begin
type QWindowList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QWindowList) Operator_equal_0() *QWindowList {
	// QWindowList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QWindowList) Operator_equal_1() *QWindowList {
	// QWindowList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QWindowList) Swap_0() {
	// QWindowList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QWindowList) Operator_equal_equal_0() bool {
	// QWindowList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QWindowList) Operator_not_equal_0() bool {
	// QWindowList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QWindowList) Size_0() int {
	// QWindowList_size_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QWindowList) Detach_0() {
	// QWindowList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QWindowList) DetachShared_0() {
	// QWindowList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QWindowList) IsDetached_0() bool {
	// QWindowList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(_Bool)
func (this *QWindowList) SetSharable_0() {
	// QWindowList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QWindowList) IsSharedWith_0() bool {
	// QWindowList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QWindowList) IsEmpty_0() bool {
	// QWindowList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QWindowList) Clear_0() {
	// QWindowList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QWindowList) At_0() *QWindow {
	// QWindowList_at_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// const T & operator[](int)
func (this *QWindowList) Operator_get_index_0() *QWindow {
	// QWindowList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// T & operator[](int)
func (this *QWindowList) Operator_get_index_1() *QWindow {
	// QWindowList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// void reserve(int)
func (this *QWindowList) Reserve_0() {
	// QWindowList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QWindowList) Append_0() {
	// QWindowList_append_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QWindowList) Append_1() {
	// QWindowList_append_1()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QWindowList) Prepend_0() {
	// QWindowList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QWindowList) Insert_0() {
	// QWindowList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QWindowList) Replace_0() {
	// QWindowList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QWindowList) RemoveAt_0() {
	// QWindowList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QWindowList) RemoveAll_0() int {
	// QWindowList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QWindowList) RemoveOne_0() bool {
	// QWindowList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QWindowList) TakeAt_0() *QWindow {
	// QWindowList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// T takeFirst()
func (this *QWindowList) TakeFirst_0() *QWindow {
	// QWindowList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// T takeLast()
func (this *QWindowList) TakeLast_0() *QWindow {
	// QWindowList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// void move(int, int)
func (this *QWindowList) Move_0() {
	// QWindowList_move_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QWindowList) Swap_1() {
	// QWindowList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QWindowList) IndexOf_0() int {
	// QWindowList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QWindowList) LastIndexOf_0() int {
	// QWindowList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QWindowList) Contains_0() bool {
	// QWindowList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QWindowList) Count_0() int {
	// QWindowList_count_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QWindowList) Begin_0() {
	// QWindowList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QWindowList) Begin_1() {
	// QWindowList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QWindowList) Cbegin_0() {
	// QWindowList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QWindowList) ConstBegin_0() {
	// QWindowList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QWindowList) End_0() {
	// QWindowList_end_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QWindowList) End_1() {
	// QWindowList_end_1()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QWindowList) Cend_0() {
	// QWindowList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QWindowList) ConstEnd_0() {
	// QWindowList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QWindowList) Rbegin_0() {
	// QWindowList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QWindowList) Rend_0() {
	// QWindowList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QWindowList) Rbegin_1() {
	// QWindowList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QWindowList) Rend_1() {
	// QWindowList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QWindowList) Crbegin_0() {
	// QWindowList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QWindowList) Crend_0() {
	// QWindowList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(class QList::iterator, const T &)
func (this *QWindowList) Insert_1() {
	// QWindowList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(class QList::iterator)
func (this *QWindowList) Erase_0() {
	// QWindowList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(class QList::iterator, class QList::iterator)
func (this *QWindowList) Erase_1() {
	// QWindowList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QWindowList) Count_1() int {
	// QWindowList_count_1()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QWindowList) Length_0() int {
	// QWindowList_length_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QWindowList) First_0() *QWindow {
	// QWindowList_first_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// const T & constFirst()
func (this *QWindowList) ConstFirst_0() *QWindow {
	// QWindowList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// const T & first()
func (this *QWindowList) First_1() *QWindow {
	// QWindowList_first_1()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// T & last()
func (this *QWindowList) Last_0() *QWindow {
	// QWindowList_last_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// const T & last()
func (this *QWindowList) Last_1() *QWindow {
	// QWindowList_last_1()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// const T & constLast()
func (this *QWindowList) ConstLast_0() *QWindow {
	// QWindowList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// void removeFirst()
func (this *QWindowList) RemoveFirst_0() {
	// QWindowList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QWindowList) RemoveLast_0() {
	// QWindowList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QWindowList) StartsWith_0() bool {
	// QWindowList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QWindowList) EndsWith_0() bool {
	// QWindowList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QWindowList) Mid_0() *QWindowList {
	// QWindowList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QWindowList) Value_0() *QWindow {
	// QWindowList_value_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// T value(int, const T &)
func (this *QWindowList) Value_1() *QWindow {
	// QWindowList_value_1()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// void push_back(const T &)
func (this *QWindowList) Push_back_0() {
	// QWindowList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QWindowList) Push_front_0() {
	// QWindowList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QWindowList) Front_0() *QWindow {
	// QWindowList_front_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// const T & front()
func (this *QWindowList) Front_1() *QWindow {
	// QWindowList_front_1()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// T & back()
func (this *QWindowList) Back_0() *QWindow {
	// QWindowList_back_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// const T & back()
func (this *QWindowList) Back_1() *QWindow {
	// QWindowList_back_1()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// void pop_front()
func (this *QWindowList) Pop_front_0() {
	// QWindowList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QWindowList) Pop_back_0() {
	// QWindowList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QWindowList) Empty_0() bool {
	// QWindowList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QWindowList) Operator_add_equal_0() *QWindowList {
	// QWindowList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QWindowList) Operator_add_0() *QWindowList {
	// QWindowList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QWindowList) Operator_add_equal_1() *QWindowList {
	// QWindowList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QWindowList) Operator_left_shift_0() *QWindowList {
	// QWindowList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QWindowList) Operator_left_shift_1() *QWindowList {
	// QWindowList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QWindowList) ToVector_0() {
	// QWindowList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QWindowList) ToSet_0() {
	// QWindowList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QWindowList) FromVector_0() *QWindowList {
	// QWindowList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QWindowList) FromSet_0() *QWindowList {
	// QWindowList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QWindowList) FromStdList_0() *QWindowList {
	// QWindowList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QWindowList) ToStdList_0() {
	// QWindowList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QWindowList) Detach_helper_grow_0() {
	// QWindowList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QWindowList) Detach_helper_0() {
	// QWindowList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QWindowList) Detach_helper_1() {
	// QWindowList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(struct QListData::Data *)
func (this *QWindowList) Dealloc_0() {
	// QWindowList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(struct QList::Node *, const T &)
func (this *QWindowList) Node_construct_0() {
	// QWindowList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(struct QList::Node *)
func (this *QWindowList) Node_destruct_0() {
	// QWindowList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(struct QList::Node *, struct QList::Node *, struct QList::Node *)
func (this *QWindowList) Node_copy_0() {
	// QWindowList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(struct QList::Node *, struct QList::Node *)
func (this *QWindowList) Node_destruct_1() {
	// QWindowList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const class QList::iterator &)
func (this *QWindowList) IsValidIterator_0() bool {
	// QWindowList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, struct QListData::NotArrayCompatibleLayout)
func (this *QWindowList) Op_eq_impl_0() bool {
	// QWindowList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, struct QListData::ArrayCompatibleLayout)
func (this *QWindowList) Op_eq_impl_1() bool {
	// QWindowList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, struct QListData::NotArrayCompatibleLayout)
func (this *QWindowList) Contains_impl_0() bool {
	// QWindowList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, struct QListData::ArrayCompatibleLayout)
func (this *QWindowList) Contains_impl_1() bool {
	// QWindowList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, struct QListData::NotArrayCompatibleLayout)
func (this *QWindowList) Count_impl_0() int {
	// QWindowList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, struct QListData::ArrayCompatibleLayout)
func (this *QWindowList) Count_impl_1() int {
	// QWindowList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QWindowList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
