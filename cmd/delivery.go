package cmd

import (
	"music-management-system/pkg/helpers"
	"music-management-system/pkg/constants"

	"music-management-system/internal/tracks"
	"music-management-system/internal/playlists"
	"music-management-system/internal/albums"
	"music-management-system/internal/artists"
	"music-management-system/internal/genres"
)

type Delivery struct {
	albumHandler albums.AlbumHandler
	artistHandler artists.ArtistHandler
	genreHandler genres.GenreHandler
	playlistHandler playlists.PlaylistHandler
	trackHandler tracks.TrackHandler
	helper helpers.Helper
}

func NewDelivery(albumHandler albums.AlbumHandler, artistHandler artists.ArtistHandler, genreHandler genres.GenreHandler, playlistHandler playlists.PlaylistHandler, trackHandler tracks.TrackHandler, helper helpers.Helper) *Delivery {
	return &Delivery{
		albumHandler: albumHandler,
		artistHandler: artistHandler,
		genreHandler: genreHandler,
		playlistHandler: playlistHandler,
		trackHandler: trackHandler,
		helper: helper,
	}
}

func (h *Delivery) HandleOption(option int8) {
	optionHandlers := map[int8]func(){
		1:  h.albumHandler.CreateAlbum,
		2:  h.albumHandler.GetAlbums,
		3:  h.albumHandler.GetAlbum,
		4:  h.albumHandler.DeleteAlbum,
		5:  h.albumHandler.UpdateAlbum,
		6:  h.albumHandler.GetTracksOfAlbum,
		7:  h.artistHandler.CreateArtist,
		8:  h.artistHandler.GetArtists,
		9:  h.artistHandler.GetArtist,
		10: h.artistHandler.GetAlbumsOfArtist,
		11: h.artistHandler.GetTracksOfArtist,
		12: h.artistHandler.DeleteArtist,
		13: h.artistHandler.UpdateArtist,
		14: h.genreHandler.CreateGenre,
		15: h.genreHandler.GetGenres,
		16: h.genreHandler.GetGenre,
		17: h.genreHandler.DeleteGenre,
		18: h.genreHandler.UpdateGenre,
		19: h.genreHandler.GetTracksOfGenre,
		20: h.trackHandler.CreateTrack,
		21: h.trackHandler.GetTracks,
		22: h.trackHandler.GetTrack,
		23: h.trackHandler.DeleteTrack,
		24: h.trackHandler.UpdateTrack,
		25: h.playlistHandler.CreatePlaylist,
		26: h.playlistHandler.GetPlaylists,
		27: h.playlistHandler.GetPlaylist,
		28: h.playlistHandler.DeletePlaylist,
		29: h.playlistHandler.UpdatePlaylist,
		30: h.playlistHandler.AddTrackToPlaylist,
		31: h.playlistHandler.DeleteTrackFromPlaylist,
		32: h.playlistHandler.GetTracksOfPlaylist,
		33: h.playlistHandler.GetPlaylistsHaveTrack,
	}

	handler, exists := optionHandlers[option]
	if exists {
		handler()
	} else {
		h.helper.OutputNomal(constants.ERROR, "Invalid option")
	}
}
