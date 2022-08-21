package main

import (
	"fmt"
	"log"
	"net"

	"github.com/OlegRaykevich/testGRPC/internal/app/api/note_v1"
	desc "github.com/OlegRaykevich/testGRPC/pkg/note_v1"
	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to mapping port: %s", err.Error())
	}

	s := grpc.NewServer()
	desc.RegisterNoteV1Server(s, note_v1.NewNote())

	fmt.Println("server is running on port: ", port)
	if err := s.Serve(list); err != nil {
		log.Fatalf("failed serve")
	}
}
