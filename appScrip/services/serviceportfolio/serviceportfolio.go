package serviceportfolio

import (
	"appScrip/app"
	"appScrip/models"

	"gopkg.in/mgo.v2/bson"
)

// portfolioDAO specifies the interface of the portfolio DAO needed by ArtistService.
type portfolioDAO interface {
	// Get returns the all portfolios.
	Get(rs app.RequestScope) ([]models.Portfolio, error)
	// Get returns the all portfolios.
	GetByCategory(rs app.RequestScope, id string) ([]models.Portfolio, error)
	// GetOne returns the portfolio with the specified portfolio ID.
	GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Portfolio, error)
	// Create saves a new portfolio in the storage.
	Create(rs app.RequestScope, model *models.Portfolio) error
	// Update updates the portfolio with given ID in the storage.
	Update(rs app.RequestScope, id bson.ObjectId, model *models.Portfolio) error
	// Delete removes the portfolio with given ID from the storage.
	Delete(rs app.RequestScope, id bson.ObjectId) error
}

// Service provides services related with portfolios.
type Service struct {
	dao portfolioDAO
}

// NewService creates a new ArtistService with the given portfolio DAO.
func NewService(dao portfolioDAO) *Service {
	return &Service{dao}
}

// Get returns the portfolio with the specified the portfolio ID.
func (s *Service) Get(rs app.RequestScope) ([]models.Portfolio, error) {
	return s.dao.Get(rs)
}

// GetByCategory returns the portfolio with the specified the portfolio ID.
func (s *Service) GetByCategory(rs app.RequestScope, id string) ([]models.Portfolio, error) {
	return s.dao.GetByCategory(rs, id)
}

// GetOne returns the portfolio with the specified the portfolio ID.
func (s *Service) GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Portfolio, error) {
	return s.dao.GetOne(rs, id)
}

// Create creates a new portfolio.
func (s *Service) Create(rs app.RequestScope, model *models.Portfolio) (*models.Portfolio, error) {
	// if err := model.Validate(); err != nil {
	// 	return nil, err
	// }
	err := s.dao.Create(rs, model)
	if err != nil {
		return nil, err
	}
	return s.dao.GetOne(rs, model.ID)
}

// Update updates the portfolio with the specified ID.
func (s *Service) Update(rs app.RequestScope, id bson.ObjectId, model *models.Portfolio) (*models.Portfolio, error) {
	// if err := model.Validate(); err != nil {
	// 	return nil, err
	// }
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.GetOne(rs, id)
}

// Delete deletes the portfolio with the specified ID.
func (s *Service) Delete(rs app.RequestScope, id bson.ObjectId) (*models.Portfolio, error) {
	portfolio, err := s.dao.GetOne(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return portfolio, err
}

// // Count returns the number of portfolios.
// func (s *Service) Count(rs app.RequestScope) (int, error) {
// 	return s.dao.Count(rs)
// }

// // Query returns the portfolios with the specified offset and limit.
// func (s *Service) Query(rs app.RequestScope, offset, limit int) ([]models.Artist, error) {
// 	return s.dao.Query(rs, offset, limit)
// }
