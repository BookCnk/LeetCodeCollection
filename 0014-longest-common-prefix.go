package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for _, str := range strs {
		for len(str) < len(prefix) || str[:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}

func main() {
	// ทดสอบฟังก์ชัน longestCommonPrefix
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // Output: "fl"
}
