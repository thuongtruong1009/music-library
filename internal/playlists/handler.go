package playlists

import (
	"music-management/pkg/helpers"
	"music-management/pkg/constants"
)

type PlaylistHandler struct {
	uc PlaylistUsecase
	helper helpers.Helper
}

func NewPlaylistHandler(uc PlaylistUsecase, helper helpers.Helper) *PlaylistHandler {
	return &PlaylistHandler{
		uc: uc,
		helper: helper,
	}
}

func (u *PlaylistHandler) GetPlaylists() {
	result, err := u.uc.GetPlaylists()
	if err != "" {
		u.helper.OutputError(constants.GET_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.GET_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Playlists", result, nil)
}

func (u *PlaylistHandler) GetPlaylist() {
	result, err := u.uc.GetPlaylist()
	if err != "" {
		u.helper.OutputError(constants.GET_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.GET_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Playlists", result, nil)
}

func (u *PlaylistHandler) CreatePlaylist() {
	result, err := u.uc.CreatePlaylist()
	if err != "" {
		u.helper.OutputError(constants.CREATE_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.CREATE_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Playlists", result, nil)
}

func (u *PlaylistHandler) DeletePlaylist() {
	err := u.uc.DeletePlaylist()
	if err != nil {
		u.helper.OutputError(constants.DELETE_FAILED, err.Error())
		return
	}

	u.helper.OutputSuccess(constants.DELETE_SUCCESS)
}

func (u *PlaylistHandler) UpdatePlaylist() {
	 result, err := u.uc.UpdatePlaylist()
	if err != "" {
		u.helper.OutputError(constants.UPDATE_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.UPDATE_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Playlists", result, nil)
}

func (u *PlaylistHandler) GetTracksOfPlaylist() {
	u.helper.OutputSuccess("GetTracksOfPlaylist")
}

func (u *PlaylistHandler) AddTrackToPlaylist() {
	u.helper.OutputSuccess("AddTrackToPlaylist")
}

func (u *PlaylistHandler) DeleteTrackFromPlaylist() {
	u.helper.OutputSuccess("DeleteTrackFromPlaylist")
}

func (u *PlaylistHandler) GetPlaylistsHaveTrack() {
	u.helper.OutputSuccess("GetPlaylistsHaveTrack")
}