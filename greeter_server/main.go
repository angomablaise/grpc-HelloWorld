package main
import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/angomablaise/grpc-HelloWorld/helloworld"
)

const (
	port = "50051"
)

?/server iis used to impement hellowor