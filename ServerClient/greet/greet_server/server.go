package main
import (
    "fmt"
    "log"
    "net"
	
    
    "google.golang.org/grpc"
)




func main() {
    fmt.Println("Hello world")
    
    lis, err := net.Listen("tcp", "0.0.0.0:50051")
    if err != nil {
        log.Fatalf("Failed")
	}
	
	s := grpc.NewServer()

	
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}