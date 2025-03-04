package admin

import (
	"time"
)

type (
	Admin struct {
		CreatedAt         time.Time
		Email             string
		EncryptedPassword string
		Id                uint32
		Token             string
		UpdatedAt         time.Time
	}
)
