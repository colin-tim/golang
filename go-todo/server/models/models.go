package models

import (
	"time"
)

type ToDoList struct {
	ID                  int64     `json:"_id,omitempty" bson:"_id,omitempty"`
	Task                string    `json:"task,omitempty"`
	Status              bool      `json:"status"`
	CreatedDateTime     time.Time `json:"createddatetime,omitempty"`
	LastUpdatedDateTime time.Time `json:"lastupdateddatetime,omitempty"`
}
