package go2v8

// #cgo CFLAGS: -I/home/link/workspace_git/clabs/src/go2v8
// #cgo LDFLAGS: -L/home/link/workspace_git/clabs/built -lgo2v8 -lbridge -L/home/link/workspace_labs/v8/v8/out.gn/x64.release_shared_nosnapshots -lv8 -lv8_libbase -lv8_libplatform -licuuc -licui18n -lrt -ldl -lpthread
// #include "go2v8.h"
import "C"

func Test() {
	C.v8_startup()
	C.vm_create()
	C.v8_shutdown()
}
