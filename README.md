# jsoncomp: compare any two JSON strings

jsoncomp is a Go package for comparing two JSON strings. I envisage this being
used primarily in tests.

	go get github.com/jcinnamond/jsoncomp

## Usage

	import (
		"jsoncomp"
	)

	func main() {
		str1 := `{"key1": "value1", "key2": [null, 4, 4.1, true]}`
		str2 := `{"key2": [null, 4, 4.1, true], "key1": "value1"}`

		eq, err := jsoncomp.Equal(str1, str2)
		if err != nil {
			// Invalid JSON
		} else if eq {
			// JSON is equal
		} else {
			// JSON is different
		}
	}

For more details see the examples in jsoncomp_test.go
