#### 题目

<p>给你一个长度为 <code>n</code> 的整数数组 <code>nums</code>。</p>

<p>子数组 <code>nums[l..r]</code>（其中 <code>0 &lt;= l &lt;= r &lt; n</code>）的 <strong>成本 </strong>定义为：</p>

<p><code>cost(l, r) = nums[l] - nums[l + 1] + ... + nums[r] * (−1)<sup>r − l</sup></code></p>

<p>你的任务是将 <code>nums</code> 分割成若干子数组，使得所有子数组的成本之和 <strong>最大化</strong>，并确保每个元素 <strong>正好 </strong>属于一个子数组。</p>

<p>具体来说，如果 <code>nums</code> 被分割成 <code>k</code> 个子数组，且分割点为索引 <code>i<sub>1</sub>, i<sub>2</sub>, ..., i<sub>k − 1</sub></code>（其中 <code>0 &lt;= i<sub>1</sub> &lt; i<sub>2</sub> &lt; ... &lt; i<sub>k - 1</sub> &lt; n - 1</code>），则总成本为：</p>

<p><code>cost(0, i<sub>1</sub>) + cost(i<sub>1</sub> + 1, i<sub>2</sub>) + ... + cost(i<sub>k − 1</sub> + 1, n − 1)</code></p>

<p>返回在最优分割方式下的子数组成本之和的最大值。</p>

<p><strong>注意：</strong>如果 <code>nums</code> 没有被分割，即 <code>k = 1</code>，则总成本即为 <code>cost(0, n - 1)</code>。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [1,-2,3,4]</span></p>

<p><strong>输出：</strong> <span class="example-io">10</span></p>

<p><strong>解释：</strong></p>

<p>一种总成本最大化的方法是将 <code>[1, -2, 3, 4]</code> 分割成子数组 <code>[1, -2, 3]</code> 和 <code>[4]</code>。总成本为 <code>(1 + 2 + 3) + 4 = 10</code>。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [1,-1,1,-1]</span></p>

<p><strong>输出：</strong> <span class="example-io">4</span></p>

<p><strong>解释：</strong></p>

<p>一种总成本最大化的方法是将 <code>[1, -1, 1, -1]</code> 分割成子数组 <code>[1, -1]</code> 和 <code>[1, -1]</code>。总成本为 <code>(1 + 1) + (1 + 1) = 4</code>。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [0]</span></p>

<p><strong>输出：</strong> 0</p>

<p><strong>解释：</strong></p>

<p>无法进一步分割数组，因此答案为 0。</p>
</div>

<p><strong class="example">示例 4：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [1,-1]</span></p>

<p><strong>输出：</strong> <span class="example-io">2</span></p>

<p><strong>解释：</strong></p>

<p>选择整个数组，总成本为 <code>1 + 1 = 2</code>，这是可能的最大成本。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>-10<sup>9</sup> &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
</ul>

#### 思路

## 一、寻找子问题

为方便描述，下文将 $\textit{nums}$ 简记为 $a$。

对于示例 1，$a[0]$ 一定不用变号（取相反数），我们从 $a[1]$ 开始思考。

分类讨论：

- 分割，把 $a[1]$ 作为子数组的第一个数，接下来需要解决的问题为：考虑 $a[2]$ 到 $a[n-1]$，且 $a[2]$ 如果不是子数组的第一个数，则需要变号，在这种情况下的最大成本和。
- 不分割，那么 $a[1]$ 需要变号，接下来需要解决的问题为：考虑 $a[2]$ 到 $a[n-1]$，且 $a[2]$ 如果不是子数组的第一个数，则无需变号，在这种情况下的最大成本和。

由于分割或不分割都会把原问题变成一个**和原问题相似的、规模更小的子问题**，所以可以用**递归**解决。

> 注：动态规划有「选或不选」和「枚举选哪个」两种基本思考方式。在做题时，可根据题目要求，选择适合题目的一种来思考。本题是「分割」这个操作「选或不选」。

## 二、状态定义与状态转移方程

因为要解决的问题都形如「考虑 $a[i]$ 到 $a[n-1]$，如果 $a[i]$ 不是子数组的第一个数，则不需要/需要变号，在这种情况下的最大成本和」，所以用它作为本题的状态定义 $\textit{dfs}(i,j)$。其中 $j=0$ 表示不变号，$j=1$ 表示变号。

定义

$$
x =
\begin{cases}
a[i],\ &j=0\\
-a[i],\ &j=1
\end{cases}
$$

分类讨论：

- 分割，把 $a[i]$ 作为子数组的第一个数，接下来需要解决的问题为：考虑 $a[i+1]$ 到 $a[n-1]$，且 $a[i+1]$ 如果不是子数组的第一个数，则需要变号，在这种情况下的最大成本和，即 $\textit{dfs}(i,j) = \textit{dfs}(i+1, 1) + a[i]$。
- 不分割，接下来需要解决的问题为：考虑 $a[i+1]$ 到 $a[n-1]$，且 $a[i+1]$ 如果不是子数组的第一个数，则变号情况与 $a[i]$ 相反，在这种情况下的最大成本和，即 $\textit{dfs}(i,j) = \textit{dfs}(i+1, j\oplus 1) + x$。

这两种情况取最大值，就得到了 $\textit{dfs}(i,j)$，即

$$
\textit{dfs}(i,j) = \max(\textit{dfs}(i+1, 1) + a[i], \textit{dfs}(i+1, j\oplus 1) + x)
$$

递归边界：$\textit{dfs}(n,j)=0$。没有元素，成本和为 $0$。

递归入口：$\textit{dfs}(0,0)$，也就是答案。

## 三、递归搜索 + 保存递归返回值 = 记忆化搜索

考虑到整个递归过程中有大量重复递归调用（递归入参相同）。由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

**注意**：$\textit{memo}$ 数组的**初始值**一定不能等于要记忆化的值！例如初始值设置为 $0$，并且要记忆化的 $\textit{dfs}(i,j)$ 也等于 $0$，那就没法判断 $0$ 到底表示第一次遇到这个状态，还是表示之前遇到过了，从而导致记忆化失效。一般把初始值设置为 $-1$ 或 $-\infty$。

> Python 用户可以无视上面这段，直接用 `@cache` 装饰器。

```
func maximumTotalCost(a []int) int64 {
	n := len(a)
	memo := make([][2]int, n)
	for i := range memo {
		memo[i] = [2]int{math.MinInt, math.MinInt}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == n {
			return 0
		}
		p := &memo[i][j]
		if *p != math.MinInt { // 之前计算过
			return *p
		}
		res := dfs(i+1, 1) + a[i] // 分割
		r := dfs(i+1, j^1) // 不分割
		if j == 0 {
			r += a[i]
		} else {
			r -= a[i]
		}
		res = max(res, r)
		*p = res // 记忆化
		return res
	}
	return int64(dfs(0, 0))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $a$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以动态规划的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。保存多少状态，就需要多少空间。

## 四、1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[i][j]$ 的定义和 $\textit{dfs}(i,j)$ 的定义是一样的，都表示考虑 $a[i]$ 到 $a[n-1]$，如果 $a[i]$ 不是子数组的第一个数，则不需要/需要变号，在这种情况下的最大成本和。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[i][j] = \max(f[i+1][1] + a[i], f[i+1][j\oplus 1] + x)
$$

即

$$
\begin{align}
&f[i][0] = f[i+1][1] + a[i]\\
&f[i][1] = \max(f[i+1][1]+a[i], f[i+1][0]-a[i])\\
\end{align}
$$

初始值 $f[n][j]=0$，翻译自递归边界 $\textit{dfs}(n,j)=0$。

答案为 $f[0][0]$，翻译自递归入口 $\textit{dfs}(0,0)$。

```
func maximumTotalCost(a []int) int64 {
	n := len(a)
	f := make([][2]int, n+1)
	for i := n - 1; i >= 0; i-- {
		f[i][0] = f[i+1][1] + a[i]
		f[i][1] = max(f[i+1][1]+a[i], f[i+1][0]-a[i])
	}
	return int64(f[0][0])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $a$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 五、空间优化

观察上面的状态转移方程，在计算 $f[i]$ 时，只会用到 $f[i+1]$，不会用到比 $i+1$ 更大的状态。

因此可以像 [买卖股票的最佳时机](https://www.bilibili.com/video/BV1ho4y1W7QK/) 那样，反复利用两个变量。

状态转移方程改为

$$
\begin{align}
&f_0 = f_1 + a[i]\\
&f_1 = \max(f_1+a[i], f_0-a[i])\\
\end{align}
$$

注意这两个式子要**同时计算**。

初始值 $f_0=f_1=0$。

答案为 $f_0$。

```
func maximumTotalCost(a []int) int64 {
	f0, f1 := 0, 0
	for i := len(a) - 1; i >= 0; i-- {
		f0, f1 = f1+a[i], max(f1+a[i], f0-a[i])
	}
	return int64(f0)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $a$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。


## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
- [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
- [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)