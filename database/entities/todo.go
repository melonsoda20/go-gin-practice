package entities

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
}
