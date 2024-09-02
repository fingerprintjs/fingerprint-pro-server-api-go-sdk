VERSION=$(grep '"version"' package.json | sed -E 's/.*"version": "([0-9]+\.[0-9]+\.[0-9]+)".*/\1/')
jq --arg version "$VERSION" '.packageVersion = $version' config.json > temp.json && mv temp.json config.json
