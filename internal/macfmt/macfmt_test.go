package macfmt

import "testing"

func TestFormatLine(t *testing.T) {
	cases := map[string]string{
		"001A2B3C4D5E":         "00:1A:2B:3C:4D:5E",
		"00-1a-2b-3c-4d-5e":    "00:1A:2B:3C:4D:5E",
		"001a.2b3c.4d5e":       "00:1A:2B:3C:4D:5E",
		"00:1a:2b:3c:4d:5e":    "00:1A:2B:3C:4D:5E",
		"  00 1A 2B 3C 4D 5E ": "00:1A:2B:3C:4D:5E",
		"":                     "",
		"   ":                  "",
		"zzzz":                 "",
		"001A2B3C4D":           ErrTooFew,
		"001A2B3C4D5E6F":       ErrTooMany,
	}
	for in, want := range cases {
		if got := FormatLine(in); got != want {
			t.Errorf("FormatLine(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestFormatAll(t *testing.T) {
	in := "001a2b3c4d5e\r\n\r\n00-11-22-33-44-55\r\nabc"
	want := "00:1A:2B:3C:4D:5E\r\n\r\n00:11:22:33:44:55\r\n" + ErrTooFew
	if got := FormatAll(in, "\r\n"); got != want {
		t.Errorf("FormatAll = %q, want %q", got, want)
	}
}
