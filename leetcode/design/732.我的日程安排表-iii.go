/*
 * @lc app=leetcode.cn id=732 lang=golang
 *
 * [732] 我的日程安排表 III
 *
 * https://leetcode.cn/problems/my-calendar-iii/description/
 *
 * algorithms
 * Hard (69.52%)
 * Likes:    113
 * Dislikes: 0
 * Total Accepted:    8.4K
 * Total Submissions: 12K
 * Testcase Example:  '["MyCalendarThree","book","book","book","book","book","book"]\n' +
  '[[],[10,20],[50,60],[10,40],[5,15],[5,10],[25,55]]'
 *
 * 当 k 个日程安排有一些时间上的交叉时（例如 k 个日程安排都在同一时间内），就会产生 k 次预订。
 *
 * 给你一些日程安排 [start, end) ，请你在每个日程安排添加后，返回一个整数 k ，表示所有先前日程安排会产生的最大 k 次预订。
 *
 * 实现一个 MyCalendarThree 类来存放你的日程安排，你可以一直添加新的日程安排。
 *
 *
 * MyCalendarThree() 初始化对象。
 * int book(int start, int end) 返回一个整数 k ，表示日历中存在的 k 次预订的最大值。
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 * ["MyCalendarThree", "book", "book", "book", "book", "book", "book"]
 * [[], [10, 20], [50, 60], [10, 40], [5, 15], [5, 10], [25, 55]]
 * 输出：
 * [null, 1, 1, 2, 3, 3, 3]
 *
 * 解释：
 * MyCalendarThree myCalendarThree = new MyCalendarThree();
 * myCalendarThree.book(10, 20); // 返回 1 ，第一个日程安排可以预订并且不存在相交，所以最大 k 次预订是 1 次预订。
 * myCalendarThree.book(50, 60); // 返回 1 ，第二个日程安排可以预订并且不存在相交，所以最大 k 次预订是 1 次预订。
 * myCalendarThree.book(10, 40); // 返回 2 ，第三个日程安排 [10, 40) 与第一个日程安排相交，所以最大 k
 * 次预订是 2 次预订。
 * myCalendarThree.book(5, 15); // 返回 3 ，剩下的日程安排的最大 k 次预订是 3 次预订。
 * myCalendarThree.book(5, 10); // 返回 3
 * myCalendarThree.book(25, 55); // 返回 3
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * 每个测试用例，调用 book 函数最多不超过 400次
 *
 *
*/

package leetcode

import "sort"

// @lc code=start
type MyCalendarThree struct {
	calendar map[int]int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{
		calendar: make(map[int]int),
	}
}

func (mct *MyCalendarThree) Book(start int, end int) int {
	mct.calendar[start]++
	mct.calendar[end]--
	calenderArr := make([][]int, 0)
	for k, v := range mct.calendar {
		calenderArr = append(calenderArr, []int{k, v})
	}
	sort.Slice(calenderArr, func(i, j int) bool {
		return calenderArr[i][0] < calenderArr[j][0]
	})
	crossCount := 1
	prefixSum := 0

	for _, v := range calenderArr {
		prefixSum += v[1]
		if prefixSum > crossCount {
			crossCount = prefixSum
		}
	}
	return crossCount
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
// @lc code=end
