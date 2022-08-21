package main

import (
	"context"
	desc "github.com/OlegRaykevich/testGRPC/pkg/note_v1"
	"google.golang.org/grpc"
	"log"
)

const address = "localhost:50051"

func main() {
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewNoteV1Client(con)

	res, err := client.CreateNote(context.Background(), &desc.CreateNoteRequest{
		Title:  "Shop List",
		Text:   "milk",
		Author: "Oleg",
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Id: ", res.Id)
}
