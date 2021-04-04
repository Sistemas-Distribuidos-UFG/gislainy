package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/gislainy/course-distributed-systems/grpc-questions/questions"
	"google.golang.org/grpc"
)

type QuestionsServer struct {
	pb.UnimplementedQuestionsServer
}

const (
	port = ":55551"
)

func (s *QuestionsServer) CalculateReajustSalary(ctx context.Context, employee *pb.Employee) (*pb.Employee, error) {
	return CalculateReajustSalary(employee), nil
}

// CalculateReajustSalary reajust the salary according to the employee's position
func CalculateReajustSalary(employee *pb.Employee) *pb.Employee {
	switch employee.Role {
	case "Operator":
		employee.Salary *= 1.2
	case "Developer":
		employee.Salary *= 1.18
	}
	return &pb.Employee{
		Name:   employee.Name,
		Role:   employee.Role,
		Salary: employee.Salary}
}

func main() {

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := QuestionsServer{}

	grpcServer := grpc.NewServer()

	pb.RegisterQuestionsServer(grpcServer, &s)

	fmt.Printf("Listen to port %s\n", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server over port %s: %v", port, err)
	}

}
