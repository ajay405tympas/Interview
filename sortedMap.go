// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
)

func main() {

	m := map[string]int{"a": 1, "b": 2, "c": 3}

	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println("keys", key, m[key])
	}

	fmt.Println("Hello, 世界")
}
