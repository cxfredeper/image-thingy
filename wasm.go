package main

import (
	"github.com/cxfredeper/image-thingy/codec"
	"syscall/js"
	"unsafe"
)


// We communicate with JS by passing pointers to the wasm memory directly.
// 1. JS request memory of size N.
// 2. Go makes a slice of size N and returns the pointer as uint32.
// 3. JS fills the memory region with payload (write to memory.buffer[idx]).
// 4. JS calls Go to start processing data.
// 5. Go reads from slice, and writes to an output region.
// 6. Go returns the output region address.
// 7. JS queries for output region length.
// 8. JS reads result from output region.
var in, out []byte
var err error

type addr = uint32


//export requestBuffer
func RequestBuffer(length uint32) addr {
	in = make([]byte, int(length))
	return uint32(uintptr(unsafe.Pointer(&in[0])))
}


//export encode
func Encode() addr {
	out, err = codec.EncodePNG(in)
	if err != nil { return 0 }
	return uint32(uintptr(unsafe.Pointer(&out[0])))
}


//export decode
func Decode() addr {
	out, err = codec.DecodePNG(in)
	if err != nil { return 0 }
	return uint32(uintptr(unsafe.Pointer(&out[0])))
}


//export getOutLen
func GetOutLen() uint32 {
	return uint32(len(out))
}

func GetError(_ js.Value, _ []js.Value) any {
	switch err {
		case nil: return js.ValueOf(nil)
		default: return js.ValueOf(err.Error())
	}
}


func init() {
	js.Global().Set("getCodecErrorMsg", js.FuncOf(GetError))
}


func main() {
	select {}
}
