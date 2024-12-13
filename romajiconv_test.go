package romajiconv

import (
	"testing"
)

func TestExampleConvertFullWidthToHalf(t *testing.T) {
	examples := []struct {
		input    string
		expected string
	}{
		// Original test cases
		{
			input:    "御殿場プレミアム・アウトレット／ｉＤ",
			expected: "御殿場プレミアム・アウトレット/iD",
		},
		{
			input:    "カフェ　ＣＡＦＥ",
			expected: "カフェ CAFE",
		},
		{
			input:    "テスト１２３テスト",
			expected: "テスト123テスト",
		},
		{
			input:    "ひらがなテスト",
			expected: "ひらがなテスト",
		},
		{
			input:    "カタカナテスト",
			expected: "カタカナテスト",
		},
		// Full-width numbers
		{
			input:    "０１２３４５６７８９",
			expected: "0123456789",
		},
		// Full-width uppercase letters
		{
			input:    "ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ",
			expected: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},
		// Full-width lowercase letters
		{
			input:    "ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ",
			expected: "abcdefghijklmnopqrstuvwxyz",
		},
		// Full-width punctuation
		{
			input:    "，．：；！？",
			expected: ",.:;!?",
		},
		// Full-width quotation marks and related symbols
		{
			input:    "＂＇｀＾～￣＿",
			expected: "\"'`^~~_",
		},
		// Full-width special characters (excluding dash which has special rules)
		{
			input:    "＆＠＃％＋＊＝＜＞",
			expected: "&@#%+*=<>",
		},
		// Full-width brackets
		{
			input:    "（）［］｛｝｟｠",
			expected: "()[]{}()",
		},
		// Full-width vertical bars and slashes
		{
			input:    "｜￤／＼￢",
			expected: "||/\\¬",
		},
		// Full-width currency symbols
		{
			input:    "＄￡￠￦￥",
			expected: "$£¢₩¥",
		},
		// Mixed content tests
		{
			input:    "テスト＠ｔｅｓｔ＃１２３",
			expected: "テスト@test#123",
		},
		{
			input:    "［ＴＥＳＴ］：テスト！",
			expected: "[TEST]:テスト!",
		},
		{
			input:    "価格：＄１２３．４５",
			expected: "価格:$123.45",
		},
		// Dash conversion tests
		{
			input:    "iKI－BA PARCO/iD",
			expected: "iKI-BA PARCO/iD",
		},
		{
			input:    "ｉＫＩ－ＢＡ　ＰＡＲＣＯ／ｉＤ",
			expected: "iKI-BA PARCO/iD",
		},
		{
			input:    "渋谷－東京",
			expected: "渋谷－東京",
		},
		{
			input:    "ひらがな－カタカナ",
			expected: "ひらがな－カタカナ",
		},
		// Complex mixed content
		{
			input:    "［テスト］＊＄１２３．４５＠ｔｅｓｔ．ｃｏｍ",
			expected: "[テスト]*$123.45@test.com",
		},
		{
			input:    "（株）テスト＆カンパニー￥１，０００",
			expected: "(株)テスト&カンパニー¥1,000",
		},
		// Special cases with multiple symbols
		{
			input:    "＜＜＜テスト＞＞＞",
			expected: "<<<テスト>>>",
		},
		{
			input:    "～～～＊＊＊～～～",
			expected: "~~~***~~~",
		},
	}

	for _, ex := range examples {
		result := ConvertFullWidthToHalf(ex.input)
		if result != ex.expected {
			t.Errorf("ConvertFullWidthToHalf(%q) = %q, want %q", ex.input, result, ex.expected)
		}
	}
}
