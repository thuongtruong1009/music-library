package artists

import (
	"fmt"
	"music-management-system/models"
	"music-management-system/database"
	"music-management-system/pkg/constants"
	"music-management-system/pkg/helpers"
)

type ArtistRepository struct {
	helper helpers.Helper
}

func NewArtistRepository(helper helpers.Helper) *ArtistRepository {
	return &ArtistRepository{
		helper: helper,
	}
}

func (a *ArtistRepository) GetArtists() ([]*models.Artist, error) {
	allArtist, err := database.ReadDB[*models.Artist](constants.ARTIST_PATH)
	if err != nil {
		return nil, err
	}

	if allArtist == nil {
		return nil, nil
	}

	return allArtist, nil
}

func (a *ArtistRepository) GetArtist(artistID string) (*models.Artist, error) {
	allArtist, _ := a.GetArtists()

	if allArtist == nil {
		return nil, nil
	}

	for _, v := range allArtist {
		if v.ID == artistID {
			return v, nil
		}
	}

	return nil, nil
}

func (a *ArtistRepository) CreateArtist(newArtist *models.Artist) (*models.Artist, error) {
	allGenres, _ := a.GetArtists()

	var artistInit []*models.Artist

	if allGenres == nil {
		artistInit = make([]*models.Artist, 0)
	} else {
		artistInit = make([]*models.Artist, len(allGenres))
		copy(artistInit, allGenres)
	}

	artistInit = append(artistInit, newArtist)

	err2 := database.SaveDB[[]*models.Artist](constants.ARTIST_PATH, artistInit)
	if err2 != nil {
		return nil, fmt.Errorf(constants.CREATE_FAILED)
	}
	return newArtist, nil
}

func (a *ArtistRepository) UpdateArtist(artlistUpdate *models.Artist) (*models.Artist, error) {
	allArtist, _ := a.GetArtists()

	if allArtist == nil {
		return nil, fmt.Errorf(constants.UPDATE_FAILED)
	}

	var artistInit []*models.Artist

	for i, v := range allArtist {
		if v.ID == artlistUpdate.ID {
			artistInit = append(allArtist[:i], artlistUpdate)
			break
		}
	}

	artistInit = append(artistInit, allArtist[len(artistInit):]...)

	err2 := database.SaveDB[[]*models.Artist](constants.ARTIST_PATH, allArtist)
	if err2 != nil {
		return nil, fmt.Errorf(constants.UPDATE_FAILED)
	}
	return artlistUpdate, nil
}

func (a *ArtistRepository) DeleteArtist(artlistID string) error {
	allArtist, _ := a.GetArtists()

	if allArtist == nil {
		return fmt.Errorf(constants.DELETE_FAILED)
	}

	for i, v := range allArtist {
		if v.ID == artlistID {
			allArtist = append(allArtist[:i], allArtist[i+1:]...)
			break
		}
	}

	err2 := database.SaveDB[[]*models.Artist](constants.ARTIST_PATH, allArtist)
	if err2 != nil {
		return fmt.Errorf(constants.DELETE_FAILED)
	}
	return nil
}
