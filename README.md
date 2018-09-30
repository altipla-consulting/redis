
# redis

[![GoDoc](https://godoc.org/github.com/altipla-consulting/redis?status.svg)](https://godoc.org/github.com/altipla-consulting/redis)
[![Build Status](https://travis-ci.org/altipla-consulting/redis.svg?branch=master)](https://travis-ci.org/altipla-consulting/redis)

Abstraction layer to access Redis with repositories and models.


### Install

```shell
go get github.com/altipla-consulting/redis
```


### Usage

Repository file:

```go
package models

import (
  "fmt"

  "github.com/altipla-consulting/redis"
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
    return err
  }
}
```


### Contributing

You can make pull requests or create issues in GitHub. Any code you send should be formatted using `gofmt`.


### Running tests

Run the tests

```shell
make test
```


### License

[MIT License](LICENSE)
