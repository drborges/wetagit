package controllers

import (
	"fmt"
	"github.com/drborges/ds"
	"strings"
)

type Resource struct {
	ds.Resource
}

func (this Resource) Path() string {
	return fmt.Sprintf("/%v/%v", strings.ToLower(this.Key().Kind()), this.UUID())
}
