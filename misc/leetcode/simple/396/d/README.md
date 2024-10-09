#### 题目

<p>给你一个整数数组&nbsp;<code>nums</code>&nbsp;和两个整数&nbsp;<code>cost1</code> 和&nbsp;<code>cost2</code>&nbsp;。你可以执行以下&nbsp;<strong>任一</strong>&nbsp;操作&nbsp;<strong>任意</strong>&nbsp;次：</p>

<ul>
	<li>从 <code>nums</code>&nbsp;中选择下标 <code>i</code>&nbsp;并且将 <code>nums[i]</code>&nbsp;<strong>增加</strong> <code>1</code>&nbsp;，开销为 <code>cost1</code>。</li>
	<li>选择 <code>nums</code>&nbsp;中两个 <strong>不同</strong>&nbsp;下标 <code>i</code>&nbsp;和 <code>j</code>&nbsp;，并且将 <code>nums[i]</code>&nbsp;和 <code>nums[j]</code>&nbsp;都&nbsp;<strong>增加</strong> <code>1</code>&nbsp;，开销为 <code>cost2</code>&nbsp;。</li>
</ul>

<p>你的目标是使数组中所有元素都 <strong>相等</strong>&nbsp;，请你返回需要的 <strong>最小开销</strong>&nbsp;之和。</p>

<p>由于答案可能会很大，请你将它对&nbsp;<code>10<sup>9</sup> + 7</code>&nbsp;<strong>取余</strong>&nbsp;后返回。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [4,1], cost1 = 5, cost2 = 2</span></p>

<p><span class="example-io"><b>输出：</b>15</span></p>

<p><strong>解释：</strong></p>

<p>执行以下操作可以使数组中所有元素相等：</p>

<ul>
	<li>将&nbsp;<code>nums[1]</code>&nbsp;增加 1 ，开销为 5 ，<code>nums</code> 变为&nbsp;<code>[4,2]</code>&nbsp;。</li>
	<li>将&nbsp;<code>nums[1]</code>&nbsp;增加 1 ，开销为 5 ，<code>nums</code> 变为&nbsp;<code>[4,3]</code>&nbsp;。</li>
	<li>将&nbsp;<code>nums[1]</code>&nbsp;增加 1 ，开销为 5 ，<code>nums</code> 变为&nbsp;<code>[4,4]</code>&nbsp;。</li>
</ul>

<p>总开销为 15 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums = [2,3,3,3,5], cost1 = 2, cost2 = 1</span></p>

<p><span class="example-io"><b>输出：</b>6</span></p>

<p><b>解释：</b></p>

<p>执行以下操作可以使数组中所有元素相等：</p>

<ul>
	<li>将&nbsp;<code>nums[0]</code> 和&nbsp;<code>nums[1]</code>&nbsp;同时增加 1 ，开销为 1 ，<code>nums</code> 变为&nbsp;<code>[3,4,3,3,5]</code>&nbsp;。</li>
	<li>将&nbsp;<code>nums[0]</code> 和&nbsp;<code>nums[2]</code>&nbsp;同时增加 1 ，开销为 1 ，<code>nums</code> 变为&nbsp;<code>[4,4,4,3,5]</code>&nbsp;。</li>
	<li>将&nbsp;<code>nums[0]</code> 和&nbsp;<code>nums[3]</code>&nbsp;同时增加 1 ，开销为 1 ，<code>nums</code> 变为&nbsp;<code>[5,4,4,4,5]</code>&nbsp;。</li>
	<li>将&nbsp;<code>nums[1]</code> 和&nbsp;<code>nums[2]</code>&nbsp;同时增加 1 ，开销为 1 ，<code>nums</code> 变为&nbsp;<code>[5,5,5,4,5]</code>&nbsp;。</li>
	<li>将&nbsp;<code>nums[3]</code>&nbsp;增加 1 ，开销为 2 ，<code>nums</code> 变为&nbsp;<code>[5,5,5,5,5]</code>&nbsp;。</li>
</ul>

<p>总开销为 6 。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [3,5,3], cost1 = 1, cost2 = 3</span></p>

<p><span class="example-io"><b>输出：</b>4</span></p>

<p><strong>解释：</strong></p>

<p>执行以下操作可以使数组中所有元素相等：</p>

<ul>
	<li>将&nbsp;<code>nums[0]</code>&nbsp;增加 1 ，开销为 1 ，<code>nums</code> 变为&nbsp;<code>[4,5,3]</code>&nbsp;。</li>
	<li>将&nbsp;<code>nums[0]</code>&nbsp;增加 1 ，开销为 1 ，<code>nums</code> 变为&nbsp;<code>[5,5,3]</code>&nbsp;。</li>
	<li>将&nbsp;<code>nums[2]</code>&nbsp;增加 1 ，开销为 1 ，<code>nums</code> 变为&nbsp;<code>[5,5,4]</code>&nbsp;。</li>
	<li>将&nbsp;<code>nums[2]</code>&nbsp;增加 1 ，开销为 1 ，<code>nums</code> 变为&nbsp;<code>[5,5,5]</code>&nbsp;。</li>
</ul>

<p>总开销为 4 。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>6</sup></code></li>
	<li><code>1 &lt;= cost1 &lt;= 10<sup>6</sup></code></li>
	<li><code>1 &lt;= cost2 &lt;= 10<sup>6</sup></code></li>
</ul>

#### 思路

为方便描述，将 $\textit{cost}_1$ 和 $\textit{cost}_2$ 简记为 $c_1$ 和 $c_2$。

设 $m = \min(\textit{nums}),\ M=\max(\textit{nums})$。本题看上去把所有数都变成 $M$ 即可，但请看 $[1,3,4,4]$ 这个例子：

- 全部变成 $4$：用操作二把数字 $1$ 和 $3$ 都加一，现在数组为 $[2,4,4,4]$；然后用操作一把数字 $2$ 加二，现在数组为 $[4,4,4,4]$。总开销为 $2c_1 + c_2$。
- 全部变成 $5$：只用操作二，把数字 $1$ 和 $3$ 加一两次，现在数组为 $[3,5,4,4]$；然后把数字 $3$ 和 $4$ 加一，现在数组为 $[4,5,5,4]$；然后把数字 $4$ 和 $4$ 加一，现在数组为 $[5,5,5,5]$。总开销为 $4c_2$。

这意味着，在 $c_1$ 很大 $c_2$ 很小的情况下（比如 $c_1=999,\ c_2=1$），变成一个比 $M$ 更大的数，可能得到一个更小的总开销。

假设都变成 $x\ (x\ge M)$，那么当 $x$ 取多少时，可以得到最小的总开销？

要枚举 $x$ 吗？假如 $x=10$，具体要如何操作呢？知道怎么操作就知道怎么算总开销。

## 情况一

如果 $n\le 2$，那么只能用操作一，最小总开销为

$$
(M-m)\cdot c_1
$$

下面的讨论中，$n\ge 3$。

## 情况二

如果 $2c_1\le c_2$，用一次操作二不如用两次操作一，所以只需用操作一，把每个数都变成 $M$ 的总开销是最小的。

总共需要执行

$$
\begin{align}
& (M - \textit{nums}[0]) + (M - \textit{nums}[1]) + \cdots + (M - \textit{nums}[n-1])\\
=\ &nM - (\textit{nums}[0] + \textit{nums}[1] + \cdots + \textit{nums}[n-1])
\end{align}
$$

次操作一。

设

$$
\textit{base} = nM - (\textit{nums}[0] + \textit{nums}[1] + \cdots + \textit{nums}[n-1])
$$

那么最小总开销为

$$
\textit{base}\cdot c_1
$$

## 情况三

如果 $2c_1> c_2$，用操作二更划算，所以应当**尽量多地使用操作二**。

假设都变成 $x\ (x\ge M)$，那么所有数都需要在 $M$ 的基础上额外增加 $x-M$，总共要增加

$$
s = \textit{base} + (x - M) \cdot n
$$

具体要如何操作呢？

为方便大家理解，想象有 $n$ 个盒子，第 $i$ 个盒子装有 $x-\textit{nums}[i]$ 个小球，这是 $\textit{nums}[i]$ 需要加一的次数。总共有 $s$ 个小球。

操作一相当于从一个非空盒子中取出一个球，操作二相当于从两个不同的非空盒子中各取一个小球。

例如有三个盒子，分别装有 $3,3,2$ 个小球：

- 从第 $1,2$ 个盒子中各取一个小球，剩余个数为 $2,2,2$。
- 从第 $1,2$ 个盒子中各取一个小球，剩余个数为 $1,1,2$。
- 从第 $1,3$ 个盒子中各取一个小球，剩余个数为 $0,1,1$。
- 从第 $2,3$ 个盒子中各取一个小球，剩余个数为 $0,0,0$。

注意，如果三个盒子分别装有 $3,3,3$ 个小球，最后还剩下一个小球，只能用操作一取出。

看上去，可以执行 $\left\lfloor\dfrac{s}{2}\right\rfloor$ 次操作二，再加上 $s\bmod 2$ 次操作一。

但这是不对的，例如有三个盒子，分别装有 $4,1,1$ 个小球：

- 从第 $1,2$ 个盒子中各取一个小球，剩余个数为 $3,0,1$。
- 从第 $1,3$ 个盒子中各取一个小球，剩余个数为 $2,0,0$。
- 无法继续用操作二，只能用操作一。

看上去，如果一个盒子装有很多小球，比其它盒子的小球总数还多，就没法执行 $\left\lfloor\dfrac{s}{2}\right\rfloor$ 次操作二了。怎么证明这个观察？

以 $3,3,2$ 为例，总共有 $3+3+2=8$ 个位置，我们可以按照 $1,3,5,7,2,4,6,8$ 的顺序，首先填入三个 $1$，得到

$$
1,\text{\_},1,\text{\_},1,\text{\_},\text{\_},\text{\_}
$$

继续按照顺序，填入三个 $2$，得到

$$
1,2,1,2,1,\text{\_},2,\text{\_}
$$

最后填入两个 $3$，得到

$$
1,2,1,2,1,3,2,3
$$

两两一组，得到

$$
(1,2),(1,2),(1,3),(2,3)
$$

这正是上面取球例子中的取法。

一个盒子最多有 $d = x - m$ 个小球，分类讨论：

- 如果 $d\le s- d$，即 $d\le \left\lfloor\dfrac{s}{2}\right\rfloor$，按照上述构造方式，同一组中不可能有相同元素。这意味着，如果 $d\le s- d$，只用操作二，最后至多剩下一个小球（如果 $s$ 是奇数的话）。
- 如果 $d> s- d$，即使每次都用操作二，从这个装有 $d$ 个小球的盒子中取球，在 $s-d$ 次操作二后，其余盒子中的总共 $s-d$ 个小球全部取完，后面只能用操作一取走这个盒子中剩余的 $d-(s-d)$ 个小球。
- 这两种情况都做到了**尽量多地使用操作二**。

按照上述讨论，计算都变成 $x$ 的总开销 $f(x)$：

- 如果 $2d\le s$，那么先执行 $\left\lfloor\dfrac{s}{2}\right\rfloor$ 次操作二，然后执行 $s\bmod 2$ 次操作一，总开销为
  $$
  f(x) = \left\lfloor\dfrac{s}{2}\right\rfloor\cdot c_2 + s\bmod 2\cdot c_1
  $$
- 如果 $2d> s$，那么先执行 $s-d$ 次操作二，然后执行 $2d-s$ 次操作一，总开销为
  $$
  f(x) = (s-d)\cdot c_2 + (2d-s)\cdot c_1
  $$

枚举 $x$，取 $f(x)$ 的最小值，即为最小总开销。

## 至多枚举到哪？

最后还剩下一个问题：$x$ 至多要枚举到哪个数为止？

注意到，$x$ 增加 $1$，$d$ 只增大 $1$，而 $s$ 增大 $n$。所以 $n>1$ 时我们会先满足 $2d>s$，然后再满足 $2d\le s$。而当 $2d\le s$ 时，总开销只和 $s$ 有关，$x$ 越大 $s$ 越大，至多再枚举一次就无需继续枚举了。至多再枚举一次是因为如果 $s$ 是奇数，额外有 $c_1$ 的开销。多枚举一次可以让 $s$ 从奇数变成偶数，可能得到更小的总开销。

$2d\le s$ 即

$$
2(x-m)\le \textit{base} + (x - M) \cdot n
$$

即

$$
(n-2)x\ge nM-2m-\textit{base}
$$

即

$$
x\ge \left\lceil\dfrac{nM-2m-\textit{base}}{n-2}\right\rceil
$$

所以 $x$ 至多枚举到

$$
\max\left(\left\lceil\dfrac{nM-2m-\textit{base}}{n-2}\right\rceil, M\right) + 1
$$

加一可以保证我们一定可以枚举到 $s$ 是偶数的情况。

> 注意当 $n\ge 3$ 时，上式 $\le 2M$，所以枚举到 $2M$ 的做法是正确的。

## 优化

![w396d-c.png](https://pic.leetcode.cn/1714913854-WbqXIt-w396d-c.png)

如上图所示：

- $2d>s$ 对应 $(s-d)\cdot c_2 + (2d-s)\cdot c_1$，可以视作一条斜率或正或负（还可能为零）的一次函数。
- $2d\le s$ 对应 $\left\lfloor\dfrac{s}{2}\right\rfloor\cdot c_2 + s\bmod 2\cdot c_1$，根据 $s$ 的奇偶性，可以视作两条斜率为正的一次函数。

如果函数交点横坐标 $\le M$，$x$ 只需枚举 $M$ 和 $M+1$；否则 $x$ 应当枚举交点横坐标附近的数（右图），以及 $M$（左图）。

设 $i = \left\lceil\dfrac{nM-2m-\textit{base}}{n-2}\right\rceil = \left\lfloor\dfrac{nM-2m-\textit{base}+n-3}{n-2}\right\rfloor$，最小总开销为

$$
\begin{cases}
\min(f(M),f(M+1)),&i\le M\\
\min(f(M),f(i-1),f(i),f(i+1)),&i> M
\end{cases}
$$


```
func minCostToEqualizeArray(nums []int, c1 int, c2 int) int {
	const mod = 1_000_000_007
	n := len(nums)
	m := slices.Min(nums)
	M := slices.Max(nums)
	base := n * M
	for _, x := range nums {
		base -= x
	}
	if n <= 2 || c1*2 <= c2 {
		return base * c1 % mod
	}

	f := func(x int) int {
		s := base + (x-M)*n
		d := x - m
		if d*2 <= s {
			return s/2*c2 + s%2*c1
		}
		return (s-d)*c2 + (d*2-s)*c1
	}

	i := (n*M - m*2 - base + n - 3) / (n - 2)
	if i <= M {
		return min(f(M), f(M+1)) % mod
	}
	return min(f(M), f(i-1), f(i), f(i+1)) % mod
}

```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)