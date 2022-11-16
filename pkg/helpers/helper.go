package helpers

import (
	"encoding/json"
	"reflect"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func unmarshal(data []byte, dst interface{}) error {

	typ := reflect.ValueOf(dst).Elem().Type()
	tmp := reflect.New(typ)
	err := json.Unmarshal(data, tmp.Interface())
	if err == nil {
		reflect.ValueOf(dst).Elem().Set(tmp.Elem())
	}
	return err
}
