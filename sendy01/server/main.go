package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"main.go/youtube"
)

type myYoutubeServer struct {
	youtube.YouTubeServer
}

func (s *myYoutubeServer) GetPlaylist(ctx context.Context, in *youtube.PlaylistRequest) (*youtube.PlaylistResponse, error) {
	// Реализовать обращение к YouTube API и возвращение данных
	// Пример запроса к YouTube API: https://developers.google.com/youtube/v3/docs/playlistItems/list
	log.Printf("Received playlist ID: %v", in.GetPlaylistId())
	// Добавьте код для обращения к YouTube API здесь
	// ...

	res := youtube.GetVideoList(ctx, in.GetPlaylistId())

	videoList := []*youtube.Video{}
	for _, v := range res {
		videoList = append(videoList, &youtube.Video{Title: v})
	}
	return &youtube.PlaylistResponse{Videos: videoList}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}
	serverRegistrat := grpc.NewServer()
	service := &myYoutubeServer{}

	youtube.RegisterYouTubeServer(serverRegistrat, service)
	err = serverRegistrat.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
