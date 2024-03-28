jq --indent 4 'with_entries(if .key == "definitions" then .key = "components" | .value = { "schemas": .value } else . end)' ./docs/swagger.json \
| sed 's/definitions\//components\/schemas\//g' > tmp \
&& mv tmp ./docs/swagger.json \