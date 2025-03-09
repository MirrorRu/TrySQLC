package models

import (
	"time"
)

type (
	MemberID = int64

	Member struct {
		ID       MemberID  `json:"member_id"`
		Name     string    `json:"member_name"`
		Birthday time.Time `json:"member_birthday"`
	}

	MembersFilter struct {
		NamePart string `json:"member_name"`
	}
)
