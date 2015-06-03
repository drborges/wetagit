package services

import (
	"appengine"
	"appengine/datastore"
)

type KeyProvider interface {
	Key(appengine.Context) *datastore.Key
}

type Datastore struct {
	context appengine.Context
}

func (this Datastore) Save(e KeyProvider) bool {
	_, err := datastore.Put(this.context, e.Key(this.context), e); ok := err == nil
	if !ok {
		this.context.Errorf("%v", err)
	}
	return ok
}

func (this Datastore) Load(e KeyProvider) bool {
	return this.Get(e.Key(this.context), e)
}

func (this Datastore) Get(key *datastore.Key, dst interface{}) bool {
	err := datastore.Get(this.context, key, dst); ok := err == nil
	if !ok {
		this.context.Errorf("%v", err)
	}
	return ok
}