package awardsapi

import (
	"appScrip/app"
	"appScrip/middleware"
	"appScrip/models"

	"gopkg.in/mgo.v2/bson"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/auth"
)

type (
	// awardsService specifies the interface for the awards service needed by awardsResource.
	awardsService interface {
		Get(rs app.RequestScope) ([]models.Awards, error)
		GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Awards, error)
		Create(rs app.RequestScope, model *models.Awards) (*models.Awards, error)
		Update(rs app.RequestScope, id bson.ObjectId, model *models.Awards) (*models.Awards, error)
		Delete(rs app.RequestScope, id bson.ObjectId) (*models.Awards, error)
	}

	// awardsResource defines the handlers for the CRUD APIs.
	awardsResource struct {
		service awardsService
	}
)

// ServeAwardsResource sets up the routing of awards endpoints and the corresponding handlers.
func ServeAwardsResource(rg *routing.RouteGroup, service awardsService) {
	r := &awardsResource{service}

	rg.Get("/awards", r.getAll)
	rg.Get("/awards/<id>", r.get)
}

// ServeAwardsResourceWithAuth sets up the routing of awards endpoints and the corresponding handlers.
func ServeAwardsResourceWithAuth(rg *routing.RouteGroup, service awardsService) {
	rg.Use(auth.JWT(app.Config.JWTVerificationKey, auth.JWTOptions{
		SigningMethod: app.Config.JWTSigningMethod,
		TokenHandler:  middleware.JWTHandler,
	}))

	r := &awardsResource{service}
	rg.Post("/awards", r.post)
	rg.Put("/awards/<id>", r.put)
	rg.Delete("/awards/<id>", r.delete)
}

func (r *awardsResource) getAll(c *routing.Context) error {
	items, err := r.service.Get(app.GetRequestScope(c))
	if err != nil {
		return err
	}
	return c.Write(items)
}

func (r *awardsResource) get(c *routing.Context) error {
	id := c.Param("id")

	response, err := r.service.GetOne(app.GetRequestScope(c), bson.ObjectIdHex(id))
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *awardsResource) post(c *routing.Context) error {
	var model models.Awards
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *awardsResource) put(c *routing.Context) error {
	id := c.Param("id")
	// var model models.Awards
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

func (r *awardsResource) delete(c *routing.Context) error {
	id := c.Param("id")

	response, err := r.service.Delete(app.GetRequestScope(c), bson.ObjectIdHex(id))
	if err != nil {
		return err
	}

	return c.Write(response)
}
