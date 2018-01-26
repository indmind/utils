package utils

import (
	"bytes"
	"sort"
	"strings"
)

// StringInSlice return true if string in slice
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// IntInSlice return true if int in slice
func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// ReverseString ...
func ReverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// ReverseSlice ...
func ReverseSlice(ss []string) {
	last := len(ss) - 1

	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

// SplitSubN split string every n number and returned slice of strings
func SplitSubN(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))

	l := len(runes)

	for i, r := range runes {
		sub = sub + string(r)

		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	return subs
}

// StringAppend append multiple strings parameters
func StringAppend(strs ...string) string {
	return strings.Join(strs, "")
}

// AppendSlice appends a string with array of string
func AppendSlice(s string, strs []string) string {
	return s + strings.Join(strs, "")
}

// IntSliceSort sort a slice of integer in ascending order.
func IntSliceSort(n []int) {
	sort.Ints(n)
}

// StringSliceSort sorts a slice of strings in ascending order lexicographically.
func StringSliceSort(s []string) {
	sort.Strings(s)
}

// Float64SliceSort sorts a slice of float64 in ascending order.
func Float64SliceSort(n []float64) {
	sort.Float64s(n)
}

// IntsAreSorted function tests whether a slice of integer is sorted in ascending order
func IntsAreSorted(n []int) bool {
	return sort.IntsAreSorted(n)
}

// StringsAreSorted function tests whether a slice of string is sorted in ascending order
func StringsAreSorted(strSlice []string) bool {
	return sort.StringsAreSorted(strSlice)
}

// Float64sAreSorted function tests whether a slice of float64s is sorted in ascending order
func Float64sAreSorted(fltSlice []float64) bool {
	return sort.Float64sAreSorted(fltSlice)
}

// SearchInts searches the position of x in a sorted slice of int and returns the index as specified by Search
func SearchInts(intSlice []int, x int) int {
	// make duplicate
	temp := make([]int, len(intSlice))
	copy(temp, intSlice)
	// sorting
	IntSliceSort(temp)

	return sort.SearchInts(temp, x)
}

// SearchStrings searches the position of x in a sorted slice of string and returns the index as specified by Search
func SearchStrings(s []string, x string) int {
	// make duplicate
	temp := make([]string, len(s))
	copy(temp, s)
	// sorting
	StringSliceSort(temp)

	return sort.SearchStrings(temp, x)
}
