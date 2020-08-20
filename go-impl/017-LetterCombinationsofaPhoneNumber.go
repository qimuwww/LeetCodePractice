package 

/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.

*/

var numLetterMap = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

// 如果是两个字符串切片进行组合,则拼接出来的新字符串个数为m*n
// 按照这种每次只拼接两组字符串切片的方法,将digits进行递归调用
// 将大于2组的字符串数组拆分成每次只操作两个字符串数组
func letterCombinations(digits string) []string {
	if len(digits) <= 1 {
		return numLetterMap[digits]
	}
	retLetter := letterCombinations(digits[1:])
	letter := make([]string, 0, len(retLetter)*len(numLetterMap[digits[0:1]]))
	for i := 0; i < len(numLetterMap[digits[0:1]]); i++ {
		for j := 0; j < len(retLetter); j++ {
			letter = append(letter, strings.join(numLetterMap[digits[0:1]][i]+retLetter[j]))
		}
	}
	return letter
}
/*
func letterCombinations(digits string) []string {
	if len(digits) <= 1 {
		return numLetterMap[digits]
	}
	retLetter := letterCombinations(digits[1:])
	letter := make([]string, 0, len(retLetter)*len(numLetterMap[digits[0:1]]))
	for i := 0; i < len(numLetterMap[digits[0:1]]); i++ {
		for j := 0; j < len(retLetter); j++ {
			var b bytes.Buffer
			b.WriteString(numLetterMap[digits[0:1]][i])
			b.WriteString(retLetter[j])
			letter = append(letter, b.String())
		}
	}
	return letter
}
*/
/*
func main() {
	letter := letterCombinations("234")
	for i := range letter {
		fmt.Println(letter[i])
	}
	// output
	// adg adh adi aeg aeh aei afg afh afi bdg bdh bdi beg beh bei bfg bfh bfi cdg cdh cdi ceg ceh cei cfg cfh cfi
}
*/