package main

import (
	"fmt"
	"github.com/prashantv/gostub"
)

var CleanUp = cleanUp

func cleanUp(val string) {
	fmt.Println(val)
}

func main() {
	stubs := gostub.StubFunc(&CleanUp)
	CleanUp("Hello go")
	defer stubs.Reset()
}

