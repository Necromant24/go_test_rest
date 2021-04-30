package main

import (
	"test/config"
	"test/startup"

	"fmt"
)

func main() {

	fmt.Println(config.GetDbConnString())

	startup.LoadApp()

}
