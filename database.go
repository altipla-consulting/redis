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

func (db *Database) StringKV(key string) *StringKV {
	return &StringKV{
		db:  db,
		key: fmt.Sprintf("%s:%s", db.app, key),
	}
}

func (db *Database) Int32KV(key string) *Int32KV {
	return &Int32KV{
		db:  db,
		key: fmt.Sprintf("%s:%s", db.app, key),
	}
}

func (db *Database) ProtoKV(key string) *ProtoKV {
	return &ProtoKV{
		db:  db,
		key: fmt.Sprintf("%s:%s", db.app, key),
	}
}

func (db *Database) ProtoHash(key string) *ProtoHash {
	return &ProtoHash{
		db:  db,
		key: fmt.Sprintf("%s:%s", db.app, key),
	}
}

func (db *Database) Queue(key string) *Queue {
	return &Queue{
		db:            db,
		queueKey:      fmt.Sprintf("%s:%s", db.app, key),
		processingKey: fmt.Sprintf("%s:%s:processing", db.app, key),
		tasksKey:      fmt.Sprintf("%s:%s:tasks", db.app, key),
	}
}
