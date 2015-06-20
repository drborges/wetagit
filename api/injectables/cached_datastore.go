package injectables

import (
	"net/http"
	"github.com/drborges/appx"
	"github.com/drborges/wetagit/api/services"
	"github.com/go-martini/martini"
)

func CachedDatastoreProvider(req *http.Request, c martini.Context) {
	c.Map(appx.NewCachedDatastore(services.Gae{req}.NewContext()))
}
