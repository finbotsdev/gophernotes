// this file was generated by gomacro command: import _b "runtime/trace"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"runtime/trace"
)

// reflection: allow interpreted code to import "runtime/trace"
func init() {
	Packages["runtime/trace"] = Package{
	Binds: map[string]Value{
		"Start":	ValueOf(trace.Start),
		"Stop":	ValueOf(trace.Stop),
	},
	}
}
