package injectables

import (
	"appengine"
	"github.com/drborges/appx"
	"github.com/go-martini/martini"
)

func DatastoreProvider(gaeContext appengine.Context, c martini.Context) {
	c.Map(appx.NewDatastore(gaeContext))
}
