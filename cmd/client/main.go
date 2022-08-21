package main

import (
	"context"
	"fmt"
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

	resCreateNote, err := client.CreateNote(context.Background(), &desc.CreateNoteRequest{
		Title:  "Shop List",
		Text:   "milk",
		Author: "Oleg",
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("created note id: ", resCreateNote.GetId())

	resGetNote, err := client.GetNote(context.Background(), &desc.GetNoteRequest{
		Id: 1,
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("get note title: ", resGetNote.GetTitle())
	log.Println("get note text: ", resGetNote.GetText())
	log.Println("get note author: ", resGetNote.GetAuthor())

	resUpdateNote, err := client.UpdateNote(context.Background(), &desc.UpdateNoteRequest{
		Id:       1,
		NewTitle: "newName",
		Text:     "New text",
	})
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("Result update: ", resUpdateNote.Result)

	resDeleteNote, err := client.DeleteNote(context.Background(), &desc.DeleteNoteRequest{
		Id: 1,
	})
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("Result delete: ", resDeleteNote.Result)

}
