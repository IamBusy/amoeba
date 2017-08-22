package amoeba

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

func clean(target string) string {
	target = strings.Trim(target," ")
	start,end := 0,len(target)-1
	for end > start && target[start] == '(' && target[end] == ')' {
		start++
		end--
	}
	return target[start:end+1]
}

func SplitAttrs(includeStr string) (res []string)  {
	var parentheses = 0
	var lastSplitPoint = 0
	for i := 0; i < len(includeStr); i++ {
		if includeStr[i] == '(' {
			parentheses++
			continue
		}

		if includeStr[i] == ')' {
			parentheses--
		}

		if includeStr[i] == ATTR_DELI && parentheses == 0 {
			res = append(res, clean(includeStr[lastSplitPoint:i]))
			lastSplitPoint = i+1
		}
	}
	res = append(res, strings.Trim(includeStr[lastSplitPoint:]," "))
	return
}


func ParseAttrs(target string) (first string, rest string) {
	var parentheses = 0
	for i := 0; i < len(target); i++ {
		if target[i] == '(' {
			parentheses++
			continue
		}

		if target[i] == ')' {
			parentheses--
		}

		if target[i] == ATTR_DOT && parentheses == 0 {
			first = target[:i]
			rest = clean(target[i+1:])
			return
		}
	}
	return
}

func IsSingular(word string) bool {
	return word == inflection.Singular(word)
}

