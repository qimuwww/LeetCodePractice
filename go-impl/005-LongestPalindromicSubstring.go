package leetcode

/*
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"
*/
/*
	遍历字符串
	每次循环时以当前字符为中心向两端递增递减  判断左右两端的每个字符是否相同
	需要同时判断字符串为奇数串和偶数串的情况
*/
func longestPalindrome(s string) string {
	var lps string
	for i := 0; i < len(s); i++ {
		s1 := palindrome(s, i, i)   // 奇数串
		s2 := palindrome(s, i, i+1) // 偶数串
		if len(s1) > len(lps) {
			lps = s1
		}
		if len(s2) > len(lps) {
			lps = s2
		}
	}
	return lps
}

func palindrome(s string, l, r int) string {
	for {
		if l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
			continue
		}
		break
	}
	return s[l+1 : r]
}
