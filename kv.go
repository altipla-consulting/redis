 package redis

import (
	"strconv"

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

func (kv *StringKV) Get() (string, error) {
	result, err := kv.db.sess.Get(kv.key).Result()
	if err != nil {
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

type Int32KV struct {
	db  *Database
	key string
}

func (kv *Int32KV) Set(value int32) error {
	return errors.Trace(kv.db.sess.Set(kv.key, value, 0).Err())
}

func (kv *Int32KV) Get() (int32, error) {
	result, err := kv.db.sess.Get(kv.key).Int64()
	if err != nil {
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

func (kv *ProtoKV) Get(value proto.Message) error {
	result, err := kv.db.sess.Get(kv.key).Result()
	if err != nil {
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

type BooleanKV struct {
	db  *Database
	key string
}

func (kv *BooleanKV) Set(value bool) error {
	return errors.Trace(kv.db.sess.Set(kv.key, value, 0).Err())
}

func (kv *BooleanKV) Get() (bool, error) {
	result, err := kv.db.sess.Get(kv.key).Result()
	if err != nil {
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
