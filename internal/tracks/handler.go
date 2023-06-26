package tracks

import (
	"music-management/pkg/helpers"
	"music-management/pkg/constants"
)

type TrackHandler struct {
	uc TrackUsecase
	helper helpers.Helper
}

func NewTrackHandler(uc TrackUsecase, helper helpers.Helper) *TrackHandler {
	return &TrackHandler{
		uc: uc,
		helper: helper,
	}
}

func (u *TrackHandler) GetTracks() {
	result, err := u.uc.GetTracks()
	if err != "" {
		u.helper.OutputError(constants.GET_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.GET_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Tracks", result, nil)
}

func (u *TrackHandler) GetTrack() {
	result, err := u.uc.GetTrack()
	if err != "" {
		u.helper.OutputError(constants.GET_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.GET_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Tracks", result, nil)
}

func (u *TrackHandler) CreateTrack() {
	result, err := u.uc.CreateTrack()
	if err != "" {
		u.helper.OutputError(constants.CREATE_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.CREATE_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Tracks", result, nil)	
}

func (u *TrackHandler) DeleteTrack() {
	err := u.uc.DeleteTrack()
	if err != nil {
		u.helper.OutputError(constants.DELETE_FAILED, err.Error())
		return
	}

	u.helper.OutputSuccess(constants.DELETE_SUCCESS)
}

func (u *TrackHandler) UpdateTrack() {
	result, err := u.uc.UpdateTrack()
	if err != "" {
		u.helper.OutputError(constants.UPDATE_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.UPDATE_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Tracks", result, nil)
}