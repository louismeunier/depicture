package lib

import (
	"fmt"
	"html/template"
	"os"
)

type SummaryData struct {
	FilePath string
	Colors   []Color
}

func GenerateHTMLSummary(filePath string, colors []Color) {
	const tmpl = `<h1>{{.FilePath}}</h1>
		<img src="{{.FilePath}}"/>
		<ul>
			{{range .Colors}}
				<li>{{.Percentage}}%: {{.Name}}</li>
			{{end}}
		</ul>`

	t, err := template.New("webpage").Parse(tmpl)

	if err != nil {
		fmt.Println("Error occurred")
	}

	data := SummaryData{
		FilePath: filePath,
		Colors:   colors,
	}

	f, err := os.Create("index.html")

	if err != nil {
		fmt.Println(err)
	} else {
		t.Execute(f, data)

	}
}
