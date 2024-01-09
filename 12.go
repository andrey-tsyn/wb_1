package main

import "fmt"

func main() {
	fmt.Printf("%v\n", getSetFromStrings("cat", "cat", "dog", "cat", "tree"))
}

func getSetFromStrings(values ...string) map[string]struct{} {
	set := make(map[string]struct{})

	for _, val := range values {
		set[val] = struct{}{}
	}

	return set
}
