package romajiconv

func isRomaji(r rune) bool {
	switch {
	case (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z'):
		return true
	case r >= 'Ａ' && r <= 'Ｚ':
		return true
	case r >= 'ａ' && r <= 'ｚ':
		return true
	}
	return false
}

// ConvertFullWidthToHalf converts full-width romaji characters to half-width
// while preserving other characters like kanji and kana
func ConvertFullWidthToHalf(input string) string {
	var result []rune
	runes := []rune(input)

	for i, r := range runes {
		switch {
		// Full-width uppercase letters (Ａ-Ｚ)
		case r >= 'Ａ' && r <= 'Ｚ':
			r -= 0xFEE0
		// Full-width lowercase letters (ａ-ｚ)
		case r >= 'ａ' && r <= 'ｚ':
			r -= 0xFEE0
		// Full-width numbers (０-９)
		case r >= '０' && r <= '９':
			r -= 0xFEE0
		// Full-width space
		case r == '　':
			r = ' '
		// Full-width forward slash
		case r == '／':
			r = '/'
		// Full-width asterisk
		case r == '＊':
			r = '*'
		// Full-width brackets
		case r == '（':
			r = '('
		case r == '）':
			r = ')'
		case r == '［':
			r = '['
		case r == '］':
			r = ']'

		// Full-length dash (－) when surrounded by romaji
		case r == '－':
			if i > 0 && i < len(runes)-1 {
				prevIsRomaji := isRomaji(runes[i-1])
				nextIsRomaji := isRomaji(runes[i+1])
				if prevIsRomaji && nextIsRomaji {
					r = '-'
				}
			}
		}
		result = append(result, r)
	}

	return string(result)
}
