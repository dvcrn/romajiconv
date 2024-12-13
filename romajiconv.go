package romajiconv

// ConvertFullWidthToHalf converts full-width romaji characters to half-width
// while preserving other characters like kanji and kana
func ConvertFullWidthToHalf(input string) string {
	var result []rune

	for _, r := range input {
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
		}
		result = append(result, r)
	}

	return string(result)
}
