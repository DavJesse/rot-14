package main

func rot14(str string) string {
	var result int

	for _, r := range str {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
			if r >= 'a' && r <= 'z' {
				result += r + 14
			}
		}
	}
}