package redis

import (
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/juju/errors"
)

type ProtoHash struct {
	db  *Database
	key string
}

func (hash *ProtoHash) PrepareInsert() *ProtoHashInsert {
	return &ProtoHashInsert{
		hash:   hash,
		fields: make(map[string]interface{}),
	}
}

// GetMulti fetchs a list of keys from the hash. Result should be a slice of proto.Message
// that will be filled with the results in the same order as the keys.
func (hash *ProtoHash) GetMulti(keys []string, result interface{}) error {
	rt := reflect.TypeOf(result)
	rv := reflect.ValueOf(result)
	msg := reflect.TypeOf((*proto.Message)(nil)).Elem()
	if rt.Kind() != reflect.Ptr || rt.Elem().Kind() != reflect.Slice || !rt.Elem().Elem().Implements(msg) {
		return errors.Errorf("redis: expected a pointer to a slice for the result, received %T", result)
	}

	dest := reflect.MakeSlice(rt.Elem(), 0, 0)

	var merr MultiError

	redisResult, err := hash.db.sess.HMGet(hash.key, keys...).Result()
	if err != nil {
		return errors.Trace(err)
	}
	for _, item := range redisResult {
		var model reflect.Value
		if item == nil {
			model = reflect.Zero(rt.Elem().Elem())
			merr = append(merr, ErrNoSuchEntity)
		} else {
			model = reflect.New(rt.Elem().Elem())
			if err := proto.Unmarshal([]byte(item.(string)), model.Interface().(proto.Message)); err != nil {
				return errors.Trace(err)
			}
			merr = append(merr, nil)
		}

		dest = reflect.Append(dest, model)
	}

	rv.Set(dest)

	if merr.HasError() {
		return merr
	}
	return nil
}

type ProtoHashInsert struct {
	hash   *ProtoHash
	fields map[string]interface{}
}

func (insert *ProtoHashInsert) Set(key string, value proto.Message) error {
	bytes, err := proto.Marshal(value)
	if err != nil {
		return errors.Trace(err)
	}

	insert.fields[key] = string(bytes)

	return nil
}

func (insert *ProtoHashInsert) Commit() error {
	return errors.Trace(insert.hash.db.sess.HMSet(insert.hash.key, insert.fields).Err())
}
