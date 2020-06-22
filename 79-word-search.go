package leetcode

/**
给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。

同一个单元格内的字母不允许被重复使用。

 

示例:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false
 

提示：

board 和 word 中只包含大写和小写英文字母。
1 <= board.length <= 200
1 <= board[i].length <= 200
1 <= word.length <= 10^3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
choice: 上下左右 (x,y)
path: [][]int{} 1,1,1,1 走过的路径 => 0

如何判断是否全部匹配, word 走到最后一位
*/
func exist(board [][]byte, word string) bool {
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			if dfs := existDFS([][]int{{x, y}}, existNewBoard(board), &word, 0); dfs {
				return true
			}
		}
	}
	return false
}

func existNewBoard(board [][]byte) [][]byte {
	var newBoard [][]byte
	for _, bytes := range board {
		newBoard = append(newBoard, []byte(string(bytes)))
	}
	return newBoard
}

func existGetChoices(path [][]byte, x, y int) [][]int {
	var result [][]int
	if x > 0 && path[x-1][y] != '0' {
		result = append(result, []int{x - 1, y})
	}
	if x < len(path)-1 && path[x+1][y] != '0' {
		result = append(result, []int{x + 1, y})
	}
	if y > 0 && path[x][y-1] != '0' {
		result = append(result, []int{x, y - 1})
	}
	if y < len(path[0])-1 && path[x][y+1] != '0' {
		result = append(result, []int{x, y + 1})
	}
	return result
}

func existDFS(choice [][]int, path [][]byte, word *string, position int) bool {
	if len(*word) == position {
		return true
	}
	//fmt.Println(position, "choice: ", choice, string(path[0]), string(path[1]), string(path[2]))
	for _, xy := range choice {
		x := xy[0]
		y := xy[1]
		if path[x][y] != (*word)[position] {
			continue
		}
		newPath := existNewBoard(path)
		newPath[x][y] = '0'
		if res := existDFS(existGetChoices(newPath, x, y), newPath, word, position+1); res {
			return true
		}
	}
	return false
}
