package parts

import (
	"fmt"
	"time"
)

const TIME_SLEEP time.Duration = 1 // SECONDS

func timer(func_ func()) {
	fmt.Printf("EXECUTING DECORATING FUNC WITH \"%d\" SECONDS SLEEP\n", TIME_SLEEP)
	st := time.Now()
	func_()
	fmt.Printf("DECORATED FUNC EXECUTED WITHIN \"%f\" SECONDS\n", float32(time.Since(st))/float32(time.Second))
}

func DecoratorsPartExecute() {
	printInfo("\"DECORATORS\"")

	timer(func() { time.Sleep(TIME_SLEEP * time.Second) })

	printInfo("\"DECORATORS\"")
}
