package injectables

import (
	"github.com/drborges/appx"
	"github.com/go-martini/martini"
	"appengine"
)

func CachedDatastoreProvider(gaeContext appengine.Context, c martini.Context) {
	c.Map(appx.NewCachedDatastore(gaeContext))
}
