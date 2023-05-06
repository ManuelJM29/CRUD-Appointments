package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Appointment struct {
	ID          primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	Title       string             `json:"tltle" bson:"title, omitempty"`
	Description string             `json:"description" bson:"description, omitempty"`
	StartDate   time.Time          `json:"start_date" bson:"start_date, omitempty"`
	EndDate     time.Time          `json:"end_date" bson:"end_date, omitempty"`
}
