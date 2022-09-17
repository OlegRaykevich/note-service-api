package note_v1

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

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

	builder := sq.
		Select("*").
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
	for row.Next() {
		var note desc.NoteInfo

		err = row.Scan(&note.Id, &note.Title, &note.Text, &note.Author)
		if err != nil {
			return nil, err
		}

		noteInfoList = append(noteInfoList, &note)
	}

	return &desc.GetNoteListResponse{
		Note: noteInfoList,
	}, nil
}
