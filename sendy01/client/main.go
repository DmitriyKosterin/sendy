package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"main.go/youtube"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatalf("error")
	}

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := youtube.NewYouTubeClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetPlaylist(ctx, &youtube.PlaylistRequest{PlaylistId: flag.Arg(0)})
	if err != nil {
		log.Fatalf("could not get playlist: %v", err)
	}

	// Выведите полученные видео
	for _, video := range r.GetVideos() {
		log.Printf("Video ID: %v, Title: %v", video.GetVideoId(), video.GetTitle())
	}
}
