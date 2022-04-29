package leetcode

func IsValid(s string) bool {

	l := len(s)
	if l%2 != 0 {
		return false
	}

	px := true

	for i := 0; i < l; i += 2 {
		if s[i] == '(' && s[i+1] != ')' {
			px = false
			break
		}
		if s[i] == '[' && s[i+1] != ']' {
			px = false
			break
		}
		if s[i] == '{' && s[i+1] != '}' {
			px = false
			break
		}

		if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			px = false
			break
		}
	}

	if px {
		return true
	}

	for i := 0; i < l/2; i++ {
		if s[i] == '(' && s[l-i-1] != ')' {
			return false
		}
		if s[i] == '[' && s[l-i-1] != ']' {
			return false
		}
		if s[i] == '{' && s[l-i-1] != '}' {
			return false
		}
		if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			return false
		}
	}

	return true

}
