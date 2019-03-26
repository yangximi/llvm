// Code generated by "string2enum -linecomment -type DLLStorageClass ../../ir/enum"; DO NOT EDIT.

package enum

import "fmt"
import "github.com/llir/llvm/ir/enum"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the string2enum command to generate them again.
	var x [1]struct{}
	_ = x[enum.DLLStorageClassNone-0]
	_ = x[enum.DLLStorageClassDLLExport-1]
	_ = x[enum.DLLStorageClassDLLImport-2]
}

const _DLLStorageClass_name = "nonedllexportdllimport"

var _DLLStorageClass_index = [...]uint8{0, 4, 13, 22}

func DLLStorageClassFromString(s string) enum.DLLStorageClass {
	if len(s) == 0 {
		return 0
	}
	for i := range _DLLStorageClass_index[:len(_DLLStorageClass_index)-1] {
		if s == _DLLStorageClass_name[_DLLStorageClass_index[i]:_DLLStorageClass_index[i+1]] {
			return enum.DLLStorageClass(i)
		}
	}
	panic(fmt.Errorf("unable to locate DLLStorageClass enum corresponding to %q", s))
}
