/*
 * @lc app=leetcode.cn id=94 lang=javascript
 *
 * [94] 二叉树的中序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[]}
 */
var inorderTraversal = function (root) {
  let res = [];
  dfs(root, res);
  return res;
};

function dfs(root, res) {
  if (!root) {
    return;
  }
  dfs(root.left, res);
  res.push(root.val);
  dfs(root.right, res);
}
// @lc code=end

var inorderTraversal1 = function (root) {
  const stack = [];
  const vals = [];
  while (stack.length > 0 || root) {
    while (root) {
      stack.push(root);
      root = root.left;
    }
    root = stack.pop();
    vals.push(root.val);
    root = root.right;
  }
  return vals;
};
