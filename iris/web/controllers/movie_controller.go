package controllers

import (
	"github.com/kataras/iris/mvc"
	"iris/repositories"
	"iris/services"
)

type MovieController struct {
}

func (c *MovieController) Get() mvc.View {
	repo := repositories.NewMovieManager()
	movie := services.NewMovieServiceManager(repo)
	movieName := movie.ShowMovieName()
	return mvc.View{
		Name: "/movie/index.html",
		Data: movieName,
	}
}
