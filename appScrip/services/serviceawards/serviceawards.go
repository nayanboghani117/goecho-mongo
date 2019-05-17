package serviceawards

import (
	"appScrip/app"
	"appScrip/models"

	"gopkg.in/mgo.v2/bson"
)

// awardsDAO specifies the interface of the award DAO needed by ArtistService.
type awardsDAO interface {
	// Get returns the all awardss.
	Get(rs app.RequestScope) ([]models.Awards, error)
	// GetOne returns the award with the specified award ID.
	GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Awards, error)
	// Create saves a new award in the storage.
	Create(rs app.RequestScope, model *models.Awards) error
	// Update updates the award with given ID in the storage.
	Update(rs app.RequestScope, id bson.ObjectId, model *models.Awards) error
	// Delete removes the award with given ID from the storage.
	Delete(rs app.RequestScope, id bson.ObjectId) error
}

// Service provides services related with awards.
type Service struct {
	dao awardsDAO
}

// NewService creates a new ArtistService with the given award DAO.
func NewService(dao awardsDAO) *Service {
	return &Service{dao}
}

// Get returns the award with the specified the award ID.
func (s *Service) Get(rs app.RequestScope) ([]models.Awards, error) {
	return s.dao.Get(rs)
}

// GetOne returns the award with the specified the award ID.
func (s *Service) GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Awards, error) {
	return s.dao.GetOne(rs, id)
}

// Create creates a new award.
func (s *Service) Create(rs app.RequestScope, model *models.Awards) (*models.Awards, error) {
	// if err := model.Validate(); err != nil {
	// 	return nil, err
	// }
	err := s.dao.Create(rs, model)
	if err != nil {
		return nil, err
	}
	return s.dao.GetOne(rs, model.ID)
}

// Update updates the award with the specified ID.
func (s *Service) Update(rs app.RequestScope, id bson.ObjectId, model *models.Awards) (*models.Awards, error) {
	// if err := model.Validate(); err != nil {
	// 	return nil, err
	// }
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.GetOne(rs, id)
}

// Delete deletes the award with the specified ID.
func (s *Service) Delete(rs app.RequestScope, id bson.ObjectId) (*models.Awards, error) {
	award, err := s.dao.GetOne(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return award, err
}

// // Count returns the number of awards.
// func (s *Service) Count(rs app.RequestScope) (int, error) {
// 	return s.dao.Count(rs)
// }

// // Query returns the awards with the specified offset and limit.
// func (s *Service) Query(rs app.RequestScope, offset, limit int) ([]models.Artist, error) {
// 	return s.dao.Query(rs, offset, limit)
// }
