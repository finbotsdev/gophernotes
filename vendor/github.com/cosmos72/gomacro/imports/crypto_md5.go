// this file was generated by gomacro command: import _b "crypto/md5"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"crypto/md5"
)

// reflection: allow interpreted code to import "crypto/md5"
func init() {
	Packages["crypto/md5"] = Package{
	Binds: map[string]Value{
		"BlockSize":	ValueOf(md5.BlockSize),
		"New":	ValueOf(md5.New),
		"Size":	ValueOf(md5.Size),
		"Sum":	ValueOf(md5.Sum),
	},Untypeds: map[string]string{
		"BlockSize":	"int:64",
		"Size":	"int:16",
	},
	}
}
