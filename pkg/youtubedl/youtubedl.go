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
	expression = `(?:youtube(?:-nocookie)?\.com/(?:[^/]+/.+/|(?:v|e(?:mbed)?)/|.*[?&]v=)|youtu\.be/)([^"&?/\s]{11})`
	itagMp3    = 140
)

var (
	ErrInvalidCharactersInVideoID = errors.New("invalid characters in video id")
	ErrVideoIDMinLength           = errors.New("the video id must be at least 11 characters long")
	ErrValidUrlVideo              = errors.New("not valid urlVideo")
)

type Downloader interface {
	DownloadAudio(ctx context.Context, urlVideo string) (string, io.ReadCloser, error)
	ExtractVideoID(urlVideo string) (string, error)
}

type Youtubedl struct {
	videoIdRegexp *regexp.Regexp
}

func NewYoutubedl() *Youtubedl {
	r := regexp.MustCompile(expression)
	return &Youtubedl{
		videoIdRegexp: r,
	}
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
	if !d.videoIdRegexp.MatchString(urlVideo) {
		return "", ErrValidUrlVideo
	}

	subs := d.videoIdRegexp.FindStringSubmatch(urlVideo)
	videoID := subs[1]

	if strings.ContainsAny(videoID, "?&/<%=") {
		return "", ErrInvalidCharactersInVideoID
	}
	if len(videoID) < 11 {
		return "", ErrVideoIDMinLength
	}

	return videoID, nil
}
