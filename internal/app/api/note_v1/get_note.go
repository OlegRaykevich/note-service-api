package note_v1

import (
	"context"
	"fmt"

	desc "github.com/OlegRaykevich/note-service-api/pkg/note_v1"
)

func (n *Implementation) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	fmt.Println("GetNote working")

	return &desc.GetNoteResponse{
		Note: &desc.NoteInfo{
			Id:     req.GetId(),
			Title:  "Old note",
			Text:   "my old note",
			Author: "same author",
		},
	}, nil
}
