package note_v1

import (
	"context"
	"fmt"

	desc "github.com/OlegRaykevich/note-service-api/pkg/note_v1"
)

func (n *Implementation) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.Empty, error) {
	fmt.Println("UpdateNote working")
	req.GetId()
	req.GetText()
	req.GetNewTitle()

	return &desc.Empty{}, nil
}
