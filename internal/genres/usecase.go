package genres

import (
	"fmt"
	"music-management-system/models"
	"music-management-system/pkg/constants"
	"music-management-system/pkg/helpers"
)

type GenreUsecase struct {
	repo GenreRepository
	helper helpers.Helper
}

func NewGenreUsecase(repo GenreRepository, helper helpers.Helper) *GenreUsecase {
	return &GenreUsecase{
		repo: repo,
		helper: helper,
	}
}

func (g *GenreUsecase) GetGenres() ([]string, string) {
	result, err := helpers.QueryTimeTwoOutput[[]*models.Genre](g.repo.GetGenres)()

	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	var items []string
	for _, item := range result {
		items = append(items, fmt.Sprintf("Id: %s, Name: %s, Description: %s", item.ID, item.Name, item.Description))
	}

	return items, ""
}

func (g *GenreUsecase) GetGenre() ([]string, string) {
	var id string
	fmt.Print("» Enter id: ")
	fmt.Scanln(&id)

	result, err := helpers.QueryTimeTwoOutputWithParams[*models.Genre, string](g.repo.GetGenre)(id)
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	newItem := []string{fmt.Sprintf("Id: %s, Name: %s, Description: %s", result.ID, result.Name, result.Description)}

	return newItem, ""
}

func (g *GenreUsecase) CreateGenre() ([]string, string) {
	var name string
	fmt.Print("» Enter name: ")
	fmt.Scanln(&name)

	var description string
	fmt.Print("» Enter description: ")
	fmt.Scanln(&description)

	newGenre := &models.Genre{
		ID:          g.helper.GenerateID(),
		Name:        name,
		Description: description,
	}

	result, err := helpers.QueryTimeTwoOutputWithParams[*models.Genre, *models.Genre](g.repo.CreateGenre)(newGenre)
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	newItem := []string{fmt.Sprintf("Id: %s", result.ID)}

	return newItem, ""
}

func (g *GenreUsecase) DeleteGenre() string {
	var id string
	fmt.Print("» Enter id: ")
	fmt.Scanln(&id)

	err := helpers.QueryTimeOneOutputWithParams[error](g.repo.DeleteGenre)(id)
	if err != nil {
		return err.Error()
	}

	return ""
}

func (g *GenreUsecase) UpdateGenre() ([]string, string) {
	var id string
	fmt.Print("» Enter id: ")
	fmt.Scanln(&id)

	var name string
	fmt.Print("» Enter name: ")
	fmt.Scanln(&name)

	var description string
	fmt.Print("» Enter description: ")
	fmt.Scanln(&description)

	newGenre := &models.Genre{
		ID:          id,
		Name:        name,
		Description: description,
	}

	result, err := helpers.QueryTimeTwoOutputWithParams[*models.Genre, *models.Genre](g.repo.UpdateGenre)(newGenre)
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	newItem := []string{fmt.Sprintf("Id: %s, Name: %s, Description: %s", result.ID, result.Name, result.Description)}

	return newItem, ""

}

func (g *GenreUsecase) GetTracksOfGenre() string {
	return "Tracks of Genre"
}