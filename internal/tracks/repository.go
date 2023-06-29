package tracks

import (
	"fmt"
	"music-management-system/models"
	"music-management-system/database"
	"music-management-system/pkg/constants"
	"music-management-system/pkg/helpers"
	"music-management-system/internal/genres"
	"music-management-system/internal/artists"
)

type TrackRepository struct {
	genreRepo genres.GenreRepository
	artistRepo artists.ArtistRepository
	helper helpers.Helper
}

func NewTrackRepository(genreRepo genres.GenreRepository, artistRepo artists.ArtistRepository, helper helpers.Helper) *TrackRepository {
	return &TrackRepository{
		genreRepo: genreRepo,
		artistRepo: artistRepo,
		helper: helper,
	}
}

func (t *TrackRepository) GetTracks() ([]*models.Track, error) {
	allTrack, err := database.ReadDB[*models.Track](constants.TRACK_PATH)
	if err != nil {
		return nil, err
	}

	if allTrack == nil {
		return nil, nil
	}

	return allTrack, nil
}

func (t *TrackRepository) GetTrack(trackID string) (*models.Track, error) {
	allTrack, _ := t.GetTracks()

	if allTrack == nil {
		return nil, nil
	}

	for _, v := range allTrack {
		if v.ID == trackID {
			return v, nil
		}
	}

	return nil, nil
}

func (t *TrackRepository) CreateTrack(newTrack *models.Track) (*models.Track, error) {
	genreExist, err1 := t.genreRepo.GetGenre(newTrack.GenreID)
	if err1 != nil {
		return nil, err1
	}

	if genreExist == nil {
		return nil, fmt.Errorf(constants.NOT_FOUND_GENRE)
	}

	existArtist, err2 := t.artistRepo.GetArtist(newTrack.ArtistID)
	if err2 != nil {
		return nil, err2
	}

	if existArtist == nil {
		return nil, fmt.Errorf(constants.NOT_FOUND_ARTIST)
	}

	allTrack, _ := t.GetTracks()

	var trackInit []*models.Track

	if allTrack == nil {
		trackInit = make([]*models.Track, 0)
	} else {
		trackInit = make([]*models.Track, len(allTrack))
		copy(trackInit, allTrack)
	}

	trackInit = append(trackInit, newTrack)

	err3 := database.SaveDB[[]*models.Track](constants.TRACK_PATH, trackInit)
	if err3 != nil {
		return nil, err3
	}

	return newTrack, nil
}

func (t *TrackRepository) UpdateTrack(updateTrack *models.Track) (*models.Track, error) {
	trackExist, _ := t.GetTrack(updateTrack.ID)

	if trackExist == nil {
		return nil, fmt.Errorf(constants.NOT_FOUND_TRACK)
	}

	allTrack, _ := t.GetTracks()
	if allTrack == nil {
		return nil, fmt.Errorf(constants.NOT_FOUND_TRACK)
	}

	var trackInit []*models.Track

	for i, v := range allTrack {
		if v.ID == updateTrack.ID {
			v.Title = updateTrack.Title
			v.Duration = updateTrack.Duration
			v.Year = updateTrack.Year
			v.ArtistID = updateTrack.ArtistID
			v.GenreID = updateTrack.GenreID
		}

		trackInit = append(trackInit[:i], updateTrack)
		break
	}

	trackInit = append(trackInit, allTrack[len(trackInit):]...)

	err := database.SaveDB[[]*models.Track](constants.TRACK_PATH, trackInit)
	if err != nil {
		return nil, fmt.Errorf(constants.UPDATE_FAILED)
	}

	return updateTrack, nil
}

func (t *TrackRepository) DeleteTrack(trackID string) error {
	trackExist, _ := t.GetTrack(trackID)

	if trackExist == nil {
		return fmt.Errorf(constants.NOT_FOUND_TRACK)
	}

	allTrack, _ := t.GetTracks()

	if allTrack == nil {
		return fmt.Errorf(constants.NOT_FOUND_TRACK)
	}

	for i, v := range allTrack {
		if v.ID == trackID {
			allTrack = append(allTrack[:i], allTrack[i+1:]...)
			break
		}
	}

	err := database.SaveDB[[]*models.Track](constants.TRACK_PATH, allTrack)
	if err != nil {
		return fmt.Errorf(constants.DELETE_FAILED)
	}

	return nil
}
