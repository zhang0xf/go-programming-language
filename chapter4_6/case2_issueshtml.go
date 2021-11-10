package chapter4_6

import (
	"exercise/chapter4_5"
	"html/template"
	"log"
	"os"
)

// usage : ./exercise repo:golang/go commenter:gopherbot json encoder > data/issues.html

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
