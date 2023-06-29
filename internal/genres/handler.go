package genres

import (
	"music-management-system/pkg/helpers"
	"music-management-system/pkg/constants"
)

type GenreHandler struct {
	uc GenreUsecase
	helper helpers.Helper
}

func NewGenreHandler(uc GenreUsecase, helper helpers.Helper) *GenreHandler {
	return &GenreHandler{
		uc: uc,
		helper: helper,
	}
}

func (u *GenreHandler) GetGenres() {
	result, err := u.uc.GetGenres()
	if err != "" {
		u.helper.OutputError(constants.GET_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.GET_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Genre", result, nil)
}

func (u *GenreHandler) GetGenre() {
	result, err := u.uc.GetGenre()
	if err != "" {
		u.helper.OutputError(constants.GET_FAILED, err)
		return
	}
	
	u.helper.OutputSuccess(constants.GET_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Genre", result, nil)
}

func (u *GenreHandler) CreateGenre() {
	result, err := u.uc.CreateGenre()
	if err != "" {
		u.helper.OutputError(constants.CREATE_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.CREATE_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Genre", result, nil)
}

func (u *GenreHandler) DeleteGenre() {
	err := u.uc.DeleteGenre()
	if err != "" {
		u.helper.OutputError(constants.DELETE_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.DELETE_SUCCESS)
}

func (u *GenreHandler) UpdateGenre() {
	result, err := u.uc.UpdateGenre()
	if err != "" {
		u.helper.OutputError(constants.UPDATE_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.UPDATE_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Genre", result, nil)
}

func (u *GenreHandler) GetTracksOfGenre() {
	u.helper.OutputSuccess(constants.GET_SUCCESS)
}