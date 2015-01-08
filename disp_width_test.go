package disp_width

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeasure(t *testing.T) {
	data := []struct {
		src   string
		width int
	}{
		{"こんにちは", 10},   // ひらがな
		{"コンニチハ", 10},   // カタカナ
		{"ｺﾝﾆﾁﾊ", 5},    // 半角カタカナ
		{"こんにちは世界", 14}, // 漢字かな混じり
		{"Hello", 5},    // 英語
		{"你好", 4},       // 中国語
		{"안녕하세요", 10},   // 韓国語
		{"Καλημερα", 8}, // ギリシャ語

		{"こんにちは ｺﾝﾆﾁﾊ Hello 你好 안녕하세요 Καλημερα 世界", 52}, // 各国語
	}

	for _, v := range data {
		assert.Equal(t, Measure(v.src), v.width, "Measure(%q)", v.src)
	}
}

func TestTruncate(t *testing.T) {
	data := []struct {
		src      string
		maxWidth int
		omission string
		dest     string
		rest     int
	}{
		{"こんにちは世界", 12, "", "こんにちは世", 0},
		{"こんにちは世界", 13, "", "こんにちは世", 1},
		{"こんにちは世界", 14, "", "こんにちは世界", 0},
		{"こんにちは世界", 15, "", "こんにちは世界", 1},

		{"こんにちは世界", 12, "...", "こんにち...", 1},
		{"こんにちは世界", 13, "...", "こんにちは...", 0},
		{"こんにちは世界", 14, "...", "こんにちは世界", 0},
		{"こんにちは世界", 15, "...", "こんにちは世界", 1},

		{"こんにちは世界", 12, "…", "こんにちは…", 1},
		{"こんにちは世界", 13, "…", "こんにちは世…", 0},
		{"こんにちは世界", 14, "…", "こんにちは世界", 0},
		{"こんにちは世界", 15, "…", "こんにちは世界", 1},

		{"Hello 世界", 8, "...", "Hello...", 0},
		{"Hello 世界", 9, "...", "Hello ...", 0},
		{"Hello 世界", 10, "...", "Hello 世界", 0},
		{"Hello 世界", 11, "...", "Hello 世界", 1},
	}

	for _, v := range data {
		dest, rest := Truncate(v.src, v.maxWidth, v.omission)
		assert.Equal(t, dest, v.dest, "Truncate(%q, %d, %q)", v.src, v.maxWidth, v.omission)
		assert.Equal(t, rest, v.rest, "Truncate(%q, %d, %q)", v.src, v.maxWidth, v.omission)
	}
}
