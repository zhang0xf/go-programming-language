package chapter4_6

import (
	"exercise/chapter04/chapter4_5"
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

// usage : ./exercise repo:golang/go is:open json decoder

// 一个模板是一个字符串或一个文件，里面包含了一个或多个由双花括号包含的{{action}}对象
// 对于每一个action，都有一个当前值的概念，对应点操作符.当前值“.”最初被初始化为调用模板时的参数，在当前例子中对应chapter4_5.IssuesSearchResult类型的变量
const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// 生成模板的输出需要两个处理步骤。第一步是要分析模板并转为内部表示，然后基于指定的输入执行模板。
// 注意方法调用链的顺序：template.New先创建并返回一个模板；Funcs方法将daysAgo等自定义函数注册到模板中，并返回模板；最后调用Parse函数分析模板。
func TextTemplate() {
	report1, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T", report1)

	// 模板通常在编译时就测试好了，如果模板解析失败将是一个致命的错误。template.Must辅助函数可以简化这个致命错误的处理.
	var report2 = template.Must(template.New("issuelist").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ))
	fmt.Printf("%T", report2)

	// 执行模板
	result, err := chapter4_5.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report2.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
