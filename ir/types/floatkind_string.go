// Code generated by "stringer -linecomment -type FloatKind"; DO NOT EDIT.

package types

import "strconv"

const _FloatKind_name = "halffloatdoublefp128x86_fp80ppc_fp128"

var _FloatKind_index = [...]uint8{0, 4, 9, 15, 20, 28, 37}

func (i FloatKind) String() string {
	if i >= FloatKind(len(_FloatKind_index)-1) {
		return "FloatKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FloatKind_name[_FloatKind_index[i]:_FloatKind_index[i+1]]
}
