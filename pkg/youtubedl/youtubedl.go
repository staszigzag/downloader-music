package youtubedl

import (
	"context"
	"errors"
	"fmt"
	"io"
	"regexp"

	"github.com/kkdai/youtube/v2"
)

const (
	// expression = "^(http(s)?:\\/\\/)?((w){3}.)?(music\\.)?youtu(be|.be)?(\\.com)?\\/.+"
	expression = "http(?:s?):\\/\\/(?:www\\.)?youtu(?:be\\.com\\/watch\\?v=|\\.be\\/)([\\w\\-\\_]*)(&(amp;)?‌​[\\w\\?‌​=]*)?"
	itagMp3    = 140
)

var ErrValidUrlVideo = errors.New("not valid urlVideo")

type Downloader interface {
	DownloadAudio(ctx context.Context, urlVideo string) (string, io.ReadCloser, error)
}

type Youtubedl struct {
	r *regexp.Regexp
}

func NewYoutubedl() *Youtubedl {
	r, _ := regexp.Compile(expression)

	return &Youtubedl{
		r: r,
	}
}

func (d *Youtubedl) DownloadAudio(ctx context.Context, urlVideo string) (string, io.ReadCloser, error) {
	if !d.r.MatchString(urlVideo) {
		return "", nil, ErrValidUrlVideo
	}
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
