package main

import (
	"fmt"
	"unicode/utf8"
)

func banner(text string, width int) {
	padding := (width - len(text)) / 2
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}

	fmt.Println(text)

	for i := 0; i < width; i++ {
		fmt.Print("-")
	}

	fmt.Println()
}

func betterBanner(text string, width int) {
	// B/c of length bug with non-latin unicode chars, use other utils to handle
	padding := (width - utf8.RuneCountInString(text)) / 2
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}

	fmt.Println(text)

	for i := 0; i < width; i++ {
		fmt.Print("-")
	}

	fmt.Println()
}

// isPalindrome("g") -> true
// isPalindrome("go") -> false
// isPalindrome("gog") -> true
// isPalindrome("gogo") -> false
func isPalindrome(text string) bool {
	// should convert to runes to avoid unicode character issues
	rs := []rune(text)
	for i := 0; i < len(rs)/2; i++ {
		if rs[i] != rs[len(rs)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	banner("go", 16)
	// notice how this isn't centered b/c the emoji is multiple bytes long
	banner("g😊", 16)

	// this one is fine though
	betterBanner("g😊", 16)

	// Testing area for string printing
	x, y := 1, "1"
	// Both of these will look the same in the console
	fmt.Printf("x=%v, y=%v\n", x, y)
	// But these will look different
	fmt.Printf("x=%#v, y=%#v\n", x, y)

	// Palindrome testing
	fmt.Println(isPalindrome("g"))
	fmt.Println(isPalindrome("go"))
	fmt.Println(isPalindrome("gog"))
	fmt.Println(isPalindrome("gogo"))
	// try with unicode chars
	fmt.Println(isPalindrome("g◉g"))
}
