package leetcode

/**
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func permute(nums []int) [][]int {
	return permuteRecursion(nums, []int{})
}

func permuteRecursion(choiceList, pathList []int) [][]int {
	var result [][]int
	if len(choiceList) == 0 {
		return [][]int{pathList}
	}
	for i := 0; i < len(choiceList); i++ {
		newChoiceList := append([]int{}, choiceList[:i]...)
		newChoiceList = append(newChoiceList, choiceList[i+1:]...)
		newPathList := append([]int{choiceList[i]}, pathList...)
		//fmt.Println("组合: ", choiceList[i], newChoiceList, newPathList)
		result = append(result, permuteRecursion(newChoiceList, newPathList)...)
	}
	return result
}
