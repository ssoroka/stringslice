package stringslice

import (
	"fmt"
	"reflect"
)

// GetKeys returns the string keys from any map type. It will panic if the type is not a string.
func GetKeys(o interface{}) []string {
	v := reflect.ValueOf(o)
	if v.Kind() != reflect.Map {
		panic("GetKeys expected a Map but received a kind " + v.Kind().String())
	}

	result := make([]string, v.Len())
	for i, k := range v.MapKeys() {
		if k.Kind() != reflect.String {
			panic(fmt.Sprintf("GetKeys failed to convert map key to a string, it's a kind %s", k.Kind().String()))
		}
		result[i] = k.String()
	}
	return result
}

// GetValues returns the string values from any map type. It will panic if the type is not a string.
func GetValues(o interface{}) []string {
	v := reflect.ValueOf(o)
	if v.Kind() != reflect.Map {
		panic("GetValues expected a Map but received a kind " + v.Kind().String())
	}

	result := make([]string, v.Len())
	for i, k := range v.MapKeys() {
		val := v.MapIndex(k)
		if val.Kind() != reflect.String {
			panic(fmt.Sprintf("GetKeys failed to convert map key to a string, it's a kind %s", val.Kind().String()))
		}
		result[i] = val.String()
	}
	return result
}
