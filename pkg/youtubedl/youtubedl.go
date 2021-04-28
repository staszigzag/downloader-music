package youtubedl

import (
	"context"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/kkdai/youtube/v2"
)

const (
	// expression = "^(http(s)?:\\/\\/)?((w){3}.)?(music\\.)?youtu(be|.be)?(\\.com)?\\/.+"
	expression = "http(?:s?):\\/\\/(?:www\\.)?youtu(?:be\\.com\\/watch\\?v=|\\.be\\/)([\\w\\-\\_]*)(&(amp;)?‌​[\\w\\?‌​=]*)?"
	itagMp3    = 140
)

var (
	ErrInvalidCharactersInVideoID = errors.New("invalid characters in video id")
	ErrVideoIDMinLength           = errors.New("the video id must be at least 10 characters long")
	ErrValidUrlVideo              = errors.New("not valid urlVideo")
)

var urlValidRegexp = regexp.MustCompile(expression)

var videoIdRegexpList = []*regexp.Regexp{
	regexp.MustCompile(`(?:v|embed|watch\?v)(?:=|/)([^"&?/=%]{11})`),
	regexp.MustCompile(`(?:=|/)([^"&?/=%]{11})`),
	regexp.MustCompile(`([^"&?/=%]{11})`),
}

type Downloader interface {
	DownloadAudio(ctx context.Context, urlVideo string) (string, io.ReadCloser, error)
	ExtractVideoID(urlVideo string) (string, error)
}

type Youtubedl struct{}

func NewYoutubedl() *Youtubedl {
	// r, _ := regexp.Compile(expression)
	return &Youtubedl{}
}

func (d *Youtubedl) DownloadAudio(ctx context.Context, urlVideo string) (string, io.ReadCloser, error) {
	client := youtube.Client{}
	video, err := client.GetVideo(urlVideo)
	if err != nil {
		return "", nil, fmt.Errorf("error getting video info: %v", err)
	}

	resp, err := client.GetStream(video, video.Formats.FindByItag(itagMp3))
	if err != nil {
		return "", nil, fmt.Errorf("error getting resp: %v", err)
	}
	// defer resp.Body.Close()

	filename := video.Title + ".mp3"
	return filename, resp.Body, nil
}

func (d *Youtubedl) ExtractVideoID(urlVideo string) (string, error) {
	if !urlValidRegexp.MatchString(urlVideo) {
		return "", ErrValidUrlVideo
	}

	var videoID string
	for _, re := range videoIdRegexpList {
		if isMatch := re.MatchString(urlVideo); isMatch {
			subs := re.FindStringSubmatch(urlVideo)
			videoID = subs[1]
			break
		}
	}

	if strings.ContainsAny(videoID, "?&/<%=") {
		return "", ErrInvalidCharactersInVideoID
	}
	if len(videoID) < 10 {
		return "", ErrVideoIDMinLength
	}

	return videoID, nil
}
