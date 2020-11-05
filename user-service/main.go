package main

import (
	"fmt"
	database "github.com/felixfong/laracom/user-service/db"
	"github.com/felixfong/laracom/user-service/handler"
	pb "github.com/felixfong/laracom/user-service/proto/user"
	repo2 "github.com/felixfong/laracom/user-service/repo"
	"github.com/micro/go-micro"
	"log"
)

func main() {
	db, err := database.CreateConnection()
	defer db.Close()
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	//数据库迁移，每次启动服务时都会检查，如果数据表不存在则创建，已存在则检查是否有修改
	db.AutoMigrate(&pb.User{})

	repo := &repo2.UserRepository{db}

	srv := micro.NewService(micro.Name("laracom.user.service"), micro.Version("latest"))
	srv.Init()

	//注册处理器
	pb.RegisterUserServiceHandler(srv.Server(), &handler.UserService{repo})

	//启动用户服务
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
