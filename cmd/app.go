package cmd

import (
	"music-management/pkg/helpers"
	"music-management/internal/tracks"
	"music-management/internal/playlists"
	"music-management/internal/artists"
	"music-management/internal/genres"
	"music-management/internal/albums"
)

func App() {
	helper := helpers.NewHelper()

	albumRepo := albums.NewAlbumRepository(*helper)
	albumUC := albums.NewAlbumUsecase(*albumRepo, *helper)
	albumHandler := albums.NewAlbumHandler(*albumUC, *helper)

	artistRepo := artists.NewArtistRepository(*helper)
	artistUC := artists.NewArtistUsecase(*artistRepo, *helper)
	artistHandler := artists.NewArtistHandler(*artistUC, *helper)

	genreRepo := genres.NewGenreRepository(*helper)
	genreUC := genres.NewGenreUsecase(*genreRepo, *helper)
	genreHandler := genres.NewGenreHandler(*genreUC, *helper)

	playlistRepo := playlists.NewPlaylistRepository(*helper)
	playlistUC := playlists.NewPlaylistUsecase(*playlistRepo, *helper)
	playlistHandler := playlists.NewPlaylistHandler(*playlistUC, *helper)

	trackRepo := tracks.NewTrackRepository(*helper)
	trackUC := tracks.NewTrackUsecase(*trackRepo, *helper)
	trackHandler := tracks.NewTrackHandler(*trackUC, *helper)

	exe := NewDelivery(*albumHandler, *artistHandler, *genreHandler, *playlistHandler, *trackHandler, *helper)
	
	exe.Execution()
}