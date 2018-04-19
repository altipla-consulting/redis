package redis

import (
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
