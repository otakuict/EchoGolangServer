package main

import (
	"myapp/src/router"
)

func main() {
	e := router.New()
	e.Start(":1323")
}
