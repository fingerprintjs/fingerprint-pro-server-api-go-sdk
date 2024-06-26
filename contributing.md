# Contributing to Fingerprint Pro Server API Go SDK

## Structure

Most files in the project are autogenerated by [swagger-codegen](https://swagger.io/tools/swagger-codegen/).

- [template](./template) - folder contains redefined templates of `swagger-codegen`. Original templates you can find in [swagger-codegen repo](https://github.com/swagger-api/swagger-codegen/tree/751e59df060b1c3ecf54921e104f2086dfa9f820/modules/swagger-codegen/src/main/resources/go).
- [docs](./docs) - generated documentation for models and [API Client](./docs/FingerprintApi.md).
- [sdk](./sdk) - API Client code - partialy generated by swagger.

## Code generation

You need `swagger-codegen` to run code generation. There are many ways described in the [readme](https://github.com/swagger-api/swagger-codegen).
In the project we use local jar file `swagger-codegen-cli.jar`.

You can just run `go run generate.go` script and it will do all the work.

To download fresh OpenAPI schema run `./sync.sh`.

Swagger generates models and API service interface for us, which we then implement manually.
You can check [./sdk/api_fingerprint.go](./sdk/api_fingerprint.go) and [./sdk/api_fingerprint_impl.go](./sdk/api_fingerprint_impl.go) for examples.

### Configuration

Project configuration is described in `config.json` file. To read about available parameters run the command below:

```shell
java -jar ./bin/swagger-codegen-cli.jar config-help -l go
```

### Running tests

Tests are located in tests.

To run tests you can use IDE instruments or just run:

```shell
cd test && go test
```

## How to test the local version of the SDK

Go inside the `example` folder to test API requests using the local version of the SDK. The [example/go.mod](./example/go.mod) file reroutes the SDK module references inside `example` to its parent folder.

Create an `.env` file inside the `example` folder according to [example.env](/example/example.env).

Run the scripts like this:

```shell
cd example
go run getEvent.go
go run getVisits.go
```

### How to publish

The library is automatically released on every push to the `main` branch if there are relevant changes using [semantic-release](https://github.com/semantic-release/semantic-release) with following plugins:

- [@semantic-release/commit-analyzer](https://github.com/semantic-release/commit-analyzer)
- [@semantic-release/release-notes-generator](https://github.com/semantic-release/release-notes-generator)
- [@semantic-release/changelog](https://github.com/semantic-release/changelog)
- [@semantic-release/github](https://github.com/semantic-release/github)

The workflow must be approved by one of the maintainers, first.
The release configuration can be found in the [.releaserc.js](./.releaserc.js) file.
