/*
 * @lc app=leetcode.cn id=483 lang=golang
 *
 * [483] 最小好进制
 *
 * https://leetcode-cn.com/problems/smallest-good-base/description/
 *
 * algorithms
 * Hard (56.72%)
 * Likes:    120
 * Dislikes: 0
 * Total Accepted:    12.1K
 * Total Submissions: 21K
 * Testcase Example:  '"13"'
 *
 * 对于给定的整数 n, 如果n的k（k>=2）进制数的所有数位全为1，则称 k（k>=2）是 n 的一个好进制。
 *
 * 以字符串的形式给出 n, 以字符串的形式返回 n 的最小好进制。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入："13"
 * 输出："3"
 * 解释：13 的 3 进制是 111。
 *
 *
 * 示例 2：
 *
 *
 * 输入："4681"
 * 输出："8"
 * 解释：4681 的 8 进制是 11111。
 *
 *
 * 示例 3：
 *
 *
 * 输入："1000000000000000000"
 * 输出："999999999999999999"
 * 解释：1000000000000000000 的 999999999999999999 进制是 11。
 *
 *
 *
 *
 * 提示：
 *
 *
 * n的取值范围是 [3, 10^18]。
 * 输入总是有效且没有前导 0。
 *
 *
 *
 *
 */

package leetcode

import (
	"math"
	"math/bits"
	"strconv"
)

// @lc code=start
// math
func smallestGoodBase(n string) string {
	nVal, _ := strconv.Atoi(n)
	mMax := bits.Len(uint(nVal)) - 1
	for m := mMax; m > 1; m-- {
		k := int(math.Pow(float64(nVal), 1/float64(m)))
		mul, sum := 1, 1
		for i := 0; i < m; i++ {
			mul *= k
			sum += mul
		}
		if sum == nVal {
			return strconv.Itoa(k)
		}
	}
	return strconv.Itoa(nVal - 1)

}

// @lc code=end
