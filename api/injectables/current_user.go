package injectables

import (
	"net/http"
	"github.com/go-martini/martini"
	"github.com/drborges/wetagit/api/models"
	"github.com/drborges/appx"
	"github.com/martini-contrib/render"
)

func CurrentUserProvider(r render.Render, req *http.Request, c martini.Context, cds *appx.CachedDatastore) {
	// TODO add proper authentication thru AuthTokens
	user := models.User{Name: req.Header.Get("X-Auth-Token")}
	if err := cds.Load(&user); err != nil {
//		r.JSON(http.StatusInternalServerError, fmt.Sprintf("Was not able to find user for auth token %v", user.Name))
//		return
	}
	c.Map(user)
}