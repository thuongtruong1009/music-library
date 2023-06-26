package albums

import (
	"fmt"
	"music-management/database"
	"music-management/models"
	"music-management/pkg/helpers"
	"music-management/pkg/constants"
)

type AlbumRepository struct {
	helper helpers.Helper
}

func NewAlbumRepository(helper helpers.Helper) *AlbumRepository {
	return &AlbumRepository {
		helper: helper,
	}
}

func (r *AlbumRepository) GetAlbums() ([]*models.Album, error) {
	allAlbums, err := database.ReadDB[*models.Album](constants.ALBUM_PATH)
	if err != nil {
		return nil, err
	}

	if allAlbums == nil {
		return nil, nil
	}

	return allAlbums, nil
}

func (r *AlbumRepository) GetAlbum(albumID string) (*models.Album, error) {
	allAlbums, _ := r.GetAlbums()
	
	if allAlbums == nil {
		return nil, nil
	}

	for _, v := range allAlbums {
		if v.ID == albumID {
			return v, nil
		}
	}

	return nil, nil
}

func (r *AlbumRepository) CreateAlbum(newAlbums *models.Album) (*models.Album, error) {
	allAlbums, _ := r.GetAlbums()

	var albumInit []*models.Album

	if allAlbums == nil {
		albumInit = make([]*models.Album, 0)
	} else {
		albumInit = make([]*models.Album, len(allAlbums))
		copy(albumInit, allAlbums)
	}

	albumInit = append(albumInit, newAlbums)

	err2 := database.SaveDB[[]*models.Album](constants.ALBUM_PATH, albumInit)
	if err2 != nil {
		return nil, err2
	}

	return newAlbums, nil
}

func (r *AlbumRepository) DeleteAlbum(albumID string) error {
	allAlbums, _ := r.GetAlbums()

	if allAlbums == nil {
		return fmt.Errorf("Album not found")
	}

	for i, v := range allAlbums {
		if v.ID == albumID {
			allAlbums = append(allAlbums[:i], allAlbums[i+1:]...)
			break
		}
	}

	err := database.SaveDB[[]*models.Album](constants.ALBUM_PATH, allAlbums)
	if err != nil {
		return fmt.Errorf("Delete album failed")
	}

	return nil
}

func (r *AlbumRepository) UpdateAlbum(albumUpdate *models.Album) (*models.Album, error) {
	allAlbums, _ := r.GetAlbums()

	if allAlbums == nil {
		return nil, fmt.Errorf(constants.UPDATE_FAILED)
	}

	var albumInit []*models.Album

	for i, v := range allAlbums {
		if v.ID == albumUpdate.ID {
			albumInit = append(albumInit[:i], albumUpdate)
			break
		}
	}

	albumInit = append(albumInit, allAlbums[len(albumInit):]...)

	err := database.SaveDB[[]*models.Album](constants.ALBUM_PATH, albumInit)
	if err != nil {
		return nil, fmt.Errorf(constants.UPDATE_FAILED)
	}

	return albumUpdate, nil
}
