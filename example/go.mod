module github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/example

go 1.20

require (
	github.com/antihax/optional v1.0.0
	github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v5 v5.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.5.1
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/oauth2 v0.17.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

replace github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v5 => ../
