package note_v1

import (
	"context"
	"fmt"

	desc "github.com/OlegRaykevich/testGRPC/pkg/note_v1"
)

func (n *Note) GetNoteList(ctx context.Context, req *desc.GetNoteListRequest) (*desc.GetNoteListResponse, error) {
	fmt.Println("GetNote working")
	fmt.Println(req.GetId()[0])

	return &desc.GetNoteListResponse{
		Title:  []string{"titles"},
		Text:   []string{"texts"},
		Author: []string{"Authors"},
	}, nil
}
