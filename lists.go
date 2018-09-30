package redis

import (
	"fmt"
	"reflect"

	"github.com/golang/protobuf/proto"
)

type ProtoList struct {
	db  *Database
	key string
}

func (list *ProtoList) Len() (int64, error) {
	return list.db.sess.LLen(list.key).Result()
}

func (list *ProtoList) GetAll(result interface{}) error {
	rt := reflect.TypeOf(result)
	rv := reflect.ValueOf(result)
	msg := reflect.TypeOf((*proto.Message)(nil)).Elem()
	if rt.Kind() != reflect.Ptr || rt.Elem().Kind() != reflect.Slice || !rt.Elem().Elem().Implements(msg) {
		return fmt.Errorf("redis: expected a pointer to a slice for the result, received %T", result)
	}

	dest := reflect.MakeSlice(rt.Elem(), 0, 0)

	redisResult, err := list.db.sess.LRange(list.key, 0, -1).Result()
	if err != nil {
		return err
	}
	for _, item := range redisResult {
		value := reflect.New(rt.Elem().Elem().Elem())
		if err := proto.Unmarshal([]byte(item), value.Interface().(proto.Message)); err != nil {
			return err
		}

		dest = reflect.Append(dest, value)
	}

	rv.Elem().Set(dest)

	return nil
}

func (list *ProtoList) Add(values ...proto.Message) error {
	members := make([]interface{}, len(values))
	for i, value := range values {
		bytes, err := proto.Marshal(value)
		if err != nil {
			return err
		}

		members[i] = string(bytes)
	}

	return list.db.sess.LPush(list.key, members...).Err()
}