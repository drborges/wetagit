package injectables

import (
	"github.com/drborges/wetagit/api/services"
	"github.com/go-martini/martini"
	"net/http"
)

func GaeContextProvider(req *http.Request, c martini.Context) {
	c.Map(services.Gae{req}.NewContext())
}
