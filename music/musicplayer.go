package music

import "fmt"

type MusicPlayer struct{}

func (MusicPlayer) MaxSongsCount(duration1 []int, duration2 []int, minutesLim int, TSongs int) int {
	var maxSongs int = 0
	// var maxDurationCount int
	// if len(duration1) > len(duration2) {
	// 	maxDurationCount = len(duration1)
	// } else {
	// 	maxDurationCount = len(duration2)
	// }
	album1Songs := NewSongs(duration1)
	album2Songs := NewSongs(duration2)
	album1 := NewAlbum(album1Songs)
	album2 := NewAlbum(album2Songs)

	album1.SortAlbumAsc()
	album2.SortAlbumAsc()

	secondsLim := minutesLim * 60 //convert to seconds
	index := 0
	for index < TSongs {
		if index != len(duration1) {
			secondsLim = secondsLim - album1.songs[index].durationSec
			if secondsLim < 0 {
				break
			}
			maxSongs++
		} else {
			return -1
		}
		if index != len(duration2) {
			secondsLim = secondsLim - album2.songs[index].durationSec
			if secondsLim < 0 {
				break
			}
			maxSongs++
		} else {
			return -1
		}
		index++
	}

	if index != TSongs {
		return -1
	}

	mixedAlbum := NewMixedAlbum(album1.songs.GetSongsStartFrom(index), album2.songs.GetSongsStartFrom(index))
	mixedAlbum.SortAlbumAsc()
	fmt.Printf("mixed %+v \n", mixedAlbum)
	for _, song := range mixedAlbum.songs {
		secondsLim = secondsLim - song.durationSec
		if secondsLim < 0 {
			break
		}
		maxSongs++
	}
	return maxSongs
}
