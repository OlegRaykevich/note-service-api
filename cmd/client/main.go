package main

import (
	"context"
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

	_, err = client.GetNote(ctx, &desc.GetNoteRequest{
		Id: 1,
	})
	if err != nil {
		log.Println(err.Error())
	}

	_, err = client.UpdateNote(ctx, &desc.UpdateNoteRequest{
		Note: &desc.NoteInfo{
			Id:     1,
			Title:  "title name",
			Text:   "some text",
			Author: "user",
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
