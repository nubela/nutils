package nutils
import (
	"io"
	"bytes"
	"unsafe"
	"net/http"
	"io/ioutil"
)

// this function converts a Reader type into string
// note: it does it unsafely, altho efficiently.
// note2: this will consumer the Reader's buffer
func ReaderToString(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	b := buf.Bytes()
	return BytesToString(b)
}

func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

// this takes a http.Request obj, removes `chunked encoding` header, and adds a Content-Length header by calculating
// the Body size.
func MakeNonChunkEncodingRequest(r *http.Request) *http.Request {
	delete(r.Header, "Transfer-Encoding")

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	i := len(buf.Bytes())
	r.ContentLength = int64(i)
	r.Body = ioutil.NopCloser(buf)

	return r
}