package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Appointment struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Patient     string             `json:"patient" bson:"patient" validate:"required"`
	Description string             `json:"description" bson:"description" validate:"required"`
	StartDate   time.Time          `json:"start_date" bson:"start_date" validate:"required"`
	UpdateDate  time.Time          `json:"update_date" bson:"update_date" validate:"required"`
}
