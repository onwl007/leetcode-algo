/*
 * @lc app=leetcode.cn id=169 lang=javascript
 *
 * [169] 多数元素
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @return {number}
 */
var majorityElement = function (nums) {
  nums.sort();
  return nums[Math.floor(nums.length / 2)];
};
// @lc code=end

var majorityElement1 = function (nums) {
  let map = new Map();
  for (const value of nums) {
    if (map.has(value)) {
      map.set(value, map.get(value) + 1);
    } else {
      map.set(value, 1);
    }
  }

  for (const [k, v] of map.entries()) {
    if (v > nums.length / 2) {
      return k;
    }
  }
  return -1;
};
