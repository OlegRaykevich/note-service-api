package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/OlegRaykevich/note-service-api/internal/app/api/note_v1"
	desc "github.com/OlegRaykevich/note-service-api/pkg/note_v1"
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
