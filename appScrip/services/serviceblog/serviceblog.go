package serviceblog

import (
	"appScrip/app"
	"appScrip/models"

	"gopkg.in/mgo.v2/bson"
)

// blogDAO specifies the interface of the blog DAO needed by ArtistService.
type blogDAO interface {
	// Get returns the all blogs.
	Get(rs app.RequestScope) ([]models.Blog, error)
	// GetOne returns the blog with the specified blog ID.
	GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Blog, error)
	// Create saves a new blog in the storage.
	Create(rs app.RequestScope, model *models.Blog) error
	// Update updates the blog with given ID in the storage.
	Update(rs app.RequestScope, id bson.ObjectId, model *models.Blog) error
	// Delete removes the blog with given ID from the storage.
	Delete(rs app.RequestScope, id bson.ObjectId) error
}

// Service provides services related with blogs.
type Service struct {
	dao blogDAO
}

// NewService creates a new ArtistService with the given blog DAO.
func NewService(dao blogDAO) *Service {
	return &Service{dao}
}

// Get returns the blog with the specified the blog ID.
func (s *Service) Get(rs app.RequestScope) ([]models.Blog, error) {
	return s.dao.Get(rs)
}

// GetOne returns the blog with the specified the blog ID.
func (s *Service) GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Blog, error) {
	return s.dao.GetOne(rs, id)
}

// Create creates a new blog.
func (s *Service) Create(rs app.RequestScope, model *models.Blog) (*models.Blog, error) {
	// if err := model.Validate(); err != nil {
	// 	return nil, err
	// }
	err := s.dao.Create(rs, model)
	if err != nil {
		return nil, err
	}
	return s.dao.GetOne(rs, model.ID)
}

// Update updates the blog with the specified ID.
func (s *Service) Update(rs app.RequestScope, id bson.ObjectId, model *models.Blog) (*models.Blog, error) {
	// if err := model.Validate(); err != nil {
	// 	return nil, err
	// }
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.GetOne(rs, id)
}

// Delete deletes the blog with the specified ID.
func (s *Service) Delete(rs app.RequestScope, id bson.ObjectId) (*models.Blog, error) {
	blog, err := s.dao.GetOne(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return blog, err
}

// // Count returns the number of blogs.
// func (s *Service) Count(rs app.RequestScope) (int, error) {
// 	return s.dao.Count(rs)
// }

// // Query returns the blogs with the specified offset and limit.
// func (s *Service) Query(rs app.RequestScope, offset, limit int) ([]models.Artist, error) {
// 	return s.dao.Query(rs, offset, limit)
// }
