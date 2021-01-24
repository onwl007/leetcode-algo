# 二分查找

## 时间复杂度

二分查找是一种非常高效的算法，我们假设数据大小是 n，每次查找后数据都会缩小为原来的一半。

> 被查找区间的大小变化
> n, n/2, n/4, 8/n, ... n/2^k

这是一个等比数列，其中 n/2^k = 1 时，k 的值就是总共缩小的次数。而每一次缩小操作只涉及两个数据的大小比较，所以，经过了 k 次区间缩小操作，时间复杂度就是 O(k)，通过 n/2^k = 1，可求出 k = log2n，所以时间复杂度就是 O(logn)。

因为 logn 是一个非常“恐怖”的数量级，即便 n 非常非常大，对应的 logn 也很小。比如 2^32，这个数很大了吧？大约是 42 亿，也就是说我们在 42 亿个数据中用二分查找一个数据，最多需要比较 32 次。

## 二分查找容易出错的 3 个地方

1. 循环退出条件
   注意 `left <= right`，而不是 `left < right`，这块判断的条件主要根据 `right`的取值不同而变化
2. mid 的取值
   `mid = (left + right)/2`这种写法是有问题的。因为 `left` 和 `right` 比较大的话，两者之和就有可能发生溢出。改进的方法是将 mid 的计算方式写成 `left + (right - left)/2`，更进一步的优化可以将除法换成位运算。`left + ((right - left) >> 1)`
3. `left` 和 `right` 的更新
   `left = mid + 1`，`right = mid - 1`

## 二分查找应用场景的局限性

1. 二分查找依赖的是顺序表结构，简单点说就是数组
2. 二分查找针对的是有序数据
3. 数据量太小不适合二分查找
   太小就跟遍历查找没啥区别了
4. 数据量太大也不适合二分查护
   底层依赖数组，数组为了支持随机访问的特性，要求内存空间连续。注意这里的“连续”二字，当内存中的空间都是零散的，那就无法申请一个连续的符合条件的内存空间，所以也就不能使用二分查找了。

## 4 种常见的二分查找变形问题

1. 查找第一个值等于给定值的元素，也就是查找左侧边界

```go
func bsearch(nums []int, target int) int {
  left, right := 0, len(nums)-1 // 右边界赋值时需注意
  for left <= right { // 注意
    mid := left + (right - left)/2 // 注意，这样写是为了防止左右边界直接相加太大导致整型溢出
    if target > nums[mid] {
      left = mid + 1
    } else if target < nums[mid] {
      right = mid - 1
    } else if target == nums[mid] {
      if mid == 0 || nums[mid-1] != target {
        return mid
      } else {
        right = mid - 1
      }
    }
  }
  return -1
}
```

2. 查找最后一个值等于给定值的元素，也就是右侧边界

```go
func bsearch(nums []int, target int) int {
  left, right := 0, len(nums)-1 // 右边界赋值时需注意
  for left <= right { // 注意
    mid := left + (right - left)/2 // 注意，这样写是为了防止左右边界直接相加太大导致整型溢出
    if target > nums[mid] {
      left = mid + 1
    } else if target < nums[mid] {
      right = mid - 1
    } else if target == nums[mid] {
      if mid == len(nums)-1 || nums[mid+1] != target {
        return mid
      } else {
        left = mid + 1
      }
    }
  }
  return -1
}
```

3. 查找第一个大于等于给定值的元素

```go
func bsearch(nums []int, target int) int {
  left, right := 0, len(nums)-1 // 右边界赋值时需注意
  for left <= right { // 注意
    mid := left + (right - left)/2 // 注意，这样写是为了防止左右边界直接相加太大导致整型溢出
    if nums[mid] >= target {
      if mid == 0 || nums[mid-1] < target {
        return mid
      } else {
        right = mid - 1
      }
    } else {
      left = mid + 1
    }
  }
  return -1
}
```

4. 查找最后一个小于等于给定值的元素

```go
func bsearch(nums []int, target int) int {
  left, right := 0, len(nums)-1 // 右边界赋值时需注意
  for left <= right { // 注意
    mid := left + (right - left)/2 // 注意，这样写是为了防止左右边界直接相加太大导致整型溢出
    if nums[mid] > target {
      right = mid - 1
    } else {
      if mid == len(nums) - 1 || nums[mid+1] > traget {
        return mid
      } else {
        left = mid + 1
      }
    }
  }
  return -1
}
```

## 标准二分查找

```go
func binarySearch(nums []int, target int) int {
  left, right := 0, len(nums)-1 // 右边界赋值时需注意
  for left <= right { // 注意
    mid := left + (right - left)/2 // 注意，这样写是为了防止左右边界直接相加太大导致整型溢出
    if target == nums[mid] {
      return mid
    } else if target > nums[mid] {
      left = mid + 1
    } else if target < nums[mid] {
      right = mid - 1
    }
  }
  return -1
}
```

> 分析二分查找的一个技巧就是：不要出现 else，而是把所有情况用 else if 写清楚，这样可以清楚地展现所有细节

1. 为什么 for 循环的条件是 <=，而不是 < ？
   因为初始化 right 的赋值是 len(nums)-1，即最后一个元素的索引，而不是 len(nums)

   这二者的区别是，前者相当于在闭区间 [left, right] 中搜索，后者相当于在左闭右开区间 [left, right)，因为索引大小为 len(nums) 是越界的。

2. 缺陷
   比如有序数组`nums = [1,2,2,2,3]`，`target` 为 2，此算法返回的索引是 2，但是如果我想得到 `target` 的左侧边界或右侧边界，这样的话此算法是无法处理的

## 寻找左侧边界的二分搜索

```go
func left_bound(nums []int, target int) int {
  left, right := 0, len(nums)-1
  // 搜索区间为 [left, right]
  for left <= right {
    mid := left + (right - left)/2
    if target == nums[mid] {
      right = mid - 1 // 收缩右侧边界
    } else if target > nums[mid] {
      // 搜索区间变为 [mid+1, right]
      left = mid + 1
    } else if target < nums[mid] {
      // 搜索区间变为 [left, mid-1]
      right = mid - 1
    }
  }
  // 由于循环退出条件是 left == right + 1，所以当 target 比 nums 中所有元素都大时，会使得索引越界
  if left >= len(nums) || nums[left] != target {
    return -1
  }
  return left
}
```

1. 为什么该算法能够搜索左侧边界
   关键在于对于 `target == nums[mid]` 这种情况的处理

   ```go
   if target == nums[mid] {
     right = mid - 1
   }
   ```

## 寻找右侧边界的二分查找

```go
func right_bound(nums []int, target int) int {
  left, right = 0, len(nums)-1
  for left <= right {
    mid := left + (right - left)/2
    if target == nums[mid] {
      // 收缩左侧边界
      left = mid + 1
    } else if target > nums[mid] {
      left = mid + 1
    } else if target < nums[mid] {
      right = mid - 1
    }
  }
  // 这里检查 right 的越界情况，当 target 比所有元素都小时，right 会被见减成 -1，所以需要在最后防止越界
  if right < 0 || nums[right] != target {
    return -1
  }
  return right
}
```

1. 为什么该算法能够找到右侧边界
   关键点还是在这里

```go
if target == nums[mid] {
  left = mid + 1
}
```

## 二分查找的另一种书写形式

```go
func binary_search(nums []int, target int) int {
  left, right := 0, len(nums)
  for left < right {
    mid := left + (right - left)/2
    if target == nums[mid] {
      return mid
    } else if target > nums[mid] {
      left = mid + 1
    } else if target < nums[mid] {
      right = mid - 1
    }
  }
  return nums[left] == target ? left : -1
}

// 寻找左侧边界
func left_bound(nums []int, target int) int {
  left, right := 0, len(nums)
  for left < right {
    mid := left + (right - left)/2
    if target == nums[mid] {
      right = mid
    } else if target > nums[mid] {
      left = mid + 1
    } else if target < nums[mid] {
      right = mid
    }
  }
  // target 比所有数都大，当不存在时，因为是向左侧搜索，left 指向最后一个元素
  if left == len(nums) {
    return -1
  }
  return nums[left] == target ? left : -1
}

// 寻找右侧边界
func right_bound(nums []int, target int) int {
  left, right := 0, len(nums)
  for left < right {
    mid := left + (right - left)/2
    if target == nums[mid] {
      left = mid + 1
    } else if target > nums[mid] {
      left = mid + 1
    } else if target < nums[mid] {
      right = mid
    }
  }
  if left == 0 {
    return -1
  }
  return nums[left-1] == target ? left - 1 : -1
}
```

## 总结

1. 分析二分查找代码时，不要出现 `else`，全部展开成 `else if` 方便理解
2. 注意「搜索区间」和循环的终止条件，如果存在漏掉的元素，记得在最后检查
3. 如需定义左闭右开的「搜索区间」搜索左右边界，只要在 `nums[mid] == target` 时做修改即可，搜索右侧时需要减一
4. 如果将「搜索区间」全部统一成两端都闭的闭区间，只要稍改 `nums[mid] == target` 条件处的代码和返回逻辑即可
