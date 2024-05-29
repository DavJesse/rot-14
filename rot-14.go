package main

func rot14(str string) string {
	var result string
	var rot rune

	for _, r := range str {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
			rot = r + 14
			if r >= 'a' && r <= 'z' {
				if rot > 'z' {
					rot -= 26
					result += string(rot)
				}
				result += string(rot)
				continue
			}
			if r >= 'A' && r <= 'Z' {
				if rot > 'Z' {
					rot -= 26
					result += string(rot)
				}
				result += string(rot)
				continue
			}

		}
		result += string(r)
	}
	return result
}
