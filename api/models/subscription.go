package models

import (
	"github.com/drborges/appx"
)

type Subscription struct {
	appx.Model
	Tag string `json:"tag"`
}

func (this Subscription) KeyMetadata() *appx.KeyMetadata {
	return &appx.KeyMetadata{
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