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
	videoId, err := d.youtubedl.ExtractVideoID(url)
	if err != nil {
		return "", err
	}

	name, file, err := d.youtubedl.DownloadAudio(ctx, url)
	if err != nil {
		return "", err
	}

	filepath, err := d.audioRepo.CreateAudioFile(name, file)
	if err != nil {
		return "", err
	}

	err = d.audioRepo.CreateAudioDb(videoId, name, filepath)
	if err != nil {
		return "", err
	}
	return filepath, nil
}
