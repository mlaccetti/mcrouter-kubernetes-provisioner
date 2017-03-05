{
  "pools": {
    "default": {
      "servers": [
        {{if .servers}}
          {{ range _, $value := .servers }}
            "{{$value}}:5000",
          {{end}}
        {{end}}
      ]
    }
  },
  "route": "PoolRoute|default"
}
