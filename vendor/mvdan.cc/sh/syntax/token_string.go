// Code generated by "stringer -type token -linecomment -trimprefix _"; DO NOT EDIT.

package syntax

import "strconv"

const _token_name = "illegalTokEOFNewlLitLitWordLitRedir'\"`&&&||||&$$'$\"${$[$($(([(((}])));;;;&;;&;|!++--***==!=<=>=+=-=*=/=%=&=|=^=<<=>>=>>><<><&>&>|<<<<-<<<&>&>><(>(+:+-:-?:?=:=%%%###^^^,,,@///:-e-f-d-c-b-p-S-L-k-g-u-G-O-N-r-w-x-s-t-z-n-o-v-R=~-nt-ot-ef-eq-ne-le-ge-lt-gt?(*(+(@(!("

var _token_index = [...]uint16{0, 10, 13, 17, 20, 27, 35, 36, 37, 38, 39, 41, 43, 44, 46, 47, 49, 51, 53, 55, 57, 60, 61, 62, 64, 65, 66, 67, 69, 70, 72, 74, 77, 79, 80, 82, 84, 85, 87, 89, 91, 93, 95, 97, 99, 101, 103, 105, 107, 109, 111, 114, 117, 118, 120, 121, 123, 125, 127, 129, 131, 134, 137, 139, 142, 144, 146, 147, 149, 150, 152, 153, 155, 156, 158, 159, 161, 162, 164, 165, 167, 168, 170, 171, 172, 174, 175, 177, 179, 181, 183, 185, 187, 189, 191, 193, 195, 197, 199, 201, 203, 205, 207, 209, 211, 213, 215, 217, 219, 221, 223, 225, 228, 231, 234, 237, 240, 243, 246, 249, 252, 254, 256, 258, 260, 262}

func (i token) String() string {
	if i >= token(len(_token_index)-1) {
		return "token(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _token_name[_token_index[i]:_token_index[i+1]]
}