package go2clabs

// #cgo CFLAGS: -I/home/link/workspace_git/clabs/src -I/home/link/workspace_labs/quickjs-2019-07-09
// #cgo LDFLAGS: -L/home/link/workspace_git/clabs/lib -L/home/link/workspace_labs/quickjs-2019-07-09 -ldl -lm -lclabs_qjs -lquickjs
// #include <qjs/qjs.h>
import "C"

import (
	"unsafe"
)

type qjs struct {
}

//EvalBinary 运行字节码
func (q *qjs) EvalBinary(args []string, raw []byte) {
	C.eval_binary(C.int(len(args)), go2cStringSlice(args), (*C.uchar)(unsafe.Pointer(&raw[0])), C.int(len(raw)))
}

//EvalSource 运行源码
func (q *qjs) EvalSource(args []string, raw []byte) {
	C.eval_source(C.int(len(args)), go2cStringSlice(args), (*C.uchar)(unsafe.Pointer(&raw[0])), C.int(len(raw)))
}
