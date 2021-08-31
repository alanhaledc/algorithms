/*
 * @lc app=leetcode.cn id=1109 lang=golang
 *
 * [1109] 航班预订统计
 *
 * https://leetcode-cn.com/problems/corporate-flight-bookings/description/
 *
 * algorithms
 * Medium (50.92%)
 * Likes:    198
 * Dislikes: 0
 * Total Accepted:    36.2K
 * Total Submissions: 67.8K
 * Testcase Example:  '[[1,2,10],[2,3,20],[2,5,25]]\n5'
 *
 * 这里有 n 个航班，它们分别从 1 到 n 进行编号。
 *
 * 有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi] 意味着在从
 * firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。
 *
 * 请你返回一个长度为 n 的数组 answer，其中 answer[i] 是航班 i 上预订的座位总数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
 * 输出：[10,55,45,25,25]
 * 解释：
 * 航班编号        1   2   3   4   5
 * 预订记录 1 ：   10  10
 * 预订记录 2 ：       20  20
 * 预订记录 3 ：       25  25  25  25
 * 总座位数：      10  55  45  25  25
 * 因此，answer = [10,55,45,25,25]
 *
 *
 * 示例 2：
 *
 *
 * 输入：bookings = [[1,2,10],[2,2,15]], n = 2
 * 输出：[10,25]
 * 解释：
 * 航班编号        1   2
 * 预订记录 1 ：   10  10
 * 预订记录 2 ：       15
 * 总座位数：      10  25
 * 因此，answer = [10,25]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 * bookings[i].length == 3
 * 1 i i
 * 1 i
 *
 *
 */

package leetcode

// @lc code=start
// force
func corpFlightBookings2(bookings [][]int, n int) []int {
	ret := make([]int, n)
	for _, booking := range bookings {
		first, last, seat := booking[0], booking[1], booking[2]
		for i := first - 1; i < last; i++ {
			ret[i] += seat
		}
	}
	return ret
}

// 差分
func corpFlightBookings(bookings [][]int, n int) []int {
	ret := make([]int, n)
	for _, booking := range bookings {
		first, last, seat := booking[0], booking[1], booking[2]
		ret[first-1] += seat
		if last < n {
			ret[last] -= seat
		}
	}
	for i := 1; i < n; i++ {
		ret[i] += ret[i-1]
	}
	return ret
}

// @lc code=end
