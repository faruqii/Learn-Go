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
	idx := 1
	for current != nil {
		fmt.Println("Song ", idx, ": ", current.Title, " by ", current.Artist, " from ", current.Album)
		idx++
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
	playlist.addSong("The Sign", "Ace of Base", "The Sign")
	playlist.addSong("Just the way you are", "Bruno Mars", "24K Magic")
	playlist.addSong("I'm yours", "Bruno Mars", "24K Magic")
	playlist.print()
	playlist.play()
	fmt.Println(playlist.nextSong())
	fmt.Println(playlist.nextSong())
	fmt.Println(playlist.prevSong())
	playlist.removeSong("The Sign")
	playlist.print()

}