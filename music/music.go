package music

import (
	"sort"
)

type song struct {
	durationSec int
}

type songs []song

func (songs songs) RemoveSong(index int) {
	// Remove the element at index index from songs.
	copy(songs[index:], songs[index+1:]) // Shift songs[index+1:] left one index.
	songs[len(songs)-1] = song{}         // Erase last element (write zero value).
	songs = songs[:len(songs)-1]         // Truncate slice.	fmt.Println(songs)

}

func (songs songs) GetSongsStartFrom(index int) songs {
	return songs[index:]
}

func NewSongs(durationSecs []int) songs {
	songs := make([]song, len(durationSecs))
	for idx, durationSec := range durationSecs {
		songs[idx] = song{
			durationSec: durationSec,
		}
	}
	return songs
}

type album struct {
	songs songs
}

func NewAlbum(songs []song) album {
	return album{
		songs: songs,
	}
}

func NewMixedAlbum(songsMultiDimension ...songs) album {
	newsongs := []song{}
	for _, songs := range songsMultiDimension {
		for _, song := range songs {
			newsongs = append(
				newsongs,
				song,
			)
		}
	}

	return album{
		songs: newsongs,
	}
}

func (a album) GetSongs() songs {
	return a.songs
}

type merge2Album []album

func NewMerge2Album(album1 album, album2 album) merge2Album {
	return merge2Album{
		album1,
		album2,
	}
}

func (songs album) SortAlbumAsc() {
	sort.SliceStable(songs.songs, func(index, j int) bool {
		return songs.songs[index].durationSec < songs.songs[j].durationSec
	})
}
