package v1

import (
	"bytes"
	"text/template"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

const queryFindUserLog = `
SELECT * FROM user_logs
WHERE 1=1
{{- if .Model.UserId }} AND user_id = {{ .Model.UserId }} {{- end}}	
{{- if .Model.DeclarationId }} AND declaration_id = {{ .Model.DeclarationId }} {{- end}}	
{{- if .Model.From }} AND created_at >= '{{ call .MakeTime .Model.From }}' {{- end}}	
{{- if .Model.To }} AND created_at <= '{{call .MakeTime .Model.To }}' {{- end}}	
{{- if .Model.Type }} AND type = {{.Model.Type}} {{- end}}	
{{- if .Model.Limit }} LIMIT {{.Model.Limit}} {{- end}}	
{{- if .Model.Offset }} OFFSET {{.Model.Offset}} {{- end}}		
`

func createQuery(queryTemplate string, model interface{}) (string, error) {
	type ViewData struct {
		Model    interface{}
		MakeTime func(*timestamp.Timestamp) string
	}

	var data bytes.Buffer
	t := template.New("")
	t.Parse(queryTemplate)
	if err := t.Execute(&data, ViewData{Model: model, MakeTime: func(t *timestamp.Timestamp) string {
		t2, _ := ptypes.Timestamp(t)
		return t2.Format(time.RFC3339)
		// return t2.Format("2006-01-02 15:04:05")
	}}); err != nil {
		return "", err
	}

	return data.String(), nil
}
