// Code generated by "string2enum -linecomment -type EmissionKind ../../ir/enum"; DO NOT EDIT.

package enum

import "fmt"
import "github.com/llir/llvm/ir/enum"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the string2enum command to generate them again.
	var x [1]struct{}
	_ = x[enum.EmissionKindNoDebug-0]
	_ = x[enum.EmissionKindFullDebug-1]
	_ = x[enum.EmissionKindLineTablesOnly-2]
}

const _EmissionKind_name = "NoDebugFullDebugLineTablesOnly"

var _EmissionKind_index = [...]uint8{0, 7, 16, 30}

func EmissionKindFromString(s string) enum.EmissionKind {
	if len(s) == 0 {
		return 0
	}
	for i := range _EmissionKind_index[:len(_EmissionKind_index)-1] {
		if s == _EmissionKind_name[_EmissionKind_index[i]:_EmissionKind_index[i+1]] {
			return enum.EmissionKind(i)
		}
	}
	panic(fmt.Errorf("unable to locate EmissionKind enum corresponding to %q", s))
}
