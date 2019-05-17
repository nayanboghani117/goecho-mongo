package serviceuser

import (
	"appScrip/app"
	"appScrip/models"

	"gopkg.in/mgo.v2/bson"
)

// userDAO specifies the interface of the user DAO needed by ArtistService.
type userDAO interface {
	// Get returns the all blogs.
	Get(rs app.RequestScope) ([]models.User, error)
	// GetOne returns the user with the specified user ID.
	GetOne(rs app.RequestScope, id bson.ObjectId) (*models.User, error)
	// Create saves a new user in the storage.
	Create(rs app.RequestScope, model *models.User) error
	// Update updates the user with given ID in the storage.
	Update(rs app.RequestScope, id bson.ObjectId, model *models.User) error
	// Delete removes the user with given ID from the storage.
	Delete(rs app.RequestScope, id bson.ObjectId) error
	// Authenticate User
	Authenticate(rs app.RequestScope, email string, pass string) (*models.User, error)
}

// Service provides services related with users.
type Service struct {
	dao userDAO
}

// NewService creates a new ArtistService with the given user DAO.
func NewService(dao userDAO) *Service {
	return &Service{dao}
}

// Get returns the user with the specified the user ID.
func (s *Service) Get(rs app.RequestScope) ([]models.User, error) {
	return s.dao.Get(rs)
}

// GetOne returns the user with the specified the user ID.
func (s *Service) GetOne(rs app.RequestScope, id bson.ObjectId) (*models.User, error) {
	return s.dao.GetOne(rs, id)
}

// Create creates a new user.
func (s *Service) Create(rs app.RequestScope, model *models.User) (*models.User, error) {
	// if err := model.Validate(); err != nil {
	// 	return nil, err
	// }
	err := s.dao.Create(rs, model)
	if err != nil {
		return nil, err
	}
	return s.dao.GetOne(rs, model.ID)
}

// Update updates the user with the specified ID.
func (s *Service) Update(rs app.RequestScope, id bson.ObjectId, model *models.User) (*models.User, error) {
	// if err := model.Validate(); err != nil {
	// 	return nil, err
	// }
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.GetOne(rs, id)
}

// Delete deletes the user with the specified ID.
func (s *Service) Delete(rs app.RequestScope, id bson.ObjectId) (*models.User, error) {
	user, err := s.dao.GetOne(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return user, err
}

// Authenticate User
func (s *Service) Authenticate(rs app.RequestScope, email string, pass string) (*models.User, error) {
	return s.dao.Authenticate(rs, email, pass)
}

// // Count returns the number of users.
// func (s *Service) Count(rs app.RequestScope) (int, error) {
// 	return s.dao.Count(rs)
// }

// // Query returns the users with the specified offset and limit.
// func (s *Service) Query(rs app.RequestScope, offset, limit int) ([]models.Artist, error) {
// 	return s.dao.Query(rs, offset, limit)
// }
