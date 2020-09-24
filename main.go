package main

import (
	"fmt"
	"ginrbac/config"
)

func main() {
	fmt.Println(config.Conf.MysqlUser)
}
