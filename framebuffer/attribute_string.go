// Code generated by "stringer -type=Attribute"; DO NOT EDIT.

package framebuffer

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Normal-1]
	_ = x[Bold-2]
	_ = x[Faint-4]
	_ = x[Italic-8]
	_ = x[Underline-16]
	_ = x[Blink-32]
	_ = x[Inverse-64]
	_ = x[Invisible-128]
	_ = x[Crossed-256]
	_ = x[Double-512]
}

const (
	_Attribute_name_0 = "NormalBold"
	_Attribute_name_1 = "Faint"
	_Attribute_name_2 = "Italic"
	_Attribute_name_3 = "Underline"
	_Attribute_name_4 = "Blink"
	_Attribute_name_5 = "Inverse"
	_Attribute_name_6 = "Invisible"
	_Attribute_name_7 = "Crossed"
	_Attribute_name_8 = "Double"
)

var (
	_Attribute_index_0 = [...]uint8{0, 6, 10}
)

func (i Attribute) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _Attribute_name_0[_Attribute_index_0[i]:_Attribute_index_0[i+1]]
	case i == 4:
		return _Attribute_name_1
	case i == 8:
		return _Attribute_name_2
	case i == 16:
		return _Attribute_name_3
	case i == 32:
		return _Attribute_name_4
	case i == 64:
		return _Attribute_name_5
	case i == 128:
		return _Attribute_name_6
	case i == 256:
		return _Attribute_name_7
	case i == 512:
		return _Attribute_name_8
	default:
		return "Attribute(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
