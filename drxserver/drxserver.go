// Тестовый сервер-модуль gRPC передачи данных в Directum RX
package main

import (
	"context"
	pb "drxproto"
	"flag"
	"fmt"
	"log"

	"net"
	"postcreate"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement drx.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// RouteId функция передачи данных Id
func (s *server) RouteId(ctx context.Context, in *pb.RequestId) (*pb.ReplyId, error) {
	cr := make(chan string) // канал функций post-запроса
	go postcreate.PostCreate(cr)
	log.Println("\n\nРезультат post-запроса: \n", <-cr) // Получение данных запроса из канала горутины
	log.Printf("Результат: Тип созданного документа: %v", in.Id)
	return &pb.ReplyId{Message: "Ответ сервера: Тип созданного документа: " + in.Id}, nil
}

// RouteStatus функция передачи данных Status
func (s *server) RouteStatus(ctx context.Context, in *pb.RequestStatus) (*pb.ReplyStatus, error) {
	log.Printf("Результат: Название созданного документа: %v", in.Status)
	return &pb.ReplyStatus{Message: "Ответ сервера: Название созданного документа: " + in.Status}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
