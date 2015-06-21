package injectables

import (
	"github.com/drborges/appx"
	"github.com/drborges/wetagit/api/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

func CurrentUserProvider(r render.Render, req *http.Request, c martini.Context, cds *appx.CachedDatastore) {
	// TODO add proper authentication thru AuthTokens
	user := &models.User{Name: req.Header.Get("X-Auth-Token")}
	if err := cds.Load(user); err != nil {
		//		r.JSON(http.StatusInternalServerError, fmt.Sprintf("Was not able to find user for auth token %v", user.Name))
		//		return
	}
	c.Map(user)
}
