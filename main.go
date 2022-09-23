package main

import (
	"go_examinations/parts"
)

var HTTP_PART_URL string = "https://httpbin.org/anything"

func main() {
	parts.StructsPartExecute()
	parts.HttpPartExecute(&HTTP_PART_URL)
	parts.PrimitivesPartExecute()
	parts.ArraysAndSlicesPartExecute()
	parts.DecoratorsPartExecute()
}
