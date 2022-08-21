package note_v1

import (
	desc "github.com/OlegRaykevich/testGRPC/pkg/note_v1"
)

type Note struct {
	desc.UnimplementedNoteV1Server
}

func NewNote() *Note {
	return &Note{}
}
