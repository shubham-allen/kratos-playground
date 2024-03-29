# Go Kratos Sample Project for ALLEN Digital

Sample code to demonstrate writing a grpc service in golang that stores data in mysql.
Follows the code organization convention borrowed from ideas at https://github.com/golang-standards/project-layout, https://developer20.com/how-to-structure-go-code/ and https://levelup.gitconnected.com/go-project-structure-5157f458c520.

## Go Coverage Report
[![Quality Gate Status](http://3.109.248.221:9000/sonar/api/project_badges/measure?project=go-kratos-sample&metric=alert_status&token=sqb_f379fbbe74501a5c05110f069a3d0cb8a77804b0)](http://3.109.248.221:9000/sonar/dashboard?id=go-kratos-sample)
[![Coverage](http://3.109.248.221:9000/sonar/api/project_badges/measure?project=go-kratos-sample&metric=coverage&token=sqb_f379fbbe74501a5c05110f069a3d0cb8a77804b0)](http://3.109.248.221:9000/sonar/dashboard?id=go-kratos-sample)
[![Technical Debt](http://3.109.248.221:9000/sonar/api/project_badges/measure?project=go-kratos-sample&metric=sqale_index&token=sqb_f379fbbe74501a5c05110f069a3d0cb8a77804b0)](http://3.109.248.221:9000/sonar/dashboard?id=go-kratos-sample)
[![Update Release Version](https://github.com/tj-actions/coverage-badge-go/workflows/Update%20release%20version./badge.svg)](https://github.com/tj-actions/coverage-badge-go/actions?query=workflow%3A%22Update+release+version.%22)
[![Build and Deploy to Amazon ECS](https://github.com/Allen-Career-Institute/go-kratos-sample/actions/workflows/main.yml/badge.svg)](https://github.com/Allen-Career-Institute/go-kratos-sample/actions/workflows/main.yml)

## Installing golang and related tools on macOSX

- Install xcode Open Terminal and type the following command: `xcode-select --install`.
    - In the new dialog window, confirm you want to install the Xcode tools. Agree to a license agreement and wait for the installation process to complete. It might take a while.
- Install homebrew on your macOSX. 
  ```
  /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
  ```
- Install golang, protobuf, grpc, mysql: 
  ```
  brew install go protobuf grpc protoc-gen-go-grpc mysql vault
  ```
- Configure MySQL on your local instance following the instructions as emitted on the console post brew installation.
- Install MySQL Workbench if you like from: https://dev.mysql.com/downloads/workbench/
- Setup go home, path following instructions here: https://go.dev/doc/gopath_code
  ```
  export PATH=$PATH:$(go env GOPATH)/bin
  export GOPATH=$(go env GOPATH)
  ```
- The GO111MODULE should be enabled
  ```
  go env -w GO111MODULE=on
  ```
- Edit: Additional guidance available at: https://docs.google.com/document/d/1QMpBCfj48NWWEZs7pNQJEMaVW2Xl1ggrp7viZ0rahZE/edit

## Install Kratos - the go web framework we use
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
## Create a service
```
# Create a template project
kratos new go-kratos-sample

cd go-kratos-sample
# Add a proto template
kratos proto add api/user/v1/user.proto
# Generate the proto code
kratos proto client api/user/v1/user.proto
# Generate the source code of service by proto file
kratos proto server api/user/v1/user.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/go-kratos-sample -conf ./configs
```

## Validation
Follow the steps described here to configure the validation middleware in the framework and leverage the proto level
validations: https://go-kratos.dev/en/docs/component/middleware/validate/

Detailed list of available validations to configure can be read here: https://github.com/bufbuild/protoc-gen-validate/blob/main/validate/validate.proto

Generate proto
```
protoc --proto_path=. \              
           --proto_path=./third_party \
           --go_out=paths=source_relative:. \
           --validate_out=paths=source_relative,lang=go:. \
           api/user/v1/user.proto
```

HTTP
```
httpSrv := http.NewServer(
  http.Address(":8000"),
  http.Middleware(
  validate.Validator(),
))
```

gRPC
```
grpcSrv := grpc.NewServer(
  grpc.Address(":9000"),
  grpc.Middleware(
  validate.Validator(),
))
```


## Errors
- Installation `go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest`
- Add error to proto file
```
enum ErrorReason {
  // Set default error code.
  option (errors.default_code) = 500;
  
  // sample custom error code
  USER_NOT_FOUND = 0 [(errors.code) = 404];
  CONTENT_MISSING = 1 [(errors.code) = 400];
  USER_UNSPECIFIED = 2 [(errors.code) = 500];
  
  // sample video processing error
  VIDEO_PROCESSING_ERROR = 3 [(errors.code) = 500];
}
```
- Generate proto using: make error or make all
- Create error by using the the code that is generated by proto
  ```
  v1.ErrorUserNotFound("user not found for the userId : %s", "36546547")
  ```
- Create error by using pre-defined set of standard errors provided in the error package present in types.go
  ```
  errors.BadRequest("FAILED_PRECONDITION", "Precondition failed")
  ```
- Add metadata to error to add any extra information
  ```
  err := errors.BadRequest("USER_NAME_EMPTY", "user name is empty")
  err = err.WithMetadata(map[string]string{
  "foo": "bar",
  })
  ```

For further info,
- https://acikota.atlassian.net/wiki/spaces/Platform/pages/1900908/Error+Handling
- https://go-kratos.dev/en/docs/component/errors

## Register the newly created service
HTTP API is an `http.Handler`, which is generated by `protoc-gen-go-http` plugin, can be registered into HTTP Server.
Edit `internal/server/http.go` file to register the new service.
```
import "github.com/go-kratos/kratos/v2/transport/http"
import v1 "go-kratos-sample/api/user/v1"

user := &UserService{}
srv := http.NewServer(http.Address(":8000"))
v1.RegisterUserHTTPServer(srv, user)
```

gRPC API is a gRPC Register, which is generated by `protoc-gen-go-grpc` plugin, can be registered into GRPC Server.
Edit `internal/server/grpc.go` file to register the new service.
```
import "github.com/go-kratos/kratos/v2/transport/grpc"
import v1 "go-kratos-sample/api/user/v1"

user := &UserService{}
srv := grpc.NewServer(grpc.Address(":9000"))
v1.RegisterUserServer(srv, user)
```

## Add code
- Add business logic handler code (like public repositories, business logic across multiple entities, format conversions), in `internal/biz/user_handler.go`
- Add entities, migrations, db access related code in 
  - `internal/data/entity/user_entity.go`
  - `internal/data/user_repository_impl.go`
  - tests in `internal/data/user_repository_impl_test.go`
- Add service code in `internal/service/user_service.go`
- Refer more at: 
  - https://go-kratos.dev/en/docs/getting-started/plugin
  - https://go-kratos.dev/en/docs/getting-started/examples

## Update the wiring 
Add and return the new service instance in the provider set in `internal/service/service.go`
```
var ProviderSet = wire.NewSet(NewUserService)
```

Add the new data access repo(s) in the provider set in `internal/data/data.go`
```
var ProviderSet = wire.NewSet(NewData, NewUserRepo)
```

Add the new business logic provider(s) in `biz` in the provider set in `internal/biz/biz.go`
```
var ProviderSet = wire.NewSet(NewUserHandler)
```

Add the new data access logic in `data` in the provider set in `internal/data/data.go`
```
var ProviderSet = wire.NewSet(NewData, NewUserRepository)
```

## Automated Initialization / Dependency Injection (wire)
After configuring the right set of providers as described in previous step, regenerate the dependency injection wiring 
using the steps below.
```
# install wire
go get github.com/google/wire/cmd/wire

# generate/regenrate wire dependency injections
cd cmd/go-kratos-sample
wire
```
The above command would update the `internal/cmd/go-kratos-sample/wire_gen.go` with the correct initialization of your 
newly added code.

## Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```

## To locally run and test your server
Run the server using the command
```
kratos run
```

Install `grpcui` and `grpcurl` using brew
```
brew install grpcurl grpcui
```

Start a browser based UI to test the grpc services using
```
grpcui -plaintext localhost:9000
```

You can test the new grpc server on browser using the interface.

You can also test the http version of the code on port `8000` using `CURL` command.

## Docker
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```

## Integrate private repos
```bash
# add the exact path which will be used as an import statement in other repos to your go.mod file in your private repo, 
# ex. module github.com/Allen-Career-Institute/go-kratos-commons

# run these commands in your importing repo
git config --global url."ssh://git@github.com".insteadOf "https://github.com"
git config --list --show-origin
go env -w GOPRIVATE="github.com/<my_user>/<my_privaterepo>"
# ex. go env -w GOPRIVATE="github.com/Allen-Career-Institute/go-kratos-commons"
go get <import path>
# ex. go get github.com/Allen-Career-Institute/go-kratos-commons

# After version v1.0 of the imported repo, use the following command to get latest commits from a branch
go get -u github.com/Allen-Career-Institute/go-kratos-commons/middleware/circuitbreaker@master
```

## Circuit Breaker Integration
```bash
CB := circuitbreaker.NewCircuitBreaker("demo_circuit_breaker", 1, 
        30*time.Second, 30*time.Second, 0.6, 
        func(err error) bool {
			if errors.IsGatewayTimeout(err) {
				return false
			}
			return true
		}, func() {}, func() {})
		
httpClient, _ := http.NewClient(context.Background(), 
                    http.WithMiddleware(circuitbreaker.Breaker(CB)))
err := httpClient.Invoke(ctx, hp.MethodGet, "localhost:8000/v1/ping",
		nil, http.EmptyCallOption{}, http.EmptyCallOption{})
```
- refer this document for detailed explanation https://acikota.atlassian.net/wiki/spaces/Platform/pages/12550266/Circuit+Breaker

## Linting
Install golangci-lint
```
brew install golangci-lint
```
Install pre-commit
```
brew install pre-commit
```
- Run golangci-lint run before raising any pull request and committing any code, it will analyse your Go code and report any linting issues found.
- Common used linters like gofmt, govet, errorcheck, and unused are enabled by default, to enable additional linters configurations needs to be added in .golanci.yml.

## Read more about -
- about wire at https://github.com/google/wire/blob/main/_tutorial/README.md
- about kratos at https://github.com/go-kratos
- about golangci-lint at https://golangci-lint.run

## Refer the Google AIP for designing APIs
Refer: https://google.aip.dev/general
These are pretty good standards to maintain consistency in our API design.
Also, these can be referred to define the `http` interface for the protos we are declaring.


## Other related reading material
- Top Go Web Frameworks: https://www.atatus.com/blog/go-web-frameworks/#gorilla
- https://gorm.io/
- https://go-kratos.dev/en/docs/component/log
- https://go-kratos.dev/en/docs/component/metrics
- https://go-kratos.dev/en/docs/component/middleware/circuitbreaker
- https://go-kratos.dev/en/docs/component/middleware/auth
- https://go-kratos.dev/en/docs/component/middleware/ratelimit
- https://github.com/marketplace/actions/go-coverage-report

That's all folks!

