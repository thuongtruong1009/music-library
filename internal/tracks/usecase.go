package tracks

import (
	"fmt"
	"music-management-system/models"
	"music-management-system/pkg/constants"
	"music-management-system/pkg/helpers"
)

type TrackUsecase struct {
	repo TrackRepository
	helper helpers.Helper
}

func NewTrackUsecase(repo TrackRepository, helper helpers.Helper) *TrackUsecase {
	return &TrackUsecase{
		repo: repo,
		helper: helper,
	}
}

func (s *TrackUsecase) GetTracks() ([]string, string) {
	result, err := helpers.QueryTimeTwoOutput[[]*models.Track](s.repo.GetTracks)()

	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	var output []string
	for _, v := range result {
		output = append(output, fmt.Sprintf("ID: %s, Year: %d, Title: %s, Duration: %.f, GenreID: %s, ArtistID: %s", v.ID, v.Year, v.Title, v.Duration, v.GenreID, v.ArtistID))
	}

	return output, ""
}

func (s *TrackUsecase) GetTrack() ([]string, string) {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	result, err := helpers.QueryTimeTwoOutputWithParams[*models.Track, string](s.repo.GetTrack)(id)
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	output := []string{fmt.Sprintf("ID: %s, Year: %d, Title: %s, Duration: %.f, GenreID: %s, ArtistID: %s", result.ID, result.Year, result.Title, result.Duration, result.GenreID, result.ArtistID)}

	return output, ""
}

func (s *TrackUsecase) CreateTrack() ([]string, string) {
	var title string
	fmt.Print("» Enter Title: ")
	fmt.Scanln(&title)

	var year int
	fmt.Print("» Enter Year: ")
	fmt.Scanln(&year)

	var duration float32
	fmt.Print("» Enter Duration: ")
	fmt.Scanln(&duration)

	var genreID string
	fmt.Print("» Enter GenreID: ")
	fmt.Scanln(&genreID)

	var artistID string
	fmt.Print("» Enter ArtistID: ")
	fmt.Scanln(&artistID)

	newTrack := &models.Track{
		ID: s.helper.GenerateID(),
		Title: title,
		Year: year,
		Duration: duration,
		GenreID: genreID,
		ArtistID: artistID,
	}

	result, err := helpers.QueryTimeTwoOutputWithParams[*models.Track, *models.Track](s.repo.CreateTrack)(newTrack)
	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	output := []string{fmt.Sprintf("ID: %s, Year: %d, Title: %s, Duration: %.f, GenreID: %s, ArtistID: %s", result.ID, result.Year, result.Title, result.Duration, result.GenreID, result.ArtistID)}

	return output, ""
}

func (s *TrackUsecase) DeleteTrack() error {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	err := helpers.QueryTimeOneOutputWithParams[error](s.repo.DeleteTrack)(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *TrackUsecase) UpdateTrack() ([]string, string){
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	var title string
	fmt.Print("» Enter Title: ")
	fmt.Scanln(&title)

	var year int
	fmt.Print("» Enter Year: ")
	fmt.Scanln(&year)

	var duration float32
	fmt.Print("» Enter Duration: ")
	fmt.Scanln(&duration)

	var genreID string
	fmt.Print("» Enter GenreID: ")
	fmt.Scanln(&genreID)
	
	var artistID string
	fmt.Print("» Enter ArtistID: ")
	fmt.Scanln(&artistID)

	newTrack := &models.Track{
		ID: id,
		Title: title,
		Year: year,
		Duration: duration,
		GenreID: genreID,
		ArtistID: artistID,
	}

	result, err := helpers.QueryTimeTwoOutputWithParams[*models.Track, *models.Track](s.repo.UpdateTrack)(newTrack)

	if err != nil {
		return nil, err.Error()
	}

	if result == nil {
		return nil, constants.NOT_FOUND_DATA
	}

	output := []string{fmt.Sprintf("ID: %s, Year: %d, Title: %s, Duration: %.f, GenreID: %s, ArtistID: %s", result.ID, result.Year, result.Title, result.Duration, result.GenreID, result.ArtistID)}

	return output, ""
}