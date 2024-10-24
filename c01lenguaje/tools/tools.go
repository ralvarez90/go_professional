package tools

import (
	"fmt"
	"math/rand"
	"strings"
)

func SystemPause() {
	fmt.Print("\nPress any key to continue . . . ")
	fmt.Scanln()
}

func RandInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func MakeTitle(s string) string {
	return fmt.Sprintf("%s %s %s", strings.Repeat("-", 30), strings.ToUpper(s), strings.Repeat("-", 30))
}
