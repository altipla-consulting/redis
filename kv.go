package redis

import (
	"fmt"

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

type Int64KV struct {
	db  *Database
	key int64
}

func (kv *Int64KV) Set(value int64) error {
	return errors.Trace(kv.db.sess.Set(kv.db.sess.Set(kv.key, value, 0).Err())
}

func (kv *Int64KV) Get() (int64, error) {
	result, err := kv.db.sess.Get(fmt.Sprintf("%d", kv.key).Int64()
	if err != nil {
		return nil, errors.Trace(err)
	}
	
	return result, nil
}

func (kv *Int64KV) Exists() (bool, error) {
	result, err := kv.db.sess.Exists(fmt.Sprintf("%d", kv.key).Result()
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
