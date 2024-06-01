package main

func stringContains(str, subStr string) bool {
	if len(str) < len(subStr) {
		return false
	}

	for i := 0; i < len(str); i++ {
		if i + len(subStr) <= len(str) {
			if str[i:i+len(subStr)] == subStr {
				return true
			}
		}
		break
	}
	return false
}
