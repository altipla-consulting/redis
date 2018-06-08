
# redis

[![GoDoc](https://godoc.org/github.com/altipla-consulting/redis?status.svg)](https://godoc.org/github.com/altipla-consulting/redis)

> Abstraction layer to access Redis with repositories and models.


### Install

```shell
go get github.com/altipla-consulting/redis
```

This library depends on the following ones:
- [github.com/go-redis/redis](github.com/go-redis/redis)
- [github.com/juju/errors](github.com/juju/errors)
- [github.com/golang/protobuf/proto](github.com/golang/protobuf/proto)
- [github.com/altipla-consulting/sentry](github.com/altipla-consulting/sentry)
- [github.com/segmentio/ksuid](github.com/segmentio/ksuid)
- [github.com/sirupsen/logrus](github.com/sirupsen/logrus)
- [golang.org/x/net/context](golang.org/x/net/context)


### Usage

Repository file:

```go
package models

import (
  "fmt"

  "github.com/altipla-consulting/redis"
  "github.com/juju/errors"
)

var Repo *MainDatabase

func ConnectRepo() error {
  Repo = &MainDatabase{
    sess: redis.Open("redis:6379", "cells"),
  }

  return nil
}

type MainDatabase struct {
  sess *redis.Database
}

func (repo *MainDatabase) Offers(hotel int64) *redis.ProtoHash {
  return repo.sess.ProtoHash(fmt.Sprintf("hotel:%d", hotel))
}

func (repo *MainDatabase) CloseOffers() *redis.Queue {
  return repo.sess.Queue("close-offers")
}
```

Usage:

```go
func run() error {
  offers := []*pbmodels.Offer{}
  if err := models.Repo.Offers(in.Hotel).GetMulti(codes, &offers); err != nil {
    return errors.Trace(err)
  }
}
```


### Queues usage

To use queues effectively a dead-letter queue must be monitorized and cleaned up regularly. For that purpose a new goroutine should be created as soon as possible after starting the application:

```go
func main() {
  // Configure a clean-up process with a Sentry notificator.
  go models.Repo.CloseOffers().CleanUpProcess(config.Settings.Sentry.DSN)
}
```


### Contributing

You can make pull requests or create issues in GitHub. Any code you send should be formatted using `gofmt`.


### License

[MIT License](LICENSE)
