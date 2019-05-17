package servicetestimonial

import (
	"appScrip/app"
	"appScrip/models"

	"gopkg.in/mgo.v2/bson"
)

// testimonialDAO specifies the interface of the testimonial DAO needed by ArtistService.
type testimonialDAO interface {
	// Get returns the all testimonials.
	Get(rs app.RequestScope) ([]models.Testimonial, error)
	// GetOne returns the testimonial with the specified testimonial ID.
	GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Testimonial, error)
	// Create saves a new testimonial in the storage.
	Create(rs app.RequestScope, model *models.Testimonial) error
	// Update updates the testimonial with given ID in the storage.
	Update(rs app.RequestScope, id bson.ObjectId, model *models.Testimonial) error
	// Delete removes the testimonial with given ID from the storage.
	Delete(rs app.RequestScope, id bson.ObjectId) error
}

// Service provides services related with testimonials.
type Service struct {
	dao testimonialDAO
}

// NewService creates a new ArtistService with the given testimonial DAO.
func NewService(dao testimonialDAO) *Service {
	return &Service{dao}
}

// Get returns the testimonial with the specified the testimonial ID.
func (s *Service) Get(rs app.RequestScope) ([]models.Testimonial, error) {
	return s.dao.Get(rs)
}

// GetOne returns the testimonial with the specified the testimonial ID.
func (s *Service) GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Testimonial, error) {
	return s.dao.GetOne(rs, id)
}

// Create creates a new testimonial.
func (s *Service) Create(rs app.RequestScope, model *models.Testimonial) (*models.Testimonial, error) {
	// if err := model.Validate(); err != nil {
	// 	return nil, err
	// }
	err := s.dao.Create(rs, model)
	if err != nil {
		return nil, err
	}
	return s.dao.GetOne(rs, model.ID)
}

// Update updates the testimonial with the specified ID.
func (s *Service) Update(rs app.RequestScope, id bson.ObjectId, model *models.Testimonial) (*models.Testimonial, error) {
	// if err := model.Validate(); err != nil {
	// 	return nil, err
	// }
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.GetOne(rs, id)
}

// Delete deletes the testimonial with the specified ID.
func (s *Service) Delete(rs app.RequestScope, id bson.ObjectId) (*models.Testimonial, error) {
	testimonial, err := s.dao.GetOne(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return testimonial, err
}

// // Count returns the number of testimonials.
// func (s *Service) Count(rs app.RequestScope) (int, error) {
// 	return s.dao.Count(rs)
// }

// // Query returns the testimonials with the specified offset and limit.
// func (s *Service) Query(rs app.RequestScope, offset, limit int) ([]models.Artist, error) {
// 	return s.dao.Query(rs, offset, limit)
// }
