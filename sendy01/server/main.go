package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"google.golang.org/grpc"
	"main.go/youtube"
)

type myYoutubeServer struct {
	youtube.YouTubeServer
}

func (s *myYoutubeServer) GetPlaylist(ctx context.Context, in *youtube.PlaylistRequest) (*youtube.PlaylistResponse, error) {
	log.Printf("Received playlist ID: %v", in.GetPlaylistId())

	res := youtube.GetVideoList(ctx, in.GetPlaylistId())

	videoList := []*youtube.Video{}
	for _, v := range res {
		videoList = append(videoList, &youtube.Video{Title: v})
	}
	return &youtube.PlaylistResponse{Videos: videoList}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Чтение содержимого файла index.html
	content, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	// Отправка содержимого как ответ на запрос
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(content)
}

func main() {
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("cannot create listener: %s", err)
		}
		serverRegistrat := grpc.NewServer()
		service := &myYoutubeServer{}

		youtube.RegisterYouTubeServer(serverRegistrat, service)
		fmt.Println("gRPC сервер запущен на http://localhost:50051")
		err = serverRegistrat.Serve(lis)
		if err != nil {
			log.Fatalf("impossible to serve: %s", err)
		}
	}()

	// Запуск HTTP сервера на порту 8080
	http.HandleFunc("/", handler)

	fmt.Println("HTTP сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
