package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	desc "github.com/OlegRaykevich/note-service-api/pkg/note_v1"
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
		Id: 5,
	})
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(resGetNote)

	var ids = []int64{3, 5}
	resGetNoteList, err := client.GetNoteList(ctx, &desc.GetNoteListRequest{
		Ids: ids,
	})
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(resGetNoteList)

	_, err = client.UpdateNote(ctx, &desc.UpdateNoteRequest{
		Note: &desc.NoteInfo{
			Id:     5,
			Title:  "new name",
			Text:   "new some text",
			Author: "user322",
		},
	})
	if err != nil {
		log.Println(err.Error())
	}

	_, err = client.DeleteNote(ctx, &desc.DeleteNoteRequest{
		Id: 1,
	})
	if err != nil {
		log.Println(err.Error())
	}
}
