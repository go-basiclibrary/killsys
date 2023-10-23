package repositories

import "iris/datamodels"

type MovieRepository interface {
	GetMovieName() string
}

type MovieManager struct {
}

func (m *MovieManager) GetMovieName() string {
	// 模拟查询复制
	movie := &datamodels.Movie{Name: "wangshao"}
	return movie.Name
}

func NewMovieManager() MovieRepository {
	return &MovieManager{}
}
