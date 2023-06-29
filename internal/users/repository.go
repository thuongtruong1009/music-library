package users

import (
	"fmt"
	"music-management-system/models"
	"music-management-system/database"
	"music-management-system/pkg/constants"
	"music-management-system/pkg/helpers"
)

type UserRepository struct {
	helper helpers.Helper
}

func NewUserRepository(helper helpers.Helper) *UserRepository {
	return &UserRepository{
		helper: helper,
	}
}

func (a *UserRepository) GetUsers() ([]*models.User, error) {
	allUser, err := database.ReadDB[*models.User](constants.USER_PATH)
	if err != nil {
		return nil, err
	}

	if allUser == nil {
		return nil, nil
	}

	return allUser, nil
}

func (a *UserRepository) GetUser(userID string) (*models.User, error) {
	allUser, _ := a.GetUsers()

	if allUser == nil {
		return nil, nil
	}

	for _, v := range allUser {
		if v.ID == userID {
			return v, nil
		}
	}

	return nil, nil
}

func (a *UserRepository) CreateUser(newUser *models.User) (*models.User, error) {
	allUsers, _ := a.GetUsers()

	var userInit []*models.User

	if allUsers == nil {
		userInit = make([]*models.User, 0)
	} else {
		userInit = make([]*models.User, len(allUsers))
		copy(userInit, allUsers)
	}

	userInit = append(userInit, newUser)

	err2 := database.SaveDB[[]*models.User](constants.USER_PATH, userInit)
	if err2 != nil {
		return nil, fmt.Errorf(constants.CREATE_FAILED)
	}
	return newUser, nil
}

func (a *UserRepository) UpdateUser(userUpdate *models.User) (*models.User, error) {
	allUsers, _ := a.GetUsers()

	if allUsers == nil {
		return nil, fmt.Errorf(constants.UPDATE_FAILED)
	}

	var userInit []*models.User

	for i, v := range allUsers {
		if v.ID == userUpdate.ID {
			userInit = append(allUsers[:i], userUpdate)
			break
		}
	}

	userInit = append(userInit, allUsers[len(userInit):]...)

	err2 := database.SaveDB[[]*models.User](constants.USER_PATH, allUsers)
	if err2 != nil {
		return nil, fmt.Errorf(constants.UPDATE_FAILED)
	}
	return userUpdate, nil
}

func (a *UserRepository) DeleteUser(userID string) error {
	allUser, _ := a.GetUsers()

	if allUser == nil {
		return fmt.Errorf(constants.DELETE_FAILED)
	}

	for i, v := range allUser{
		if v.ID == userID {
			allUser = append(allUser[:i], allUser[i+1:]...)
			break
		}
	}

	err2 := database.SaveDB[[]*models.User](constants.USER_PATH, allUser)
	if err2 != nil {
		return fmt.Errorf(constants.DELETE_FAILED)
	}
	return nil
}
