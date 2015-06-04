package models

import (
	"github.com/drborges/datastore-model"
)

type Tag struct {
	db.Model
	Value string `json:"value",db:"id"`
}
