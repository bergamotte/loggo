package print

import (
    "fmt"
)

func PrintFiles (str []string) {
    for i, s := range str {
        _ = i //blank identifier to avoid "declared but not used"
        fmt.Println(s)
    }
}
