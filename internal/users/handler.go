package users

import (
	"music-management-system/pkg/helpers"
	"music-management-system/pkg/constants"
)

type UserHandler struct {
	uc UserUsecase
	helper helpers.Helper
}

func NewUserHandler(uc UserUsecase, helper helpers.Helper) *UserHandler {
	return &UserHandler{
		uc: uc,
		helper: helper,
	}
}

func (u *UserHandler) GetUsers() {
	result, err := u.uc.GetUsers()
	if err != "" {
		u.helper.OutputError(constants.GET_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.GET_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("Users", result, nil)
}

func (u *UserHandler) GetUser() {
	result, err := u.uc.GetUser()
	if err != "" {
		u.helper.OutputError(constants.GET_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.GET_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("User", result, nil)
}

func (u *UserHandler) CreateUser() {
	result, err := u.uc.CreateUser()
	if err != "" {
		u.helper.OutputError(constants.CREATE_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.CREATE_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("User", result, nil)
}

func (u *UserHandler) DeleteUser() {
	err := u.uc.DeleteUser()
	if err != nil {
		u.helper.OutputError(constants.DELETE_FAILED, err.Error())
		return
	}

	u.helper.OutputSuccess(constants.DELETE_SUCCESS)
}

func (u *UserHandler) UpdateUser() {
	result, err := u.uc.UpdateUser()
	if err != "" {
		u.helper.OutputError(constants.UPDATE_FAILED, err)
		return
	}

	u.helper.OutputSuccess(constants.UPDATE_SUCCESS)

	helpers.TableOutput[string, string, interface{}]("User", result, nil)
}
