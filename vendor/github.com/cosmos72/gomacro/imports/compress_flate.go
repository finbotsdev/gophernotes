// this file was generated by gomacro command: import _b "compress/flate"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"compress/flate"
	"io"
)

// reflection: allow interpreted code to import "compress/flate"
func init() {
	Packages["compress/flate"] = Package{
	Binds: map[string]Value{
		"BestCompression":	ValueOf(flate.BestCompression),
		"BestSpeed":	ValueOf(flate.BestSpeed),
		"DefaultCompression":	ValueOf(flate.DefaultCompression),
		"HuffmanOnly":	ValueOf(flate.HuffmanOnly),
		"NewReader":	ValueOf(flate.NewReader),
		"NewReaderDict":	ValueOf(flate.NewReaderDict),
		"NewWriter":	ValueOf(flate.NewWriter),
		"NewWriterDict":	ValueOf(flate.NewWriterDict),
		"NoCompression":	ValueOf(flate.NoCompression),
	},Types: map[string]Type{
		"CorruptInputError":	TypeOf((*flate.CorruptInputError)(nil)).Elem(),
		"InternalError":	TypeOf((*flate.InternalError)(nil)).Elem(),
		"ReadError":	TypeOf((*flate.ReadError)(nil)).Elem(),
		"Reader":	TypeOf((*flate.Reader)(nil)).Elem(),
		"Resetter":	TypeOf((*flate.Resetter)(nil)).Elem(),
		"WriteError":	TypeOf((*flate.WriteError)(nil)).Elem(),
		"Writer":	TypeOf((*flate.Writer)(nil)).Elem(),
	},Proxies: map[string]Type{
		"Reader":	TypeOf((*Reader_compress_flate)(nil)).Elem(),
		"Resetter":	TypeOf((*Resetter_compress_flate)(nil)).Elem(),
	},Untypeds: map[string]string{
		"BestCompression":	"int:9",
		"BestSpeed":	"int:1",
		"DefaultCompression":	"int:-1",
		"HuffmanOnly":	"int:-2",
		"NoCompression":	"int:0",
	},
	}
}

// --------------- proxy for compress/flate.Reader ---------------
type Reader_compress_flate struct {
	Object	interface{}
	Read_	func(_proxy_obj_ interface{}, p []byte) (n int, err error)
	ReadByte_	func(interface{}) (byte, error)
}
func (Proxy *Reader_compress_flate) Read(p []byte) (n int, err error) {
	return Proxy.Read_(Proxy.Object, p)
}
func (Proxy *Reader_compress_flate) ReadByte() (byte, error) {
	return Proxy.ReadByte_(Proxy.Object)
}

// --------------- proxy for compress/flate.Resetter ---------------
type Resetter_compress_flate struct {
	Object	interface{}
	Reset_	func(_proxy_obj_ interface{}, r io.Reader, dict []byte) error
}
func (Proxy *Resetter_compress_flate) Reset(r io.Reader, dict []byte) error {
	return Proxy.Reset_(Proxy.Object, r, dict)
}
