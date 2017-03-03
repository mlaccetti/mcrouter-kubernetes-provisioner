{
  "pools": {
    "default": {
      "servers": [
        {{if servers}}
          {{ range $key, $value := . }}
            "{{$key}}:{{$value}}",
          {{end}}
        {{end}}
      ]
    }
  },
  "route": "PoolRoute|default"
}
