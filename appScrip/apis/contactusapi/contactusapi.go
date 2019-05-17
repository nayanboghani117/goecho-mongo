package contactusapi

import (
	"appScrip/app"
	"appScrip/middleware"
	"appScrip/models"

	"gopkg.in/mgo.v2/bson"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/auth"
)

type (
	// contactusService specifies the interface for the contactus service needed by contactusResource.
	contactusService interface {
		Get(rs app.RequestScope) ([]models.ContactUs, error)
		GetOne(rs app.RequestScope, id bson.ObjectId) (*models.ContactUs, error)
		Create(rs app.RequestScope, model *models.ContactUs) (*models.ContactUs, error)
		Update(rs app.RequestScope, id bson.ObjectId, model *models.ContactUs) (*models.ContactUs, error)
		Delete(rs app.RequestScope, id bson.ObjectId) (*models.ContactUs, error)
	}

	// contactusResource defines the handlers for the CRUD APIs.
	contactusResource struct {
		service contactusService
	}
)

// ServeContactUsResource sets up the routing of contactus endpoints and the corresponding handlers.
func ServeContactUsResource(rg *routing.RouteGroup, service contactusService) {
	r := &contactusResource{service}

	rg.Get("/contactus", r.getAll)
	rg.Post("/contactus", r.post)
	rg.Get("/contactus/<id>", r.get)
}

// ServeContactUsResourceWithAuth sets up the routing of contactus endpoints and the corresponding handlers.
func ServeContactUsResourceWithAuth(rg *routing.RouteGroup, service contactusService) {
	rg.Use(auth.JWT(app.Config.JWTVerificationKey, auth.JWTOptions{
		SigningMethod: app.Config.JWTSigningMethod,
		TokenHandler:  middleware.JWTHandler,
	}))

	r := &contactusResource{service}
	rg.Put("/contactus/<id>", r.put)
	rg.Delete("/contactus/<id>", r.delete)
}

func (r *contactusResource) getAll(c *routing.Context) error {
	items, err := r.service.Get(app.GetRequestScope(c))
	if err != nil {
		return err
	}
	return c.Write(items)
}

func (r *contactusResource) get(c *routing.Context) error {
	id := c.Param("id")

	response, err := r.service.GetOne(app.GetRequestScope(c), bson.ObjectIdHex(id))
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *contactusResource) post(c *routing.Context) error {
	var model models.ContactUs
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *contactusResource) put(c *routing.Context) error {
	id := c.Param("id")
	// var model models.ContactUs
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

func (r *contactusResource) delete(c *routing.Context) error {
	id := c.Param("id")

	response, err := r.service.Delete(app.GetRequestScope(c), bson.ObjectIdHex(id))
	if err != nil {
		return err
	}

	return c.Write(response)
}
