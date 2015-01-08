package disp_width

import (
	"unicode/utf8"

	"github.com/moznion/go-unicode-east-asian-width"
)

// 文字列の表示幅
func Measure(str string) (w int) {
	for _, r := range str {
		if eastasianwidth.IsFullwidth(r) {
			w += 2
		} else {
			w++
		}
	}

	return
}

// 文字列の表示幅で切り詰める
func Truncate(str string, width int, omission string) (s string, rest int) {
	omissionLen := Measure(omission)
	omissionRest := width
	tp := 0
	i := 0

	rest = width
	s = str

	for _, r := range str {
		n := 1
		if eastasianwidth.IsFullwidth(r) {
			n = 2
		}

		if tp == 0 && rest < n+omissionLen {
			// omission を追加できる位置を覚えておく
			tp = i
			omissionRest = rest
		}

		if rest < n {
			s = str[0:tp] + omission
			rest = omissionRest - omissionLen
			break
		}

		rest -= n
		i += utf8.RuneLen(r)
	}

	return
}
