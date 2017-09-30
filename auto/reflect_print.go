package auto

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

type printTypeX func(t reflect.Type, lv int)
type printValueX func(v reflect.Value, lv int)

var printTypeXes map[reflect.Kind]printTypeX
var printValueXes map[reflect.Kind]printValueX

func init() {
	printTypeXes = map[reflect.Kind]printTypeX{
		reflect.Invalid:       printTypeInvalid,
		reflect.Bool:          printTypeBool,
		reflect.Int:           printTypeInt,
		reflect.Int8:          printTypeInt8,
		reflect.Int16:         printTypeInt16,
		reflect.Int32:         printTypeInt32,
		reflect.Int64:         printTypeInt64,
		reflect.Uint:          printTypeUint,
		reflect.Uint8:         printTypeUint8,
		reflect.Uint16:        printTypeUint16,
		reflect.Uint32:        printTypeUint32,
		reflect.Uint64:        printTypeUint64,
		reflect.Uintptr:       printTypeUintptr,
		reflect.Float32:       printTypeFloat32,
		reflect.Float64:       printTypeFloat64,
		reflect.Complex64:     printTypeComplex64,
		reflect.Complex128:    printTypeComplex128,
		reflect.Array:         printTypeArray,
		reflect.Chan:          printTypeChan,
		reflect.Func:          printTypeFunc,
		reflect.Interface:     printTypeInterface,
		reflect.Map:           printTypeMap,
		reflect.Ptr:           printTypePtr,
		reflect.Slice:         printTypeSlice,
		reflect.String:        printTypeString,
		reflect.Struct:        printTypeStruct,
		reflect.UnsafePointer: printTypeUnsafePointer,
	}
	printValueXes = map[reflect.Kind]printValueX{
		reflect.Invalid:       printValueInvalid,
		reflect.Bool:          printValueBool,
		reflect.Int:           printValueInt,
		reflect.Int8:          printValueInt8,
		reflect.Int16:         printValueInt16,
		reflect.Int32:         printValueInt32,
		reflect.Int64:         printValueInt64,
		reflect.Uint:          printValueUint,
		reflect.Uint8:         printValueUint8,
		reflect.Uint16:        printValueUint16,
		reflect.Uint32:        printValueUint32,
		reflect.Uint64:        printValueUint64,
		reflect.Uintptr:       printValueUintptr,
		reflect.Float32:       printValueFloat32,
		reflect.Float64:       printValueFloat64,
		reflect.Complex64:     printValueComplex64,
		reflect.Complex128:    printValueComplex128,
		reflect.Array:         printValueArray,
		reflect.Chan:          printValueChan,
		reflect.Func:          printValueFunc,
		reflect.Interface:     printValueInterface,
		reflect.Map:           printValueMap,
		reflect.Ptr:           printValuePtr,
		reflect.Slice:         printValueSlice,
		reflect.String:        printValueString,
		reflect.Struct:        printValueStruct,
		reflect.UnsafePointer: printValueUnsafePointer,
	}
}

/*PrintType 自动打印Instance type info*/
func PrintType(instance interface{}) {
	shuntTypePrint(reflect.TypeOf(instance), 0)
}

func printType(it reflect.Type, lv int) {
	tab := string(bytes.Repeat([]byte("\t"), lv))

	infos := make([]string, 0, 64)
	infos = append(infos, fmt.Sprint("Type: ", it))
	infos = append(infos, fmt.Sprint("Type.Kind: ", it.Kind()))

	fmt.Println(tab + strings.Join(infos, ", "))
}

func shuntTypePrint(it reflect.Type, lv int) {
	printTypeXes[it.Kind()](it, lv)
}

func printTypeInvalid(it reflect.Type, lv int)    { printType(it, lv) }
func printTypeBool(it reflect.Type, lv int)       { printType(it, lv) }
func printTypeInt(it reflect.Type, lv int)        { printType(it, lv) }
func printTypeInt8(it reflect.Type, lv int)       { printTypeInt(it, lv) }
func printTypeInt16(it reflect.Type, lv int)      { printTypeInt(it, lv) }
func printTypeInt32(it reflect.Type, lv int)      { printTypeInt(it, lv) }
func printTypeInt64(it reflect.Type, lv int)      { printTypeInt(it, lv) }
func printTypeUint(it reflect.Type, lv int)       { printType(it, lv) }
func printTypeUint8(it reflect.Type, lv int)      { printTypeUint(it, lv) }
func printTypeUint16(it reflect.Type, lv int)     { printTypeUint(it, lv) }
func printTypeUint32(it reflect.Type, lv int)     { printTypeUint(it, lv) }
func printTypeUint64(it reflect.Type, lv int)     { printTypeUint(it, lv) }
func printTypeUintptr(it reflect.Type, lv int)    { printType(it, lv) }
func printTypeFloat32(it reflect.Type, lv int)    { printType(it, lv) }
func printTypeFloat64(it reflect.Type, lv int)    { printTypeFloat32(it, lv) }
func printTypeComplex64(it reflect.Type, lv int)  { printType(it, lv) }
func printTypeComplex128(it reflect.Type, lv int) { printTypeComplex64(it, lv) }
func printTypeArray(it reflect.Type, lv int)      { printType(it, lv); shuntTypePrint(it.Elem(), lv) }
func printTypeChan(it reflect.Type, lv int)       { printType(it, lv); shuntTypePrint(it.Elem(), lv) }
func printTypeFunc(it reflect.Type, lv int) {
	printType(it, lv)

	for i := 0; i < it.NumIn(); i++ {
		shuntTypePrint(it.In(i), lv+1)
	}
	for i := 0; i < it.NumOut(); i++ {
		shuntTypePrint(it.Out(i), lv+1)
	}
}
func printTypeInterface(it reflect.Type, lv int) { printType(it, lv) }
func printTypeMap(it reflect.Type, lv int) {
	printType(it, lv)

	shuntTypePrint(it.Key(), lv)
	shuntTypePrint(it.Elem(), lv)
}
func printTypePtr(it reflect.Type, lv int)    { printType(it, lv); shuntTypePrint(it.Elem(), lv) }
func printTypeSlice(it reflect.Type, lv int)  { printType(it, lv); shuntTypePrint(it.Elem(), lv) }
func printTypeString(it reflect.Type, lv int) { printType(it, lv) }
func printTypeStruct(it reflect.Type, lv int) {
	printType(it, lv)

	for i := 0; i < it.NumField(); i++ {
		shuntTypePrint(it.Field(i).Type, lv+1)
	}
}
func printTypeUnsafePointer(it reflect.Type, lv int) {}

/*PrintValue 自动打印Instance value info*/
func PrintValue(instance interface{}) {
	shuntValuePrint(reflect.ValueOf(instance), 0)
}

func printValue(iv reflect.Value, lv int) {
	tab := string(bytes.Repeat([]byte("\t"), lv))

	infos := make([]string, 0, 64)

	infos = append(infos, fmt.Sprint("Type.CanAddr: ", iv.CanAddr()))
	infos = append(infos, fmt.Sprint("Type.CanSet: ", iv.CanSet()))
	infos = append(infos, fmt.Sprint("Type.CanInterface: ", iv.CanInterface()))
	infos = append(infos, fmt.Sprint("Type: ", iv.String()))
	infos = append(infos, fmt.Sprint("Type.Kind: ", iv.Kind()))

	fmt.Println(tab + strings.Join(infos, ", "))
}

func shuntValuePrint(iv reflect.Value, lv int) {
	printValueXes[iv.Kind()](iv, lv)
}

func printValueInvalid(iv reflect.Value, lv int)    { printValue(iv, lv) }
func printValueBool(iv reflect.Value, lv int)       { printValue(iv, lv) }
func printValueInt(iv reflect.Value, lv int)        { printValue(iv, lv) }
func printValueInt8(iv reflect.Value, lv int)       { printValueInt(iv, lv) }
func printValueInt16(iv reflect.Value, lv int)      { printValueInt(iv, lv) }
func printValueInt32(iv reflect.Value, lv int)      { printValueInt(iv, lv) }
func printValueInt64(iv reflect.Value, lv int)      { printValueInt(iv, lv) }
func printValueUint(iv reflect.Value, lv int)       { printValue(iv, lv) }
func printValueUint8(iv reflect.Value, lv int)      { printValueUint(iv, lv) }
func printValueUint16(iv reflect.Value, lv int)     { printValueUint(iv, lv) }
func printValueUint32(iv reflect.Value, lv int)     { printValueUint(iv, lv) }
func printValueUint64(iv reflect.Value, lv int)     { printValueUint(iv, lv) }
func printValueUintptr(iv reflect.Value, lv int)    { printValue(iv, lv) }
func printValueFloat32(iv reflect.Value, lv int)    { printValue(iv, lv) }
func printValueFloat64(iv reflect.Value, lv int)    { printValueFloat32(iv, lv) }
func printValueComplex64(iv reflect.Value, lv int)  { printValue(iv, lv) }
func printValueComplex128(iv reflect.Value, lv int) { printValueComplex64(iv, lv) }
func printValueArray(iv reflect.Value, lv int)      { printValue(iv, lv) }
func printValueChan(iv reflect.Value, lv int)       {}
func printValueFunc(iv reflect.Value, lv int)       {}
func printValueInterface(iv reflect.Value, lv int)  { printValue(iv, lv); shuntValuePrint(iv.Elem(), lv) }
func printValueMap(iv reflect.Value, lv int)        { printValue(iv, lv) }
func printValuePtr(iv reflect.Value, lv int)        { printValue(iv, lv); shuntValuePrint(iv.Elem(), lv) }
func printValueSlice(iv reflect.Value, lv int)      { printValue(iv, lv) }
func printValueString(iv reflect.Value, lv int)     { printValue(iv, lv) }
func printValueStruct(iv reflect.Value, lv int) {
	printValue(iv, lv)

	for i := 0; i < iv.NumField(); i++ {
		shuntValuePrint(iv.Field(i), lv+1)
	}
}
func printValueUnsafePointer(iv reflect.Value, lv int) {}
