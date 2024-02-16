package identifier

type Identifier struct{}

func (id Identifier) ValidateIdentifier(s string) bool {
	var achar rune
	var valid_id bool = false

	if len(s) == 0 {
		return false
	}

	achar = rune(s[0])
	valid_id = id.valid_s(achar)
	if len(s) > 1 {
		for i := 1; i < len(s); i++ {
			achar = rune(s[i])
			if !id.valid_f(achar) {
				valid_id = false
			}
		}
	}
	return valid_id && (len(s) >= 1) && (len(s) <= 6)
}

func (id Identifier) valid_s(ch rune) bool {
	return ((ch >= 'A') && (ch <= 'Z')) ||
		((ch >= 'a') && (ch <= 'z'))
}

func (id Identifier) valid_f(ch rune) bool {
	return ((ch >= 'A') && (ch <= 'Z')) ||
		((ch >= 'a') && (ch <= 'z')) ||
		((ch >= '0') && (ch <= '9'))
}
