# romajiconv

A Go package for converting full-width romaji characters to half-width while preserving other characters like kanji and kana.

## Features

- Converts full-width Latin letters (Ａ-Ｚ, ａ-ｚ) to half-width
- Converts full-width numbers (０-９) to half-width
- Converts full-width spaces and special characters (／, ＊) to half-width
- Converts full-length dash (－) to half-width (-) when surrounded by romaji
- Preserves full-length dash when used with kanji or kana
- Preserves kanji, hiragana, katakana, and other characters

## Installation

```
go get github.com/dvcrn/romajiconv
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/dvcrn/romajiconv"
)

func main() {
    input := "Ｈｅｌｌｏ　世界！"
    result := romajiconv.ConvertFullWidthToHalf(input)
    fmt.Println(result) // Output: Hello 世界！

    // Numbers
    input = "１２３４５"
    result = romajiconv.ConvertFullWidthToHalf(input)
    fmt.Println(result) // Output: 12345

    // Mixed content
    input = "ＡＢＣ　あいうえお　１２３"
    result = romajiconv.ConvertFullWidthToHalf(input)
    fmt.Println(result) // Output: ABC あいうえお 123

    // Full-length dash
    input = "ＡＢＣ－ＤＥＦ"
    result = romajiconv.ConvertFullWidthToHalf(input)
    fmt.Println(result) // Output: ABC-DEF

    // Preserved full-length dash
    input = "東京－大阪"
    result = romajiconv.ConvertFullWidthToHalf(input)
    fmt.Println(result) // Output: 東京－大阪
}
