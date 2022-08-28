package note_v1

import (
	"context"
	"fmt"

	desc "github.com/OlegRaykevich/note-service-api/pkg/note_v1"
)

func (n *Implementation) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.Empty, error) {
	fmt.Println("DeleteNote working")
	req.GetId()

	return &desc.Empty{}, nil
}
