package redis

import (
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/golang/protobuf/proto"
	"github.com/juju/errors"
)

type StringKV struct {
	db  *Database
	key string
}

func (kv *StringKV) Set(value string) error {
	return errors.Trace(kv.db.sess.Set(kv.key, value, 0).Err())
}

func (kv *StringKV) SetTTL(value string, ttl time.Duration) error {
	return errors.Trace(kv.db.sess.Set(kv.key, value, ttl).Err())
}

func (kv *StringKV) Get() (string, error) {
	result, err := kv.db.sess.Get(kv.key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", ErrNoSuchEntity
		}

		return "", errors.Trace(err)
	}

	return result, nil
}

func (kv *StringKV) Exists() (bool, error) {
	result, err := kv.db.sess.Exists(kv.key).Result()
	if err != nil {
		return false, errors.Trace(err)
	}

	return result == 1, nil
}

func (kv *StringKV) Delete() error {
	return errors.Trace(kv.db.sess.Del(kv.key).Err())
}

type Int32KV struct {
	db  *Database
	key string
}

func (kv *Int32KV) Set(value int32) error {
	return errors.Trace(kv.db.sess.Set(kv.key, value, 0).Err())
}

func (kv *Int32KV) SetTTL(value int32, ttl time.Duration) error {
	return errors.Trace(kv.db.sess.Set(kv.key, value, ttl).Err())
}

func (kv *Int32KV) Get() (int32, error) {
	result, err := kv.db.sess.Get(kv.key).Int64()
	if err != nil {
		if err == redis.Nil {
			return 0, ErrNoSuchEntity
		}

		return 0, errors.Trace(err)
	}

	return int32(result), nil
}

func (kv *Int32KV) Exists() (bool, error) {
	result, err := kv.db.sess.Exists(kv.key).Result()
	if err != nil {
		return false, errors.Trace(err)
	}

	return result == 1, nil
}

func (kv *Int32KV) Delete() error {
	return errors.Trace(kv.db.sess.Del(kv.key).Err())
}

type Int64KV struct {
	db  *Database
	key string
}

func (kv *Int64KV) Set(value int64) error {
	return errors.Trace(kv.db.sess.Set(kv.key, value, 0).Err())
}

func (kv *Int64KV) SetTTL(value int64, ttl time.Duration) error {
	return errors.Trace(kv.db.sess.Set(kv.key, value, ttl).Err())
}

func (kv *Int64KV) Get() (int64, error) {
	result, err := kv.db.sess.Get(kv.key).Int64()
	if err != nil {
		if err == redis.Nil {
			return 0, ErrNoSuchEntity
		}

		return 0, errors.Trace(err)
	}

	return int64(result), nil
}

func (kv *Int64KV) Exists() (bool, error) {
	result, err := kv.db.sess.Exists(kv.key).Result()
	if err != nil {
		return false, errors.Trace(err)
	}

	return result == 1, nil
}

func (kv *Int64KV) Delete() error {
	return errors.Trace(kv.db.sess.Del(kv.key).Err())
}

type ProtoKV struct {
	db  *Database
	key string
}

func (kv *ProtoKV) Set(value proto.Message) error {
	bytes, err := proto.Marshal(value)
	if err != nil {
		return errors.Trace(err)
	}

	return errors.Trace(kv.db.sess.Set(kv.key, string(bytes), 0).Err())
}

func (kv *ProtoKV) SetTTL(value proto.Message, ttl time.Duration) error {
	bytes, err := proto.Marshal(value)
	if err != nil {
		return errors.Trace(err)
	}

	return errors.Trace(kv.db.sess.Set(kv.key, string(bytes), ttl).Err())
}

func (kv *ProtoKV) Get(value proto.Message) error {
	result, err := kv.db.sess.Get(kv.key).Result()
	if err != nil {
		if err == redis.Nil {
			return ErrNoSuchEntity
		}

		return errors.Trace(err)
	}

	return errors.Trace(proto.Unmarshal([]byte(result), value))
}

func (kv *ProtoKV) Exists() (bool, error) {
	result, err := kv.db.sess.Exists(kv.key).Result()
	if err != nil {
		return false, errors.Trace(err)
	}

	return result == 1, nil
}

func (kv *ProtoKV) Delete() error {
	return errors.Trace(kv.db.sess.Del(kv.key).Err())
}

type BooleanKV struct {
	db  *Database
	key string
}

func (kv *BooleanKV) Set(value bool) error {
	return errors.Trace(kv.db.sess.Set(kv.key, value, 0).Err())
}

func (kv *BooleanKV) SetTTL(value bool, ttl time.Duration) error {
	return errors.Trace(kv.db.sess.Set(kv.key, value, ttl).Err())
}

func (kv *BooleanKV) Get() (bool, error) {
	result, err := kv.db.sess.Get(kv.key).Result()
	if err != nil {
		if err == redis.Nil {
			return false, ErrNoSuchEntity
		}

		return false, errors.Trace(err)
	}

	return strconv.ParseBool(result)
}

func (kv *BooleanKV) Exists() (bool, error) {
	result, err := kv.db.sess.Exists(kv.key).Result()
	if err != nil {
		return false, errors.Trace(err)
	}

	return result == 1, nil
}

func (kv *BooleanKV) Delete() error {
	return errors.Trace(kv.db.sess.Del(kv.key).Err())
}
