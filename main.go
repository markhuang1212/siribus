package main

import (
	"github.com/markhuang1212/siribus/internal/router"
)

func main() {
	r := router.NewRouter()
	r.Run()
}
