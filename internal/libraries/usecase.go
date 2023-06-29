package libraries

import (
	"fmt"
	"music-management-system/models"
	"music-management-system/pkg/helpers"
	"music-management-system/pkg/constants"
)

type LibraryUsecase struct {
	repo LibraryRepository
	helper helpers.Helper
}

func NewLibraryUsecase(repo LibraryRepository, helper helpers.Helper) *LibraryUsecase {
	return &LibraryUsecase{
		repo: repo,
		helper: helper,
	}
}

func (p *LibraryUsecase) GetLibraries() ([]string, string) {
	result, err := helpers.QueryTimeTwoOutput[[]*models.Library](p.repo.GetLibraries)()
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	var output []string
	for _, v := range result {
		output = append(output, fmt.Sprintf("Track ID: %s, Playlist ID: %s", v.TrackID, v.PlaylistID))
	}

	return output, ""
}

func (p *LibraryUsecase) GetLibrary() ([]string, string) {
	var trackID string
	fmt.Print("» Enter track id: ")
	fmt.Scanln(&trackID)

	var playlistID string
	fmt.Print("» Enter playlist id: ")
	fmt.Scanln(&playlistID)

	id := &models.Library{
		TrackID: trackID,
		PlaylistID: playlistID,
	}

	result, err := helpers.QueryTimeTwoOutputWithParams[*models.Library, *models.Library](p.repo.GetLibrary)(id)
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	output := []string{fmt.Sprintf("Track ID: %s, Playlist ID: %s", result.TrackID, result.PlaylistID)}

	return output, ""
}

func (p *LibraryUsecase) CreateLibrary() ([]string, string) {
	var trackID string
	fmt.Print("» Enter track id: ")
	fmt.Scanln(&trackID)

	var playlistID string
	fmt.Print("» Enter playlist id: ")
	fmt.Scanln(&playlistID)

	var userID string
	fmt.Print("» Enter user id: ")
	fmt.Scanln(&userID)


	newLibrary := &models.Library{
		TrackID: trackID,
		PlaylistID: playlistID,
		UserID: userID,
	}

	result, err := helpers.QueryTimeTwoOutputWithParams[*models.Library](p.repo.CreateLibrary)(newLibrary)
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	output := []string{fmt.Sprintf("Track ID: %s, User ID: %s, Playlist ID: %s", result.TrackID, result.UserID, result.PlaylistID)}

	return output, ""
}

func (p *LibraryUsecase) DeleteLibrary() error {
	var trackID string
	fmt.Print("» Enter track id: ")
	fmt.Scanln(&trackID)

	var playlistID string
	fmt.Print("» Enter playlist id: ")
	fmt.Scanln(&playlistID)

	id := &models.Library{
		TrackID: trackID,
		PlaylistID: playlistID,
	}

	err := helpers.QueryTimeOneOutputWithParams[error](p.repo.DeleteLibrary)(id)
	if err != nil {
		return err
	}

	return nil
}

func (p *LibraryUsecase) AddSongToLibrary() string {
	return "Add Song to Library"
}

func (p *LibraryUsecase) RemoveSongFromLibrary() string {
	return "Remove Song from Library"
}

func (p *LibraryUsecase) GetSongsFromLibrary() string {
	return "Get Songs from Library"
}