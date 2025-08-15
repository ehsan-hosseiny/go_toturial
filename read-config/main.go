package main

import (
	"fmt"

	"github.com/ehsanhossini/go/read-config/config"
)

func main() {

	config.LoadEnv()

	fmt.Println(config.GlobalConfig)

}
