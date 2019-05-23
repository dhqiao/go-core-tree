package main

import (
	"fmt"

	"github.com/prashantv/gostub"
	"io/ioutil"
)

var counter = 100

func stubGlobalVariable() {
	stubs := gostub.Stub(&counter, 888)
	defer stubs.Reset()
	fmt.Println("Counter:", counter)
}

var configFile = "config.json"

func GetConfig() ([]byte, error) {
	return ioutil.ReadFile(configFile)
}



func main() {
	stubGlobalVariable()
}

// output:
// Counter: 200

