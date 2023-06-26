package playlists

import (
	"fmt"
	"music-management/models"
	"music-management/pkg/helpers"
	"music-management/pkg/constants"
)

type PlaylistUsecase struct {
	repo PlaylistRepository
	helper helpers.Helper
}

func NewPlaylistUsecase(repo PlaylistRepository, helper helpers.Helper) *PlaylistUsecase {
	return &PlaylistUsecase{
		repo: repo,
		helper: helper,
	}
}

func (p *PlaylistUsecase) GetPlaylists() ([]string, string) {
	result, err := helpers.QueryTimeTwoOutput[[]*models.Playlist](p.repo.GetPlaylists)()
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	var output []string
	for _, v := range result {
		output = append(output, fmt.Sprintf("ID: %s, Name: %s, Duration: %.2f", v.ID, v.Name, v.Duration))
	}

	return output, ""
}

func (p *PlaylistUsecase) GetPlaylist() ([]string, string) {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	result, err := helpers.QueryTimeTwoOutputWithParams[*models.Playlist, string](p.repo.GetPlaylist)(id)
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	output := []string{fmt.Sprintf("ID: %s, Name: %s, Duration: %.2f", result.ID, result.Name, result.Duration)}

	return output, ""
}

func (p *PlaylistUsecase) CreatePlaylist() ([]string, string) {
	var name string
	fmt.Print("» Enter Name: ")
	fmt.Scanln(&name)

	var duration float32
	fmt.Print("» Enter Duration: ")
	fmt.Scanln(&duration)

	newPlaylist := &models.Playlist{
		ID: p.helper.GenerateID(),
		Name: name,
		Duration: duration,
	}

	result, err := helpers.QueryTimeTwoOutputWithParams[*models.Playlist](p.repo.CreatePlaylist)(newPlaylist)
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	output := []string{fmt.Sprintf("ID: %s, Name: %s, Duration: %.2f", result.ID, result.Name, result.Duration)}

	return output, ""
}

func (p *PlaylistUsecase) DeletePlaylist() error {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	err := helpers.QueryTimeOneOutputWithParams[error](p.repo.DeletePlaylist)(id)
	if err != nil {
		return err
	}

	return nil
}

func (p *PlaylistUsecase) UpdatePlaylist() ([]string, string) {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	var name string
	fmt.Print("» Enter Name: ")
	fmt.Scanln(&name)

	var duration float32
	fmt.Print("» Enter Duration: ")
	fmt.Scanln(&duration)

	newPlaylist := &models.Playlist{
		ID: id,
		Name: name,
		Duration: duration,
	}

	result, err := helpers.QueryTimeTwoOutputWithParams[*models.Playlist](p.repo.UpdatePlaylist)(newPlaylist)
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	output := []string{fmt.Sprintf("ID: %s, Name: %s, Duration: %.2f", result.ID, result.Name, result.Duration)}

	return output, ""
}

func (p *PlaylistUsecase) AddSongToPlaylist() string {
	return "Add Song to Playlist"
}

func (p *PlaylistUsecase) RemoveSongFromPlaylist() string {
	return "Remove Song from Playlist"
}

func (p *PlaylistUsecase) GetSongsFromPlaylist() string {
	return "Get Songs from Playlist"
}