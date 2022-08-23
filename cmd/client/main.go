package main

import (
	"context"
	"fmt"
	"log"

	desc "github.com/OlegRaykevich/testGRPC/pkg/note_v1"
	"google.golang.org/grpc"
)

const address = "localhost:50051"

func main() {
	ctx := context.Background()

	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to grpc connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewNoteV1Client(con)

	resCreateNote, err := client.CreateNote(ctx, &desc.CreateNoteRequest{
		Title:  "Shop List",
		Text:   "milk",
		Author: "Oleg",
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("created note id: ", resCreateNote.GetId())

	resGetNote, err := client.GetNote(ctx, &desc.GetNoteRequest{
		Id: 1,
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("get note title: ", resGetNote.GetTitle())
	log.Println("get note text: ", resGetNote.GetText())
	log.Println("get note author: ", resGetNote.GetAuthor())

	resUpdateNote, err := client.UpdateNote(ctx, &desc.UpdateNoteRequest{
		Id:       1,
		NewTitle: "newName",
		Text:     "New text",
	})
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("Result update: ", resUpdateNote.Result)

	resDeleteNote, err := client.DeleteNote(ctx, &desc.DeleteNoteRequest{
		Id: 1,
	})
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("Result delete: ", resDeleteNote.Result)
}
