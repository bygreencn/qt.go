package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtCore/qcommandlineparser.h
// dst-file: /src/core/qcommandlineparser.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "runtime"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QCommandLineParser::addOption(const QCommandLineOption & commandLineOption);
extern bool C_ZN18QCommandLineParser9addOptionERK18QCommandLineOption(void* qthis, void* arg0); // 4
  // proto:  void QCommandLineParser::process(const QStringList & arguments);
extern void C_ZN18QCommandLineParser7processERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  void QCommandLineParser::process(const QCoreApplication & app);
extern void C_ZN18QCommandLineParser7processERK16QCoreApplication(void* qthis, void* arg0); // 4
  // proto:  QString QCommandLineParser::errorText();
extern void* C_ZNK18QCommandLineParser9errorTextEv(void* qthis); // 4
  // proto:  void QCommandLineParser::addPositionalArgument(const QString & name, const QString & description, const QString & syntax);
extern void C_ZN18QCommandLineParser21addPositionalArgumentERK7QStringS2_S2_(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  bool QCommandLineParser::parse(const QStringList & arguments);
extern bool C_ZN18QCommandLineParser5parseERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  void QCommandLineParser::~QCommandLineParser();
extern void C_ZN18QCommandLineParserD2Ev(void* qthis); // 4
  // proto:  void QCommandLineParser::showHelp(int exitCode);
extern void C_ZN18QCommandLineParser8showHelpEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QCommandLineParser::applicationDescription();
extern void* C_ZNK18QCommandLineParser22applicationDescriptionEv(void* qthis); // 4
  // proto:  QCommandLineOption QCommandLineParser::addVersionOption();
extern void* C_ZN18QCommandLineParser16addVersionOptionEv(void* qthis); // 4
  // proto:  QString QCommandLineParser::helpText();
extern void* C_ZNK18QCommandLineParser8helpTextEv(void* qthis); // 4
  // proto:  void QCommandLineParser::showVersion();
extern void C_ZN18QCommandLineParser11showVersionEv(void* qthis); // 4
  // proto:  void QCommandLineParser::QCommandLineParser();
extern void* C_ZN18QCommandLineParserC2Ev(); // 3
  // proto:  void QCommandLineParser::setApplicationDescription(const QString & description);
extern void C_ZN18QCommandLineParser25setApplicationDescriptionERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QCommandLineParser::clearPositionalArguments();
extern void C_ZN18QCommandLineParser24clearPositionalArgumentsEv(void* qthis); // 4
  // proto:  QCommandLineOption QCommandLineParser::addHelpOption();
extern void* C_ZN18QCommandLineParser13addHelpOptionEv(void* qthis); // 4
  // proto:  QStringList QCommandLineParser::optionNames();
extern void C_ZNK18QCommandLineParser11optionNamesEv(void* qthis); // 4
  // proto:  QStringList QCommandLineParser::unknownOptionNames();
extern void C_ZNK18QCommandLineParser18unknownOptionNamesEv(void* qthis); // 4
  // proto:  QString QCommandLineParser::value(const QString & name);
extern void* C_ZNK18QCommandLineParser5valueERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QCommandLineParser::value(const QCommandLineOption & option);
extern void* C_ZNK18QCommandLineParser5valueERK18QCommandLineOption(void* qthis, void* arg0); // 4
  // proto:  bool QCommandLineParser::isSet(const QString & name);
extern bool C_ZNK18QCommandLineParser5isSetERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QCommandLineParser::isSet(const QCommandLineOption & option);
extern bool C_ZNK18QCommandLineParser5isSetERK18QCommandLineOption(void* qthis, void* arg0); // 4
  // proto:  QStringList QCommandLineParser::values(const QCommandLineOption & option);
extern void C_ZNK18QCommandLineParser6valuesERK18QCommandLineOption(void* qthis, void* arg0); // 4
  // proto:  QStringList QCommandLineParser::values(const QString & name);
extern void C_ZNK18QCommandLineParser6valuesERK7QString(void* qthis, void* arg0); // 4
  // proto:  QStringList QCommandLineParser::positionalArguments();
extern void C_ZNK18QCommandLineParser19positionalArgumentsEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QCommandLineParser)=8
type QCommandLineParser struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// addOption(const class QCommandLineOption &)
func (this *QCommandLineParser) AddOption(args ...interface{}) (ret interface{}) {
  // addOption(const class QCommandLineOption &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCommandLineOption{}) // "const QCommandLineOption &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineParser9addOptionERK18QCommandLineOption
    // invoke: bool addOption(const class QCommandLineOption &)
    var arg0 = args[0].(*QCommandLineOption).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN18QCommandLineParser9addOptionERK18QCommandLineOption(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCommandLineParser", "addOption", args)
  }

  return
}

// process(const class QStringList &)
func (this *QCommandLineParser) Process(args ...interface{}) () {
  // process(const class QStringList &)
  // process(const class QCoreApplication &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QCoreApplication{}) // "const QCoreApplication &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineParser7processERK11QStringList
    // invoke: void process(const class QStringList &)
    var arg0 = args[0].(*QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QCommandLineParser7processERK11QStringList(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN18QCommandLineParser7processERK16QCoreApplication
    // invoke: void process(const class QCoreApplication &)
    var arg0 = args[0].(*QCoreApplication).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QCommandLineParser7processERK16QCoreApplication(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "process", args)
  }

  return
}

// errorText()
func (this *QCommandLineParser) ErrorText(args ...interface{}) (ret interface{}) {
  // errorText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineParser9errorTextEv
    // invoke: QString errorText()
    var ret0 = C.C_ZNK18QCommandLineParser9errorTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCommandLineParser", "errorText", args)
  }

  return
}

// addPositionalArgument(const class QString &, const class QString &, const class QString &)
func (this *QCommandLineParser) AddPositionalArgument(args ...interface{}) () {
  // addPositionalArgument(const class QString &, const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineParser21addPositionalArgumentERK7QStringS2_S2_
    // invoke: void addPositionalArgument(const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QString).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN18QCommandLineParser21addPositionalArgumentERK7QStringS2_S2_(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "addPositionalArgument", args)
  }

  return
}

// parse(const class QStringList &)
func (this *QCommandLineParser) Parse(args ...interface{}) (ret interface{}) {
  // parse(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineParser5parseERK11QStringList
    // invoke: bool parse(const class QStringList &)
    var arg0 = args[0].(*QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN18QCommandLineParser5parseERK11QStringList(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCommandLineParser", "parse", args)
  }

  return
}

// ~QCommandLineParser()
func (this *QCommandLineParser) Free(args ...interface{}) () {
  // ~QCommandLineParser()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineParserD0Ev
    // invoke: void ~QCommandLineParser()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN18QCommandLineParserD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QCommandLineParser", "~QCommandLineParser", args)
  }

  return
}

// showHelp(int)
func (this *QCommandLineParser) ShowHelp(args ...interface{}) () {
  // showHelp(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineParser8showHelpEi
    // invoke: void showHelp(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN18QCommandLineParser8showHelpEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "showHelp", args)
  }

  return
}

// applicationDescription()
func (this *QCommandLineParser) ApplicationDescription(args ...interface{}) (ret interface{}) {
  // applicationDescription()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineParser22applicationDescriptionEv
    // invoke: QString applicationDescription()
    var ret0 = C.C_ZNK18QCommandLineParser22applicationDescriptionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCommandLineParser", "applicationDescription", args)
  }

  return
}

// addVersionOption()
func (this *QCommandLineParser) AddVersionOption(args ...interface{}) (ret interface{}) {
  // addVersionOption()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineParser16addVersionOptionEv
    // invoke: QCommandLineOption addVersionOption()
    var ret0 = C.C_ZN18QCommandLineParser16addVersionOptionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QCommandLineOption{}) // "QCommandLineOption"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCommandLineParser", "addVersionOption", args)
  }

  return
}

// helpText()
func (this *QCommandLineParser) HelpText(args ...interface{}) (ret interface{}) {
  // helpText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineParser8helpTextEv
    // invoke: QString helpText()
    var ret0 = C.C_ZNK18QCommandLineParser8helpTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCommandLineParser", "helpText", args)
  }

  return
}

// showVersion()
func (this *QCommandLineParser) ShowVersion(args ...interface{}) () {
  // showVersion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineParser11showVersionEv
    // invoke: void showVersion()
    C.C_ZN18QCommandLineParser11showVersionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "showVersion", args)
  }

  return
}

// QCommandLineParser()
func GcfreeQCommandLineParser(this *QCommandLineParser) {
  qtrt.UniverseFree(this)
}
func NewQCommandLineParser(args ...interface{}) *QCommandLineParser {
  // QCommandLineParser()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineParserC1Ev
    // invoke: void QCommandLineParser()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QCommandLineParserC2Ev()
    this := &QCommandLineParser{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQCommandLineParser)
    return this
  default:
    qtrt.ErrorResolve("QCommandLineParser", "QCommandLineParser", args)
  }

  return nil // QCommandLineParser{Qclsinst:qthis}
}

// setApplicationDescription(const class QString &)
func (this *QCommandLineParser) SetApplicationDescription(args ...interface{}) () {
  // setApplicationDescription(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineParser25setApplicationDescriptionERK7QString
    // invoke: void setApplicationDescription(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QCommandLineParser25setApplicationDescriptionERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "setApplicationDescription", args)
  }

  return
}

// clearPositionalArguments()
func (this *QCommandLineParser) ClearPositionalArguments(args ...interface{}) () {
  // clearPositionalArguments()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineParser24clearPositionalArgumentsEv
    // invoke: void clearPositionalArguments()
    C.C_ZN18QCommandLineParser24clearPositionalArgumentsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "clearPositionalArguments", args)
  }

  return
}

// addHelpOption()
func (this *QCommandLineParser) AddHelpOption(args ...interface{}) (ret interface{}) {
  // addHelpOption()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineParser13addHelpOptionEv
    // invoke: QCommandLineOption addHelpOption()
    var ret0 = C.C_ZN18QCommandLineParser13addHelpOptionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QCommandLineOption{}) // "QCommandLineOption"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCommandLineParser", "addHelpOption", args)
  }

  return
}

// optionNames()
func (this *QCommandLineParser) OptionNames(args ...interface{}) () {
  // optionNames()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineParser11optionNamesEv
    // invoke: QStringList optionNames()
    C.C_ZNK18QCommandLineParser11optionNamesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "optionNames", args)
  }

  return
}

// unknownOptionNames()
func (this *QCommandLineParser) UnknownOptionNames(args ...interface{}) () {
  // unknownOptionNames()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineParser18unknownOptionNamesEv
    // invoke: QStringList unknownOptionNames()
    C.C_ZNK18QCommandLineParser18unknownOptionNamesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "unknownOptionNames", args)
  }

  return
}

// value(const class QString &)
func (this *QCommandLineParser) Value(args ...interface{}) (ret interface{}) {
  // value(const class QString &)
  // value(const class QCommandLineOption &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QCommandLineOption{}) // "const QCommandLineOption &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineParser5valueERK7QString
    // invoke: QString value(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK18QCommandLineParser5valueERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK18QCommandLineParser5valueERK18QCommandLineOption
    // invoke: QString value(const class QCommandLineOption &)
    var arg0 = args[0].(*QCommandLineOption).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK18QCommandLineParser5valueERK18QCommandLineOption(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCommandLineParser", "value", args)
  }

  return
}

// isSet(const class QString &)
func (this *QCommandLineParser) IsSet(args ...interface{}) (ret interface{}) {
  // isSet(const class QString &)
  // isSet(const class QCommandLineOption &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QCommandLineOption{}) // "const QCommandLineOption &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineParser5isSetERK7QString
    // invoke: bool isSet(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK18QCommandLineParser5isSetERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK18QCommandLineParser5isSetERK18QCommandLineOption
    // invoke: bool isSet(const class QCommandLineOption &)
    var arg0 = args[0].(*QCommandLineOption).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK18QCommandLineParser5isSetERK18QCommandLineOption(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCommandLineParser", "isSet", args)
  }

  return
}

// values(const class QCommandLineOption &)
func (this *QCommandLineParser) Values(args ...interface{}) () {
  // values(const class QCommandLineOption &)
  // values(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCommandLineOption{}) // "const QCommandLineOption &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineParser6valuesERK18QCommandLineOption
    // invoke: QStringList values(const class QCommandLineOption &)
    var arg0 = args[0].(*QCommandLineOption).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK18QCommandLineParser6valuesERK18QCommandLineOption(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZNK18QCommandLineParser6valuesERK7QString
    // invoke: QStringList values(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK18QCommandLineParser6valuesERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "values", args)
  }

  return
}

// positionalArguments()
func (this *QCommandLineParser) PositionalArguments(args ...interface{}) () {
  // positionalArguments()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineParser19positionalArgumentsEv
    // invoke: QStringList positionalArguments()
    C.C_ZNK18QCommandLineParser19positionalArgumentsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "positionalArguments", args)
  }

  return
}

// <= body block end
