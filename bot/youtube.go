package bot

import (
	"flag"
	"log"
	"net/http"

	"github.com/rylio/ytdl"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var (
	query      = flag.String("query", "Google", "Search term")
	maxResults = flag.Int64("max-results", 1, "Max YouTube results")
)

// Search эта функция возвращает id видеофайла
// найденного в ютубе (первого совпавщего)
func Search(searchText string) string {

	*query = searchText

	flag.Parse()

	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Make the API call to YouTube.
	call := service.Search.List("id,snippet").
		Q(*query).
		MaxResults(*maxResults)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making search API call: %v", err)
	}

	// Group video results in separate lists.
	videos := make(map[string]string)

	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title
		}
	}

	keys := make([]string, 0, len(videos))
	for k := range videos {
		keys = append(keys, k)
	}

	return keys[0]
}

/* Эта функция возвращает прямую ссылку на видео по ID */

func GetDownloadUrl(idVideo string) (string, string, error) {

	infoFromId, err := ytdl.GetVideoInfoFromID(idVideo)
	if err != nil {
		return "", "", err
	}

	bestFormats := infoFromId.Formats.Extremes(ytdl.FormatAudioBitrateKey, true)

	downloadUrl, err := infoFromId.GetDownloadURL(bestFormats[0])
	if err != nil {
		return "", "", err
	}
	return downloadUrl.String(), infoFromId.Title, err
}
