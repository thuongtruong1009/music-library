package tracks

import (
	"fmt"
	"music-management/models"
	"music-management/database"
	"music-management/pkg/constants"
	"music-management/pkg/helpers"
)

type TrackRepository struct {
	helper helpers.Helper
}

func NewTrackRepository(helper helpers.Helper) *TrackRepository {
	return &TrackRepository{
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
	allTrack, _ := t.GetTracks()

	var trackInit []*models.Track

	if allTrack == nil {
		trackInit = make([]*models.Track, 0)
	} else {
		trackInit = make([]*models.Track, len(allTrack))
		copy(trackInit, allTrack)
	}

	trackInit = append(trackInit, newTrack)

	err := database.SaveDB[[]*models.Track](constants.TRACK_PATH, trackInit)
	if err != nil {
		return nil, err
	}

	return newTrack, nil
}

func (t *TrackRepository) UpdateTrack(updateTrack *models.Track) (*models.Track, error) {
	allTrack, _ := t.GetTracks()

	if allTrack == nil {
		return nil, fmt.Errorf(constants.NOT_FOUND_DATA)
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
	allTrack, _ := t.GetTracks()

	if allTrack == nil {
		return fmt.Errorf(constants.NOT_FOUND_DATA)
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
