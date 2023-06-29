package libraries

import (
	"fmt"
	"music-management-system/models"
	"music-management-system/database"
	"music-management-system/pkg/constants"
	"music-management-system/pkg/helpers"
	"music-management-system/internal/tracks"
	"music-management-system/internal/users"
)

type LibraryRepository struct {
	trackRepo tracks.TrackRepository
	userRepo users.UserRepository
	helper helpers.Helper
}

func NewLibraryRepository(trackRepo tracks.TrackRepository, userRepo users.UserRepository, helper helpers.Helper) *LibraryRepository {
	return &LibraryRepository{
		trackRepo: trackRepo,
		userRepo: userRepo,
		helper: helper,
	}
}

func (p *LibraryRepository) GetLibraries() ([]*models.Library, error) {
	allLibrary, err := database.ReadDB[*models.Library](constants.LIBRARY_PATH)
	if err != nil {
		return nil, err
	}

	if allLibrary == nil {
		return nil, nil
	}

	return allLibrary, nil
}

func (p *LibraryRepository) GetLibrary(library *models.Library) (*models.Library, error) {
	allLibrary, _ := p.GetLibraries()

	if allLibrary == nil {
		return nil, nil
	}

	for _, v := range allLibrary {
		if v.TrackID == library.TrackID && v.PlaylistID == library.PlaylistID {
			return v, nil
		}
	}

	return nil, nil
}

func (p *LibraryRepository) CreateLibrary(newLibrary *models.Library) (*models.Library, error) {
	trackExist, err1 := p.trackRepo.GetTrack(newLibrary.TrackID)
	if err1 != nil {
		return nil, fmt.Errorf("failed to get track: %w", err1)
	}

	if trackExist == nil {
		return nil, fmt.Errorf("track not found")
	}

	userExist, err2 := p.userRepo.GetUser(newLibrary.UserID)
	if err2 != nil {
		return nil, fmt.Errorf("failed to get user: %w", err2)
	}

	if userExist == nil {
		return nil, fmt.Errorf("user not found")
	}

	allLibrary, _ := p.GetLibraries()

	var LibraryInit []*models.Library

	if allLibrary == nil {
		LibraryInit = make([]*models.Library, 0)
	} else {
		LibraryInit = make([]*models.Library, len(allLibrary))
		copy(LibraryInit, allLibrary)
	}

	LibraryInit = append(LibraryInit, newLibrary)

	err3 := database.SaveDB[[]*models.Library](constants.LIBRARY_PATH, LibraryInit)
	if err3 != nil {
		return nil, fmt.Errorf("failed to save Library: %w", err3)
	}

	return newLibrary, nil
}

func (p *LibraryRepository) DeleteLibrary(library *models.Library) error {
	allLibrary, _ := p.GetLibraries()

	if allLibrary == nil {
		return fmt.Errorf("Library not found")
	}

	for i, v := range allLibrary {
		if v.TrackID == library.TrackID && v.PlaylistID == library.PlaylistID {
			allLibrary = append(allLibrary[:i], allLibrary[i+1:]...)
			break
		}
	}

	err := database.SaveDB[[]*models.Library](constants.LIBRARY_PATH, allLibrary)
	if err != nil {
		return fmt.Errorf("failed to save Library: %w", err)
	}

	return nil
}