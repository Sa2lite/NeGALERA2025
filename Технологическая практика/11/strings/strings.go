package main

import "fmt"

func findLongest(strings ...string) string {
	if len(strings) == 0 {
		return ""
	}

	longest := strings[0]
	for _, s := range strings[1:] {
		if len(s) > len(longest) {
			longest = s
		}
	}
	return longest
}

func main() {
	result := findLongest(
		"ААА",
		"АААААААААА",
		"Широта",
		"УЗОТААААААААААААААААААААА",
	)
	fmt.Println(result)
}
