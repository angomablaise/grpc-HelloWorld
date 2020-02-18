package main
import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "github.com/angomablaise/grpc-HelloWorld/helloworld"
)

const (
	address ="localhost:50051"
	defaultName = "world"
)

func main() {
	conn,err := grpc.Dial(address,grpc.WithInsecure(),grpc.WithBlock())
	if err !=nil {
		log.Fatalf("did not connect: %v",err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	//contact the server and print out its response.
	name := defaultName
	if  len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx,&pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}