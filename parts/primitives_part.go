package parts

import "fmt"

func PrimitivesPartExecute() {
	printInfo("PRIMITIVES PART")

	var a int8 = 10
	var b int16 = 12

	fmt.Println("fmt.Println(a + b) // error couse any operation only working with the same data type")
	fmt.Printf("CAST ONE OF THE VAR INTO SAME DATA TYPE \"(int16)(a) + b\" AND THEN SUCCESS: => %d\n", (int16)(a)+b)

	c := 10 // 1010
	d := 3  // 0011

	fmt.Printf("c = %d // 1010\nd = %d // 0011\n\n", c, d)
	fmt.Printf("c & d  => \"%d\" // 0010\n", c&d)
	fmt.Printf("c | d  => \"%d\" // 1011\n", c|d)
	fmt.Printf("c ^ d  => \"%d\" // 1001\n", c^d)
	fmt.Printf("c &^ d => \"%d\" // 0100\n", c&^d)

	printInfo("PRIMITIVES PART")
}
