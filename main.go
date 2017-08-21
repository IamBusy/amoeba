package amoeba

import (
	"reflect"
)

const PREFIX = "transformer-"

func RegisterTransformer(name string, transformer Transformer) {
	transformer.RegisterIncluder()
	app.Set(PREFIX+name, transformer)
}

/**
 * Entrance
 */
func Collection(entities interface{}, transformerName string, includeStr string, args ...interface{}) []interface{} {
	transformer := app.MustGet(PREFIX + transformerName).(Transformer)
	v := reflect.ValueOf(entities)
	var res []interface{}
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		} else {
			panic("entities should be slice, but receive:" + v.Kind().String())
		}
	}
	for i := 0; i < v.Len(); i++ {
		res = append(res, transformer.Apply(v.Index(i).Interface(), includeStr, args...))
	}
	return res
}

func Item(entity interface{}, transformerName string, includeStr string, args ...interface{}) interface{} {
	transformer := app.MustGet(PREFIX + transformerName).(Transformer)
	return transformer.Apply(entity, includeStr, args...)
}

func getKeyByName(name string) string {
	return PREFIX + name
}
