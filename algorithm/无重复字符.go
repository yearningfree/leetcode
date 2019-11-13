// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

// 示例 1:

// 输入: "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:

// 输入: "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:

// 输入: "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package algorithm

func lengthOfLongestSubstring(s string) int {
	lenStr := len(s)
	var subString []byte //将不重复字符加入切片
	var max int
	for i := 0; i < lenStr; i++ {
		index, flag := 0, false //重复位置、重复标志
		if len(subString) == 0 {
			subString = append(subString, s[i])
		} else {

			for k, v := range subString {
				if s[i] == v {
					index = k
					flag = true
					break
				}
			}
			subString = append(subString, s[i])

			if flag {
				subString = subString[index+1 : len(subString)]
			}
		}
		if max < len(subString) {
			max = len(subString)
		}
	}

	return max
}

// func main() {
// 	// fmt.Println("请输入一串字符：")
// 	// inputReader := bufio.NewReader(os.Stdin)
// 	// input, err := inputReader.ReadString('\n')
// 	// if err == nil {
// 	// 	fmt.Println(lengthOfLongestSubstring(input))
// 	// }
// 	fmt.Println(lengthOfLongestSubstring(" "))
// }
