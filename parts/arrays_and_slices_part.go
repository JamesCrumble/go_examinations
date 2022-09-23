package parts

import (
	"container/list"
	"fmt"
	"unsafe"
)

type SimpleStruct struct {
	a int
	b string
	c float32
}

func ArraysAndSlicesPartExecute() {
	printInfo("ARRAYS AND SLICES AND SO ON")

	fmt.Println("CREATING LIST")

	list := list.New()

	var simpleStruct SimpleStruct = SimpleStruct{
		a: 2, b: "simple_struct", c: 12.0,
	}

	fmt.Println("CREATING SIMPLE STRUCT FOR TEST")
	fmt.Println("PUSHING ELEMENTS INTO LIST")

	list.PushFront("first_element")
	list.PushFront(2)
	list.PushFront(simpleStruct)

	fmt.Printf("LENGTH OF THE LIST: => \"%d\"\n", list.Len())

	for i, element := 0, list.Front(); element != nil; i, element = i+1, element.Next() {
		fmt.Printf("ELEMENT \"%d\" OF THE LIST IS => ", i)
		fmt.Print(element.Value)
		fmt.Println()
	}

	fmt.Println("CREATING ARRAY OF STRINGS WITH LENGTH OF 5")

	var stringsArray [5]string

	fmt.Println("ADDING ELEMENT INTO ARRAY")

	stringsArray[0] = "FIRST VALUE"
	stringsArray[1] = "SECOND VALUE"
	stringsArray[2] = "THIRD VALUE"
	stringsArray[3] = "FOUTH VALUE"
	stringsArray[4] = "FIFTH VALUE"

	fmt.Printf("SIZE OF FILLED ARRAY: => \"%d\"\n", unsafe.Sizeof(stringsArray))

	// for i := 0; i < len(stringsArray); i = i + 1 {
	// 	fmt.Printf("ELEMENT NUMBER \"%d\" OF \"strignsArray\" WITH \"%s\" ELEMENT\n", i, stringsArray[i])
	// }

	for i := range stringsArray {
		fmt.Printf("ELEMENT NUMBER \"%d\" OF \"strignsArray\" WITH \"%s\" ELEMENT\n", i, stringsArray[i])
	}
}
