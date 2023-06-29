package artists

import (
	"music-management-system/pkg/helpers"
	"music-management-system/pkg/constants"
)

type ArtistHandler struct {
	uc ArtistUsecase
	helper helpers.Helper
}

func NewArtistHandler(uc ArtistUsecase, helper helpers.Helper) *ArtistHandler {
	return &ArtistHandler{
		uc: uc,
		helper: helper,
	}
}

func (u *ArtistHandler) GetArtists() {
	result, err := u.uc.GetArtists()
	if err != "" {
		u.helper.OutputError(constants.GET_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.GET_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Artists", result, nil)
}

func (u *ArtistHandler) GetArtist() {
	result, err := u.uc.GetArtist()
	if err != "" {
		u.helper.OutputError(constants.GET_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.GET_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Artists", result, nil)
}

func (u *ArtistHandler) CreateArtist() {
	result, err := u.uc.CreateArtist()
	if err != "" {
		u.helper.OutputError(constants.CREATE_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.CREATE_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Artists", result, nil)
}

func (u *ArtistHandler) DeleteArtist() {
	err := u.uc.DeleteArtist()
	if err != nil {
		u.helper.OutputError(constants.DELETE_FAILED, err.Error())
		return
	}

	u.helper.OutputSuccess(constants.DELETE_SUCCESS)
}

func (u *ArtistHandler) UpdateArtist() {
	result, err := u.uc.UpdateArtist()
	if err != "" {
		u.helper.OutputError(constants.UPDATE_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.UPDATE_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Artists", result, nil)
}

func (u *ArtistHandler) GetAlbumsOfArtist() {
	u.helper.OutputSuccess("GetAlbumsOfArtist")
}

func (u *ArtistHandler) GetTracksOfArtist() {
	u.helper.OutputSuccess("GetSongsOfArtist")
}
