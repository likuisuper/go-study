package main

//1、监听端口
//2、实例化grpc服务端
//3、将微服务注册到gprc
//4、启动服务
import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc/protos"
	"net"
)

// 定义空接口
type UserInfoService struct {
	//这里必须添加下面这行代码，不然在pb.RegisterUserInfoServiceServer(server, &u)这里会提示报错
	//参见 https://github.com/grpc/grpc-go/issues/3794
	//因为UserInfoServiceServer是一个接口，它有两个方法需要实现，而我们在下面已经实现了GetUserInfo，
	//所以还需要实现mustEmbedUnimplementedUserInfoServiceServer方法，但是这个方法如果直接实现也是会报错，因为它是
	//non-exported method，查看user_grpc.pb.go可以发现，UnimplementedUserInfoServiceServer已经实现了它，
	//所以我们在这里继承UnimplementedUserInfoServiceServer即可
	pb.UnimplementedUserInfoServiceServer
}

var u = UserInfoService{}

// 实现方法
func (s *UserInfoService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	name := req.Name
	if name == "lk" {
		resp = &pb.UserResponse{
			Id:    1,
			Name:  name,
			Age:   26,
			Hobby: []string{"sing", "run"},
		}
	}
	return
}

func main() {
	addr := "127.0.0.1:8080"
	//1、监听
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("监听异常:%s\n", err)
	}
	fmt.Printf("监听端口:%s\n", addr)
	//2、实例化gprc
	server := grpc.NewServer()
	//3、在gprc上注册微服务
	pb.RegisterUserInfoServiceServer(server, &u)
	if err := server.Serve(listen); err != nil {
		fmt.Printf("服务启动失败:%v", err)
	}
}
