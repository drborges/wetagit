package services

import (
	"appengine"
	"fmt"
	"net/http"
)

type Gae struct {
	Request *http.Request
}

func (this Gae) NewContext() appengine.Context {
	gae := appengine.NewContext(this.Request)
	namespace := appengine.ModuleName(gae)
	context, err := appengine.Namespace(gae, namespace)
	if err != nil {
		panic(fmt.Sprintf("Could not create GAE context: %v", err))
	}

	return context
}
