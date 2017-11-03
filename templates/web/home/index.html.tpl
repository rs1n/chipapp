{{ define "web/home/index" }}

{{ template "web/layout/header" . }}

<h1>Hello, {{ .name }}, from web/home/index</h1>

{{ template "web/layout/footer" . }}

{{ end }}
