package main

func stringContains(str, subStr string) bool {
	if len(str) < len(subStr) {
		return false
	}

	for i := 0; i < len(str); i++ {
		if str[i:i+len(subStr)] == subStr {
			return true
		}
	}
	return false
}
