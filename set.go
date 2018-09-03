package redis

import (
	"github.com/go-redis/redis"
)

type StringsSet struct {
	db  *Database
	key string
}

func (set *StringsSet) Members() ([]string, error) {
	result, err := set.db.sess.SMembers(set.key).Result()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (set *StringsSet) Add(values ...string) error {
	members := make([]interface{}, len(values))
	for i := range values {
		members[i] = values[i]
	}

	return set.db.sess.SAdd(set.key, members...).Err()
}

func (set *StringsSet) Remove(values ...string) error {
	members := make([]interface{}, len(values))
	for i := range values {
		members[i] = values[i]
	}

	return set.db.sess.SRem(set.key, members...).Err()
}

func (set *StringsSet) SortAlpha() ([]string, error) {
	result, err := set.sort(&redis.Sort{Alpha: true})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (set *StringsSet) sort(sort *redis.Sort) ([]string, error) {
	result, err := set.db.sess.Sort(set.key, sort).Result()
	if err != nil {
		return nil, err
	}

	return result, nil
}
