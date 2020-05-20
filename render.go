package main

import (
	"bytes"
	"text/template"
)

const card = `
{{range .Results}}
<div>res:</div>
<div>
	<div>Word: {{.Word}}</div>
	<div>Language: {{.Language}}</div>
</div>
{{end}}
`

func renderCard(data *Response) (string, error) {
	tmpl, err := template.New("anki").Parse(card)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
