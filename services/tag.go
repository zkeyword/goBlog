package services

import (
	"BLOG/model"
	"BLOG/repository"
)

// TagService 文章服务
type TagService interface {
	Get(id int64) *model.Tag
	GetList() []model.Tag
	Create(Tag *model.Tag)
}

type tagService struct {
	repo *repository.TagRepository
}

// NewTagService 实例化TagService
var NewTagService = newTagService()

func newTagService() TagService {
	return &tagService{
		repo: repository.NewTagRepository(),
	}
}

func (s *tagService) Get(id int64) *model.Tag {
	return s.repo.Get(id)
}

func (s *tagService) GetList() []model.Tag {
	return s.repo.GetList()
}

func (s *tagService) Create(Tag *model.Tag) {
	s.repo.Create(Tag)
}
