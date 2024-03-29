/*
 * @lc app=leetcode.cn id=1723 lang=golang
 *
 * [1723] 完成所有工作的最短时间
 *
 * https://leetcode-cn.com/problems/find-minimum-time-to-finish-all-jobs/description/
 *
 * algorithms
 * Hard (50.96%)
 * Likes:    215
 * Dislikes: 0
 * Total Accepted:    19.5K
 * Total Submissions: 38.3K
 * Testcase Example:  '[3,2,3]\n3'
 *
 * 给你一个整数数组 jobs ，其中 jobs[i] 是完成第 i 项工作要花费的时间。
 *
 * 请你将这些工作分配给 k 位工人。所有工作都应该分配给工人，且每项工作只能分配给一位工人。工人的 工作时间
 * 是完成分配给他们的所有工作花费时间的总和。请你设计一套最佳的工作分配方案，使工人的 最大工作时间 得以 最小化 。
 *
 * 返回分配方案中尽可能 最小 的 最大工作时间 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：jobs = [3,2,3], k = 3
 * 输出：3
 * 解释：给每位工人分配一项工作，最大工作时间是 3 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：jobs = [1,2,4,7,8], k = 2
 * 输出：11
 * 解释：按下述方式分配工作：
 * 1 号工人：1、2、8（工作时间 = 1 + 2 + 8 = 11）
 * 2 号工人：4、7（工作时间 = 4 + 7 = 11）
 * 最大工作时间是 11 。
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 *
 *
 */

package main

import "sort"

// @lc code=start
// binary search + backtracking
func minimumTimeRequired(jobs []int, k int) int {
	n := len(jobs)
	sort.Sort(sort.Reverse(sort.IntSlice(jobs)))
	l, r := jobs[0], 0
	for _, v := range jobs {
		r += v
	}

	return l + sort.Search(r-l, func(limit int) bool {
		limit += l
		workloads := make([]int, k)
		var backtrack func(int) bool
		backtrack = func(idx int) bool {
			if idx == n {
				return true
			}
			cur := jobs[idx]
			for i := range workloads {
				if workloads[i]+cur <= limit {
					workloads[i] += cur
					if backtrack(idx + 1) {
						return true
					}
					workloads[i] -= cur
				}

				if workloads[i] == 0 || workloads[i]+cur == limit {
					break
				}
			}
			return false
		}

		return backtrack(0)
	})
}

// @lc code=end
