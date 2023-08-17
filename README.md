# go_leetcode

## 二叉树难题
### 1. 迭代法
257. 二叉树的所有路径



### 2. 递归法
110. 平衡二叉树
513. 找树左下角的值
113. 路径总和 II

### 3. 注意点

```go
package main

import "fmt"

func main() {
	res := []int{1, 2}
	fmt.Printf("%p", res)  // 0xc00011c020
	res = append(res, 1)
	fmt.Printf("%p", res)  // 0xc000130000
	// append过后的res地址是不一样的，要想操作同一个res，必须通过指针操作
	// 见 113. 路径总和 II
}
```
