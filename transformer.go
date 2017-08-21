package amoeba

import (
	"fmt"
	"reflect"
	"strings"
)

type ParseFunc func(transformer Transformer, entity interface{}, includeStr string, args ...interface{}) interface{}

type Transformer interface {
	Collection(entities interface{}, transformerName string, includeStr string, args ...interface{}) interface{}
	Item(entities interface{}, transformerName string, includeStr string, args ...interface{}) interface{}
	Apply(entity interface{}, includeStr string, args ...interface{}) interface{}
	Transform(entity interface{}, args ...interface{}) map[string]interface{}
	Include(entityName string, parser ParseFunc)
	RegisterIncluder()
}

type Tran struct {
	AvailableInclude []string
	ParseFuncs       map[string]ParseFunc
}

func (t *Tran) Collection(entities interface{}, transformerName string, includeStr string, args ...interface{}) interface{} {
	transformer := app.MustGet(getKeyByName(transformerName)).(Transformer)
	v := reflect.ValueOf(entities)
	var res []interface{}
	if v.Kind() != reflect.Slice {
		fmt.Println(v.Kind())
		panic("entity should be slice")
	}
	for i := 0; i < v.Len(); i++ {
		res = append(res, transformer.Apply(v.Index(i).Interface(), includeStr, args))
	}
	return res
}

func (t *Tran) Item(entity interface{}, transformerName string, includeStr string, args ...interface{}) interface{} {
	transformer := app.MustGet(getKeyByName(transformerName)).(Transformer)
	return transformer.Apply(entity, includeStr, args)
}

func (t *Tran) Apply(entity interface{}, includeStr string, args ...interface{}) interface{} {
	res := t.Transform(entity, args)
	for _, str := range strings.Split(includeStr, ";") {
		first, rest := SplitAttr(str)
		parser, exist := t.ParseFuncs[first]
		if exist {
			res[first] = parser(t, entity, rest, args)
		}
	}
	return res
}

func (t *Tran) Transform(entity interface{}, args ...interface{}) map[string]interface{} {
	return Struct2Map(entity)
}

/**
 * Register including func
 */
func (t *Tran) Include(entityName string, parser ParseFunc) {
	if t.ParseFuncs == nil {
		t.ParseFuncs = make(map[string]ParseFunc)
	}
	t.ParseFuncs[entityName] = parser
}

func (t *Tran) RegisterIncluder() {
}
