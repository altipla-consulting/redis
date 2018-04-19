package redis

import (
	"github.com/go-redis/redis"
	"github.com/juju/errors"
)

type Set struct {
	db  *Database
	key string
}

func (set *Set) Members() ([]string, error) {
	result, err := set.db.sess.SMembers(set.key).Result()
	if err != nil {
		return nil, errors.Trace(err)
	}

	return result, nil
}

func (set *Set) Add(values ...string) error {
	members := make([]interface{}, len(values))
	for i := range values {
		members[i] = values[i]
	}

	return errors.Trace(set.db.sess.SAdd(set.key, members...).Err())
}

func (set *Set) Remove(values ...string) error {
	members := make([]interface{}, len(values))
	for i := range values {
		members[i] = values[i]
	}

	return errors.Trace(set.db.sess.SRem(set.key, members...).Err())
}

func (set *Set) SortAlpha() ([]string, error) {
	result, err := set.sort(&redis.Sort{Alpha: true})
	if err != nil {
		return nil, errors.Trace(err)
	}

	return result, nil
}

func (set *Set) sort(sort *redis.Sort) ([]string, error) {
	result, err := set.db.sess.Sort(set.key, sort).Result()
	if err != nil {
		return nil, errors.Trace(err)
	}

	return result, nil
}
