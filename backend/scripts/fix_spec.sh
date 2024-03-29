jq 'with_entries(if .key == "definitions" then .key = "components" | .value = { "schemas": .value } else . end)' ./docs/swagger.json \
| jq 'walk(if type == "object" and has("schema") then ( .content."application/json" = {"schema": ."schema"} | del(."schema") ) else . end)' \
| jq 'del(.swagger)' \
| jq  --indent 4 '. as $obj | { "openapi": "3.0.0" } + $obj'\
| sed 's/definitions\//components\/schemas\//g' > tmp \
&& mv tmp ./docs/swagger.json \