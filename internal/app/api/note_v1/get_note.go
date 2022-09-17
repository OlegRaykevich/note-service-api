package note_v1

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	desc "github.com/OlegRaykevich/note-service-api/pkg/note_v1"
)

func (n *Implementation) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {

	dbDsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		host, port, dbName, dbUser, dbPassword, sslMode)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.
		Select("*").
		From("note").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	row.Next()
	var id int64
	var title, text, author string
	err = row.Scan(&id, &title, &text, &author)
	if err != nil {
		return nil, err
	}

	return &desc.GetNoteResponse{
		Note: &desc.NoteInfo{
			Id:     id,
			Title:  title,
			Text:   text,
			Author: author,
		},
	}, nil
}
