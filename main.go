package main

import "mini_project/route"

func main() {
	route := route.StartRoute()
	route.Start(":8000")
}
