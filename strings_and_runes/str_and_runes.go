package main

import (
	"fmt"
	"unicode/utf8"
)

// https://go.dev/blog/strings

func main() {
	const s = "สวัสดี"
	fmt.Println(s)

	fmt.Println("Len:", len(s))

	// the hex values of all the bytes that
	// constitute the code points in s.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println("\nRune count:", utf8.RuneCountInString(s))

	for idx, runeVal := range s {
		fmt.Printf("%#U starts at byte position %d\n", runeVal, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString:")
	for i, w := 0, 0; i < len(s); i += w {
		r, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at byte position %d\n", r, i)
		w = width
		examineRune(r)
	}
}

// Values enclosed in single quotes are rune literals.
// We can compare a rune value to a rune literal directly.
func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
