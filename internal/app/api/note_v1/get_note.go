package note_v1

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"

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

	builder := sq.Select("note.id, note.title, note.text, note.author, note.created_at, note.updated_at").
		From("note").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()}).
		Limit(1)

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
	var createdAt time.Time
	var updatedAt sql.NullTime
	err = row.Scan(&id, &title, &text, &author, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	var updatedAtPb *timestamppb.Timestamp
	if updatedAt.Valid {
		updatedAtPb = timestamppb.New(updatedAt.Time)
	}

	return &desc.GetNoteResponse{
		Note: &desc.NoteInfo{
			Id:        id,
			Title:     title,
			Text:      text,
			Author:    author,
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: updatedAtPb,
		},
	}, nil
}
