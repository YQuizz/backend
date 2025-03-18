package main

import (
	"yquiz_back/internal"
	"yquiz_back/internal/pkg"
)

func main() {
	pkg.LoadEnv()
	internal.Init_server()

}
