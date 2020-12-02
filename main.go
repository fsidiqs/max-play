package main

import (
	"fmt"
	"musicplayer/music"
)

func main() {
	musicPlayer := music.MusicPlayer{}
	maxSongs := musicPlayer.MaxSongsCount([]int{300, 200, 100}, []int{400, 500, 600}, 20, 2)
	fmt.Println(maxSongs)
	maxSongs2 := musicPlayer.MaxSongsCount([]int{60, 400, 60}, []int{60, 60, 60}, 5, 1)
	fmt.Println(maxSongs2)
	maxSongs3 := musicPlayer.MaxSongsCount([]int{200, 100, 500}, []int{400, 50, 20, 20}, 5, 2)
	fmt.Println(maxSongs3)

}
