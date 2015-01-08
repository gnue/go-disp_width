# disp_width package

Go言語で半角全角の混じった文字列の文字幅を計算する

* [moznion/go-unicode-east-asian-width](https://github.com/moznion/go-unicode-east-asian-width) を使用

## Example

```go
package main

import (
	"fmt"
	"strings"

	"github.com/gnue/go-disp_width"
)

func main() {
	src := "国境の長いトンネルを抜けると雪国であった"

	n := disp_width.Measure(src)
	fmt.Printf("　表示幅: 半角%d文字\n", n)

	t, _ := disp_width.Truncate(src, 20, "...")
	fmt.Printf("切り詰め %q\n", t)

	s, rest := disp_width.Truncate(src, 50, "...")
	fmt.Printf("　左寄せ: %q\n", s+strings.Repeat(" ", rest))
	fmt.Printf("中央寄せ: %q\n", strings.Repeat(" ", rest/2)+s+strings.Repeat(" ", rest-rest/2))
	fmt.Printf("　右寄せ: %q\n", strings.Repeat(" ", rest)+s)
}
```


