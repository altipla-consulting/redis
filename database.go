package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

type Database struct {
	app  string
	sess *redis.Client
}

func Open(hostname, applicationName string) *Database {
	return &Database{
		app:  applicationName,
		sess: redis.NewClient(&redis.Options{Addr: hostname}),
	}
}

func (db *Database) Set(key string) *Set {
	return &Set{
		db:  db,
		key: fmt.Sprintf("%s:%s", db.app, key),
	}
}
