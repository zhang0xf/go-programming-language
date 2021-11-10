package chapter4_6

import (
	"exercise/chapter4_5"
	"html/template"
	"log"
	"os"
)

// usage : ./exercise repo:golang/go commenter:gopherbot json encoder > data/issues.html
// usage : ./exercise repo:golang/go 3133 10535 > data/issues2.html

// html/template包已经自动将特殊字符转义，因此我们依然可以看到正确的字面值。
// 如果我们使用text/template包的话，这2个issue将会产生错误，其中“&lt;”四个字符将会被当作小于字符“<”处理，同时“<link>”字符串将会被当作一个链接元素处理.

func HtmlTemplate() {
	var issueList = template.Must(template.New("issuelist").Parse(`	
	<h1>{{.TotalCount}} issues</h1>
	<table>
	<tr style='text-align: left'>
	  <th>#</th>
	  <th>State</th>
	  <th>User</th>
	  <th>Title</th>
	</tr>
	{{range .Items}}
	<tr>
	  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	  <td>{{.State}}</td>
	  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
	</table>
	`))

	result, err := chapter4_5.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
