// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	. "fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"path/filepath"
	"problems/testutil"
	"problems/testutil/codeforces"
	"strings"
	"testing"
)

var customTestCases = [][2]string{
	{
		`3
		1 2 3`,
		`12`,
	},
}

func Test(t *testing.T) {
	dir, _ := filepath.Abs(".")
	t.Logf("Current problem is [%s]", filepath.Base(dir))

	if len(customTestCases) > 0 && strings.TrimSpace(customTestCases[0][0]) != "" {
		tarCase := 0 // -1
		codeforces.AssertEqualStringCase(t, run, customTestCases, tarCase)
		t.Log("======= custom =======")
		return
	}

	tarCase := 0 // -1
	codeforces.AssertEqualFileCaseWithName(t, run, dir, "in*.txt", "ans*.txt", tarCase)
}

// 无尽对拍 / 构造 hack 数据
// 提醒：如果对拍正确但没有 AC，请检查
// 1. 你真的真的真的读对题了吗？
// 2. 是否搞错了输入输出（重新阅读题目中有关输入输出格式的说明）
// 3. 是否有边界情况没考虑到
// 4. 数据是否超过 int 最大值
func TestCompare(t *testing.T) {
	return
	codeforces.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		rg.One() // 若不是多测则 remove
		n := rg.Int(1, 3)
		rg.NewLine()
		rg.IntSlice(n, 1, 5)
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		solve := func(Case int) {

		}

		T := 1
		Fscan(in, &T) //
		for Case := 1; Case <= T; Case++ {
			solve(Case)
		}

		_leftData, _ := ioutil.ReadAll(in)
		if _s := strings.TrimSpace(string(_leftData)); _s != "" {
			panic("有未读入的数据：\n" + _s)
		}
	}

	// 先用 runBF 跑下样例，检查 runBF 是否正确
	dir, _ := filepath.Abs(".")
	codeforces.AssertEqualFileCaseWithName(t, runBF, dir, "in*.txt", "ans*.txt", 0)
	// codeforces.AssertEqualStringCase(t, runBF, customTestCases, 0,)
	return

	codeforces.AssertEqualRunResultsInf(t, inputGenerator, run, runBF)

	// for hacking, write the hacked codes in runBF
	// codeforces.AssertEqualRunResultsInf(t, inputGenerator, run, runBF)
}

// 无尽检查输出是否正确 / 构造 hack 数据
// 通常用于 special judge 题目
func TestCheck(_t *testing.T) {
	return
	assert := assert.New(_t)
	_ = assert

	codeforces.DebugTLE = 0

	inputGenerator := func() (string, codeforces.OutputChecker) {
		rg := testutil.NewRandGenerator()
		rg.One() // 若不是多测则 remove
		n := rg.Int(1, 5)
		rg.NewLine()
		a := rg.IntSlice(n, 0, 5)
		return rg.String(), func(myOutput string) (_b bool) {
			// 检查 myOutput 是否符合题目要求
			// * 最好重新看一遍题目描述以免漏判 *
			// 对于 special judge 的题目，可能还需要额外跑个暴力来检查 myOutput 是否满足最优解等
			in := strings.NewReader(myOutput)

			myA := make([]int, n)
			for i := range myA {
				Fscan(in, &myA[i])
			}
			if !assert.EqualValues(a, myA) {
				return
			}

			return true
		}
	}

	target := 0
	codeforces.CheckRunResultsInfWithTarget(_t, inputGenerator, target, run)

	//runHack := func(in io.Reader, out io.Writer) { }
	//codeforces.CheckRunResultsInf(_t, inputGenerator, runHack)
}

func TestRE(_t *testing.T) {
	return
	codeforces.DebugTLE = 0

	inputGenerator := func() (string, codeforces.OutputChecker) {
		rg := testutil.NewRandGenerator()
		rg.One() // 若不是多测则 remove
		n := rg.Int(1, 5)
		rg.NewLine()
		rg.IntSlice(n, 0, 5)
		return rg.String(), func(myOutput string) bool { return true }
	}
	codeforces.CheckRunResultsInfWithTarget(_t, inputGenerator, 0, run)
}
