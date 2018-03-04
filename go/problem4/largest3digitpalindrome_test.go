package problem4_test

import (
	"testing"

	"github.com/sks/euler/go/problem4"
)

func TestIsPalindrome(t *testing.T) {
	isPalindrom := problem4.IsPalindrom(1001)
	if !isPalindrom {
		t.Errorf("1001 should be a palindrome")
	}
}

func TestFalseIsPalindrome(t *testing.T) {
	isPalindrom := problem4.IsPalindrom(123123)
	if isPalindrom {
		t.Errorf("123123 is not a palindrome")
	}
}
