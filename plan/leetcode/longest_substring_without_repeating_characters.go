// Longest Substring Without Repeating Characters
//
// [Medium] [AC:29.0% 1.1M of 3.9M] [filetype:golang]
//
// Given a string, find the length of the longest substring without
// repeating characters.
//
// Example 1:
//
// Input: "abcabcbb"
//
// Output: 3
//
// Explanation: The answer is "abc", with the length of 3.
//
// Example 2:
//
// Input: "bbbbb"
//
// Output: 1
//
// Explanation: The answer is "b", with the length of 1.
//
// Example 3:
//
// Input: "pwwkew"
//
// Output: 3
//
// Explanation: The answer is "wke", with the length of 3.
//
//              Note that the answer must be a substring, "pwke" is a
//              subsequence and not a substring.
//
// [End of Description]
//abcdefajklh
package main

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var start, end int
	for i, str := range s {
		if strings.Index(s[start:end], string(str)) {
			start += 1
		}
		end += 1
	}
	return end - start
}
