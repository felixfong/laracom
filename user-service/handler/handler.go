package handler

import (
	"context"
	pb "github.com/felixfong/laracom/user-service/proto/user"
	"github.com/felixfong/laracom/user-service/repo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo repo.Repository
}

func (srv *UserService) Get(ctx context.Context, req *pb.User, rsp *pb.Response) error {
	user, err := srv.Repo.Get(req.Id)
	if err != nil {
		return err
	}
	rsp.User = user
	return nil
}

func (srv *UserService) GetAll(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	users, err := srv.Repo.GetAll()
	if err != nil {
		return err
	}
	rsp.Users = users
	return nil
}

func (srv *UserService) Create(ctx context.Context, req *pb.User, rsp *pb.Response) error {
	//对密码进行加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPassword)
	if err := srv.Repo.Create(req); err != nil {
		return err
	}
	rsp.User = req
	return nil
}


