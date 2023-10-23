package services

import (
	"fmt"
	"iris/repositories"
)

type MovieService interface {
	ShowMovieName() string
}

type MovieServiceManager struct {
	repo repositories.MovieRepository
}

func (m *MovieServiceManager) ShowMovieName() string {
	fmt.Println("获取到的视频名称为:" + m.repo.GetMovieName())
	return m.repo.GetMovieName()
}

func NewMovieServiceManager(repo repositories.MovieRepository) MovieService {
	return &MovieServiceManager{repo: repo}
}
