/*
 * @lc app=leetcode.cn id=20 lang=javascript
 *
 * [20] 有效的括号
 */

// @lc code=start
/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function (s) {
  let stack = [];
  let map = {
    ")": "(",
    "]": "[",
    "}": "{",
  };
  for (const c of s) {
    if (map[c]) {
      if (stack.length === 0 || stack[stack.length - 1] !== map[c]) {
        return false;
      }
      stack.pop();
    } else {
      stack.push(c);
    }
  }
  return stack.length === 0;
};
// @lc code=end

var isValid1 = function (s) {
  let stack = [];
  for (const c of s) {
    if (c === "(" || c === "[" || c === "{") {
      [];
      stack.push(c);
    } else if (
      (c === ")" && stack.length > 0 && stack[stack.length - 1] === "(") ||
      (c === "]" && stack.length > 0 && stack[stack.length - 1] === "[") ||
      (c === "}" && stack.length > 0 && stack[stack.length - 1] === "{")
    ) {
      stack.pop();
    } else {
      return false;
    }
  }
  return stack.length === 0;
};
