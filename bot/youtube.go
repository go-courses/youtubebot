package bot

import (
	"flag"
	"log"
	"net/http"
	"net/url"

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

func GetDownloadUrl(idVideo string) (*url.URL, error) {

	infoFromId, _ := ytdl.GetVideoInfoFromID(idVideo)

	foundFormat := func(formats ytdl.FormatList) ytdl.Format {
		var foundFormat ytdl.Format
		bestFormats := formats.Filter(ytdl.FormatResolutionKey, []interface{}{"360p", "720p"}).Filter(ytdl.FormatExtensionKey, []interface{}{"mp4"}).Filter(ytdl.FormatAudioEncodingKey, []interface{}{"aac"}).Extremes(ytdl.FormatResolutionKey, true).Extremes(ytdl.FormatAudioBitrateKey, true)
		for _, format := range bestFormats {
			if format.Extension == "mp4" {
				foundFormat = format
			}
		}
		return foundFormat
	}
	downloadUrl, err := infoFromId.GetDownloadURL(foundFormat(infoFromId.Formats))
	return downloadUrl, err
}
