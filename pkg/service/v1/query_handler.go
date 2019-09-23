package v1

import (
	"bytes"
	"log"
	"text/template"
	"time"

	"github.com/davecgh/go-spew/spew"

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
{{- if .Model.Type }} AND type = '{{.Model.Type}}' {{- end}}	
{{- if .Model.Limit }} LIMIT {{.Model.Limit}} {{- end}}	
{{- if .Model.Offset }} OFFSET {{.Model.Offset}} {{- end}}		
`

const queryFindExchangeLog = `
SELECT * FROM exchanges
WHERE 1=1
{{- if .Model.From }} AND created_at >= '{{ call .MakeTime .Model.From }}' {{- end}}	
{{- if .Model.To }} AND created_at <= '{{call .MakeTime .Model.To }}' {{- end}}	
{{- if .Model.TypeId }} AND type_id = {{ .Model.TypeId }} {{- end}}	
{{- if .Model.StateId }} AND state_id = {{ .Model.StateId }} {{- end}}	
{{- if .Model.RequestId }} AND request_id = {{ .Model.RequestId }} {{- end}}	
{{- if .Model.DeclarationId }} AND declaration_id = {{ .Model.DeclarationId }} {{- end}}	
{{- if .Model.RegisterId }} AND register_id = '{{ .Model.RegisterId }}' {{- end}}	
{{- if .Model.Content }} AND content = '{{ .Model.Content }}' {{- end}}	
{{- if .Model.OrderBy }} ORDER BY {{ .Model.OrderBy }} {{- end}}		
{{- if .Model.Limit }} LIMIT {{ .Model.Limit }} {{- end}}	
{{- if .Model.Offset }} OFFSET {{ .Model.Offset }} {{- end}}		
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
		spew.Dump(model)
		log.Println(model)
		return "", err
	}

	return data.String(), nil
}
