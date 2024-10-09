### 题目

给你三个整数 `n`, `k`, `m` 。定义 $S=\{i\mid 1\le i\le nm+k,i\in \mathbb Z\}$ 。

请返回一个下标从 **0** 开始、长度为 `m` 的数组 `answer`，其中 `answer[i]` 表示符合下列条件集合 `T` 的个数。

- 集合 `T` 是集合 `S` 的子集。
- 集合 `T` 中所有元素的和对 `m` 取余的值恰好为 `i` 。

由于答案可能很大，你只需要求出它模 `998244353` 的结果。


**示例 1：**
>输入：`n = 1, k = 1, m = 2`
>输出：`[4,4]`

**示例 2：**
>输入：`n = 1919, k = 8, m = 10`
>输出：`[577613260,577613260,822345879,577613260,822345879,577613260,577613260,822345879,577613260,822345879]`

**提示：**
- `0 <= n <= 10^13`
- `0 <= k < min(m, 500)`
- `1 <= m <= 10^5`
- `n * m + k > 0`

### 思路


```go 

```

### 复杂度分析