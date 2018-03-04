package bot

import (
	"github.com/pkg/errors"
	"github.com/rylio/ytdl"
)

const (
	maxResults = 1
)

func (b *Bot) search(searchText string) (string, error) {
	// Make the API call to YouTube.
	call := b.yClient.Search.List("id,snippet").
		Q(searchText).
		MaxResults(maxResults)
	response, err := call.Do()
	if err != nil {
		return "", errors.Wrap(err, "could not find videos on youtube")
	}

	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			return item.Id.VideoId, nil
		}
	}

	return "", errors.New("unknown error for youtube")
}

// GetDownloadURL Эта функция возвращает прямую ссылку на видео по ID
func GetDownloadURL(idVideo string) (string, string, error) {
	infoFromID, err := ytdl.GetVideoInfoFromID(idVideo)
	if err != nil {
		return "", "", err
	}
	bestFormats := infoFromID.Formats.Extremes(ytdl.FormatAudioBitrateKey, true)
	downloadURL, err := infoFromID.GetDownloadURL(bestFormats[0])
	return downloadURL.String(), infoFromID.Title, err
}
