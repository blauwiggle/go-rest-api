package comment

import "github.com/jinzhu/gorm"

// Service - the comment service
type Service struct {
	DB *gorm.DB
}

// Comment - defines our comment structure
type Comment struct {
	gorm.Model
	Slug   string
	Body   string
	Author string
}

// CommentService - the comment service interface
type CommentService interface {
	GetComment(id uint) (Comment, error)
	GetCommentsBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(id uint, comment Comment) (Comment, error)
	DeleteComment(id uint) error
	GetAllComments() ([]Comment, error)
}

// NewService - creates a new comment service
func NewService(db *gorm.DB) *Service {
	return &Service{DB: db}
}

// GetComment - gets a comment by id
func (s *Service) GetComment(id uint) (Comment, error) {
	var comment Comment
	if err := s.DB.First(&comment, id).Error; err != nil {
		return comment, err
	}
	return comment, nil
}

// GetCommentsBySlug - gets all comments for a slug
func (s *Service) GetCommentsBySlug(slug string) ([]Comment, error) {
	var comments []Comment
	if err := s.DB.Where("slug = ?", slug).Find(&comments).Error; err != nil {
		return comments, err
	}
	return comments, nil
}

// PostComment - posts a comment
func (s *Service) PostComment(comment Comment) (Comment, error) {
	if err := s.DB.Save(&comment).Error; err != nil {
		return comment, err
	}
	return comment, nil
}

// UpdateComment - updates a comment
func (s *Service) UpdateComment(id uint, comment Comment) (Comment, error) {
	if err := s.DB.Model(&comment).Where("id = ?", id).Updates(comment).Error; err != nil {
		return comment, err
	}
	return comment, nil
}

// DeleteComment - deletes a comment
func (s *Service) DeleteComment(id uint) error {
	if err := s.DB.Delete(&Comment{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetAllComments - gets all comments
func (s *Service) GetAllComments() ([]Comment, error) {
	var comments []Comment
	if err := s.DB.Find(&comments).Error; err != nil {
		return comments, err
	}
	return comments, nil
}
