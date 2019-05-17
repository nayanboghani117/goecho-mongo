package serviceabout

import (
	"appScrip/app"
	"appScrip/models"

	"gopkg.in/mgo.v2/bson"
)

// aboutDAO specifies the interface of the about DAO needed by ArtistService.
type aboutDAO interface {
	// Get returns the all abouts.
	Get(rs app.RequestScope) ([]models.About, error)
	// GetOne returns the about with the specified about ID.
	GetOne(rs app.RequestScope) (*models.About, error)
	// Create saves a new about in the storage.
	Create(rs app.RequestScope, model *models.About) error
	// Update updates the about with given ID in the storage.
	Update(rs app.RequestScope, id bson.ObjectId, model *models.About) error
	// Delete removes the about with given ID from the storage.
	Delete(rs app.RequestScope, id bson.ObjectId) error
}

// Service provides services related with about.
type Service struct {
	dao aboutDAO
}

// NewService creates a new ArtistService with the given about DAO.
func NewService(dao aboutDAO) *Service {
	return &Service{dao}
}

// Get returns the about with the specified the about ID.
func (s *Service) Get(rs app.RequestScope) ([]models.About, error) {
	return s.dao.Get(rs)
}

// GetOne returns the about with the specified the about ID.
func (s *Service) GetOne(rs app.RequestScope) (*models.About, error) {
	return s.dao.GetOne(rs)
}

// Create creates a new about.
func (s *Service) Create(rs app.RequestScope, model *models.About) (*models.About, error) {
	// if err := model.Validate(); err != nil {
	// 	return nil, err
	// }
	err := s.dao.Create(rs, model)
	if err != nil {
		return nil, err
	}
	return s.dao.GetOne(rs)
}

// Update updates the about with the specified ID.
func (s *Service) Update(rs app.RequestScope, id bson.ObjectId, model *models.About) (*models.About, error) {
	// if err := model.Validate(); err != nil {
	// 	return nil, err
	// }
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.GetOne(rs)
}

// Delete deletes the about with the specified ID.
func (s *Service) Delete(rs app.RequestScope, id bson.ObjectId) (*models.About, error) {
	about, err := s.dao.GetOne(rs)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return about, err
}

// // Count returns the number of about.
// func (s *Service) Count(rs app.RequestScope) (int, error) {
// 	return s.dao.Count(rs)
// }

// // Query returns the about with the specified offset and limit.
// func (s *Service) Query(rs app.RequestScope, offset, limit int) ([]models.Artist, error) {
// 	return s.dao.Query(rs, offset, limit)
// }
