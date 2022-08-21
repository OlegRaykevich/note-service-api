package note_v1

import (
	"context"
	"fmt"

	desc "github.com/OlegRaykevich/testGRPC/pkg/note_v1"
)

func (n *Note) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	fmt.Println("GetNote working")
	fmt.Println(req.GetId())

	return &desc.GetNoteResponse{
		Title:  "Old note",
		Text:   "my old note",
		Author: "same author",
	}, nil
}
