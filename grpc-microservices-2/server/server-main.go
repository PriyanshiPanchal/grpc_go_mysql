package main

import (
	"context"
	// "log"
	"net"
	"grpc-microservices-2/proto"
	// "grpc-microservices-2/db"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// type employeeItem struct {
// 	ID      int    `json:"id"`
// 	Name    string `json:"name"`
// 	City    string `json:"city"`
// }

type server struct{}


func main() {

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	proto.RegisterEmployeeServiceServer(srv, &server{})

	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
func (s *server) Employee(ctx context.Context, request *proto.CreateEmployeeReq) (*proto.CreateEmployeeRes, error) {

	// db, err := db.getconn()
	// if err != nil {
	// 	log.Println("Something Went Wrong")
	// }

	var id = request.GetId()
	var emp_name = request.GetName()
	var emp_city = request.GetCity()

	// result, err := db.Query("insert into employee_table values(?,?,?)", id, emp_name, emp_city)
	// if err != nil {
	// 	panic(err.Error())
	// }
	
	return &proto.CreateEmployeeRes{Id: id,
		Name: emp_name,
		City: emp_city}, nil
	// defer db.Close()
	// return response, nil
}

