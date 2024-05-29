package main

func rot14(str string) string {
	var result string
	var rot rune

	for _, r := range str {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
			if r >= 'a' && r <= 'z' {
				rot = r + 14
				if rot < 'z' {
					rot -= 26
					result += string(rot)
				}
				result += string(rot)
			} else if r >= 'A' && r <= 'Z' {
				rot = r + 14
				if rot < 'Z' {
					rot -= 26
					result += string(rot)
				}
				result += string(rot)
			} else {
				result += string(r)
			}

		}
	}
	return result
}