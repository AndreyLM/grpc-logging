package v2

//___________________________________________________________________________________________________
//----------------------------------------<< TOTAL COUNT QUERIES >> ----------------------------------------|
//___________________________________________________________________________________________________|

const queryTotalCountUsers = `
SELECT COUNT(*) as count FROM users
WHERE 1=1
{{- if .Model.CreatedAtFrom }} AND created_at >= '{{ call .MakeTime .Model.CreatedAtFrom }}' {{- end}}	
{{- if .Model.CreatedAtTo }} AND created_at <= '{{call .MakeTime .Model.CreatedAtTo }}' {{- end}}	
{{- if .Model.UserId }} AND user_id = {{ .Model.UserId }} {{- end}}	
{{- if .Model.TypeId }} AND type_id = {{ .Model.TypeId }} {{- end}}	
{{- if .Model.Content }} AND LOWER(content) LIKE '%' || LOWER('{{.Model.Content}}') || '%' {{- end}}	
`
const queryTotalCountRules = `
SELECT COUNT(*) as count FROM rules
WHERE 1=1
{{- if .Model.CreatedAtFrom }} AND created_at >= '{{ call .MakeTime .Model.CreatedAtFrom }}' {{- end}}	
{{- if .Model.CreatedAtTo }} AND created_at <= '{{call .MakeTime .Model.CreatedAtTo }}' {{- end}}	
{{- if .Model.RuleId }} AND rule_id = {{ .Model.RuleId }} {{- end}}	
{{- if .Model.CreatedBy }} AND created_by = {{ .Model.CreatedBy }} {{- end}}	
{{- if .Model.RuleNumber }} AND rule_number = {{ .Model.RuleNumber }} {{- end}}
{{- if .Model.Content }} AND LOWER(content) LIKE '%' || LOWER('{{.Model.Content}}') || '%' {{- end}}	
`
const queryTotalCountExchanges = `
SELECT COUNT(*) as count FROM exchanges
WHERE 1=1
{{- if .Model.CreatedAtFrom }} AND created_at >= '{{ call .MakeTime .Model.CreatedAtFrom }}' {{- end}}	
{{- if .Model.CreatedAtTo }} AND created_at <= '{{call .MakeTime .Model.CreatedAtTo }}' {{- end}}	
{{- if .Model.TypeId }} AND type_id = {{ .Model.TypeId }} {{- end}}	
{{- if .Model.StateId }} AND state_id = {{ .Model.StateId }} {{- end}}	
{{- if .Model.RequestId }} AND request_id = {{ .Model.RequestId }} {{- end}}	
{{- if .Model.DeclarationId }} AND declaration_id = {{ .Model.DeclarationId }} {{- end}}	
{{- if .Model.RegisterId }} AND register_id = {{ .Model.RegisterId }} {{- end}}
{{- if .Model.Content }} AND LOWER(content) LIKE '%' || LOWER('{{.Model.Content}}') || '%' {{- end}}	
`
const queryTotalCountDeclarations = `
SELECT COUNT(*) as count FROM declarations
WHERE 1=1
{{- if .Model.CreatedAtFrom }} AND created_at >= '{{ call .MakeTime .Model.CreatedAtFrom }}' {{- end}}	
{{- if .Model.CreatedAtTo }} AND created_at <= '{{call .MakeTime .Model.CreatedAtTo }}' {{- end}}	
{{- if .Model.DeclarationId }} AND declaration_id = {{ .Model.DeclarationId }} {{- end}}	
{{- if .Model.UserId }} AND user_id = {{ .Model.UserId }} {{- end}}	
{{- if .Model.UserIp }} AND user_ip = '{{ .Model.UserIp }}' {{- end}}
{{- if .Model.Content }} AND LOWER(content) LIKE '%' || LOWER('{{.Model.Content}}') || '%' {{- end}}	
`

//___________________________________________________________________________________________________
//----------------------------------------<< FIND QUERIES >> ----------------------------------------|
//___________________________________________________________________________________________________|

const queryFindUsers = `
SELECT created_at, user_id, type_id, content FROM users
WHERE 1=1
{{- if .Model.CreatedAtFrom }} AND created_at >= '{{ call .MakeTime .Model.CreatedAtFrom }}' {{- end}}	
{{- if .Model.CreatedAtTo }} AND created_at <= '{{call .MakeTime .Model.CreatedAtTo }}' {{- end}}	
{{- if .Model.UserId }} AND user_id = {{ .Model.UserId }} {{- end}}	
{{- if .Model.TypeId }} AND type_id = {{ .Model.TypeId }} {{- end}}
{{- if .Model.Content }} AND LOWER(content) LIKE '%' || LOWER('{{.Model.Content}}') || '%' {{- end}}	
{{- if .Model.OrderBy }} ORDER BY {{ .Model.OrderBy }} {{- end}}
{{- if .Model.Limit }} LIMIT {{ .Model.Limit }} {{- end}}	
{{- if .Model.Offset }} OFFSET {{ .Model.Offset }} {{- end}}		
`

const queryFindRules = `
SELECT created_at, rule_id, created_by, content, rule_number FROM rules
WHERE 1=1
{{- if .Model.CreatedAtFrom }} AND created_at >= '{{ call .MakeTime .Model.CreatedAtFrom }}' {{- end}}	
{{- if .Model.CreatedAtTo }} AND created_at <= '{{call .MakeTime .Model.CreatedAtTo }}' {{- end}}	
{{- if .Model.RuleId }} AND rule_id = {{ .Model.RuleId }} {{- end}}	
{{- if .Model.CreatedBy }} AND created_by = {{ .Model.CreatedBy }} {{- end}}	
{{- if .Model.RuleNumber }} AND rule_number = {{ .Model.RuleNumber }} {{- end}}
{{- if .Model.Content }} AND LOWER(content) LIKE '%' || LOWER('{{.Model.Content}}') || '%' {{- end}}	
{{- if .Model.OrderBy }} ORDER BY {{ .Model.OrderBy }} {{- end}}
{{- if .Model.Limit }} LIMIT {{ .Model.Limit }} {{- end}}	
{{- if .Model.Offset }} OFFSET {{ .Model.Offset }} {{- end}}		
`

const queryFindExchanges = `
SELECT created_at, type_id, state_id, request_id, declaration_id, register_id, content FROM exchanges
WHERE 1=1
{{- if .Model.CreatedAtFrom }} AND created_at >= '{{ call .MakeTime .Model.CreatedAtFrom }}' {{- end}}	
{{- if .Model.CreatedAtTo }} AND created_at <= '{{call .MakeTime .Model.CreatedAtTo }}' {{- end}}	
{{- if .Model.TypeId }} AND type_id = {{ .Model.TypeId }} {{- end}}	
{{- if .Model.StateId }} AND state_id = {{ .Model.StateId }} {{- end}}	
{{- if .Model.RequestId }} AND request_id = {{ .Model.RequestId }} {{- end}}	
{{- if .Model.DeclarationId }} AND declaration_id = {{ .Model.DeclarationId }} {{- end}}	
{{- if .Model.RegisterId }} AND register_id = {{ .Model.RegisterId }} {{- end}}
{{- if .Model.Content }} AND LOWER(content) LIKE '%' || LOWER('{{.Model.Content}}') || '%' {{- end}}	
{{- if .Model.OrderBy }} ORDER BY {{ .Model.OrderBy }} {{- end}}
{{- if .Model.Limit }} LIMIT {{ .Model.Limit }} {{- end}}	
{{- if .Model.Offset }} OFFSET {{ .Model.Offset }} {{- end}}		
`

const queryFindDeclarations = `
SELECT created_at, declaration_id, content, user_id, user_ip FROM declarations
WHERE 1=1
{{- if .Model.CreatedAtFrom }} AND created_at >= '{{ call .MakeTime .Model.CreatedAtFrom }}' {{- end}}	
{{- if .Model.CreatedAtTo }} AND created_at <= '{{call .MakeTime .Model.CreatedAtTo }}' {{- end}}	
{{- if .Model.DeclarationId }} AND declaration_id = {{ .Model.DeclarationId }} {{- end}}	
{{- if .Model.UserId }} AND user_id = {{ .Model.UserId }} {{- end}}	
{{- if .Model.UserIp }} AND user_ip = '{{ .Model.UserIp }}' {{- end}}
{{- if .Model.Content }} AND LOWER(content) LIKE '%' || LOWER('{{.Model.Content}}') || '%' {{- end}}	
{{- if .Model.OrderBy }} ORDER BY {{ .Model.OrderBy }} {{- end}}
{{- if .Model.Limit }} LIMIT {{ .Model.Limit }} {{- end}}	
{{- if .Model.Offset }} OFFSET {{ .Model.Offset }} {{- end}}		
`
