package models

import (
	"github.com/drborges/ds"
)

type Subscription struct {
	ds.Model
	Tag string `json:"tag"`
}

func (this Subscription) KeyMetadata() *ds.KeyMetadata {
	return &ds.KeyMetadata{
		Kind:      "Subscriptions",
		StringID:  this.Tag,
		HasParent: true, // Belongs to a User
	}
}

type Subscriptions []*Subscription

func (this Subscriptions) BelongTo(user User) {
	for _, subscription := range this {
		subscription.SetParentKey(user.Key())
	}
}