// Code generated by 'github.com/traefik/yaegi/extract index/suffixarray'. DO NOT EDIT.

// +build go1.15,!go1.16

package stdlib

import (
	"index/suffixarray"
	"reflect"
)

func init() {
	Symbols["index/suffixarray"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"New": reflect.ValueOf(suffixarray.New),

		// type definitions
		"Index": reflect.ValueOf((*suffixarray.Index)(nil)),
	}
}
