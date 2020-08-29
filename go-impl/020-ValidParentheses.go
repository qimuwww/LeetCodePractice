package leetcode

/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true
*/
/*
	合法括号判断
	个人感觉和消消乐有点像
	循环遍历s字符串,遇到'(','[','{'则将该字符append到slice中,遇到')',']','}'则判断
	slice中的最后一个元素是否与其匹配,如果匹配则移除该元素,然后继续循环判断,如果不匹配则
	直接返回.
	用stack更好
*/

func isValid(s string) bool {
	l := len(s)
	if l < 1 {
		return true
	}
	var tmp = make([]byte, 0, l)
	var index int
	for i := 0; i < l; i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			//tmp[index] = s[i]
			tmp = append(tmp, s[i])
			index++
			continue
		} else {
			if len(tmp) < 1 {
				return false
			}
			switch {
			case s[i] == ')':
				{
					if tmp[index-1] == '(' {
						index--
						tmp = tmp[:index]
						continue
					}
					return false
				}
			case s[i] == ']':
				{
					if tmp[index-1] == '[' {
						index--
						tmp = tmp[:index]
						continue
					}
					return false
				}
			case s[i] == '}':
				{
					if tmp[index-1] == '{' {
						index--
						tmp = tmp[:index]
						continue
					}
					return false
				}
			}
		}
	}
	return index == 0
}
