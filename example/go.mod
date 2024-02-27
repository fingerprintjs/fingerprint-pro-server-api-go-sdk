module github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/example

require (
	github.com/antihax/optional v1.0.0
	github.com/joho/godotenv v1.5.1
)

require (
	github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v5 v5.0.0-00010101000000-000000000000 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/oauth2 v0.17.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)

go 1.20

replace github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v5 => ../
