VERSION=$(jq -r '.version' package.json)
jq --arg version "$VERSION" '.packageVersion = $version' config.json > temp.json && mv temp.json config.json
