package members

import (
	"context"
	"database/sql"

	"github.com/mirrorru/trysqlc/internal/models"
	sqlc_members "github.com/mirrorru/trysqlc/internal/repository/members/sqlc_gen"
)

const (
	pkgName = "pg_members"
)

type DbHandler interface {
	Connection() *sql.DB
	Exec(ctx context.Context, query string, args ...any) (result sql.Result, err error)
	QueryRow(ctx context.Context, query string, args ...any) (result *sql.Row, err error)
	Query(ctx context.Context, query string, args ...any) (result *sql.Rows, err error)
	Close() (err error)
}

type membersRepo struct {
	writer  DbHandler
	reader  DbHandler
	queries sqlc_members.Queries
}

func NewMembersRepo(writer DbHandler, reader DbHandler) *membersRepo {
	return &membersRepo{
		writer: writer,
		reader: reader,
	}
}

func (r *membersRepo) Create(ctx context.Context, src *models.Member) (err error) {
	_, err = r.queries.CreateMember(ctx, r.writer.Connection(), toCreateMemberParams(src))

	return err
}

func (r *membersRepo) Update(ctx context.Context, src *models.Member) (err error) {
	_, err = r.queries.UpdateMember(ctx, r.writer.Connection(), toUpdateMemberParams(src))

	return err
}

func (r *membersRepo) Delete(ctx context.Context, src models.MemberID) (err error) {
	_, err = r.queries.DeleteMember(ctx, r.writer.Connection(), src)

	return err
}

func (r *membersRepo) GetByID(ctx context.Context, src models.MemberID) (err error) {
	_, err = r.queries.GetMember(ctx, r.reader.Connection(), src)

	return err
}

func (r *membersRepo) GetList(ctx context.Context, src *models.MembersFilter) (err error) {
	_, err = r.queries.GetMembersList(
		ctx,
		r.reader.Connection(),
		sql.NullString{String: src.NamePart, Valid: src.NamePart != ""},
	)

	return err
}
