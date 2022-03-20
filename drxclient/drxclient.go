// Тестовый клиент-модуль gRPC передачи данных в Directum RX
package drxclient

import (
	"context"
	"flag"
	"log"
	"time"

	pb "drxproto"

	"google.golang.org/grpc"
)

const (
	defaultName = "SAP"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func PostGrpc(Id, Status string, cc chan string) {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Запрос передачи значения Id
	pri, err := c.RouteId(ctx, &pb.RequestId{Id: Id})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	// Запрос передачи значения Status
	prs, err := c.RouteStatus(ctx, &pb.RequestStatus{Status: Status})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	//log.Printf("Greeting: %s", r.GetMessage())
	cc <- string(pri.GetMessage()) // передача результата в main-программу
	cc <- string(prs.GetMessage())
}
