package main

import (
	_ "blog/dao"
	_ "blog/model"
	"blog/router"
)

func main() {
	router.Start()
}
