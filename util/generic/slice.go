package generic

import (
	"reflect"
	"sort"
)

func IsSliceable(value interface{}) bool {
	if value == nil {
		return false
	}
	return reflect.TypeOf(value).Kind() == reflect.Slice
}

func SortAndUniquifyStringSlice(slice []string) []string {
	sort.Strings(slice)

	var newSlice []string

	for _, str := range slice {
		currentLength := len(newSlice)
		if currentLength == 0 || newSlice[currentLength-1] != str {
			newSlice = append(newSlice, str)
		}
	}

	return newSlice
}
