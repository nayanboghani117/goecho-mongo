package categoryapi

import (
	"appScrip/app"
	"appScrip/middleware"
	"appScrip/models"

	"gopkg.in/mgo.v2/bson"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/auth"
)

type (
	// categoryService specifies the interface for the category service needed by categoryResource.
	categoryService interface {
		Get(rs app.RequestScope) ([]models.Category, error)
		GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Category, error)
		Create(rs app.RequestScope, model *models.Category) (*models.Category, error)
		Update(rs app.RequestScope, id bson.ObjectId, model *models.Category) (*models.Category, error)
		Delete(rs app.RequestScope, id bson.ObjectId) (*models.Category, error)
	}

	// awardService specifies the interface for the award service needed by categoryResource.
	awardService interface {
		Get(rs app.RequestScope) ([]models.Awards, error)
		GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Awards, error)
		Create(rs app.RequestScope, model *models.Awards) (*models.Awards, error)
		Update(rs app.RequestScope, id bson.ObjectId, model *models.Awards) (*models.Awards, error)
		Delete(rs app.RequestScope, id bson.ObjectId) (*models.Awards, error)
	}

	// aboutService specifies the interface for the about service needed by categoryResource.
	aboutService interface {
		Get(rs app.RequestScope) ([]models.About, error)
		GetOne(rs app.RequestScope) (*models.About, error)
		Create(rs app.RequestScope, model *models.About) (*models.About, error)
		Update(rs app.RequestScope, id bson.ObjectId, model *models.About) (*models.About, error)
		Delete(rs app.RequestScope, id bson.ObjectId) (*models.About, error)
	}

	// testimonialService specifies the interface for the about service needed by categoryResource.
	testimonialService interface {
		Get(rs app.RequestScope) ([]models.Testimonial, error)
		GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Testimonial, error)
		Create(rs app.RequestScope, model *models.Testimonial) (*models.Testimonial, error)
		Update(rs app.RequestScope, id bson.ObjectId, model *models.Testimonial) (*models.Testimonial, error)
		Delete(rs app.RequestScope, id bson.ObjectId) (*models.Testimonial, error)
	}

	// serviceService specifies the interface for the about service needed by categoryResource.
	serviceService interface {
		Get(rs app.RequestScope) ([]models.Service, error)
		GetOne(rs app.RequestScope) (*models.Service, error)
		Create(rs app.RequestScope, model *models.Service) (*models.Service, error)
		Update(rs app.RequestScope, id bson.ObjectId, model *models.Service) (*models.Service, error)
		Delete(rs app.RequestScope, id bson.ObjectId) (*models.Service, error)
	}

	// categoryResource defines the handlers for the CRUD APIs.
	categoryResource struct {
		service     categoryService
		award       awardService
		about       aboutService
		testimonial testimonialService
		ourservice  serviceService
	}
)

// ServeCategoryResource sets up the routing of category endpoints and the corresponding handlers.
func ServeCategoryResource(
	rg *routing.RouteGroup,
	service categoryService,
	award awardService,
	about aboutService,
	testimonial testimonialService,
	ourservice serviceService,
) {
	r := &categoryResource{service, award, about, testimonial, ourservice}

	rg.Get("/category", r.getAll)
	rg.Get("/category/<id>", r.get)
	rg.Get("/home", r.home)
}

// ServeCategoryResourceWithAuth sets up the routing of category endpoints and the corresponding handlers.
func ServeCategoryResourceWithAuth(
	rg *routing.RouteGroup,
	service categoryService,
	award awardService,
	about aboutService,
	testimonial testimonialService,
	ourservice serviceService,
) {
	r := &categoryResource{service, award, about, testimonial, ourservice}

	rg.Use(auth.JWT(app.Config.JWTVerificationKey, auth.JWTOptions{
		SigningMethod: app.Config.JWTSigningMethod,
		TokenHandler:  middleware.JWTHandler,
	}))
	rg.Post("/category", r.post)
	rg.Put("/category/<id>", r.put)
	rg.Delete("/category/<id>", r.delete)
}

func (r *categoryResource) getAll(c *routing.Context) error {
	items, err := r.service.Get(app.GetRequestScope(c))
	if err != nil {
		return err
	}
	return c.Write(items)
}

func (r *categoryResource) get(c *routing.Context) error {
	id := c.Param("id")

	response, err := r.service.GetOne(app.GetRequestScope(c), bson.ObjectIdHex(id))
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *categoryResource) post(c *routing.Context) error {
	var model models.Category
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *categoryResource) put(c *routing.Context) error {
	id := c.Param("id")
	// var model models.Category
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

func (r *categoryResource) delete(c *routing.Context) error {
	id := c.Param("id")

	response, err := r.service.Delete(app.GetRequestScope(c), bson.ObjectIdHex(id))
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *categoryResource) home(c *routing.Context) error {
	type Response struct {
		About       *models.About        `json:"about"`
		Category    []models.Category    `json:"category"`
		Awards      []models.Awards      `json:"awards"`
		Testimonial []models.Testimonial `json:"testimonial"`
		Service     *models.Service      `json:"Service"`
	}

	var respo Response
	var err error

	respo.About, err = r.about.GetOne(app.GetRequestScope(c))
	if err != nil {
		return err
	}

	respo.Category, err = r.service.Get(app.GetRequestScope(c))
	if err != nil {
		return err
	}

	respo.Awards, err = r.award.Get(app.GetRequestScope(c))
	if err != nil {
		return err
	}

	respo.Testimonial, err = r.testimonial.Get(app.GetRequestScope(c))
	if err != nil {
		return err
	}

	respo.Service, err = r.ourservice.GetOne(app.GetRequestScope(c))
	if err != nil {
		return err
	}

	return c.Write(respo)
}
