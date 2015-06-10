package controllers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/drborges/wetagit/api/controllers"
	"github.com/drborges/wetagit/api/models"
	"github.com/drborges/datastore-model"
	"appengine/datastore"
	"net/http"
	"github.com/martini-contrib/render"
)

type FakeDatastore struct {
	db.Datastore
}

func (this FakeDatastore) Create(e db.Entity) error {
	e.SetKey(&datastore.Key{})
	return nil
}

type FakeRenderer struct {
	status int
	body   interface{}
	render.Render
}

func (this *FakeRenderer) JSON(status int, v interface{}) {
	this.status = status
	this.body = v
}

var _ = Describe("Tags Controller", func() {
	Context("Tag creation", func() {
		renderer := &FakeRenderer{}
		ctrl := controllers.Tags
		ctrl.Headers = make(http.Header)
		ctrl.Renderer = renderer
		ctrl.Datastore = FakeDatastore{}

		It("should return 200 with the created tag in the response body", func() {
			tag := models.Tag{Name: "golang"}
			ctrl.Create(tag)
			Expect(renderer.status).To(Equal(201))
			Expect(renderer.body.(models.Tag).Name).To(Equal("golang"))
		})
	})
})
