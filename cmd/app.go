package cmd

import (
	"music-management-system/pkg/helpers"
	"music-management-system/internal/tracks"
	"music-management-system/internal/playlists"
	"music-management-system/internal/artists"
	"music-management-system/internal/genres"
	"music-management-system/internal/albums"
	"music-management-system/internal/libraries"
	"music-management-system/internal/users"
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

	trackRepo := tracks.NewTrackRepository(*genreRepo, *artistRepo, *helper)
	trackUC := tracks.NewTrackUsecase(*trackRepo, *helper)
	trackHandler := tracks.NewTrackHandler(*trackUC, *helper)

	userRepo := users.NewUserRepository(*helper)
	userUC := users.NewUserUsecase(*userRepo, *helper)
	userHandler := users.NewUserHandler(*userUC, *helper)
	
	libraryRepo := libraries.NewLibraryRepository(*trackRepo, *userRepo, *helper)
	libraryUC := libraries.NewLibraryUsecase(*libraryRepo, *helper)
	libraryHandler := libraries.NewLibraryHandler(*libraryUC, *helper)


	exe := NewDelivery(*albumHandler, *artistHandler, *genreHandler, *playlistHandler, *trackHandler, *libraryHandler, *userHandler, *helper)
	
	exe.Execution()
}