{
  "pools": {
    "default": {
      "servers": [ {{if .servers}}
        {{ range $key, $value := .servers }}"{{$value}}:5000",{{end}}
    {{end}}]
    }
  },
  "route": "PoolRoute|default"
}
