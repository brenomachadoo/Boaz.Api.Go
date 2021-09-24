package main

import (
	"fmt"

	"bmachado/Boaz.Api.Go/cmd/boazApi/api/server"
)

func main() {
	fmt.Println("Start application GoLang")

	fmt.Println("Start new server Company")
	server := server.NewServerCompany()
	server.RunServerCompany()
}
