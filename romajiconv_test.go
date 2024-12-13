package romajiconv

import (
	"testing"
)

func TestExampleConvertFullWidthToHalf(t *testing.T) {
	examples := []struct {
		input    string
		expected string
	}{
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
		{
			input:    "ＡBＣDｅFｇH",
			expected: "ABCDeFgH",
		},
		{
			input:    "ｈｅｌｌｏHELLO",
			expected: "helloHELLO",
		},
		{
			input:    "テスト　　　テスト",
			expected: "テスト   テスト",
		},
		{
			input:    "Ａ　Ｂ　Ｃ",
			expected: "A B C",
		},
		{
			input:    "１２３456７８９",
			expected: "123456789",
		},
		{
			input:    "テスト１２３test４５６TEST",
			expected: "テスト123test456TEST",
		},
		{
			input:    "東京ＴＯＫＹＯトウキョウ",
			expected: "東京TOKYOトウキョウ",
		},
		{
			input:    "スーパー　ＳＵＰＥＲ　１２３",
			expected: "スーパー SUPER 123",
		},
		// Store names with mixed scripts and symbols
		{
			input:    "フアミリ―マ―トタイシドウサンチ／ｉＤ",
			expected: "フアミリ―マ―トタイシドウサンチ/iD",
		},
		{
			input:    "ＰａｙＰａｙ＊ロイヤールタリー",
			expected: "PayPay*ロイヤールタリー",
		},
		{
			input:    "ＥＩＧＨＴ　ｓａｎｇｅｎｊａｙａ／ｉＤ",
			expected: "EIGHT sangenjaya/iD",
		},
		// Half-width katakana cases
		{
			input:    "ﾗｸﾃﾝﾍﾟｲ*ﾋﾟｻﾞﾊｯﾄ",
			expected: "ﾗｸﾃﾝﾍﾟｲ*ﾋﾟｻﾞﾊｯﾄ", // Should preserve half-width katakana
		},
		// Complex store names with mixed width
		{
			input:    "ダイニングバー　ＴＡＫＥＮＯＴＳＵ／ｉＤ",
			expected: "ダイニングバー TAKENOTSU/iD",
		},
		{
			input:    "ＳＨＥＷＯＬＦＤＩＮＥＲ　渋谷／ｉＤ",
			expected: "SHEWOLFDINER 渋谷/iD",
		},
		{
			input:    "ＬＩＱＵＩＤ　ＦＡＣＴＯＲＹ／ｉＤ",
			expected: "LIQUID FACTORY/iD",
		},
		// Special characters and brand names
		{
			input:    "ＳＱ＊ＶＥＮＴ",
			expected: "SQ*VENT",
		},
		{
			input:    "Ｅｄｙチャージ",
			expected: "Edyチャージ",
		},
		// Mixed Japanese and English with numbers
		{
			input:    "ロ―ソンＮＬ　コマザワ５チヨウメ／ｉＤ",
			expected: "ロ―ソンNL コマザワ5チヨウメ/iD",
		},
		// International brand names
		{
			input:    "UBER *EATS HELP.UBER.COM",
			expected: "UBER *EATS HELP.UBER.COM", // Already half-width
		},
		{
			input:    "テスト（ＴＥＳＴ）",
			expected: "テスト(TEST)",
		},
		{
			input:    "［テスト］＊（ＴＥＳＴ）",
			expected: "[テスト]*(TEST)",
		},
		{
			input:    "スーパー［ＳＵＰＥＲ］（１２３）",
			expected: "スーパー[SUPER](123)",
		},
		{
			input:    "（株）テスト［ＴＥＳＴ］",
			expected: "(株)テスト[TEST]",
		},
		// New test cases for full-length dash
		{
			input:    "iKI－BA PARCO/iD",
			expected: "iKI-BA PARCO/iD",
		},
		{
			input:    "ｉＫＩ－ＢＡ　ＰＡＲＣＯ／ｉＤ",
			expected: "iKI-BA PARCO/iD",
		},
		{
			input:    "渋谷－東京", // Dash between kanji should be preserved
			expected: "渋谷－東京",
		},
		{
			input:    "ひらがな－カタカナ", // Dash between kana should be preserved
			expected: "ひらがな－カタカナ",
		},
		{
			input:    "ABC－DEF",
			expected: "ABC-DEF",
		},
		{
			input:    "ＡＢＣ－ＤＥＦ",
			expected: "ABC-DEF",
		},
		{
			input:    "渋谷－TOKYO", // Mixed case, should preserve dash
			expected: "渋谷－TOKYO",
		},
		{
			input:    "SHIBUYA－渋谷", // Mixed case, should preserve dash
			expected: "SHIBUYA－渋谷",
		},
	}

	for _, ex := range examples {
		result := ConvertFullWidthToHalf(ex.input)
		if result != ex.expected {
			t.Errorf("ConvertFullWidthToHalf(%q) = %q, want %q", ex.input, result, ex.expected)
		}
	}
}
