package blogapi

import (
	"appScrip/app"
	"appScrip/middleware"
	"appScrip/models"

	"gopkg.in/mgo.v2/bson"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/auth"
)

type (
	// blogService specifies the interface for the blog service needed by blogResource.
	blogService interface {
		Get(rs app.RequestScope) ([]models.Blog, error)
		GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Blog, error)
		Create(rs app.RequestScope, model *models.Blog) (*models.Blog, error)
		Update(rs app.RequestScope, id bson.ObjectId, model *models.Blog) (*models.Blog, error)
		Delete(rs app.RequestScope, id bson.ObjectId) (*models.Blog, error)
	}

	// blogResource defines the handlers for the CRUD APIs.
	blogResource struct {
		service blogService
	}
)

// ServeBlogResource sets up the routing of blog endpoints and the corresponding handlers.
func ServeBlogResource(rg *routing.RouteGroup, service blogService) {
	r := &blogResource{service}

	rg.Get("/blog", r.getAll)
	rg.Get("/blog/<id>", r.get)
}

// ServeBlogResourceWithAuth sets up the routing of blog endpoints and the corresponding handlers.
func ServeBlogResourceWithAuth(rg *routing.RouteGroup, service blogService) {
	rg.Use(auth.JWT(app.Config.JWTVerificationKey, auth.JWTOptions{
		SigningMethod: app.Config.JWTSigningMethod,
		TokenHandler:  middleware.JWTHandler,
	}))

	r := &blogResource{service}
	rg.Post("/blog", r.post)
	rg.Put("/blog/<id>", r.put)
	rg.Delete("/blog/<id>", r.delete)
}

func (r *blogResource) getAll(c *routing.Context) error {
	items, err := r.service.Get(app.GetRequestScope(c))
	if err != nil {
		return err
	}
	return c.Write(items)
}

func (r *blogResource) get(c *routing.Context) error {
	id := c.Param("id")

	response, err := r.service.GetOne(app.GetRequestScope(c), bson.ObjectIdHex(id))
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *blogResource) post(c *routing.Context) error {
	var model models.Blog
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *blogResource) put(c *routing.Context) error {
	id := c.Param("id")
	// var model models.Blog
	model, err := r.service.GetOne(app.GetRequestScope(c), bson.ObjectIdHex(id))
	if err != nil {
		return err
	}
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Update(app.GetRequestScope(c), bson.ObjectIdHex(id), model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *blogResource) delete(c *routing.Context) error {
	id := c.Param("id")

	response, err := r.service.Delete(app.GetRequestScope(c), bson.ObjectIdHex(id))
	if err != nil {
		return err
	}

	return c.Write(response)
}
