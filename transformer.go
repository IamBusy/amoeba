package amoeba

import (
	"encoding/json"
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
	return Collection(entities, transformerName, includeStr, args...)
}

func (t *Tran) Item(entity interface{}, transformerName string, includeStr string, args ...interface{}) interface{} {
	return Item(entity, transformerName, includeStr, args...)
}

func (t *Tran) Apply(entity interface{}, includeStr string, args ...interface{}) interface{} {
	res := t.Transform(entity, args...)
	for _, str := range SplitAttrs(includeStr) {
		first, rest := ParseAttrs(str)
		parser, exist := t.ParseFuncs[first]
		if exist {
			res[first] = parser(t, entity, rest, args...)
		}
	}
	return res
}

func (t *Tran) Transform(entity interface{}, args ...interface{}) map[string]interface{} {
	jsonBytes, err := json.Marshal(entity)
	if err != nil {
		return nil
	}
	var res map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &res); err != nil {
		return nil
	}
	return res
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
