package libraries

import (
	"music-management-system/pkg/helpers"
	"music-management-system/pkg/constants"
)

type LibraryHandler struct {
	uc LibraryUsecase
	helper helpers.Helper
}

func NewLibraryHandler(uc LibraryUsecase, helper helpers.Helper) *LibraryHandler {
	return &LibraryHandler{
		uc: uc,
		helper: helper,
	}
}

func (u *LibraryHandler) GetLibraries() {
	result, err := u.uc.GetLibraries()
	if err != "" {
		u.helper.OutputError(constants.GET_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.GET_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Playlist", result, nil)
}

func (u *LibraryHandler) GetTracksOfPlaylist() {
	result, err := u.uc.GetLibrary()
	if err != "" {
		u.helper.OutputError(constants.GET_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.GET_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Playlist", result, nil)
}

func (u *LibraryHandler) AddTrackToPlaylist() {
	result, err := u.uc.CreateLibrary()
	if err != "" {
		u.helper.OutputError(constants.CREATE_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.CREATE_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Playlist", result, nil)
}

func (u *LibraryHandler) DeleteTrackFromPlaylist() {
	err := u.uc.DeleteLibrary()
	if err != nil {
		u.helper.OutputError(constants.DELETE_FAILED, err.Error())
		return
	}

	u.helper.OutputSuccess(constants.DELETE_SUCCESS)
}

func (u *LibraryHandler) GetPlaylistsContainTrack() {
	u.helper.OutputSuccess("GetPlaylistsHaveTrack")
}