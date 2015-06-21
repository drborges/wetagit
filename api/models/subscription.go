package models

import (
	"github.com/drborges/appx"
)

type Subscription struct {
	appx.Model
	Feed       string `json:"feed"`
}

func (this *Subscription) KeyMetadata() *appx.KeyMetadata {
	return &appx.KeyMetadata{
		Kind:      "Subscriptions",
		StringID:  this.Feed,
		HasParent: true, // Belongs to a user
	}
}

func (this *Subscription) CacheID() string {
	return this.ResourceID()
}

func (this *Subscription) BelongsTo(user *User) {
	this.SetParentKey(user.Key())
}

type Subscriptions []*Subscription
