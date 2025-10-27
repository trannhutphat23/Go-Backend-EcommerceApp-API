package main

import "github.com/trannhutphat23/Go-Backend-EcommerceApp-API/internal/routers"

func main() {
	r := routers.NewRouter()

	r.Run(":8008")
}
