package videosapi

import (
	"appScrip/app"
	"appScrip/middleware"
	"appScrip/models"

	"gopkg.in/mgo.v2/bson"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/auth"
)

type (
	// videosService specifies the interface for the videos service needed by videosResource.
	videosService interface {
		Get(rs app.RequestScope) ([]models.Videos, error)
		GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Videos, error)
		Create(rs app.RequestScope, model *models.Videos) (*models.Videos, error)
		Update(rs app.RequestScope, id bson.ObjectId, model *models.Videos) (*models.Videos, error)
		Delete(rs app.RequestScope, id bson.ObjectId) (*models.Videos, error)
	}

	// videosResource defines the handlers for the CRUD APIs.
	videosResource struct {
		service videosService
	}
)

// ServeVideosResource sets up the routing of videos endpoints and the corresponding handlers.
func ServeVideosResource(rg *routing.RouteGroup, service videosService) {
	r := &videosResource{service}

	rg.Get("/videos", r.getAll)
	rg.Get("/videos/<id>", r.get)
}

// ServeVideosResourceWithAuth sets up the routing of videos endpoints and the corresponding handlers.
func ServeVideosResourceWithAuth(rg *routing.RouteGroup, service videosService) {
	rg.Use(auth.JWT(app.Config.JWTVerificationKey, auth.JWTOptions{
		SigningMethod: app.Config.JWTSigningMethod,
		TokenHandler:  middleware.JWTHandler,
	}))

	r := &videosResource{service}
	rg.Post("/videos", r.post)
	rg.Put("/videos/<id>", r.put)
	rg.Delete("/videos/<id>", r.delete)
}

func (r *videosResource) getAll(c *routing.Context) error {
	items, err := r.service.Get(app.GetRequestScope(c))
	if err != nil {
		return err
	}
	return c.Write(items)
}

func (r *videosResource) get(c *routing.Context) error {
	id := c.Param("id")

	response, err := r.service.GetOne(app.GetRequestScope(c), bson.ObjectIdHex(id))
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *videosResource) post(c *routing.Context) error {
	var model models.Videos
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *videosResource) put(c *routing.Context) error {
	id := c.Param("id")
	// var model models.Videos
	model, err := r.service.GetOne(app.GetRequestScope(c), bson.ObjectIdHex(id))
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Update(app.GetRequestScope(c), bson.ObjectIdHex(id), model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *videosResource) delete(c *routing.Context) error {
	id := c.Param("id")

	response, err := r.service.Delete(app.GetRequestScope(c), bson.ObjectIdHex(id))
	if err != nil {
		return err
	}

	return c.Write(response)
}
