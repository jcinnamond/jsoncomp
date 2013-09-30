// Package jsoncomp contains methods for comparing JSON strings.
package jsoncomp

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// CompObjects takes two JSON objects (e.g., obtained from Json.Unmarshal)
// and returns true if they are equivalent. Simple objects (strings, bool,
// float64 and nil) are compared directly. Complex types (structs and
// arrays) are compared by comparing all of their elements.
func CompObjects(a, b interface{}) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false
	}

	switch a.(type) {
	case map[string]interface{}:
		x := a.(map[string]interface{})
		y := b.(map[string]interface{})

		if len(x) != len(y) {
			return false
		}

		eq := true
		for k, v := range x {
			v2, found := y[k]
			eq = eq && found && CompObjects(v, v2)
		}
		return eq
	case []interface{}:
		x := a.([]interface{})
		y := b.([]interface{})

		if len(x) != len(y) {
			return false
		}

		eq := true
		for i, v := range x {
			eq = eq && CompObjects(v, y[i])
		}
		return eq
	case string, float64, bool, nil:
		return a == b
	default:
		// Don't know how to compare these types
		return false
	}
}

// Equal compares two strings and returns true if they contain equivalent
// JSON. An error is returned if either string does not contain valid JSON.
func Equal(a, b string) (bool, error) {
	var json1, json2 interface{}

	if err := json.Unmarshal([]byte(a), &json1); err != nil {
		return false, fmt.Errorf("`%v` is not valid JSON", a)
	}

	if err := json.Unmarshal([]byte(b), &json2); err != nil {
		return false, fmt.Errorf("`%v` is not valid JSON", b)
	}

	return CompObjects(json1, json2), nil
}
