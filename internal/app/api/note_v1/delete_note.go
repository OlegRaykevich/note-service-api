package note_v1

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	desc "github.com/OlegRaykevich/note-service-api/pkg/note_v1"
)

func (n *Implementation) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.Empty, error) {
	fmt.Println("DeleteNote working")

	dbDsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		host, port, dbName, dbUser, dbPassword, sslMode)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	//builder := sq.Insert(noteTable).
	//	PlaceholderFormat(sq.Dollar).
	//	Columns("title, text, author").
	//	Values(req.GetTitle(), req.GetText(), req.GetAuthor()).
	//	Suffix("returning id")

	builder := sq.Delete(noteTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return &desc.Empty{}, nil
}
