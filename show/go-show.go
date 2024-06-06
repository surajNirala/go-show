package show

import (
	"fmt"
	"reflect"
)

func Show(input interface{}) {
	// Get the type and value of the input
	t := reflect.TypeOf(input)
	v := reflect.ValueOf(input)

	// Print the type
	fmt.Printf("Type: %s\n", t)

	// Check if the input has a length (for slices, arrays, maps, and strings)
	var length int
	switch t.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.String:
		length = v.Len()
		fmt.Printf("Length: %d\n", length)
	default:
		fmt.Println("Length: N/A")
	}

	// Print the value
	fmt.Printf("Value: %v\n", v)
}
