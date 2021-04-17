package service

import (
	"context"

	"github.com/staszigzag/downloader-music/internal/repository"

	"github.com/staszigzag/downloader-music/pkg/youtubedl"
)

type DownloaderService struct {
	youtubedl youtubedl.Downloader
	audioRepo repository.Audio
}

func NewDownloaderService(dl youtubedl.Downloader, repo repository.Audio) *DownloaderService {
	return &DownloaderService{
		youtubedl: dl,
		audioRepo: repo,
	}
}

func (d *DownloaderService) Download(ctx context.Context, url string) (string, error) {
	name, file, err := d.youtubedl.DownloadAudio(ctx, url)
	if err != nil {
		return "", err
	}
	filepath, err := d.audioRepo.CreateAudio(name, file)
	if err != nil {
		return "", err
	}
	return filepath, nil
}
