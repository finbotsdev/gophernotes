// this file was generated by gomacro command: import _b "fmt"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"fmt"
)

// reflection: allow interpreted code to import "fmt"
func init() {
	Packages["fmt"] = Package{
	Binds: map[string]Value{
		"Errorf":	ValueOf(fmt.Errorf),
		"Fprint":	ValueOf(fmt.Fprint),
		"Fprintf":	ValueOf(fmt.Fprintf),
		"Fprintln":	ValueOf(fmt.Fprintln),
		"Fscan":	ValueOf(fmt.Fscan),
		"Fscanf":	ValueOf(fmt.Fscanf),
		"Fscanln":	ValueOf(fmt.Fscanln),
		"Print":	ValueOf(fmt.Print),
		"Printf":	ValueOf(fmt.Printf),
		"Println":	ValueOf(fmt.Println),
		"Scan":	ValueOf(fmt.Scan),
		"Scanf":	ValueOf(fmt.Scanf),
		"Scanln":	ValueOf(fmt.Scanln),
		"Sprint":	ValueOf(fmt.Sprint),
		"Sprintf":	ValueOf(fmt.Sprintf),
		"Sprintln":	ValueOf(fmt.Sprintln),
		"Sscan":	ValueOf(fmt.Sscan),
		"Sscanf":	ValueOf(fmt.Sscanf),
		"Sscanln":	ValueOf(fmt.Sscanln),
	},Types: map[string]Type{
		"Formatter":	TypeOf((*fmt.Formatter)(nil)).Elem(),
		"GoStringer":	TypeOf((*fmt.GoStringer)(nil)).Elem(),
		"ScanState":	TypeOf((*fmt.ScanState)(nil)).Elem(),
		"Scanner":	TypeOf((*fmt.Scanner)(nil)).Elem(),
		"State":	TypeOf((*fmt.State)(nil)).Elem(),
		"Stringer":	TypeOf((*fmt.Stringer)(nil)).Elem(),
	},Proxies: map[string]Type{
		"Formatter":	TypeOf((*Formatter_fmt)(nil)).Elem(),
		"GoStringer":	TypeOf((*GoStringer_fmt)(nil)).Elem(),
		"ScanState":	TypeOf((*ScanState_fmt)(nil)).Elem(),
		"Scanner":	TypeOf((*Scanner_fmt)(nil)).Elem(),
		"State":	TypeOf((*State_fmt)(nil)).Elem(),
		"Stringer":	TypeOf((*Stringer_fmt)(nil)).Elem(),
	},
	}
}

// --------------- proxy for fmt.Formatter ---------------
type Formatter_fmt struct {
	Object	interface{}
	Format_	func(_proxy_obj_ interface{}, f fmt.State, c rune) 
}
func (Proxy *Formatter_fmt) Format(f fmt.State, c rune)  {
	Proxy.Format_(Proxy.Object, f, c)
}

// --------------- proxy for fmt.GoStringer ---------------
type GoStringer_fmt struct {
	Object	interface{}
	GoString_	func(interface{}) string
}
func (Proxy *GoStringer_fmt) GoString() string {
	return Proxy.GoString_(Proxy.Object)
}

// --------------- proxy for fmt.ScanState ---------------
type ScanState_fmt struct {
	Object	interface{}
	Read_	func(_proxy_obj_ interface{}, buf []byte) (n int, err error)
	ReadRune_	func(interface{}) (r rune, size int, err error)
	SkipSpace_	func(interface{}) 
	Token_	func(_proxy_obj_ interface{}, skipSpace bool, f func(rune) bool) (token []byte, err error)
	UnreadRune_	func(interface{}) error
	Width_	func(interface{}) (wid int, ok bool)
}
func (Proxy *ScanState_fmt) Read(buf []byte) (n int, err error) {
	return Proxy.Read_(Proxy.Object, buf)
}
func (Proxy *ScanState_fmt) ReadRune() (r rune, size int, err error) {
	return Proxy.ReadRune_(Proxy.Object)
}
func (Proxy *ScanState_fmt) SkipSpace()  {
	Proxy.SkipSpace_(Proxy.Object)
}
func (Proxy *ScanState_fmt) Token(skipSpace bool, f func(rune) bool) (token []byte, err error) {
	return Proxy.Token_(Proxy.Object, skipSpace, f)
}
func (Proxy *ScanState_fmt) UnreadRune() error {
	return Proxy.UnreadRune_(Proxy.Object)
}
func (Proxy *ScanState_fmt) Width() (wid int, ok bool) {
	return Proxy.Width_(Proxy.Object)
}

// --------------- proxy for fmt.Scanner ---------------
type Scanner_fmt struct {
	Object	interface{}
	Scan_	func(_proxy_obj_ interface{}, state fmt.ScanState, verb rune) error
}
func (Proxy *Scanner_fmt) Scan(state fmt.ScanState, verb rune) error {
	return Proxy.Scan_(Proxy.Object, state, verb)
}

// --------------- proxy for fmt.State ---------------
type State_fmt struct {
	Object	interface{}
	Flag_	func(_proxy_obj_ interface{}, c int) bool
	Precision_	func(interface{}) (prec int, ok bool)
	Width_	func(interface{}) (wid int, ok bool)
	Write_	func(_proxy_obj_ interface{}, b []byte) (n int, err error)
}
func (Proxy *State_fmt) Flag(c int) bool {
	return Proxy.Flag_(Proxy.Object, c)
}
func (Proxy *State_fmt) Precision() (prec int, ok bool) {
	return Proxy.Precision_(Proxy.Object)
}
func (Proxy *State_fmt) Width() (wid int, ok bool) {
	return Proxy.Width_(Proxy.Object)
}
func (Proxy *State_fmt) Write(b []byte) (n int, err error) {
	return Proxy.Write_(Proxy.Object, b)
}

// --------------- proxy for fmt.Stringer ---------------
type Stringer_fmt struct {
	Object	interface{}
	String_	func(interface{}) string
}
func (Proxy *Stringer_fmt) String() string {
	return Proxy.String_(Proxy.Object)
}
