package mytemplate

import (
	"fmt"
	httpTemplate "html/template"
	"log"
	"os"
	"text/template"
	"time"

	"example.com/practice/packages/issuesearch"
)

var templ = template.Must(template.New("issuelist").Funcs(template.FuncMap{"daysAgo": daysAgo}). Parse(`{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:
{{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age:
{{.CreatedAt | daysAgo}} days
{{end}}`))

var issueList = httpTemplate.Must(httpTemplate.New("issuelist").Parse(`
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
<td><a href='{{.HTMLURL}}'>{{.Number}}</td>
<td>{{.State}}</td>
<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func init() {
  fmt.Println("TEMPLATE")
  Test()
}


func IssueSearch() {
  result, err := issuesearch.SearchIssues(os.Args[1:])
  if err != nil {
    log.Fatal(err)
  }
  if err := templ.Execute(os.Stdout, result); err != nil {
    log.Fatal(err)
  }
}

func IssueSearchHtml() {
  result, err := issuesearch.SearchIssues(os.Args[1:])
  if err != nil {
    log.Fatal(err)
  }
  if err := issueList.Execute(os.Stdout, result); err != nil {
    log.Fatal(err)
  }
}

func daysAgo(t time.Time) int {
  return int(time.Since(t).Hours() / 24)
}