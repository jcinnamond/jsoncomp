package jsoncomp

import (
	"fmt"
)

func comp(a, b string) {
	eq, err := Equal(a, b)
	if err != nil {
		fmt.Println("Can't compare JSON:", err)
	} else if eq {
		fmt.Println("JSON is equal")
	} else {
		fmt.Println("JSON is different")
	}
}

func ExampleInvalidJson() {
	str1 := `{"key1":`
	str2 := `{"key1": "value1", "key2": [null, 4, 4.1, true]}`
	comp(str1, str2)
	// Output: Can't compare JSON: `{"key1":` is not valid JSON
}

func ExampleEqual() {
	str1 := `{"key1": "value1", "key2": [null, 4, 4.1, true]}`
	str2 := `{"key1": "value1", "key2": [null, 4, 4.1, true]}`
	comp(str1, str2)
	// Output: JSON is equal
}

func ExampleUnorderedEqual() {
	str1 := `{"key1": "value1", "key2": [null, 4, 4.1, true]}`
	str2 := `{"key2": [null, 4, 4.1, true], "key1": "value1"}`
	comp(str1, str2)
	// Output: JSON is equal
}

func ExampleDifferentKeys() {
	str1 := `{"key1": "value"}`
	str2 := `{"key2": "value"}`
	comp(str1, str2)
	// Output: JSON is different
}

func ExampleDifferentStrings() {
	str1 := `{"key1": "value"}`
	str2 := `{"key1": "different value"}`
	comp(str1, str2)
	// Output: JSON is different
}

func ExampleDifferentNumbers() {
	str1 := `{"key1": 1}`
	str2 := `{"key1": 2.3}`
	comp(str1, str2)
	// Output: JSON is different
}

func ExampleDifferentBoolean() {
	str1 := `{"key1": true}`
	str2 := `{"key1": false}`
	comp(str1, str2)
	// Output: JSON is different
}

func ExampleDifferentArrays() {
	str1 := `{"key1": [1, 2]}`
	str2 := `{"key1": [2, 1]}`
	comp(str1, str2)
	// Output: JSON is different
}

func ExampleDifferentObjects() {
	str1 := `{"key1": "value1"}`
	str2 := `{"key1": "value1", "key2": "value2"}`
	comp(str1, str2)
	// Output: JSON is different
}

func ExampleDifferentTypes() {
	str1 := `{"key1": "value1"}`
	str2 := `{"key1": 2}`
	comp(str1, str2)
	// Output: JSON is different
}
