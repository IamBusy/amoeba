package go_transformer

import (
	"github.com/gin-gonic/gin"
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
func Collection(entities interface{}, transformerName string, ctx *gin.Context) []interface{} {
	transformer := app.MustGet(PREFIX + transformerName).(Transformer)
	includeStr := "roles.permissions" //ctx.Param("include")
	v := reflect.ValueOf(entities)
	var res []interface{}
	if v.Kind() != reflect.Slice {
		panic("entity should be slice")
	}
	for i := 0; i < v.Len(); i++ {
		res = append(res, transformer.Apply(v.Index(i).Interface(), includeStr, ctx))
	}
	return res
}

func getKeyByName(name string) string {
	return PREFIX + name
}
