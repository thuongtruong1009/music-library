package albums

import (
	"fmt"
	"music-management-system/models"
	"music-management-system/pkg/constants"
	"music-management-system/pkg/helpers"
)

type AlbumUsecase struct {
	repo AlbumRepository
	helper helpers.Helper
}

func NewAlbumUsecase(repo AlbumRepository, helper helpers.Helper) *AlbumUsecase {
	return &AlbumUsecase{
		repo: repo,
		helper: helper,
	}
}

func (a *AlbumUsecase) GetAlbums() ([]string, string){
	result, err := helpers.QueryTimeTwoOutput[[]*models.Album](a.repo.GetAlbums)()
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	if len(result) == 0 {
		return nil, constants.EMPTY_DATA
	}

	var output []string
	for _, v := range result {
		output = append(output, fmt.Sprintf("ID: %s, Title: %s, Duration: %.2f, Year: %d, Artist: %s", v.ID, v.Title, v.Duration, v.Year, v.ArtistID))
	}

	return output, ""
}

func (a *AlbumUsecase) GetAlbum() ([]string, string){
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	result, err := helpers.QueryTimeTwoOutputWithParams[*models.Album, string](a.repo.GetAlbum)(id)
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	output := []string{fmt.Sprintf("ID: %s, Title: %s, Duration: %.2f, Year: %d, Artist: %s", result.ID, result.Title, result.Duration, result.Year, result.ArtistID)}

	return output, ""
}

func (a *AlbumUsecase) CreateAlbum() ([]string, string) {
	var title string
	fmt.Print("» Enter Title: ")
	fmt.Scanln(&title)

	var duration float32
	fmt.Print("» Enter Duration: ")
	fmt.Scanln(&duration)

	var year int
	fmt.Print("» Enter Year: ")
	fmt.Scanln(&year)

	var artistID string
	fmt.Print("» Enter Artist ID: ")
	fmt.Scanln(&artistID)

	newAlbum := &models.Album{
		ID: a.helper.GenerateID(),
		Title: title,
		Duration: duration,
		Year: year,
		ArtistID: artistID,
	}

	result, err := helpers.QueryTimeTwoOutputWithParams[*models.Album, *models.Album](a.repo.CreateAlbum)(newAlbum)
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	output := []string{fmt.Sprintf("ID: %s, Title: %s, Duration: %.2f, Year: %d, Artist: %s", result.ID, result.Title, result.Duration, result.Year, result.ArtistID)}

	return output, ""
}

func (a *AlbumUsecase) DeleteAlbum() error {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	err := helpers.QueryTimeOneOutputWithParams[error](a.repo.DeleteAlbum)(id)
	if err != nil {
		return err
	}

	return nil
}

func (a *AlbumUsecase) UpdateAlbum() ([]string, string) {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	var title string
	fmt.Print("» Enter Title: ")
	fmt.Scanln(&title)

	var duration float32
	fmt.Print("» Enter Duration: ")
	fmt.Scanln(&duration)

	var year int
	fmt.Print("» Enter Year: ")
	fmt.Scanln(&year)

	var artistID string
	fmt.Print("» Enter Artist ID: ")
	fmt.Scanln(&artistID)

	newAlbum := &models.Album{
		ID: id,
		Title: title,
		Duration: duration,
		Year: year,
		ArtistID: artistID,
	}

	result, err := helpers.QueryTimeTwoOutputWithParams[*models.Album, *models.Album](a.repo.UpdateAlbum)(newAlbum)
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	output := []string{fmt.Sprintf("ID: %s, Title: %s, Duration: %.2f, Year: %d, Artist: %s", result.ID, result.Title, result.Duration, result.Year, result.ArtistID)}

	return output, ""
}

func (a *AlbumUsecase) GetTracksOfAlbum() string {
	return "Tracks of Album"
}