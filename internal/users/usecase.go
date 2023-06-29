package users

import (
	"fmt"
	"music-management-system/models"
	"music-management-system/pkg/helpers"
	"music-management-system/pkg/constants"
)

type UserUsecase struct {
	repo UserRepository
	helper helpers.Helper
}

func NewUserUsecase(repo UserRepository, helper helpers.Helper) *UserUsecase {
	return &UserUsecase{
		repo: repo,
		helper: helper,
	}
}

func (a *UserUsecase) GetUsers() ([]string, string) {
	result, err := helpers.QueryTimeTwoOutput[[]*models.User](a.repo.GetUsers)()
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	var output []string
	for _, v := range result {
		output = append(output, fmt.Sprintf("ID: %s, Name: %s", v.ID, v.Name))
	}

	return output, ""
}

func (a *UserUsecase) GetUser() ([]string, string) {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	result, err := helpers.QueryTimeTwoOutputWithParams[*models.User, string](a.repo.GetUser)(id)
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	output := []string{fmt.Sprintf("ID: %s, Name: %s", result.ID, result.Name)}

	return output, ""
}

func (a *UserUsecase) CreateUser() ([]string, string) {
	var name string
	fmt.Print("» Enter name: ")
	fmt.Scanln(&name)

	user := &models.User{
		ID: a.helper.GenerateID(),
		Name: name,
	}

	result, err := helpers.QueryTimeTwoOutputWithParams[*models.User, *models.User](a.repo.CreateUser)(user)
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	output := []string{fmt.Sprintf("ID: %s", result.ID)}

	return output, ""
}

func (a *UserUsecase) DeleteUser() error {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	err := helpers.QueryTimeOneOutputWithParams[error](a.repo.DeleteUser)(id)
	if err != nil {
		return err
	}

	return nil
}

func (a *UserUsecase) UpdateUser() ([]string, string) {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)
	
	var name string
	fmt.Print("» Enter name: ")
	fmt.Scanln(&name)

	newUser := &models.User{
		ID: id,
		Name: name,
	}

	result, err := helpers.QueryTimeTwoOutputWithParams[*models.User, *models.User](a.repo.UpdateUser)(newUser)
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	output := []string{fmt.Sprintf("ID: %s", result.ID)}

	return output, ""
}