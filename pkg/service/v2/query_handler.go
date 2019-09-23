package v2

import (
	"bytes"
	"text/template"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type viewData struct {
	Model    interface{}
	MakeTime func(*timestamp.Timestamp) string
}

func createQuery(queryTemplate string, model interface{}) (string, error) {
	var data bytes.Buffer
	t := template.New("")
	t.Parse(queryTemplate)
	if err := t.Execute(&data, viewData{Model: model, MakeTime: func(t *timestamp.Timestamp) string {
		t2, _ := ptypes.Timestamp(t)
		return t2.Format(time.RFC3339)
		// return t2.Format("2006-01-02 15:04:05")
	}}); err != nil {
		return "", err
	}

	return data.String(), nil
}

// FIND QUERIES
const queryFindUsers = `
SELECT created_at, user_id, type_id, content FROM users
WHERE 1=1
{{- if .Model.CreatedAtFrom }} AND created_at >= '{{ call .MakeTime .Model.CreatedAtFrom }}' {{- end}}	
{{- if .Model.CreatedAtTo }} AND created_at <= '{{call .MakeTime .Model.CreatedAtTo }}' {{- end}}	
{{- if .Model.UserId }} AND user_id = {{ .Model.UserId }} {{- end}}	
{{- if .Model.Type }} AND type = {{.Model.Type}} {{- end}}	
{{- if .Model.Limit }} LIMIT {{.Model.Limit}} {{- end}}	
{{- if .Model.Offset }} OFFSET {{.Model.Offset}} {{- end}}		
{{- if .Model.OrderBy }} ORDER BY {{.Model.OrderBy}} {{- end}}		
`
