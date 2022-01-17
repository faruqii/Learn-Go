package main

import (
	"fmt"
)

type Song struct {
	Title  string
	Artist string
	Album  string
	next   *Song
	prev   *Song
}

type Playlist struct {
	head *Song
	tail *Song
	current *Song
}

// check if the playlist is empty
func (p *Playlist) isEmpty() bool {
	return p.head == nil && p.tail == nil
}

// add a song to the playlist
func (p *Playlist) addSong(title string, artist string, album string) {

	if p.isEmpty() {
		newSong := &Song{title, artist, album, nil, nil}
		p.head = newSong
		p.tail = newSong
		p.current = newSong
	} else {
		newSong := &Song{title, artist, album, nil, nil}
		newSong.prev = p.tail
		p.tail.next = newSong
		p.tail = newSong
	}

}

// play current song
func (p *Playlist) play(){

	p.current = p.head
	if p.isEmpty() {
		fmt.Println("Playlist is empty")
	}
	fmt.Println("Now playing: ", p.current.Title, " by ", p.current.Artist, " from ", p.current.Album) 
	
}

// play next song
func (p *Playlist) nextSong() string {

	if p.isEmpty() {
		fmt.Println("Playlist is empty")
	}
	
	if p.current.next == nil {
		p.current = p.head
		return "End of Playlist"
	}

	p.current = p.current.next
	return "Now playing: " + p.current.Title + " by " + p.current.Artist + " from " + p.current.Album
}

// play previous song
func (p *Playlist) prevSong() string {
	
	if p.isEmpty() {
		fmt.Println("Playlist is empty")
	}
	
	if p.current.prev == nil {
		p.current = p.tail
		return "At the start of the Playlist"
	}

	p.current = p.current.prev
	return "Now playing: " + p.current.Title + " by " + p.current.Artist + " from " + p.current.Album
}

// print all songs in playlist
func (p *Playlist) print() {
	if p.isEmpty() {
		fmt.Println("Playlist is empty")
	}
	
	current := p.head
	for current != nil {
		fmt.Println(current.Title, " by ", current.Artist, " from ", current.Album)
		current = current.next
	}
}

// remove a song from the playlist
func (p *Playlist) removeSong(title string) {
	if p.isEmpty() {
		fmt.Println("Playlist is empty")
	}
	
	current := p.head
	for current != nil {
		if current.Title == title {
			if current == p.head {
				p.head = current.next
				p.head.prev = nil
				return
			} else if current == p.tail {
				p.tail = current.prev
				p.tail.next = nil
				return
			} else {
				current.prev.next = current.next
				current.next.prev = current.prev
				return
			}
		}
		current = current.next
	}
	fmt.Println("Song not found")
}

// main function
func main() {
	playlist := &Playlist{}
	fmt.Println("Welcome to the Go Playlist")
	fmt.Println("==========================")
	fmt.Println("1. Add a song")
	fmt.Println("2. Play current song")
	fmt.Println("3. Play next song")
	fmt.Println("4. Play previous song")
	fmt.Println("5. Print playlist")
	fmt.Println("6. Remove song")
	fmt.Println("7. Exit")
	fmt.Println("==========================")
	fmt.Print("Enter your choice: ")
	var choice int
	fmt.Scanf("%d", &choice)
	for {
		if choice == 1 {
			title := ""
			artist := ""
			album := ""
			fmt.Print("Enter song title: ")
			fmt.Scanf("%s", &title)
			fmt.Print("Enter song artist: ")
			fmt.Scanf("%s", &artist)
			fmt.Print("Enter song album: ")
			fmt.Scanf("%s", &album)
			playlist.addSong(title, artist, album)
			fmt.Println("Song added")
		} else if choice == 2 {
			playlist.play()
		} else if choice == 3 {
			fmt.Println(playlist.nextSong())
		} else if choice == 4 {
			fmt.Println(playlist.prevSong())
		} else if choice == 5 {
			playlist.print()
		} else if choice == 6 {
			title := ""
			fmt.Print("Enter song title: ")
			fmt.Scanf("%s", &title)
			playlist.removeSong(title)
		} else if choice == 7 {
			fmt.Println("Goodbye")
			break
		} else {
			fmt.Println("Invalid choice")
		}
	}
}