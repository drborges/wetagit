package app

import (
	"github.com/drborges/wetagit/api"
	"net/http"
)

func init() {
	http.Handle("/", api.Router())
}
