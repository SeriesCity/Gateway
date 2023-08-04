package ports

import "github.com/SeriesCity/Gateway/internal/core/entities"

type MovieServiceContract interface {
	GetAllMovies() []entities.Movie
	GetPinnedMovies() []entities.Movie
}
type MovieHandlerContract interface {
	GetAllMovies() []entities.Movie
}
