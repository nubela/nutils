package nutils
import (
	"io"
	"bytes"
	"unsafe"
)

// this function converts a Reader type into string
// note: it does it unsafely, altho efficiently.
func ReaderToString(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	b := buf.Bytes()
	return *(*string)(unsafe.Pointer(&b))
}