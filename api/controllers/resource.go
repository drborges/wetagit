package controllers

import (
	"fmt"
	"github.com/drborges/appx"
	"strings"
)

type Resource struct {
	appx.Resource
}

func (this Resource) Path() string {
	return fmt.Sprintf("/%v/%v", strings.ToLower(this.Key().Kind()), this.ResourceID())
}
