package artists

import (
	"fmt"
	"music-management/models"
	"music-management/pkg/helpers"
	"music-management/pkg/constants"
)

type ArtistUsecase struct {
	repo ArtistRepository
	helper helpers.Helper
}

func NewArtistUsecase(repo ArtistRepository, helper helpers.Helper) *ArtistUsecase {
	return &ArtistUsecase{
		repo: repo,
		helper: helper,
	}
}

func (a *ArtistUsecase) GetArtists() ([]string, string) {
	result, err := helpers.QueryTimeTwoOutput[[]*models.Artist](a.repo.GetArtists)()
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	var output []string
	for _, v := range result {
		output = append(output, fmt.Sprintf("ID: %s, Name: %s", v.ID, v.Name))
	}

	return output, ""
}

func (a *ArtistUsecase) GetArtist() ([]string, string) {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	result, err := helpers.QueryTimeTwoOutputWithParams[*models.Artist, string](a.repo.GetArtist)(id)
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	output := []string{fmt.Sprintf("ID: %s, Name: %s", result.ID, result.Name)}

	return output, ""
}

func (a *ArtistUsecase) CreateArtist() ([]string, string) {
	var name string
	fmt.Print("» Enter name: ")
	fmt.Scanln(&name)

	artist := &models.Artist{
		ID: a.helper.GenerateID(),
		Name: name,
	}

	result, err := helpers.QueryTimeTwoOutputWithParams[*models.Artist, *models.Artist](a.repo.CreateArtist)(artist)
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	output := []string{fmt.Sprintf("ID: %s", result.ID)}

	return output, ""
}

func (a *ArtistUsecase) DeleteArtist() error {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	err := helpers.QueryTimeOneOutputWithParams[error](a.repo.DeleteArtist)(id)
	if err != nil {
		return err
	}

	return nil
}

func (a *ArtistUsecase) UpdateArtist() ([]string, string) {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)
	
	var name string
	fmt.Print("» Enter name: ")
	fmt.Scanln(&name)

	newArtist := &models.Artist{
		ID: id,
		Name: name,
	}

	result, err := helpers.QueryTimeTwoOutputWithParams[*models.Artist, *models.Artist](a.repo.UpdateArtist)(newArtist)
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	output := []string{fmt.Sprintf("ID: %s", result.ID)}

	return output, ""
}

func (a *ArtistUsecase) GetAlbumsOfArtist() string {
	return "Albums of Artist"
}

func (a *ArtistUsecase) GetTracksOfArtist() string {
	return "Tracks of Artist"
}