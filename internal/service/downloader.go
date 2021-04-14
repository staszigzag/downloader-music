package service

import (
	"context"

	"github.com/staszigzag/downloader-music/pkg/youtubedl"
)

type DownloaderService struct {
	youtubedl youtubedl.Downloader
}

func NewDownloaderService(dl youtubedl.Downloader) *DownloaderService {
	return &DownloaderService{youtubedl: dl}
}

func (d *DownloaderService) Download(ctx context.Context, url string) (string, error) {
	// something
	return d.youtubedl.DownloadAudio(ctx, url)
}
