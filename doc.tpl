{{range $data := .}}
# {{$data.Title}}
{{range $val := $data.Endpoints}}

<details><summary style="margin:3px"> 
{{$val.EndpointName}}
</summary>
<p style="margin:10px">
<span style="background-color:rgb(0, 122, 202); color:white; padding:2px"> 
{{$val.Method}}
</span>
<code> {{$val.URL}} </code>
</p>

---
{{ with $val.Params }}
### Query Parameters: 
{{- range $param := $val.Params}}
`{{$param -}}` |
{{- end}}
{{end}}
{{ with $val.RequestBody }}
### Request Body Schema: 
```json
{{$val.RequestBody}}
```
{{end}}
{{ with $val.RequestFields }}
### Request Fields: 
| Field | Required | Description |
| ------ | ------ | ------ |
{{- range $reqField := $val.RequestFields }}
|{{$reqField.Field}}|{{$reqField.Required}}|{{$reqField.Description}}|
{{- end}}
{{end}}

### Response

```json
{{$val.Response}}
```
{{ with $val.ResponseFields }}
### Response Fields: 
| Field | Required | Description |
| ------ | ------ | ------ |
{{- range $respField := $val.ResponseFields }}
|{{$respField.Field}}|{{$respField.Required}}|{{$respField.Description}}|
{{- end}}
{{end}}

</details>

---

{{ end }}
{{ end }}