package servicevideos

import (
	"appScrip/app"
	"appScrip/models"

	"gopkg.in/mgo.v2/bson"
)

// videosDAO specifies the interface of the video DAO needed by ArtistService.
type videosDAO interface {
	// Get returns the all videoss.
	Get(rs app.RequestScope) ([]models.Videos, error)
	// GetOne returns the video with the specified video ID.
	GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Videos, error)
	// Create saves a new video in the storage.
	Create(rs app.RequestScope, model *models.Videos) error
	// Update updates the video with given ID in the storage.
	Update(rs app.RequestScope, id bson.ObjectId, model *models.Videos) error
	// Delete removes the video with given ID from the storage.
	Delete(rs app.RequestScope, id bson.ObjectId) error
}

// Service provides services related with videos.
type Service struct {
	dao videosDAO
}

// NewService creates a new ArtistService with the given video DAO.
func NewService(dao videosDAO) *Service {
	return &Service{dao}
}

// Get returns the video with the specified the video ID.
func (s *Service) Get(rs app.RequestScope) ([]models.Videos, error) {
	return s.dao.Get(rs)
}

// GetOne returns the video with the specified the video ID.
func (s *Service) GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Videos, error) {
	return s.dao.GetOne(rs, id)
}

// Create creates a new video.
func (s *Service) Create(rs app.RequestScope, model *models.Videos) (*models.Videos, error) {
	// if err := model.Validate(); err != nil {
	// 	return nil, err
	// }
	err := s.dao.Create(rs, model)
	if err != nil {
		return nil, err
	}
	return s.dao.GetOne(rs, model.ID)
}

// Update updates the video with the specified ID.
func (s *Service) Update(rs app.RequestScope, id bson.ObjectId, model *models.Videos) (*models.Videos, error) {
	// if err := model.Validate(); err != nil {
	// 	return nil, err
	// }
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.GetOne(rs, id)
}

// Delete deletes the video with the specified ID.
func (s *Service) Delete(rs app.RequestScope, id bson.ObjectId) (*models.Videos, error) {
	video, err := s.dao.GetOne(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return video, err
}

// // Count returns the number of videos.
// func (s *Service) Count(rs app.RequestScope) (int, error) {
// 	return s.dao.Count(rs)
// }

// // Query returns the videos with the specified offset and limit.
// func (s *Service) Query(rs app.RequestScope, offset, limit int) ([]models.Artist, error) {
// 	return s.dao.Query(rs, offset, limit)
// }
