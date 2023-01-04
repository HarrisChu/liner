package liner

import "testing"

type tc struct {
	pattern   string
	str       string
	runeIndex int
}

func TestRuneIndex(t *testing.T) {
	cases := []tc{
		{"我的", "我是我的", 2},
		{"的", "我是我的", 3},
		{"我", "我是我的", 0},
	}
	for _, c := range cases {
		if indexRune([]rune(c.str), []rune(c.pattern)) != c.runeIndex {
			t.Errorf("indexRune(%s, %s) = %d, want %d", c.str, c.pattern, indexRune([]rune(c.str), []rune(c.pattern)), c.runeIndex)
		}
	}

}
