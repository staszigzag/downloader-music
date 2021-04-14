package youtubedl

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
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
	DownloadAudio(ctx context.Context, url string) (string, error)
}

type Youtubedl struct {
	path string
	r    *regexp.Regexp
}

func NewDownloader(path string) *Youtubedl {
	r, _ := regexp.Compile(expression)

	return &Youtubedl{
		path: path,
		r:    r,
	}
}

func (d *Youtubedl) DownloadAudio(ctx context.Context, urlVideo string) (string, error) {
	if !d.r.MatchString(urlVideo) {
		return "", ErrValidUrlVideo
	}
	client := youtube.Client{}
	video, err := client.GetVideo(urlVideo)
	if err != nil {
		return "", fmt.Errorf("error getting video info: %v", err)
	}
	filename := video.Title + ".mp3"

	resp, err := client.GetStream(video, video.Formats.FindByItag(itagMp3))
	if err != nil {
		return "", fmt.Errorf("error getting resp: %v", err)
	}
	defer resp.Body.Close()

	if d.path != "" {
		err := os.MkdirAll(d.path, 0o777)
		if err != nil {
			return "", fmt.Errorf("error creating dir : %v", err)
		}
	}
	filepath := d.path + "/" + filename
	file, err := os.Create(filepath)
	if err != nil {
		return "", fmt.Errorf("error сreating file %s: %v", filename, err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		os.Remove(filename)
		return "", fmt.Errorf("error downloading and write in file %s: %v", filename, err)
	}

	return filepath, nil
}
