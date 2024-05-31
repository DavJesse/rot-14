package main

func stringContains(str, subStr string) bool {
	var status bool

	for i := 0; i < len(str); i++ {
		if str[i:i+len(subStr)] == subStr {
			status = true
			break
		}
	}
	return status
}
