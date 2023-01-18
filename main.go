package main

import "go-pgx-loco/api"

func main() {
	r := api.InitRoutes()
	r.Run(":3333")
}
