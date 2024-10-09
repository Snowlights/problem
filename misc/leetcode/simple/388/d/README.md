#### 题目

<p>给你一个长度为 <code>n</code> 下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> 和一个 <strong>正奇数</strong> 整数 <code>k</code> 。</p>

<p><code>x</code> 个子数组的能量值定义为 <code>strength = sum[1] * x - sum[2] * (x - 1) + sum[3] * (x - 2) - sum[4] * (x - 3) + ... + sum[x] * 1</code> ，其中 <code>sum[i]</code> 是第 <code>i</code> 个子数组的和。更正式的，能量值是满足 <code>1 <= i <= x</code> 的所有 <code>i</code> 对应的 <code>(-1)<sup>i+1</sup> * sum[i] * (x - i + 1)</code> 之和。</p>

<p>你需要在 <code>nums</code> 中选择 <code>k</code> 个 <strong>不相交</strong><strong>子数组</strong> ，使得 <strong>能量值最大</strong> 。</p>

<p>请你返回可以得到的 <strong>最大</strong><strong>能量值</strong> 。</p>

<p><strong>注意</strong>，选出来的所有子数组 <strong>不</strong> 需要覆盖整个数组。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [1,2,3,-1,2], k = 3
<b>输出：</b>22
<b>解释：</b>选择 3 个子数组的最好方式是选择：nums[0..2] ，nums[3..3] 和 nums[4..4] 。能量值为 (1 + 2 + 3) * 3 - (-1) * 2 + 2 * 1 = 22 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>nums = [12,-2,-2,-2,-2], k = 5
<b>输出：</b>64
<b>解释：</b>唯一一种选 5 个不相交子数组的方案是：nums[0..0] ，nums[1..1] ，nums[2..2] ，nums[3..3] 和 nums[4..4] 。能量值为 12 * 5 - (-2) * 4 + (-2) * 3 - (-2) * 2 + (-2) * 1 = 64 。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<b>输入：</b>nums = [-1,-2,-3], k = 1
<b>输出：</b>-1
<b>解释：</b>选择 1 个子数组的最优方案是：nums[0..0] 。能量值为 -1 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= n <= 10<sup>4</sup></code></li>
	<li><code>-10<sup>9</sup> <= nums[i] <= 10<sup>9</sup></code></li>
	<li><code>1 <= k <= n</code></li>
	<li><code>1 <= n * k <= 10<sup>6</sup></code></li>
	<li><code>k</code> 是奇数。</li>
</ul>

#### 思路

## 题意补充

2. 子数组指的是非空连续子数组。
3. 这 $k$ 个子数组是从左到右的顺序，即题目中的 `sum[1]` 指的是最左边第一个子数组的元素和，`sum[2]` 指的是最左边第二个子数组的元素和，依此类推。

## 前置知识：前缀和

对于数组 $\textit{nums}$，定义它的前缀和 $\textit{s}[0]=0$，$\textit{s}[i+1] = \sum\limits_{j=0}^{i}\textit{nums}[j]$。

根据这个定义，有 $s[i+1]=s[i]+\textit{nums}[i]$。
例如 $\textit{nums}=[1,2,1,2]$，对应的前缀和数组为 $s=[0,1,3,4,6]$。
通过前缀和，我们可以把**连续子数组的元素和转换成两个前缀和的差**，$\textit{nums}[\textit{left}]$ 到 $\textit{nums}[\textit{right}]$ 的元素和等于

$$
\sum_{j=\textit{left}}^{\textit{right}}\textit{nums}[j] = \sum\limits_{j=0}^{\textit{right}}\textit{nums}[j] - \sum\limits_{j=0}^{\textit{left}-1}\textit{nums}[j] = \textit{s}[\textit{right}+1] - \textit{s}[\textit{left}]

$$

例如 $\textit{nums}$ 的子数组 $[2,1,2]$ 的和就可以用 $s[4]-s[1]=6-1=5$ 算出来。

**注**：为方便计算，常用左闭右开区间 $[\textit{left},\textit{right})$ 来表示从 $\textit{nums}[\textit{left}]$ 到 $\textit{nums}[\textit{right}-1]$ 的子数组，此时子数组的和为 $\textit{s}[\textit{right}] - \textit{s}[\textit{left}]$。
**注 2**：$s[0]=0$ 表示一个空数组的元素和。为什么要额外定义它？想一想，如果要计算的子数组恰好是一个前缀（从 $\textit{nums}[0]$ 到 $\textit{nums}[\textit{right}]$），你要用 $s[\textit{right}+1]$ 减去谁呢？通过定义 $s[0]=0$，任意子数组（包括前缀）都可以表示为两个前缀和的差。

本题是标准的划分型 DP（见文末的题单），通常来说，状态要包含**前缀长度**和**划分个数**。
对于本题，定义 $f[i][j]$ 表示从 $\textit{nums}[0]$ 到 $\textit{nums}[j-1]$ 中选出 $i$ 个不相交非空连续子数组的最大能量值。
设 $\textit{nums}$ 的前缀和数组为 $s$。

分类讨论：

- 不选 $\textit{nums}[j-1]$，问题变成从前 $j-1$ 个数中选 $i$ 个子数组，即 $f[i][j] = f[i][j-1]$。
- 把 $\textit{nums}[j-1]$ 作为第 $i$ 个子数组的最右元素，枚举第 $i$ 个子数组最左元素的下标，记作 $L$，那么子数组元素和为 $s[j] - s[L]$，所以有 $f[i][j] = f[i-1][L] + (s[j] - s[L])\cdot w$，其中 $w = (-1)^{i+1}(k-i+1)$。注意 $L$ 不能低于 $i-1$，因为前面 $i-1$ 个子数组至少要有 $1$ 个元素。$L$ 最大是 $j-1$，表示只选一个元素。

取转移来源的最大值，有

$$
f[i][j] = \max(f[i][j-1], \max\limits_{L=i-1}^{j-1} f[i-1][L] +  (s[j] - s[L])\cdot w)

$$

如果直接用上式计算，时间复杂度是 $\mathcal{O}(n^2k)$，会超时。
将上式变形为

$$
f[i][j] = \max(f[i][j-1], s[j]\cdot w + \max\limits_{L=i-1}^{j-1} f[i-1][L] - s[L]\cdot w)

$$

其中 $\max\limits_{L=i-1}^{j-1} f[i-1][L] - s[L]\cdot w$ 可以在枚举 $j$ 计算 $f[i][j]$ 的同时，用一个变量 $\textit{mx}$ 维护，即维护 $f[i - 1][j - 1] - s[j - 1] \cdot w$ 的最大值。

于是转移方程为

$$
f[i][j] = \max(f[i][j-1], s[j]\cdot w + \textit{mx})

$$

初始值：$f[0][j] = 0,\ f[i][i-1]=-\infty$。前者是因为不选子数组，能量值为 $0$，后者是因为无法从 $i-1$ 个数中选出 $i$ 个子数组。

答案：$f[k][n]$。

```go [sol]
func maximumStrength(nums []int, k int) int64 {
	n := len(nums)
	s := make([]int, n+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}
	f := make([][]int, k+1)
	f[0] = make([]int, n+1)
	for i := 1; i <= k; i++ {
		f[i] = make([]int, n+1)
		f[i][i-1] = math.MinInt
		mx := math.MinInt
		w := (k - i + 1) * (i%2*2 - 1)
		// j 不能太小也不能太大，要给前面留 i-1 个数，后面留 k-i 个数
		for j := i; j <= n-k+i; j++ {
			mx = max(mx, f[i-1][j-1]-s[j-1]*w)
			f[i][j] = max(f[i][j-1], s[j]*w+mx)
		}
	}
	return int64(f[k][n])
}
```

## 空间优化

可以去掉 $f$ 数组的第一个维度。为避免覆盖，用一个变量 $\textit{pre}$ 记录 $f[i-1][j-1]$。

```go [sol]
func maximumStrength(nums []int, k int) int64 {
	n := len(nums)
	s := make([]int, n+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}
	f := make([]int, n+1)
	for i := 1; i <= k; i++ {
		pre := f[i-1]
		f[i-1] = math.MinInt
		mx := math.MinInt
		w := (k - i + 1) * (i%2*2 - 1)
		for j := i; j <= n-k+i; j++ {
			mx = max(mx, pre-s[j-1]*w)
			pre = f[j]
			f[j] = max(f[j-1], s[j]*w+mx)
		}
	}
	return int64(f[n])
}
```


#### 复杂度分析

- 时间复杂度：$\mathcal{O}(k\cdot (n-k))$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。


## 题单：划分型 DP ①

将序列分成（恰好/至多）$k$ 个连续区间，求解与这些区间有关的最优值。
- [410. 分割数组的最大值](https://leetcode.cn/problems/split-array-largest-sum/)
- [813. 最大平均值和的分组](https://leetcode.cn/problems/largest-sum-of-averages/) 1937
- [1278. 分割回文串 III](https://leetcode.cn/problems/palindrome-partitioning-iii/) 1979
- [1335. 工作计划的最低难度](https://leetcode.cn/problems/minimum-difficulty-of-a-job-schedule/) 2035
- [2478. 完美分割的方案数](https://leetcode.cn/problems/number-of-beautiful-partitions/) 2344
- [2911. 得到 K 个半回文串的最少修改次数](https://leetcode.cn/problems/minimum-changes-to-make-k-semi-palindromes/) 2608

## 题单：划分型 DP ②

最小化/最大化分割出的区间个数等。
- [132. 分割回文串 II](https://leetcode.cn/problems/palindrome-partitioning-ii/)
- [2707. 字符串中的额外字符](https://leetcode.cn/problems/extra-characters-in-a-string/) 1736
- [2767. 将字符串分割为最少的美丽子字符串](https://leetcode.cn/problems/partition-string-into-minimum-beautiful-substrings/) 1865
- [1105. 填充书架](https://leetcode.cn/problems/filling-bookcase-shelves/) 2014
- [2547. 拆分数组的最小代价](https://leetcode.cn/problems/minimum-cost-to-split-an-array/) 2020
- [2463. 最小移动总距离](https://leetcode.cn/problems/minimum-total-distance-traveled/) 2454
- [2977. 转换字符串的最小成本 II](https://leetcode.cn/problems/minimum-cost-to-convert-string-ii/) 2696
