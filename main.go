package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
)

var user_url string
var ytdlp_path string
var user_home string
var is_playlist bool

// Download all playlists of YouTube channel/user keeping each playlist in separate directory:
// var playlist_command string = "yt-dlp -o '%(uploader)s/%(playlist)s/%(playlist_index)s - %(title)s.%(ext)s"

func ytdlp_check() (string, bool) {
	path, err := exec.LookPath("yt-dlp")
	if err != nil {
		log.Fatal("yt-dlp was not found in your path. Is it installed?")
		return path, false
	}
	fmt.Printf("\nyt-dlp is available at %s\n", path)
	return path, true
}

func check_playlist_arguments() bool {
	is_playlist, err := regexp.MatchString("playlist", os.Args[1])
	if err != nil {
		log.Fatalf("regex for 'playlists' in provided URL failed with %s\n", err)
	}
	return is_playlist
}

func check_playlist_prompt() bool {
	is_playlist, err := regexp.MatchString("playlist", user_url)
	if err != nil {
		log.Fatalf("regex for 'playlists' in provided URL failed with %s\n", err)
	}
	return is_playlist
}

func get_input() {
	fmt.Println("Enter the URL of the YouTube Video to download:")
	fmt.Scanf("%s", &user_url)
	is_playlist = check_playlist_prompt()
	downloader_prompt()
}

func downloader_prompt() {
	user_home, _ := os.UserHomeDir()
	if is_playlist == true {
		fmt.Println("\n\n\n\n\n\nThis is a playlist. Buckleup buckaroo.")
		// Download all playlists of YouTube channel/user keeping each playlist in separate directory:
		cmd := exec.Command("yt-dlp", "-o", "%(uploader)s/%(playlist)s/%(playlist_index)s - %(title)s.%(ext)s", " "+user_url)
		cmd.Dir = user_home + "/Videos/"
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
	} else {
		fmt.Println("\n\n\n\n\n\nPLAYLIST NOT DETECTED.")
		cmd := exec.Command("yt-dlp", " "+user_url)
		cmd.Dir = user_home + "/Videos/"
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
	}

}

func downloader_arguments() {

	user_home, _ := os.UserHomeDir()
	if is_playlist == true {
		fmt.Println("\n\n\n\n\n\nThis is a playlist. Buckleup buckaroo.")
		// Download all playlists of YouTube channel/user keeping each playlist in separate directory:
		cmd := exec.Command("yt-dlp", "-o", "%(uploader)s/%(playlist)s/%(playlist_index)s - %(title)s.%(ext)s", " "+os.Args[1])
		cmd.Dir = user_home + "/Videos/"
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
	} else {
		fmt.Println("\n\n\n\n\n\nPLAYLIST NOT DETECTED.")
		cmd := exec.Command("yt-dlp", " "+os.Args[1])
		cmd.Dir = user_home + "/Videos/"
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
	}

}

func main() {
	if _, ytdlp_exists := ytdlp_check(); ytdlp_exists == true {
		if len(os.Args[1:]) == 0 {
			get_input()
		} else if len(os.Args[1:]) == 1 {
			is_playlist = check_playlist_arguments()
			downloader_arguments()
		} else {
			fmt.Println("\nToo many arguments.")
		}

	}

}
