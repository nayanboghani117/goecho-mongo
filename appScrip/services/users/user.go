package users

import (
	"appScrip/app"
	"appScrip/models"

	"gopkg.in/mgo.v2/bson"
)

// userDAO specifies the interface of the artist DAO needed by ArtistService.
type userDAO interface {
	// Get returns the all blogs.
	Get(rs app.RequestScope) ([]models.User, error)
	// GetOne returns the artist with the specified artist ID.
	GetOne(rs app.RequestScope, id bson.ObjectId) (*models.User, error)
	// Create saves a new artist in the storage.
	Create(rs app.RequestScope, model *models.User) error
	// Update updates the artist with given ID in the storage.
	Update(rs app.RequestScope, id bson.ObjectId, model *models.User) error
	// Delete removes the artist with given ID from the storage.
	Delete(rs app.RequestScope, id bson.ObjectId) error
	// Authenticate User
	Authenticate(rs app.RequestScope, email string, pass string) (*models.User, error)
}

// UserService provides services related with artists.
type UserService struct {
	dao userDAO
}

// NewUserService creates a new ArtistService with the given artist DAO.
func NewUserService(dao userDAO) *UserService {
	return &UserService{dao}
}

// Get returns the artist with the specified the artist ID.
func (s *UserService) Get(rs app.RequestScope) ([]models.User, error) {
	return s.dao.Get(rs)
}

// GetOne returns the artist with the specified the artist ID.
func (s *UserService) GetOne(rs app.RequestScope, id bson.ObjectId) (*models.User, error) {
	return s.dao.GetOne(rs, id)
}

// Create creates a new artist.
func (s *UserService) Create(rs app.RequestScope, model *models.User) (*models.User, error) {
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
func (s *UserService) Update(rs app.RequestScope, id bson.ObjectId, model *models.User) (*models.User, error) {
	// if err := model.Validate(); err != nil {
	// 	return nil, err
	// }
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.GetOne(rs, id)
}

// Delete deletes the artist with the specified ID.
func (s *UserService) Delete(rs app.RequestScope, id bson.ObjectId) (*models.User, error) {
	artist, err := s.dao.GetOne(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return artist, err
}

// Authenticate User
func (s *UserService) Authenticate(rs app.RequestScope, email string, pass string) (*models.User, error) {
	return s.dao.Authenticate(rs, email, pass)
}

// // Count returns the number of artists.
// func (s *ArtistService) Count(rs app.RequestScope) (int, error) {
// 	return s.dao.Count(rs)
// }

// // Query returns the artists with the specified offset and limit.
// func (s *ArtistService) Query(rs app.RequestScope, offset, limit int) ([]models.Artist, error) {
// 	return s.dao.Query(rs, offset, limit)
// }
