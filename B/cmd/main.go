package main

import (
	"B/api/handler_grpc"
	"B/core/ports"
	"B/core/usecase"
	postGresql2 "B/infr/adapters/database/postGresql"
	"B/pb"
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		return
		fmt.Println(err)
	}
	defer lis.Close()
	connect := Connect()
	postGresql := postGresql2.NewPostGreSql(connect)
	portUser := ports.RepositoryUser(postGresql)
	userCaesUser := usecase.NewUserCase(portUser)
	hanlderGrpc := handler_grpc.NewHandlerGrpc(userCaesUser)
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, hanlderGrpc)
	err = s.Serve(lis)
	if err != nil {
		fmt.Println("err ", err)
	}
}
func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=1234 dbname=Users port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connect fail")
	} else {
		fmt.Print("Connect successfully")
	}
	return db
}
