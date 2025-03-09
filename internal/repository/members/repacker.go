package members

import (
	"database/sql"

	"github.com/mirrorru/trysqlc/internal/models"
	sqlc_members "github.com/mirrorru/trysqlc/internal/repository/members/sqlc_gen"
)

func toCreateMemberParams(src *models.Member) *sqlc_members.CreateMemberParams {
	return &sqlc_members.CreateMemberParams{
		Name:     src.Name,
		Birthday: sql.NullTime{Time: src.Birthday, Valid: true},
	}
}

func toUpdateMemberParams(src *models.Member) *sqlc_members.UpdateMemberParams {
	return &sqlc_members.UpdateMemberParams{
		ID:       src.ID,
		Name:     src.Name,
		Birthday: sql.NullTime{Time: src.Birthday, Valid: true},
	}
}

func toMember(src *sqlc_members.MembersMembersTable) (result *models.Member) {
	result = &models.Member{
		ID:       src.ID,
		Name:     src.Name,
		Birthday: src.Birthday.Time,
	}

	return result
}

func toMembersList(src []*sqlc_members.MembersMembersTable) (result []*models.Member) {
	result = make([]*models.Member, len(src))
	for i := range src {
		result[i] = toMember(src[i])
	}
	return result
}
