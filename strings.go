package main

func stringContains(str, subStr string) bool {
	lenStr := len(str)
	lenSubStr := len(subStr)

	if lenStr < lenSubStr {
		return false
	}

	for i := 0; i < lenStr; i++ {
		j := i + lenSubStr

		if j <= lenStr && str[i:j] == subStr {
				return true
			}
	}
	return false
}
