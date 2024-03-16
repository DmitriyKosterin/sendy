package youtube

import (
	"context"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func GetVideoList(ctx context.Context, playlistID string) []string {
	const apiKey = "API_KEY"
	//playlistID := "PL7TobRzJvhX0kcjEaPIJ66X5TluE31t1Z"

	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Unable to create YouTube service: %v", err)
	}

	st := []string{"snippet"}
	playlistItems := service.PlaylistItems.List(st).PlaylistId(playlistID).MaxResults(50)
	response, err := playlistItems.Do()
	if err != nil {
		log.Fatalf("Unable to retrieve playlist items: %v", err)
	}

	res := []string{}
	for _, playlistItem := range response.Items {
		title := playlistItem.Snippet.Title
		res = append(res, title)
	}
	return res
}
