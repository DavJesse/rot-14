package main

func toRot14(str string) string {
	var result string
	var rot rune

	for _, r := range str {

		// Capture and convert charaters in lowercase
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
			rot = r + 14
			if r >= 'a' && r <= 'z' {
				if rot > 'z' {
					rot -= 26
					result += string(rot)
					continue
				}
				result += string(rot)
				continue
			}

			// Capture and convert charaters in uppercase
			if r >= 'A' && r <= 'Z' {
				if rot > 'Z' {
					rot -= 26
					result += string(rot)
					continue
				}
				result += string(rot)
				continue
			}

		}

		result += string(r)
	}
	return result
}

func fromRot14(str string) string {
	var result string
	var unRot rune

	for _, r := range str {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
			unRot = r - 14

			// Capture and convert charaters in lowercase
			if r >= 'a' && r <= 'z' {
				if unRot < 'a' {
					unRot += 26
					result += string(unRot)
					continue
				}
				result += string(unRot)
				continue
			}

			// Capture and convert charaters in uppercase
			if r >= 'A' && r <= 'Z' {
				if unRot < 'A' {
					unRot += 26
					result += string(unRot)
					continue
				}
				result += string(unRot)
				continue
			}

		}
		result += string(r)
	}

	return result
}
