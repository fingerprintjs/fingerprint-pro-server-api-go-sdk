#!/bin/bash

# jar was downloaded from here https://repo1.maven.org/maven2/io/swagger/codegen/v3/swagger-codegen-cli/3.0.34/
java -jar ./bin/swagger-codegen-cli.jar generate -t ./template -l go -i https://fingerprintjs.github.io/fingerprint-pro-server-api-openapi/schemes/fingerprint-server-api.yaml -o ./sdk -c config.json

examplesList=(
  'visits_limit_1.json'
  'visits_limit_500.json'
  'webhook.json'
)

for example in ${examplesList[*]}; do
  curl -o ./test/mocks/"$example" https://fingerprintjs.github.io/fingerprint-pro-server-api-openapi/examples/"$example"
done