package servicecontactus

import (
	"appScrip/app"
	"appScrip/models"

	"gopkg.in/mgo.v2/bson"
)

// contactusDAO specifies the interface of the contactus DAO needed by ArtistService.
type contactusDAO interface {
	// Get returns the all contactuss.
	Get(rs app.RequestScope) ([]models.ContactUs, error)
	// GetOne returns the contactus with the specified contactus ID.
	GetOne(rs app.RequestScope, id bson.ObjectId) (*models.ContactUs, error)
	// Create saves a new contactus in the storage.
	Create(rs app.RequestScope, model *models.ContactUs) error
	// Update updates the contactus with given ID in the storage.
	Update(rs app.RequestScope, id bson.ObjectId, model *models.ContactUs) error
	// Delete removes the contactus with given ID from the storage.
	Delete(rs app.RequestScope, id bson.ObjectId) error
}

// Service provides services related with contactuss.
type Service struct {
	dao contactusDAO
}

// NewService creates a new ArtistService with the given contactus DAO.
func NewService(dao contactusDAO) *Service {
	return &Service{dao}
}

// Get returns the contactus with the specified the contactus ID.
func (s *Service) Get(rs app.RequestScope) ([]models.ContactUs, error) {
	return s.dao.Get(rs)
}

// GetOne returns the contactus with the specified the contactus ID.
func (s *Service) GetOne(rs app.RequestScope, id bson.ObjectId) (*models.ContactUs, error) {
	return s.dao.GetOne(rs, id)
}

// Create creates a new contactus.
func (s *Service) Create(rs app.RequestScope, model *models.ContactUs) (*models.ContactUs, error) {
	// if err := model.Validate(); err != nil {
	// 	return nil, err
	// }
	err := s.dao.Create(rs, model)
	if err != nil {
		return nil, err
	}
	return s.dao.GetOne(rs, model.ID)
}

// Update updates the contactus with the specified ID.
func (s *Service) Update(rs app.RequestScope, id bson.ObjectId, model *models.ContactUs) (*models.ContactUs, error) {
	// if err := model.Validate(); err != nil {
	// 	return nil, err
	// }
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.GetOne(rs, id)
}

// Delete deletes the contactus with the specified ID.
func (s *Service) Delete(rs app.RequestScope, id bson.ObjectId) (*models.ContactUs, error) {
	contactus, err := s.dao.GetOne(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return contactus, err
}

// // Count returns the number of contactuss.
// func (s *Service) Count(rs app.RequestScope) (int, error) {
// 	return s.dao.Count(rs)
// }

// // Query returns the contactuss with the specified offset and limit.
// func (s *Service) Query(rs app.RequestScope, offset, limit int) ([]models.Artist, error) {
// 	return s.dao.Query(rs, offset, limit)
// }
