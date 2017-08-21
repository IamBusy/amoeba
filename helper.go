package go_transformer

import (
	"fmt"
	"github.com/jinzhu/inflection"
	"reflect"
	"strings"
)

func isIteratable(args interface{}) {
	val := reflect.ValueOf(args)
	fmt.Println(val)
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func StrFirstToUpper(str string) string {
	temp := strings.Split(str, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				vv[i] -= 32
				upperStr += string(vv[i]) // + string(vv[i+1])
			} else {
				upperStr += string(vv[i])
			}
		}
	}
	return upperStr
}

func SplitAttr(target string) (first string, rest string) {
	for i := 0; i < len(target); i++ {
		if target[i] == '.' {
			rest = target[i+1:]
			return
		}
		first += string(target[i])
	}
	return
}

func IsSingular(word string) bool {
	return word == inflection.Singular(word)
}
