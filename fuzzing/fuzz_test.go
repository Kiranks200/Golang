package main

import (
	"testing"
	"unicode/utf8"
)

// func TestReverse(t *testing.T) {
// 	inputs := []string{"Hello, world", " ", "!12345", "", "😀👍🚀", "こんにちは"}
// 	for _, orig := range inputs {
// 		rev, err1 := Reverse(orig)
// 		if err1 != nil {
// 			t.Logf("Error reversing string %q: %v", orig, err1)
// 			continue
// 		}
// 		t.Logf("Original: %q, Reversed: %q", orig, rev)

// 		doubleRev, err2 := Reverse(rev)
// 		if err2 != nil {
// 			t.Logf("Error reversing string %q: %v", rev, err2)
// 			continue
// 		}
// 		t.Logf("Double Reversed: %q", doubleRev)

// 		if orig != doubleRev {
// 			t.Errorf("Mismatch: Before: %q, After: %q", orig, doubleRev)
// 		}
// 	}
// }

func FuzzReverse(f *testing.F) {
	testcases := []string{
		"Hello, world",
		" ",
		"!12345",
		"",
		"😀👍🚀",                                    // Emoji
		"こんにちは",                                  // Japanese
		"longstringwithmorethanthirtycharacters", // Long string
		"1234567890abcdefghijklmnopqrstuvwxyz",   // Alphanumeric
	}

	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			t.Logf("Error reversing string %q: %v", orig, err1)
			return // Skip further checks if there's an error
		}

		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			t.Logf("Error reversing string %q: %v", rev, err2)
			return // Skip further checks if there's an error
		}

		t.Logf("Input: %q, First Reverse: %q, Second Reverse: %q", orig, rev, doubleRev)

		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
