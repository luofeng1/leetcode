package leetcode

/**
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

 

示例：

输入：n = 3
输出：[
   "((()))",
   "(()())",
   "(())()",
   "()(())",
   "()()()"
     ]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func generateParenthesis(n int) []string {
	dfs := generateParenthesisDFS(n*2, generateParenthesisGetChoiceList(n*2, []byte{}), []byte{})
	var result []string
	for _, v := range dfs {
		//fmt.Println(string(v))
		result = append(result, string(v))
	}
	return result
}

// 获取备选 列表 通过已选路径推断
func generateParenthesisGetChoiceList(remainderLevel int, pathList []byte) []byte {
	all := (remainderLevel + len(pathList)) / 2
	left := 0
	right := 0
	for _, b := range pathList {
		if b == byte('(') {
			left++
		} else {
			right++
		}
	}

	if left > right && left < all && left > 0 {
		//fmt.Println(remainderLevel, string(pathList), "()")
		return []byte{'(', ')'}
	}
	if left == all {
		//fmt.Println(remainderLevel, string(pathList), ")")
		return []byte{')'}
	}
	//fmt.Println(remainderLevel, string(pathList), "(")
	return []byte{'('}
}

func generateParenthesisDFS(remainderLevel int, choiceList, pathList []byte) [][]byte {
	var result [][]byte
	if remainderLevel == 0 {
		return [][]byte{pathList}
	}
	for _, choice := range choiceList {
		newPathList := append([]byte{}, pathList...)
		newPathList = append(newPathList, choice)
		newChoiceList := generateParenthesisGetChoiceList(remainderLevel-1, newPathList)
		//fmt.Println(remainderLevel, "choiceList:", string(choiceList), "pathList:", string(pathList), "newChoiceList:", string(newChoiceList), "newPathList:", string(newPathList))
		result = append(result, generateParenthesisDFS(remainderLevel-1, newChoiceList, newPathList)...)
	}

	return result
}
