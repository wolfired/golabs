package go2clabs

import "C"

import "unsafe"

// NewJS 新建一个客户端
func NewJS() JS{
	return &qjs{}
}

// JS 客户端
type JS interface {
	EvalBinary(args []string, raw []byte)
	EvalSource(args []string, raw []byte)
}

func go2cStringSlice(strs []string) **C.char {
	length := len(strs)
	buf := make([]*C.char, length)

	for i := 0; i < length; i++ {
		buf[i] = (*C.char)(unsafe.Pointer(C.CString(strs[i])))
	}

	return (**C.char)(unsafe.Pointer(&buf[0]))
}
