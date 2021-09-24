package main

import (
	"fmt"

	"bmachado/Boaz.Api.Go/cmd/boazApi/api/server"

	dbMysqlDrive "bmachado/Boaz.Api.Go/infra/database/Mysql"
)

func main() {
	fmt.Println("Start application GoLang")

	fmt.Println("Start database architecture")
	dbMysqlDrive.MysqlTest()

	fmt.Println("Start new server Company")
	server := server.NewServerCompany()
	server.RunServerCompany()
}
