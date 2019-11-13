// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

// 示例:

// 给定 nums = [2, 7, 11, 15], target = 9

// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/two-sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import "fmt"

//1.
// func twoSum(nums []int, target int) []int {
// 	var index []int
// 	for i, _ := range nums {
// 		for j, _ := range nums {
// 			if i == j && nums[i] == nums[j] {
// 				continue
// 			}
// 			if nums[i]+nums[j] == target {
// 				index = append(index, i)
// 				index = append(index, j)
// 				break
// 			}
// 		}
// 		if len(index) == 2 {
// 			break
// 		}
// 	}
// 	return index
// }

//2.
func twoSum(nums []int, target int) []int {
	for i := 1; i <= len(nums); i++ {
		val := nums[i-1]
		sli := nums[i:]
		for k, v := range sli {
			if target-v == val {
				return []int{i - 1, i + k}
			}
		}
	}
	return []int{}
}

func main() {
	array := [4]int{2, 7, 11, 15}
	result := twoSum(array[:], 9)
	fmt.Println(result)
}
