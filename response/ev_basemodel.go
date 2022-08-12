package response

import (
	"github.com/google/uuid"
	"time"
)

type EvBaseModel struct {
	ID        int64      `json:"id"`
	UUID      uuid.UUID  `json:"uuid"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
