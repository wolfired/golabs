package auto

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

type printType func(iv reflect.Value, it reflect.Type, lv int)

var printfns map[reflect.Kind]printType

func init() {
	printfns = map[reflect.Kind]printType{
		reflect.Invalid:       printInvalid,
		reflect.Bool:          printBool,
		reflect.Int:           printInt,
		reflect.Int8:          printInt8,
		reflect.Int16:         printInt16,
		reflect.Int32:         printInt32,
		reflect.Int64:         printInt64,
		reflect.Uint:          printUint,
		reflect.Uint8:         printUint8,
		reflect.Uint16:        printUint16,
		reflect.Uint32:        printUint32,
		reflect.Uint64:        printUint64,
		reflect.Uintptr:       printUintptr,
		reflect.Float32:       printFloat32,
		reflect.Float64:       printFloat64,
		reflect.Complex64:     printComplex64,
		reflect.Complex128:    printComplex128,
		reflect.Array:         printArray,
		reflect.Chan:          printChan,
		reflect.Func:          printFunc,
		reflect.Interface:     printInterface,
		reflect.Map:           printMap,
		reflect.Ptr:           printPtr,
		reflect.Slice:         printSlice,
		reflect.String:        printString,
		reflect.Struct:        printStruct,
		reflect.UnsafePointer: printUnsafePointer,
	}
}

/*PrintInstance 自动打印Instance*/
func PrintInstance(instance interface{}) {
	shuntPrint(reflect.ValueOf(instance), reflect.TypeOf(instance), 0)
}

func shuntPrint(iv reflect.Value, it reflect.Type, lv int) {
	printfns[it.Kind()](iv, it, lv)
}

func printInvalid(iv reflect.Value, it reflect.Type, lv int)    { printInfo(iv, it, lv) }
func printBool(iv reflect.Value, it reflect.Type, lv int)       { printInfo(iv, it, lv) }
func printInt(iv reflect.Value, it reflect.Type, lv int)        { printInfo(iv, it, lv) }
func printInt8(iv reflect.Value, it reflect.Type, lv int)       { printInt(iv, it, lv) }
func printInt16(iv reflect.Value, it reflect.Type, lv int)      { printInt(iv, it, lv) }
func printInt32(iv reflect.Value, it reflect.Type, lv int)      { printInt(iv, it, lv) }
func printInt64(iv reflect.Value, it reflect.Type, lv int)      { printInt(iv, it, lv) }
func printUint(iv reflect.Value, it reflect.Type, lv int)       { printInfo(iv, it, lv) }
func printUint8(iv reflect.Value, it reflect.Type, lv int)      { printUint(iv, it, lv) }
func printUint16(iv reflect.Value, it reflect.Type, lv int)     { printUint(iv, it, lv) }
func printUint32(iv reflect.Value, it reflect.Type, lv int)     { printUint(iv, it, lv) }
func printUint64(iv reflect.Value, it reflect.Type, lv int)     { printUint(iv, it, lv) }
func printUintptr(iv reflect.Value, it reflect.Type, lv int)    { printInfo(iv, it, lv) }
func printFloat32(iv reflect.Value, it reflect.Type, lv int)    { printInfo(iv, it, lv) }
func printFloat64(iv reflect.Value, it reflect.Type, lv int)    { printFloat32(iv, it, lv) }
func printComplex64(iv reflect.Value, it reflect.Type, lv int)  { printInfo(iv, it, lv) }
func printComplex128(iv reflect.Value, it reflect.Type, lv int) { printComplex64(iv, it, lv) }
func printArray(iv reflect.Value, it reflect.Type, lv int)      {}
func printChan(iv reflect.Value, it reflect.Type, lv int)       {}
func printFunc(iv reflect.Value, it reflect.Type, lv int)       {}
func printInterface(iv reflect.Value, it reflect.Type, lv int)  {}
func printMap(iv reflect.Value, it reflect.Type, lv int)        {}
func printPtr(iv reflect.Value, it reflect.Type, lv int) {
	printInfo(iv, it, lv)

	shuntPrint(iv.Elem(), it.Elem(), lv)
}
func printSlice(iv reflect.Value, it reflect.Type, lv int)  {}
func printString(iv reflect.Value, it reflect.Type, lv int) { printInfo(iv, it, lv) }
func printStruct(iv reflect.Value, it reflect.Type, lv int) {
	printInfo(iv, it, lv)

	if !iv.IsValid() {
		return
	}

	for i := 0; i < it.NumField(); i++ {
		shuntPrint(iv.Field(i), it.Field(i).Type, lv+1)
	}
}
func printUnsafePointer(iv reflect.Value, it reflect.Type, lv int) {}

func printInfo(iv reflect.Value, it reflect.Type, lv int) {
	tab := string(bytes.Repeat([]byte{'\t'}, lv))

	infos := make([]string, 0, 64)
	infos = append(infos, fmt.Sprint("Type: ", it))
	infos = append(infos, fmt.Sprint("Type.Kind: ", it.Kind()))
	infos = append(infos, fmt.Sprint("Value.CanAddr: ", iv.CanAddr()))
	infos = append(infos, fmt.Sprint("Value.CanSet: ", iv.CanSet()))

	fmt.Println(tab + strings.Join(infos, ", "))
}
