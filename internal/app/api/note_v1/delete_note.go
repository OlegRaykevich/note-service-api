package note_v1

import (
	"context"
	"fmt"

	desc "github.com/OlegRaykevich/testGRPC/pkg/note_v1"
)

func (n *Note) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.DeleteNoteResponse, error) {
	fmt.Println("DeleteNote working")
	req.GetId()

	fmt.Println("Delete success")
	return &desc.DeleteNoteResponse{
		Result: true,
	}, nil
}
