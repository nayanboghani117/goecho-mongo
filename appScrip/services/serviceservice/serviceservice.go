package serviceservice

import (
	"appScrip/app"
	"appScrip/models"

	"gopkg.in/mgo.v2/bson"
)

// serviceDAO specifies the interface of the service DAO needed by ArtistService.
type serviceDAO interface {
	// Get returns the all services.
	Get(rs app.RequestScope) ([]models.Service, error)
	// GetOne returns the service with the specified service ID.
	GetOne(rs app.RequestScope) (*models.Service, error)
	// Create saves a new service in the storage.
	Create(rs app.RequestScope, model *models.Service) error
	// Update updates the service with given ID in the storage.
	Update(rs app.RequestScope, id bson.ObjectId, model *models.Service) error
	// Delete removes the service with given ID from the storage.
	Delete(rs app.RequestScope, id bson.ObjectId) error
}

// Service provides services related with services.
type Service struct {
	dao serviceDAO
}

// NewService creates a new ArtistService with the given service DAO.
func NewService(dao serviceDAO) *Service {
	return &Service{dao}
}

// Get returns the service with the specified the service ID.
func (s *Service) Get(rs app.RequestScope) ([]models.Service, error) {
	return s.dao.Get(rs)
}

// GetOne returns the service with the specified the service ID.
func (s *Service) GetOne(rs app.RequestScope) (*models.Service, error) {
	return s.dao.GetOne(rs)
}

// Create creates a new service.
func (s *Service) Create(rs app.RequestScope, model *models.Service) (*models.Service, error) {
	// if err := model.Validate(); err != nil {
	// 	return nil, err
	// }
	err := s.dao.Create(rs, model)
	if err != nil {
		return nil, err
	}
	return s.dao.GetOne(rs)
}

// Update updates the service with the specified ID.
func (s *Service) Update(rs app.RequestScope, id bson.ObjectId, model *models.Service) (*models.Service, error) {
	// if err := model.Validate(); err != nil {
	// 	return nil, err
	// }
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.GetOne(rs)
}

// Delete deletes the service with the specified ID.
func (s *Service) Delete(rs app.RequestScope, id bson.ObjectId) (*models.Service, error) {
	service, err := s.dao.GetOne(rs)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return service, err
}

// // Count returns the number of services.
// func (s *Service) Count(rs app.RequestScope) (int, error) {
// 	return s.dao.Count(rs)
// }

// // Query returns the services with the specified offset and limit.
// func (s *Service) Query(rs app.RequestScope, offset, limit int) ([]models.Artist, error) {
// 	return s.dao.Query(rs, offset, limit)
// }
