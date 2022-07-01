/*
 * @lc app=leetcode.cn id=241 lang=golang
 *
 * [241] 为运算表达式设计优先级
 *
 * https://leetcode.cn/problems/different-ways-to-add-parentheses/description/
 *
 * algorithms
 * Medium (73.67%)
 * Likes:    624
 * Dislikes: 0
 * Total Accepted:    50K
 * Total Submissions: 67.1K
 * Testcase Example:  '"2-1-1"'
 *
 * 给你一个由数字和运算符组成的字符串 expression ，按不同优先级组合数字和运算符，计算并返回所有可能组合的结果。你可以 按任意顺序 返回答案。
 *
 * 生成的测试用例满足其对应输出值符合 32 位整数范围，不同结果的数量不超过 10^4 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：expression = "2-1-1"
 * 输出：[0,2]
 * 解释：
 * ((2-1)-1) = 0
 * (2-(1-1)) = 2
 *
 *
 * 示例 2：
 *
 *
 * 输入：expression = "2*3-4*5"
 * 输出：[-34,-14,-10,-10,10]
 * 解释：
 * (2*(3-(4*5))) = -34
 * ((2*3)-(4*5)) = -14
 * ((2*(3-4))*5) = -10
 * (2*((3-4)*5)) = -10
 * (((2*3)-4)*5) = 10
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= expression.length <= 20
 * expression 由数字和算符 '+'、'-' 和 '*' 组成。
 * 输入表达式中的所有整数值在范围 [0, 99]
 *
 *
 */

package leetcode

import (
	"strconv"
)

// @lc code=start
func diffWaysToCompute(expression string) (res []int) {
	for i, ch := range expression {
		if ch == '+' || ch == '-' || ch == '*' {
			res1 := diffWaysToCompute(expression[:i])
			res2 := diffWaysToCompute(expression[i+1:])
			for _, r1 := range res1 {
				for _, r2 := range res2 {
					if ch == '+' {
						res = append(res, r1+r2)
					} else if ch == '-' {
						res = append(res, r1-r2)
					} else if ch == '*' {
						res = append(res, r1*r2)
					}
				}
			}
		}
	}
	if len(res) == 0 {
		num, err := strconv.Atoi(expression)
		if err == nil {
			res = append(res, num)
		}
	}
	return
}

// @lc code=end
