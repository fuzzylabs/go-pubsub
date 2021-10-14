# PubSub
This library provides generic methods for decoding and publishing messages to Google Pub/Sub.

## Usage
You can install the module directly from GitHub

```shell
go get -u github.com/fuzzylabs/go-pubsub@<version>
```

Where version can point to a commit hash or a branch, for example:

```shell
go get -u github.com/fuzzylabs/go-pubsub@9302e1d
```

or 

```shell
go get -u github.com/fuzzylabs/go-pubsub@master
```

You can then import the library as follows:
```go
import (
	hubspot "github.com/fuzzylabs/go-pubsub"
)
```

## Examples
### Decoding Push Subscription Messages
```go
package main

import (
	pubsub "github.com/fuzzylabs/go-pubsub"
	"net/http"
)

func Entrypoint(w http.ResponseWriter, r *http.Request) {
	projectID := "test"
	pubsubApi, err := pubsub.NewPubSub(projectID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	message, err := pubsubApi.DecodeBody(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
```

### Publishing Messages to Pub/Sub Topic
```go
package main

import (
	pubsub "github.com/fuzzylabs/go-pubsub"
	"github.com/golang/protobuf/ptypes/empty"
)

func Entrypoint() {
	projectID := "test"
	pubsubApi, err := pubsub.NewPubSub(projectID)
	if err != nil {
		panic("")
	}
	
	// Replace with your message that satisfies `proto.Message` interface
	message := &empty.Empty{}
	pubsubApi.PublishMessage("test-topic", message)
}
```

## Mocking
`moq` is used to generate mocks:
* Mocks for external interfaces to use within unit tests
* Mocks for `go-pubsub` API interfaces, for to make testing of applications that use the library easier

```
go generate
```

## Testing
```
go vet
go test -coverprofile=coverage.out
go tool cover -html=coverage.out # To view test coverage
```