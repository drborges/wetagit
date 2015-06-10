package controllers

import (
	"strings"
	"fmt"
	"github.com/drborges/datastore-model"
)

type Resource struct {
	db.Entity
}

func (this Resource) Path() string {
	return fmt.Sprintf("/%v/%v", strings.ToLower(this.Key().Kind()), this.StringId())
}