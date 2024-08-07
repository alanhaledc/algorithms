/*
 * @lc app=leetcode.cn id=3115 lang=golang
 *
 * [3115] 质数的最大距离
 *
 * https://leetcode.cn/problems/maximum-prime-difference/description/
 *
 * algorithms
 * Medium (57.75%)
 * Likes:    17
 * Dislikes: 0
 * Total Accepted:    13.6K
 * Total Submissions: 21.4K
 * Testcase Example:  '[4,2,9,5,3]'
 *
 * 给你一个整数数组 nums。
 *
 * 返回两个（不一定不同的）质数在 nums 中 下标 的 最大距离。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入： nums = [4,2,9,5,3]
 *
 * 输出： 3
 *
 * 解释： nums[1]、nums[3] 和 nums[4] 是质数。因此答案是 |4 - 1| = 3。
 *
 *
 * 示例 2：
 *
 *
 * 输入： nums = [4,8,2,8]
 *
 * 输出： 0
 *
 * 解释： nums[2] 是质数。因为只有一个质数，所以答案是 |2 - 2| = 0。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 3 * 10^5
 * 1 <= nums[i] <= 100
 * 输入保证 nums 中至少有一个质数。
 *
 *
 */

package leetcode

// @lc code=start
func maximumPrimeDifference(nums []int) int {
	first, last := -1, -1
	for i, num := range nums {
		if isPrime(num) {
			if first == -1 {
				first = i
			}
			last = i
		}
	}
	return last - first
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// @lc code=end
