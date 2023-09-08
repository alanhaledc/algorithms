/*
 * @lc app=leetcode.cn id=2651 lang=rust
 *
 * [2651] 计算列车到站时间
 *
 * https://leetcode.cn/problems/calculate-delayed-arrival-time/description/
 *
 * algorithms
 * Easy (83.34%)
 * Likes:    39
 * Dislikes: 0
 * Total Accepted:    18.8K
 * Total Submissions: 21.2K
 * Testcase Example:  '15\n5'
 *
 * 给你一个正整数 arrivalTime 表示列车正点到站的时间（单位：小时），另给你一个正整数 delayedTime 表示列车延误的小时数。
 *
 * 返回列车实际到站的时间。
 *
 * 注意，该问题中的时间采用 24 小时制。
 *
 *
 *
 * 示例 1：
 *
 * 输入：arrivalTime = 15, delayedTime = 5
 * 输出：20
 * 解释：列车正点到站时间是 15:00 ，延误 5 小时，所以列车实际到站的时间是 15 + 5 = 20（20:00）。
 *
 *
 * 示例 2：
 *
 * 输入：arrivalTime = 13, delayedTime = 11
 * 输出：0
 * 解释：列车正点到站时间是 13:00 ，延误 11 小时，所以列车实际到站的时间是 13 + 11 = 24（在 24 小时制中表示为 00:00
 * ，所以返回 0）。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= arrivaltime < 24
 * 1 <= delayedTime <= 24
 *
 *
 */

// @lc code=start
impl Solution {
    pub fn find_delayed_arrival_time(arrival_time: i32, delayed_time: i32) -> i32 {
        (arrival_time + delayed_time) % 24
    }
}
// @lc code=end
