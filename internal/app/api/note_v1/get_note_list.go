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

func (n *Implementation) GetNoteList(ctx context.Context, req *desc.GetNoteListRequest) (*desc.GetNoteListResponse, error) {
	dbDsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		host, port, dbName, dbUser, dbPassword, sslMode)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	ids := sq.Or{}
	for _, id := range req.GetIds() {
		ids = append(ids, sq.Eq{"id": id})
	}

	builder := sq.Select("note.id, note.title, note.text, note.author, note.created_at, note.updated_at").
		From("note").
		PlaceholderFormat(sq.Dollar).
		Where(ids)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var noteInfoList []*desc.NoteInfo
	var createdAt time.Time
	var updatedAt sql.NullTime

	for row.Next() {
		var note desc.NoteInfo
		var updatedAtPb *timestamppb.Timestamp
		err = row.Scan(&note.Id, &note.Title, &note.Text, &note.Author, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		note.CreatedAt = timestamppb.New(createdAt)
		if updatedAt.Valid {
			updatedAtPb = timestamppb.New(updatedAt.Time)
		}
		note.UpdatedAt = updatedAtPb

		noteInfoList = append(noteInfoList, &note)
	}

	return &desc.GetNoteListResponse{
		Note: noteInfoList,
	}, nil
}
