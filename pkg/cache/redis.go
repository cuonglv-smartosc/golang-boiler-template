package cache

import (
	"bytes"
	"context"
	"encoding/gob"
	"reflect"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	*redis.Client
}

var RedisClient *Cache

// Serialize returns a []byte representing the passed value
func Serialize(value interface{}) ([]byte, error) {
	if bytes, ok := value.([]byte); ok {
		return bytes, nil
	}

	switch v := reflect.ValueOf(value); v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []byte(strconv.FormatInt(v.Int(), 10)), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return []byte(strconv.FormatUint(v.Uint(), 10)), nil
	}

	var b bytes.Buffer
	encoder := gob.NewEncoder(&b)
	if err := encoder.Encode(value); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// Deserialize deserialices the passed []byte into a the passed ptr interface{}
func Deserialize(byt []byte, ptr interface{}) (err error) {
	if bytes, ok := ptr.(*[]byte); ok {
		*bytes = byt
		return nil
	}

	if v := reflect.ValueOf(ptr); v.Kind() == reflect.Ptr {
		switch p := v.Elem(); p.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			var i int64
			i, err = strconv.ParseInt(string(byt), 10, 64)
			if err != nil {
				return err
			}

			p.SetInt(i)
			return nil

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			var i uint64
			i, err = strconv.ParseUint(string(byt), 10, 64)
			if err != nil {
				return err
			}

			p.SetUint(i)
			return nil
		}
	}

	b := bytes.NewBuffer(byt)
	decoder := gob.NewDecoder(b)
	if err = decoder.Decode(ptr); err != nil {
		return err
	}
	return nil
}

// GetValue function get value from cache and return: exists and error
func (r *Cache) GetValue(key string, resource interface{}) (bool, error) {
	response, err := r.Get(context.Background(), key).Bytes()
	if err != nil && err.Error() != "redis: nil" {
		return false, err
	}

	if response == nil {
		return false, nil
	}

	// cache contain
	err = Deserialize(response, resource)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *Cache) CleanCacheWithKeyLike(key string) error {
	var cursor uint64
	keys, cursor, err := r.Scan(context.Background(), cursor, "*"+key+"*", 10000000).Result()
	if err != nil {
		return err
	}

	for _, k := range keys {
		err = r.Del(context.Background(), k).Err()
		if err != nil {
			return err
		}
	}
	return nil
}
