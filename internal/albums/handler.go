package albums

import (
	"music-management/pkg/helpers"
	"music-management/pkg/constants"
)

type AlbumHandler struct {
	uc AlbumUsecase
	helper helpers.Helper
}

func NewAlbumHandler(uc AlbumUsecase, helper helpers.Helper) *AlbumHandler {
	return &AlbumHandler{
		uc: uc,
		helper: helper,
	}
}

func (u *AlbumHandler) GetAlbums() {
	result, err := u.uc.GetAlbums()
	if err != "" {
		u.helper.OutputError(constants.GET_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.GET_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Albums", result, nil)
}

func (u *AlbumHandler) GetAlbum() {
	result, err := u.uc.GetAlbum()
	if err != "" {
		u.helper.OutputError(constants.GET_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.GET_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Albums", result, nil)
}

func (u *AlbumHandler) CreateAlbum() {
	result, err := u.uc.CreateAlbum()
	if err != "" {
		u.helper.OutputError(constants.CREATE_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.CREATE_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Albums", result, nil)
}

func (u *AlbumHandler) DeleteAlbum() {
	err := u.uc.DeleteAlbum()
	if err != nil {
		u.helper.OutputError(constants.DELETE_FAILED, err.Error())
		return
	}

	u.helper.OutputSuccess(constants.DELETE_SUCCESS)
}

func (u *AlbumHandler) UpdateAlbum() {
	result, err := u.uc.UpdateAlbum()
	if err != "" {
		u.helper.OutputError(constants.UPDATE_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.UPDATE_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Albums", result, nil)
}

func (u *AlbumHandler) GetTracksOfAlbum() {
	u.helper.OutputSuccess("GetTracksOfAlbum")
}
