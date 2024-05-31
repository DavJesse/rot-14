package main

func stringContains(str, subStr string) bool {
	var status bool

	for i := 0; i < len(str); i++ {
		for j := i + len(subStr) + 1; j < len(str); j++ {
			if str[i:j] == subStr {
				status = true
				break
			}
		}
	}
	return status
}
