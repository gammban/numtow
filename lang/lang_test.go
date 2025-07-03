package lang

import "testing"

//nolint:gochecknoglobals // test cases for Lang.String method
var testCaseString = []struct {
	give Lang
	want string
}{
	{give: 0, want: ""},
	{give: 1, want: "KZ"},
	{give: 2, want: "RU"},
	{give: 3, want: "EN"},
	{give: 32000, want: ""},
}

func TestLang_String(t *testing.T) {
	for _, v := range testCaseString {
		if got := v.give.String(); got != v.want {
			t.Errorf("got %s, expected %s", got, v.want)
		}
	}
}
