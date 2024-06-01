package main

func stringContains(str, subStr string) bool {
	lenStr := len(str)
	lenSubStr := len(subStr)

	// Exit function when substring is incaompatible(is longer)
	if lenStr < lenSubStr {
		return false
	}

	for i := 0; i < lenStr; i++ {
		j := i + lenSubStr

		// Compare chunks of str to substring
		if j <= lenStr && str[i:j] == subStr {
			return true
		}
	}
	return false
}
