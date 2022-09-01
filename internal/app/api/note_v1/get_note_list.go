package note_v1

import (
	"context"
	"fmt"

	desc "github.com/OlegRaykevich/note-service-api/pkg/note_v1"
)

func (n *Implementation) GetNoteList(ctx context.Context, req *desc.GetNoteListRequest) (*desc.GetNoteListResponse, error) {
	fmt.Println("GetNote working")
	for _, id := range req.GetIds() {
		fmt.Println(id)
	}

	return &desc.GetNoteListResponse{
		Note: []*desc.NoteInfo{
			{
				Id:     1,
				Title:  "title",
				Text:   "text",
				Author: "Author1",
			},
			{
				Id:     2,
				Title:  "second Title",
				Text:   "second text",
				Author: "Author2",
			},
		},
	}, nil
}
