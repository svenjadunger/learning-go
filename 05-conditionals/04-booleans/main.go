package conditionals

func In20thCentury(year int) bool {
	if year > 1900 && year < 2001 {
		return true
	} else {
		return false
	}
}
