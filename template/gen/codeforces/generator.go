package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/levigross/grequests"
	"github.com/skratchdot/open-golang/open"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

const (
	CmdCodeforces = "codeforces" // https://github.com/xalanq/cf-tool
	CmdAtcoder    = "atc"        // https://github.com/sempr/cf-tool rename
)

// 生成 CF 比赛模板（需要先 codeforces race，以确认题目数量）
func GenCodeforcesContestTemplates(cmdName, rootPath, contestID string, overwrite bool) error {
	if contestID == "" {
		fmt.Println("contest ID is empty")
		return nil
	}

	problems, err := fetchProblems(contestID)
	if err != nil {
		return err
	}
	fmt.Printf("共 %d 道题目\n", len(problems))

	var wg sync.WaitGroup
	defer wg.Wait()
	for _, problem := range problems {
		wg.Add(1)
		// we don't want spent too much time on waiting responses one by one, so we use goroutine!
		go func(problem string) {
			defer wg.Done()
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("[error]", problem, err)
				}
			}()
			problemURL := fmt.Sprintf("https://codeforces.com/contest/%s/problem/%s", contestID, problem)
			if err := genTemplates(rootPath, contestID, problem, problemURL, true); err != nil {
				fmt.Println("[error]", problem, err)
				return
			}
			fmt.Println("[ok]", problem, problemURL)
		}(problem)
	}

	return nil
}

func fetchProblems(contestID string) (problems []string, err error) {
	contestHome := "https://codeforces.com/contest/" + contestID
	resp, err := grequests.Get(contestHome, &grequests.RequestOptions{})
	if err != nil {
		return
	}
	if !resp.Ok {
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	pIDList := doc.Find("table.problems").Find("td.id")
	for _, v := range pIDList.Nodes {
		problems = append(problems, strings.ReplaceAll(
			strings.ReplaceAll(v.FirstChild.NextSibling.FirstChild.Data, " ", ""), "\n", ""))
	}
	return
}

func genTemplates(rootPath, contestID, problem, problemURL string, isContest bool) (err error) {
	resp, err := grequests.Get(problemURL, &grequests.RequestOptions{})
	if err != nil {
		return
	}
	if !resp.Ok {
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	sampleTest := doc.Find("div.sample-test")
	input := sampleTest.Find("div.input").Find("pre").Find("div")
	output := sampleTest.Find("div.output").Find("pre")
	inputs, outputs := []string{}, []string{}
	for _, v := range input.Nodes {
		inputs = append(inputs, v.FirstChild.Data)
	}
	for _, v := range output.Nodes {
		outputs = append(outputs, v.FirstChild.Data)
	}

	dirPath := filepath.Join(rootPath, problem) + "/"
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return err
	}

	// 创建 x.go
	mainFileContent := `package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)

}

func main() { run(os.Stdin, os.Stdout) }
`
	mainFilePath := dirPath + problem + ".go"
	if !isContest || problem == "A" {
		// 比赛时，在 IDE 中打开 A 题
		// defer open.Run(absPath(mainFilePath))
	}
	if err := os.WriteFile(mainFilePath, []byte(mainFileContent), 0644); err != nil {
		return err
	}

	// 创建 x_test.go
	examples := ""
	examples += "\n\t\t{\n"
	examples += "\t\t\t`" + strings.ReplaceAll(strings.Join(inputs, "\n"), "\n", "\n\t\t\t") + "`,\n"
	examples += "\t\t\t`" + strings.ReplaceAll(strings.Join(outputs, "\n"), "\n", "\n\t\t\t") + "`,\n"
	examples += "\t\t},"

	testFileContent := fmt.Sprintf(`// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [%s]")
	testCases := [][2]string{%s
		
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}
// %s
`, problem, examples, problemURL)
	testFilePath := dirPath + problem + "_test.go"
	if err := os.WriteFile(testFilePath, []byte(testFileContent), 0644); err != nil {
		return err
	}

	return nil

	return
}

// 生成单道题目的模板（Codeforces）
func GenCodeforcesProblemTemplates(problemURL string, openWebsite bool) (err error) {
	urlObj, err := url.Parse(problemURL)
	if err != nil {
		return err
	}

	contestID, problemID, isGYM := parseCodeforcesProblemURL(problemURL)
	if _, err := strconv.Atoi(contestID); err != nil {
		return fmt.Errorf("invalid URL: %v", err)
	}

	var statusURL string
	if isGYM {
		statusURL = fmt.Sprintf("https://%s/gym/%s/status/%s", urlObj.Host, contestID, problemID)
	} else {
		statusURL = fmt.Sprintf("https://%s/problemset/status/%s/problem/%s", urlObj.Host, contestID, problemID)
	}

	if openWebsite {
		luoguURL := fmt.Sprintf("https://www.luogu.com.cn/problem/CF%s%s", contestID, problemID)
		open.Run(luoguURL)
		open.Run(statusURL)
		open.Run(problemURL)
	}

	luoguURL := fmt.Sprintf("https://www.luogu.com.cn/problem/CF%s%s", contestID, problemID)
	example, err := parseExamples(luoguURL)
	if err != nil {
		fmt.Println(err)
	}
	if len(example) == 0 {
		fmt.Println("未获取到样例，请手动添加")
	}

	if !isGYM {
		problemID = contestID + problemID
	}
	mainStr := fmt.Sprintf(`//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF%[1]s(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	
}

func main() { CF%[1]s(os.Stdin, os.Stdout) }
`, problemID)
	examples := ""
	for _, p := range example {
		in := strings.TrimSpace(p[0])
		out := strings.TrimSpace(p[1])
		examples += "\n\t\t{\n"
		examples += "\t\t\t`" + in + "`,\n"
		examples += "\t\t\t`" + out + "`,\n"
		examples += "\t\t},"
	}
	mainTestStr := fmt.Sprintf(`//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// %s
// %s
func TestCF%[3]s(t *testing.T) {
	t.Log("Current test is [CF%[3]s]")
	testCases := [][2]string{%[4]s
		
	}
	codeforces.AssertEqualStringCase(t, CF%[3]s, testCases, 0)
}
`, problemURL, statusURL, problemID, examples)

	var dir string
	if isGYM {
		dir = "../../../misc/cf/main/gym/" + contestID + "/"
	} else {
		dir = "../../../misc/cf/main/" + genDirName(contestID) + "/"
	}
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	mainFilePath := dir + problemID + ".go"
	if _, err := os.Stat(mainFilePath); !os.IsNotExist(err) {
		// open.Run(absPath(mainFilePath))
		return fmt.Errorf("文件已存在！")
	}
	if err := ioutil.WriteFile(mainFilePath, []byte(mainStr), 0644); err != nil {
		return err
	}
	// open.Run(absPath(mainFilePath))
	testFilePath := dir + problemID + "_test.go"
	if err := ioutil.WriteFile(testFilePath, []byte(mainTestStr), 0644); err != nil {
		return err
	}
	// open.Run(absPath(testFilePath))
	return nil
}

// 在某一路径下批量生成模板
func GenTemplates(problemNum int, rootPath string, overwrite bool) error {
	for i := 'a'; i < 'a'+int32(problemNum); i++ {
		dir := rootPath + string(i) + "/"
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
		for j, fileName := range []string{"main.go", "main_test.go"} {
			goFilePath := dir + strings.Replace(fileName, "main", string(i), 1)
			if !overwrite {
				if _, err := os.Stat(goFilePath); !os.IsNotExist(err) {
					continue
				}
			}
			if err := copyFile(goFilePath, fileName); err != nil {
				return err
			}
			if i == 'a' && j == 0 {
				// open.Run(absPath(goFilePath))
			}
		}
	}
	return nil
}
