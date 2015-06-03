package db

import (
	"appengine/datastore"
	"appengine"
)

type entity interface {
	Kind() string
	HasKey() bool
	Key() *datastore.Key
	setKey(*datastore.Key)
	NewKey(appengine.Context) *datastore.Key
	UUID() string
	SetUUID(uuid string) error
}

type Entity struct {
	key *datastore.Key `json:"-",datastore:"-"`
}

func (this *Entity) HasKey() bool {
	return this.key != nil
}

func (this *Entity) setKey(k *datastore.Key) {
	this.key = k
}

func (this *Entity) Key() *datastore.Key {
	return this.key
}

func (this *Entity) UUID() string {
	return this.key.Encode()
}

func (this *Entity) SetUUID(uuid string) error {
	key, err := datastore.DecodeKey(uuid)
	this.setKey(key)
	return err
}