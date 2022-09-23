package parts

import (
	"fmt"
	"strings"
)

func printInfo(label string) {
	fmt.Printf("%s %s %s\n", strings.Repeat("*", 20), label, strings.Repeat("*", 20))
}
