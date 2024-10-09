### 题目

<p>给你一个长度为 <code>n</code>&nbsp;的整数数组&nbsp;<code>nums</code>&nbsp;，<code>n</code>&nbsp;是 <strong>偶数</strong>&nbsp;，同时给你一个整数&nbsp;<code>k</code>&nbsp;。</p>

<p>你可以对数组进行一些操作。每次操作中，你可以将数组中 <strong>任一</strong>&nbsp;元素替换为 <code>0</code>&nbsp;到 <code>k</code>&nbsp;之间的<strong>&nbsp;任一</strong>&nbsp;整数。</p>

<p>执行完所有操作以后，你需要确保最后得到的数组满足以下条件：</p>

<ul>
	<li>存在一个整数 <code>X</code>&nbsp;，满足对于所有的&nbsp;<code>(0 &lt;= i &lt; n)</code>&nbsp;都有&nbsp;<code>abs(a[i] - a[n - i - 1]) = X</code>&nbsp;。</li>
</ul>

<p>请你返回满足以上条件 <strong>最少</strong>&nbsp;修改次数。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1,0,1,2,4,3], k = 4</span></p>

<p><span class="example-io"><b>输出：</b>2</span></p>

<p><strong>解释：</strong><br />
我们可以执行以下操作：</p>

<ul>
	<li>将&nbsp;<code>nums[1]</code>&nbsp;变为 2 ，结果数组为&nbsp;<code>nums = [1,<em><strong>2</strong></em>,1,2,4,3]</code>&nbsp;。</li>
	<li>将&nbsp;<code>nums[3]</code>&nbsp;变为 3 ，结果数组为&nbsp;<code>nums = [1,2,1,<em><strong>3</strong></em>,4,3]</code>&nbsp;。</li>
</ul>

<p>整数&nbsp;<code>X</code>&nbsp;为 2 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [0,1,2,3,3,6,5,4], k = 6</span></p>

<p><span class="example-io"><b>输出：</b>2</span></p>

<p><strong>解释：</strong><br />
我们可以执行以下操作：</p>

<ul>
	<li>将&nbsp;<code>nums[3]</code>&nbsp;变为 0 ，结果数组为&nbsp;<code>nums = [0,1,2,<em><strong>0</strong></em>,3,6,5,4]</code>&nbsp;。</li>
	<li>将&nbsp;<code>nums[4]</code>&nbsp;变为 4 ，结果数组为&nbsp;<code>nums = [0,1,2,0,<em><strong>4</strong></em>,6,5,4]</code>&nbsp;。</li>
</ul>

<p>整数 <code>X</code>&nbsp;为 4 。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= n == nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>n</code>&nbsp;是偶数。</li>
	<li><code>0 &lt;= nums[i] &lt;= k &lt;= 10<sup>5</sup></code></li>
</ul>

### 思路

## 方法一：枚举 X

### 提示 1

想一想，什么情况下答案是 $0$？什么情况下答案是 $1$？

如果答案是 $0$，意味着所有 $|\textit{nums}[i]-\textit{nums}[n-1-i]|$ 都等于同一个数 $X$。

如果答案是 $1$，意味着有 $n/2-1$ 个 $|\textit{nums}[i]-\textit{nums}[n-1-i]|$ 都等于同一个数 $X$。我们只需要修改那对不相等的，设这两个数分别为 $p=\textit{nums}[i],\ q=\textit{nums}[n-1-i]$。

不妨设 $p\le q$，分类讨论：

- 如果修改 $p$，那么把 $p$ 改成 $0$ 可以让差值尽量大，此时差值为 $q$。
- 如果修改 $q$，那么把 $q$ 改成 $k$ 可以让差值尽量大，此时差值为 $k-p$。
- 如果 $\max(q, k-p)\ge X$，那么改一个数就行。否则要改两个数。

注意题目保证 $n$ 是偶数。

### 提示 2

枚举 $X=0,1,2,\cdots,k$，计算至少要修改多少个数。

设 $p=\textit{nums}[i],\ q=\textit{nums}[n-1-i]$，且 $p\le q$（如果 $p>q$ 则交换 $p$ 和 $q$）。

统计 $q-p$ 的出现次数，记录到一个数组 $\textit{cnt}$ 中。

统计 $\max(q, k-p)$ 的出现次数，记录到另一个数组 $\textit{cnt}_2$ 中。

讨论哪些数对无需修改，哪些数对要改一个数，哪些数对要改两个数：

1. 有 $\textit{cnt}[X]$ 对 $(p,q)$ 无需修改。
2. 有 $n/2-\textit{cnt}[X]$ 对 $(p,q)$ 至少要改一个数。
3. 在 2 的基础上，有额外的 $\textit{cnt}_2[0] + \textit{cnt}_2[1] + \cdots + \textit{cnt}_2[X-1]$ 对 $(p,q)$ **还要再改一个数**（根据提示 1）。这可以在枚举 $X$ 的同时，维护一个变量 $\textit{sum}_2$ 表示这些 $\textit{cnt}_2[i]$ 的和。

综上，至少要修改

$$
\dfrac{n}{2} - \textit{cnt}[X] + \textit{sum}_2
$$

个数，用它更新答案的最小值。

```
func minChanges(nums []int, k int) int {
	cnt := make([]int, k+1)
	cnt2 := make([]int, k+1)
	n := len(nums)
	for i := 0; i < n/2; i++ {
		p, q := nums[i], nums[n-1-i]
		if p > q { // 保证 p <= q
			p, q = q, p
		}
		cnt[q-p]++
		cnt2[max(q, k-p)]++
	}

	ans := n
	sum2 := 0 // 统计有多少对 (p,q) 都要改
	for x, c := range cnt {
		// 其他 n/2-c 对 (p,q) 至少要改一个数，在此基础上，有额外的 sum2 对 (p,q) 还要再改一个数
		ans = min(ans, n/2-c+sum2)
		// 对于后面的更大的 x，当前的这 cnt2[x] 对 (p,q) 都要改
		sum2 += cnt2[x]
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+k)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(k)$。

## 方法二：差分数组

方法一相当于计算出了一个 $\textit{minModify}$ 数组，其中 $\textit{minModify}[X]$ 表示把所有数对 $(p,q)$ 的差值都改成 $X$ 所需的最少操作次数。

如何计算 $\textit{minModify}$ 数组，我们还有一种方法。

枚举数对 $(p,q)$ 的同时，想一想，把这两个数的差值改成 $X=0$ 要操作几次？改成 $X=1$ 要操作几次？改成 $X=2$ 要操作几次？…… 改成 $X=k$ 要操作几次？

设 $x=q-p,\ \textit{mx}=\max(q, k-p)$。讨论把差值改成多少：

- 改成 $[0,x-1]$ 中的数。这意味着 $p$ 和 $q$ 的距离变小，两者靠近即可，只需改其中一个数。
- 改成 $x$。无需修改。
- 改成 $[x,\textit{mx}]$ 中的数。根据方法一中的分析，只需改其中一个数。
- 改成 $[\textit{mx}+1,k]$ 中的数。根据方法一中的分析，两个数都需要改。

把这些操作次数加到 $\textit{minModify}$ 数组中。这可以用**差分数组**实现，具体见 [差分数组原理讲解](https://leetcode.cn/problems/car-pooling/solution/suan-fa-xiao-ke-tang-chai-fen-shu-zu-fu-9d4ra/)，推荐和[【图解】从一维差分到二维差分](https://leetcode.cn/problems/stamping-the-grid/solution/wu-nao-zuo-fa-er-wei-qian-zhui-he-er-wei-zwiu/) 一起看。

遍历这 $n/2$ 个数对，都按上述方法更新 $\textit{minModify}$ 数组，最终我们得到的 $\textit{minModify}$ 数组就和方法一一样了。答案为 $\min(\textit{minModify})$。

注意差分数组需要开 $k+2$ 大小，因为 $\textit{mx}+1$ 可以等于 $k+1$。

```
func minChanges(nums []int, k int) int {
	n := len(nums)
	d := make([]int, k+2)
	for i := 0; i < n/2; i++ {
		p, q := nums[i], nums[n-1-i]
		if p > q { // 保证 p <= q
			p, q = q, p
		}
		x := q - p
		mx := max(q, k-p)
		// [0, x-1] 全部 +1：把 q-p 改成小于 x 的，只需要改 p 或 q 中的一个数
		d[0]++
		d[x]--
		// [x+1, mx] 全部 +1：把 q-p 改成大于 x 小于等于 mx 的，也只需要改 p 或 q 中的一个数
		d[x+1]++
		d[mx+1]--
		// [mx+1, k] 全部 +2：把 q-p 改成大于 mx 的，p 和 q 都需要改
		d[mx+1] += 2
	}

	ans := n
	minModify := 0
	for _, v := range d {
		minModify += v
		ans = min(ans, minModify)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+k)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(k)$。

## 题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
- [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
- [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
- [贪心（基本贪心策略/反悔/区间/字典序/数学/思维/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)