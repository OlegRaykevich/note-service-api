package note_v1

import (
	"context"
	"fmt"

	desc "github.com/OlegRaykevich/testGRPC/pkg/note_v1"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {
	fmt.Println("UpdateNote working")
	req.GetId()
	req.GetText()
	req.GetNewTitle()

	fmt.Println("Update success")
	return &desc.UpdateNoteResponse{
		Result: true,
	}, nil
}
