package main

import (
	"fmt"
	"reflect"

	"github.com/OtusGolang/webinars_practical_part/20-reflection/user"
)

type Student struct {
	Name string
	Age  int
}

func structToMap(iv interface{}) (map[string]interface{}, error) {
	v := reflect.ValueOf(iv)
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct, but received %T", iv)
	}

	t := v.Type()
	mp := make(map[string]interface{}, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i) // reflect.StructField
		fv := v.Field(i)    // reflect.Value
		if fv.CanInterface() {
			mp[field.Name] = fv.Interface()
		}
	}
	return mp, nil
}

func main() {
	st := Student{"Bob", 18}
	mp, err := structToMap(st)
	fmt.Printf("MAP: %#v\nERR: %s\n", mp, err)

	st1, err := structToMap(user.User{Name: "Alex", Age: 100})
	fmt.Println(st1, err)
}
