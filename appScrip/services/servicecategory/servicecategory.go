package servicecategory

import (
	"appScrip/app"
	"appScrip/models"

	"gopkg.in/mgo.v2/bson"
)

// categoryDAO specifies the interface of the category DAO needed by ArtistService.
type categoryDAO interface {
	// Get returns the all category.
	Get(rs app.RequestScope) ([]models.Category, error)
	// GetOne returns the category with the specified category ID.
	GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Category, error)
	// Create saves a new category in the storage.
	Create(rs app.RequestScope, model *models.Category) error
	// Update updates the category with given ID in the storage.
	Update(rs app.RequestScope, id bson.ObjectId, model *models.Category) error
	// Delete removes the category with given ID from the storage.
	Delete(rs app.RequestScope, id bson.ObjectId) error
}

// Service provides services related with categorys.
type Service struct {
	dao categoryDAO
}

// NewService creates a new ArtistService with the given category DAO.
func NewService(dao categoryDAO) *Service {
	return &Service{dao}
}

// Get returns the category with the specified the category ID.
func (s *Service) Get(rs app.RequestScope) ([]models.Category, error) {
	return s.dao.Get(rs)
}

// GetOne returns the category with the specified the category ID.
func (s *Service) GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Category, error) {
	return s.dao.GetOne(rs, id)
}

// Create creates a new category.
func (s *Service) Create(rs app.RequestScope, model *models.Category) (*models.Category, error) {
	// if err := model.Validate(); err != nil {
	// 	return nil, err
	// }
	err := s.dao.Create(rs, model)
	if err != nil {
		return nil, err
	}
	return s.dao.GetOne(rs, model.ID)
}

// Update updates the category with the specified ID.
func (s *Service) Update(rs app.RequestScope, id bson.ObjectId, model *models.Category) (*models.Category, error) {
	// if err := model.Validate(); err != nil {
	// 	return nil, err
	// }
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.GetOne(rs, id)
}

// Delete deletes the category with the specified ID.
func (s *Service) Delete(rs app.RequestScope, id bson.ObjectId) (*models.Category, error) {
	category, err := s.dao.GetOne(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return category, err
}

// // Count returns the number of categorys.
// func (s *Service) Count(rs app.RequestScope) (int, error) {
// 	return s.dao.Count(rs)
// }

// // Query returns the categorys with the specified offset and limit.
// func (s *Service) Query(rs app.RequestScope, offset, limit int) ([]models.Artist, error) {
// 	return s.dao.Query(rs, offset, limit)
// }
