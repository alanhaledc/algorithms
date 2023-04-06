/*
 * @lc app=leetcode.cn id=1017 lang=rust
 *
 * [1017] 负二进制转换
 *
 * https://leetcode.cn/problems/convert-to-base-2/description/
 *
 * algorithms
 * Medium (57.26%)
 * Likes:    91
 * Dislikes: 0
 * Total Accepted:    8.9K
 * Total Submissions: 14.4K
 * Testcase Example:  '2'
 *
 * 给你一个整数 n ，以二进制字符串的形式返回该整数的 负二进制（base -2）表示。
 *
 * 注意，除非字符串就是 "0"，否则返回的字符串中不能含有前导零。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 2
 * 输出："110"
 * 解释：(-2)^2 + (-2)^1 = 2
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 3
 * 输出："111"
 * 解释：(-2)^2 + (-2)^1 + (-2)^0 = 3
 *
 *
 * 示例 3：
 *
 *
 * 输入：n = 4
 * 输出："100"
 * 解释：(-2)^2 = 4
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= n <= 10^9
 *
 *
 */

// @lc code=start
impl Solution {
    pub fn base_neg2(mut n: i32) -> String {
        if n == 0 || n == 1 {
            return n.to_string();
        }
        let mut res = String::new();
        while n != 0 {
            let remainder = n & 1;
            res += &remainder.to_string();
            n -= remainder;
            n /= -2;
        }
        res.chars().rev().collect()
    }
}
// @lc code=end
