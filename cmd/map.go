package main

import (
	"errors"
	"fmt"
	"reflect"
)

func foo0() int {
	fmt.Println("running foo0: ")
	return 100
}

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("the number of params is not adapted")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

func main() {
	funcs := map[string]interface{}{
		"foo0": foo0,
	}
	if result, err := Call(funcs, "foo0"); err == nil {
		for _, r := range result {
			fmt.Printf("  return: type=%v, value=[%d]\n", r.Type(), r.Int())
		}
	}
}
