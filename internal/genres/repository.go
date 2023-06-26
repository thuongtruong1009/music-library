package genres

import (
	"fmt"
	"music-management/database"
	"music-management/models"
	"music-management/pkg/constants"
	"music-management/pkg/helpers"
)

type GenreRepository struct {
	helper helpers.Helper
}

func NewGenreRepository(helper helpers.Helper) *GenreRepository {
	return &GenreRepository{
		helper: helper,
	}
}

func (g *GenreRepository) GetGenres() ([]*models.Genre, error) {
	genres, err := database.ReadDB[*models.Genre](constants.GENRE_PATH)

	if err != nil {
		return nil, err
	}

	if genres == nil {
		return nil, nil
	}

	return genres, nil
}

func (g *GenreRepository) GetGenre(genreID string) (*models.Genre, error) {
	allGenres, err := g.GetGenres()
	if err != nil {
		return nil, err
	}

	if allGenres == nil {
		return nil, nil
	}

	for _, genre := range allGenres {
		if genre.ID == genreID {
			return genre, nil
		}
	}

	return nil, nil
}

func (g *GenreRepository) CreateGenre(newGenre *models.Genre) (*models.Genre, error) {
	allGenres, _ := g.GetGenres()

	var genresInit []*models.Genre

	if allGenres == nil {
		genresInit = make([]*models.Genre, 0)
	} else {
		genresInit = make([]*models.Genre, len(allGenres))
		copy(genresInit, allGenres)
	}

	genresInit = append(genresInit, newGenre)

	err2 := database.SaveDB[[]*models.Genre](constants.GENRE_PATH, genresInit)
	if err2 != nil {
		return nil, fmt.Errorf(constants.CREATE_FAILED)
	}
	return newGenre, nil
}

func (g *GenreRepository) UpdateGenre(updateGenre *models.Genre) (*models.Genre, error) {
	allGenres, err := g.GetGenres()
	if err != nil {
		return nil, err
	}

	if allGenres == nil {
		return nil, nil
	}

	var genresInit []*models.Genre

	for i, genre := range allGenres {
		if genre.ID == updateGenre.ID {
			genresInit = append(allGenres[:i], updateGenre)
			break
		}
	}

	genresInit = append(genresInit, allGenres[len(genresInit):]...)

	err2 := database.SaveDB[[]*models.Genre](constants.GENRE_PATH, genresInit)
	if err2 != nil {
		return nil, fmt.Errorf(constants.UPDATE_FAILED)
	}
	return updateGenre, nil
}

func (g *GenreRepository) DeleteGenre(genreID string) error {
	allGenres, err := g.GetGenres()
	if err != nil {
		return err
	}

	if allGenres == nil {
		return nil
	}

	var genresInit []*models.Genre

	for i, genre := range allGenres {
		if genre.ID == genreID {
			genresInit = append(allGenres[:i], allGenres[i+1:]...)
			break
		}
	}

	err2 := database.SaveDB[[]*models.Genre](constants.GENRE_PATH, genresInit)
	if err2 != nil {
		return fmt.Errorf(constants.DELETE_FAILED)
	}

	return nil
}