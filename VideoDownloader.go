// Download Youtube video using golang
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/Andreychik32/ytdl"
)

func downloadVideo(videoURL string) error {
	ctx := context.Background()
	client := ytdl.DefaultClient

	videoInfo, err := client.GetVideoInfo(ctx, videoURL)
	if err != nil {
		return err
	}

	format := videoInfo.Formats.Best(ytdl.FormatAudioEncodingKey)[0]

	// Generate download URL
	downloadURL, err := client.GetDownloadURL(ctx, videoInfo, format)
	if err != nil {
		return err
	}

	resp, err := client.HTTPClient.Get(downloadURL.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fileName := strings.Join(strings.Split(videoInfo.Title, " "), "_") + "." + format.Extension
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Video downloaded successfully to: %s\n", fileName)
	return nil
}

func main() {
	videoURL := "https://youtu.be/jOiqcDt5gAc"

	err := downloadVideo(videoURL)
	if err != nil {
		log.Fatalf("Error downloading video: %v", err)
	}
}
