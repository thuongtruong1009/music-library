package playlists

import (
	"fmt"
	"music-management-system/models"
	"music-management-system/database"
	"music-management-system/pkg/constants"
	"music-management-system/pkg/helpers"
)

type PlaylistRepository struct {
	helper helpers.Helper
}

func NewPlaylistRepository(helper helpers.Helper) *PlaylistRepository {
	return &PlaylistRepository{
		helper: helper,
	}
}

func (p *PlaylistRepository) GetPlaylists() ([]*models.Playlist, error) {
	allPlaylist, err := database.ReadDB[*models.Playlist](constants.PLAYLIST_PATH)
	if err != nil {
		return nil, err
	}

	if allPlaylist == nil {
		return nil, nil
	}

	return allPlaylist, nil
}

func (p *PlaylistRepository) GetPlaylist(playlistID string) (*models.Playlist, error) {
	allPlaylist, _ := p.GetPlaylists()

	if allPlaylist == nil {
		return nil, nil
	}

	for _, v := range allPlaylist {
		if v.ID == playlistID {
			return v, nil
		}
	}

	return nil, nil
}

func (p *PlaylistRepository) CreatePlaylist(newPlaylist *models.Playlist) (*models.Playlist, error) {
	allPlaylist, _ := p.GetPlaylists()

	var playlistInit []*models.Playlist

	if allPlaylist == nil {
		playlistInit = make([]*models.Playlist, 0)
	} else {
		playlistInit = make([]*models.Playlist, len(allPlaylist))
		copy(playlistInit, allPlaylist)
	}

	playlistInit = append(playlistInit, newPlaylist)

	err := database.SaveDB[[]*models.Playlist](constants.PLAYLIST_PATH, playlistInit)
	if err != nil {
		return nil, fmt.Errorf("failed to save playlist: %w", err)
	}

	return newPlaylist, nil
}

func (p *PlaylistRepository) UpdatePlaylist(playlistUpdate *models.Playlist) (*models.Playlist, error) {
	playlist, err1 := p.GetPlaylist(playlistUpdate.ID)
	if err1 != nil {
		return nil, fmt.Errorf("failed to get playlist: %w", err1)
	}

	if playlist == nil {
		return nil, fmt.Errorf("playlist not found")
	}

	allPlaylist, _ := p.GetPlaylists()

	if allPlaylist == nil {
		return nil, fmt.Errorf("playlist not found")
	}

	var playlistInit []*models.Playlist

	for i, v := range allPlaylist {
		if v.ID == playlistUpdate.ID {
			playlistInit = append(allPlaylist[:i], playlistUpdate)
			break;
		}
	}

	playlistInit = append(playlistInit, allPlaylist[len(allPlaylist):]...)

	err2 := database.SaveDB[[]*models.Playlist](constants.PLAYLIST_PATH, playlistInit)
	if err2 != nil {
		return nil, fmt.Errorf("failed to save playlist: %w", err2)
	}

	return playlistUpdate, nil
}

func (p *PlaylistRepository) DeletePlaylist(playlistID string) error {
	allPlaylist, _ := p.GetPlaylists()

	if allPlaylist == nil {
		return fmt.Errorf("playlist not found")
	}

	for i, v := range allPlaylist {
		if v.ID == playlistID {
			allPlaylist = append(allPlaylist[:i], allPlaylist[i+1:]...)
			break
		}
	}

	err := database.SaveDB[[]*models.Playlist](constants.PLAYLIST_PATH, allPlaylist)
	if err != nil {
		return fmt.Errorf("failed to save playlist: %w", err)
	}

	return nil
}