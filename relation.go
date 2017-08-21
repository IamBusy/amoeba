package go_transformer

import (
	"github.com/jinzhu/gorm"
	"reflect"
)

func init() {
	SetRelation(&relation{})
}

type Relation interface {
	Get(target interface{}) Relation
	From(entity interface{}, args ...interface{}) interface{}
}

type relation struct {
	attr string
}

func (r *relation) Get(target interface{}) Relation {
	return &relation{
		attr: target.(string),
	}
}

/**
 * args: gin.Context, gorm.DB
 */
func (r *relation) From(entity interface{}, args ...interface{}) interface{} {
	value := reflect.ValueOf(entity)
	valueField := value.Elem().FieldByName(StrFirstToUpper(r.attr))
	db := args[1].(*gorm.DB)
	db.Model(entity).Association(StrFirstToUpper(r.attr)).Find(valueField.Interface())
	return entity
}

func SetRelation(r Relation) {
	app.Set("_relation", r)
}
