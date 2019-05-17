package blogs

import (
	"appScrip/app"
	"appScrip/models"

	"gopkg.in/mgo.v2/bson"
)

// blogDAO specifies the interface of the artist DAO needed by ArtistService.
type blogDAO interface {
	// Get returns the all blogs.
	Get(rs app.RequestScope) ([]models.Blog, error)
	// GetOne returns the artist with the specified artist ID.
	GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Blog, error)
	// Create saves a new artist in the storage.
	Create(rs app.RequestScope, model *models.Blog) error
	// Update updates the artist with given ID in the storage.
	Update(rs app.RequestScope, id bson.ObjectId, model *models.Blog) error
	// Delete removes the artist with given ID from the storage.
	Delete(rs app.RequestScope, id bson.ObjectId) error
}

// BlogService provides services related with artists.
type BlogService struct {
	dao blogDAO
}

// NewBlogService creates a new ArtistService with the given artist DAO.
func NewBlogService(dao blogDAO) *BlogService {
	return &BlogService{dao}
}

// Get returns the artist with the specified the artist ID.
func (s *BlogService) Get(rs app.RequestScope) ([]models.Blog, error) {
	return s.dao.Get(rs)
}

// GetOne returns the artist with the specified the artist ID.
func (s *BlogService) GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Blog, error) {
	return s.dao.GetOne(rs, id)
}

// Create creates a new artist.
func (s *BlogService) Create(rs app.RequestScope, model *models.Blog) (*models.Blog, error) {
	// if err := model.Validate(); err != nil {
	// 	return nil, err
	// }
	err := s.dao.Create(rs, model)
	if err != nil {
		return nil, err
	}
	return s.dao.GetOne(rs, model.ID)
}

// Update updates the artist with the specified ID.
func (s *BlogService) Update(rs app.RequestScope, id bson.ObjectId, model *models.Blog) (*models.Blog, error) {
	// if err := model.Validate(); err != nil {
	// 	return nil, err
	// }
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.GetOne(rs, id)
}

// Delete deletes the artist with the specified ID.
func (s *BlogService) Delete(rs app.RequestScope, id bson.ObjectId) (*models.Blog, error) {
	artist, err := s.dao.GetOne(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return artist, err
}

// // Count returns the number of artists.
// func (s *ArtistService) Count(rs app.RequestScope) (int, error) {
// 	return s.dao.Count(rs)
// }

// // Query returns the artists with the specified offset and limit.
// func (s *ArtistService) Query(rs app.RequestScope, offset, limit int) ([]models.Artist, error) {
// 	return s.dao.Query(rs, offset, limit)
// }
