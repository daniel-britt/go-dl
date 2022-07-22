package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var user_url string
var ytdlp_path string
var user_home string

func ytdlp_check() (string, bool) {
	path, err := exec.LookPath("yt-dlp")
	if err != nil {
		log.Fatal("yt-dlp was not found in your path. Is it installed?")
		return path, false
	}
	fmt.Printf("yt-dlp is available at %s\n", path)
	return path, true
}

func get_input() {
	fmt.Println("Enter the URL of the YouTube Video to download:")
	fmt.Scanf("%s", &user_url)
	downloader()
}

func downloader() {
	user_home, _ := os.UserHomeDir()
	cmd := exec.Command("yt-dlp", " "+user_url)
	cmd.Dir = user_home + "/Videos/"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func main() {
	if _, ytdlp_exists := ytdlp_check(); ytdlp_exists == true {
		get_input()
	}

}
