package injectables

import (
	"net/http"
	"github.com/drborges/ds"
	"github.com/drborges/wetagit/api/services"
	"github.com/go-martini/martini"
)

func DatastoreProvider(req *http.Request, c martini.Context) {
	c.Map(ds.Datastore{services.Gae{req}.NewContext()})
}
