package cmd

import (
	"music-management-system/pkg/helpers"
	"music-management-system/pkg/constants"

	"music-management-system/internal/tracks"
	"music-management-system/internal/playlists"
	"music-management-system/internal/albums"
	"music-management-system/internal/artists"
	"music-management-system/internal/genres"
	"music-management-system/internal/libraries"
	"music-management-system/internal/users"
)

type Delivery struct {
	albumHandler albums.AlbumHandler
	artistHandler artists.ArtistHandler
	genreHandler genres.GenreHandler
	playlistHandler playlists.PlaylistHandler
	trackHandler tracks.TrackHandler
	libraryHandler libraries.LibraryHandler
	userHandler users.UserHandler
	helper helpers.Helper
}

func NewDelivery(albumHandler albums.AlbumHandler, artistHandler artists.ArtistHandler, genreHandler genres.GenreHandler, playlistHandler playlists.PlaylistHandler, trackHandler tracks.TrackHandler, libraryHandler libraries.LibraryHandler, userHandler users.UserHandler, helper helpers.Helper) *Delivery {
	return &Delivery{
		albumHandler: albumHandler,
		artistHandler: artistHandler,
		genreHandler: genreHandler,
		playlistHandler: playlistHandler,
		trackHandler: trackHandler,
		libraryHandler: libraryHandler,
		userHandler: userHandler,
		helper: helper,
	}
}

func (d *Delivery) HandleOption(option int8) {
	optionHandlers := map[int8]func(){
		1:  d.albumHandler.CreateAlbum,
		2:  d.albumHandler.GetAlbums,
		3:  d.albumHandler.GetAlbum,
		4:  d.albumHandler.DeleteAlbum,
		5:  d.albumHandler.UpdateAlbum,
		6:  d.albumHandler.GetTracksOfAlbum,

		7:  d.artistHandler.CreateArtist,
		8:  d.artistHandler.GetArtists,
		9:  d.artistHandler.GetArtist,
		10: d.artistHandler.GetAlbumsOfArtist,
		11: d.artistHandler.GetTracksOfArtist,
		12: d.artistHandler.DeleteArtist,
		13: d.artistHandler.UpdateArtist,

		14: d.genreHandler.CreateGenre,
		15: d.genreHandler.GetGenres,
		16: d.genreHandler.GetGenre,
		17: d.genreHandler.DeleteGenre,
		18: d.genreHandler.UpdateGenre,
		19: d.genreHandler.GetTracksOfGenre,

		20: d.trackHandler.CreateTrack,
		21: d.trackHandler.GetTracks,
		22: d.trackHandler.GetTrack,
		23: d.trackHandler.DeleteTrack,
		24: d.trackHandler.UpdateTrack,

		25: d.playlistHandler.CreatePlaylist,
		26: d.playlistHandler.GetPlaylists,
		27: d.playlistHandler.GetPlaylist,
		28: d.playlistHandler.DeletePlaylist,
		29: d.playlistHandler.UpdatePlaylist,

		30: d.libraryHandler.AddTrackToPlaylist,
		31: d.libraryHandler.GetTracksOfPlaylist,
		32: d.libraryHandler.DeleteTrackFromPlaylist,
		33: d.libraryHandler.GetPlaylistsContainTrack,

		34: d.userHandler.CreateUser,
		35: d.userHandler.GetUsers,
		36: d.userHandler.GetUser,
		37: d.userHandler.DeleteUser,
		38: d.userHandler.UpdateUser,
	}

	handler, exists := optionHandlers[option]
	if exists {
		handler()
	} else {
		d.helper.OutputNomal(constants.ERROR, "Invalid option")
	}
}
