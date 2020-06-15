package leetcode

import (
	"strconv"
	"strings"
)

/**
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

有效的 IP 地址正好由四个整数（每个整数位于 0 到 255 之间组成），整数之间用 '.' 分隔。

 

示例:

输入: "25525511135"
输出: ["255.255.11.135", "255.255.111.35"]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/restore-ip-addresses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func restoreIpAddresses(s string) []string {
	var result []string
	res := restoreIpAddressesRecursion(&s, 0, restoreIpAddressesIp2Parts(&s, 0), []string{})
	for _, v := range res {
		result = append(result, strings.Join(v, "."))
	}
	return result
}

// 判断ip每段是否合法
func restoreIpAddressesIpIsLegal(s string) bool {
	// 解决 01 这种以0开头的问题
	if len(s) > 1 && string(s[0]) == "0" {
		return false
	}
	if part, err := strconv.Atoi(s); err == nil && part >= 0 && part <= 255 {
		return true
	}
	return false
}

// 判断某段ip所有的可能性 255 => [2, 25, 255]
func restoreIpAddressesIp2Parts(s *string, position int) []string {
	var result []string
	for ; position < len(*s); position++ {
		prePart := ""
		if len(result) > 0 {
			prePart = result[len(result)-1]
		}

		part := prePart + string((*s)[position])
		if !restoreIpAddressesIpIsLegal(part) {
			break
		}
		result = append(result, part)
	}
	return result
}

// 获取所有组成类型 递归
// s 公共字符串
// position 当前遍历到第几位
// choiceList 下一次可供选择的部分即255 则 可选择 [2, 25, 255]
// pathList 存储已经走过的路径
func restoreIpAddressesRecursion(s *string, position int, choiceList, pathList []string) [][]string {
	var result [][]string
	// 存储路径到达4个,停止递归
	if len(pathList) == 4 {
		// 此时没有候选路径说明,此数据合法
		if len(choiceList) == 0 {
			return [][]string{pathList}
		}
		// 否则返回空值
		return [][]string{}
	}
	for i := 0; i < len(choiceList); i++ {
		// 新的候选路径
		newChoiceList := restoreIpAddressesIp2Parts(s, position+len(choiceList[i]))
		// 记录走过的路径
		newPathList := append([]string{}, pathList...)
		newPathList = append(newPathList, choiceList[i])
		//fmt.Println(position, "choiceList", choiceList, "pathList", pathList, "newChoiceList", newChoiceList, "newPathList", newPathList)
		res := restoreIpAddressesRecursion(s, position+len(choiceList[i]), newChoiceList, newPathList)
		// 如果返回的值大于0组,则记录
		if len(res) > 0 {
			result = append(result, res...)
		}
	}
	return result
}
