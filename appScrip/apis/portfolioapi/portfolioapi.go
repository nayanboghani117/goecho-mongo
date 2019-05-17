package portfolioapi

import (
	"appScrip/app"
	"appScrip/middleware"
	"appScrip/models"

	"gopkg.in/mgo.v2/bson"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/auth"
)

type (
	// portfolioService specifies the interface for the portfolio service needed by portfolioResource.
	portfolioService interface {
		Get(rs app.RequestScope) ([]models.Portfolio, error)
		GetByCategory(rs app.RequestScope, id string) ([]models.Portfolio, error)
		GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Portfolio, error)
		Create(rs app.RequestScope, model *models.Portfolio) (*models.Portfolio, error)
		Update(rs app.RequestScope, id bson.ObjectId, model *models.Portfolio) (*models.Portfolio, error)
		Delete(rs app.RequestScope, id bson.ObjectId) (*models.Portfolio, error)
	}

	// portfolioResource defines the handlers for the CRUD APIs.
	portfolioResource struct {
		service portfolioService
	}
)

// ServePortfolioResource sets up the routing of portfolio endpoints and the corresponding handlers.
func ServePortfolioResource(rg *routing.RouteGroup, service portfolioService) {
	r := &portfolioResource{service}

	rg.Get("/portfolio", r.getAll)
	rg.Get("/portfolioByCategory/<id>", r.getByCategory)
	rg.Get("/portfolio/<id>", r.get)
}

// ServePortfolioResourceWithAuth sets up the routing of portfolio endpoints and the corresponding handlers.
func ServePortfolioResourceWithAuth(rg *routing.RouteGroup, service portfolioService) {
	rg.Use(auth.JWT(app.Config.JWTVerificationKey, auth.JWTOptions{
		SigningMethod: app.Config.JWTSigningMethod,
		TokenHandler:  middleware.JWTHandler,
	}))

	r := &portfolioResource{service}
	rg.Post("/portfolio", r.post)
	rg.Put("/portfolio/<id>", r.put)
	rg.Delete("/portfolio/<id>", r.delete)
}

func (r *portfolioResource) getAll(c *routing.Context) error {
	items, err := r.service.Get(app.GetRequestScope(c))
	if err != nil {
		return err
	}
	return c.Write(items)
}

func (r *portfolioResource) getByCategory(c *routing.Context) error {
	id := c.Param("id")

	response, err := r.service.GetByCategory(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *portfolioResource) get(c *routing.Context) error {
	id := c.Param("id")

	response, err := r.service.GetOne(app.GetRequestScope(c), bson.ObjectIdHex(id))
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *portfolioResource) post(c *routing.Context) error {
	var model models.Portfolio
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *portfolioResource) put(c *routing.Context) error {
	id := c.Param("id")
	// var model models.Portfolio
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

func (r *portfolioResource) delete(c *routing.Context) error {
	id := c.Param("id")

	response, err := r.service.Delete(app.GetRequestScope(c), bson.ObjectIdHex(id))
	if err != nil {
		return err
	}

	return c.Write(response)
}
