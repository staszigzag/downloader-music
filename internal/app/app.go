package app

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/staszigzag/downloader-music/pkg/youtubedl"

	"github.com/staszigzag/downloader-music/internal/config"
	"github.com/staszigzag/downloader-music/internal/delivery/http"
	"github.com/staszigzag/downloader-music/internal/delivery/telegram"
	"github.com/staszigzag/downloader-music/internal/repository"
	"github.com/staszigzag/downloader-music/internal/server"
	"github.com/staszigzag/downloader-music/internal/service"
	"github.com/staszigzag/downloader-music/pkg/logger"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		logger.Fatal(err)
	}

	//// Dependencies
	log := logger.NewLogrus(cfg.Debug)
	log.Debug(fmt.Sprintf("%+v\n", cfg))
	// Instruction exec
	// sh := shell.NewShell()
	// Downloader audio with youtube
	ydl := youtubedl.NewDownloader(cfg.FileStorage.Path)
	db := "test"

	//// Services, Repos & API Handlers/Bot
	repos := repository.NewRepository(db)
	services := service.NewServices(service.Deps{
		Repos:      repos,
		Downloader: ydl,
	})
	handlers := http.NewHandler(services)

	wg := sync.WaitGroup{}

	//// HTTP Server
	srv := server.NewServer(cfg, handlers.Init())
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := srv.Run(); err != nil {
			log.Error("error occurred while running http server: ", err)
		}
	}()

	//// Bot
	bot := telegram.NewBot(services, cfg, log)
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := bot.Start(); err != nil {
			log.Fatal("error occurred while running bot: ", err)
		}
	}()

	//// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Stop(); err != nil {
		log.Error("failed to stop server: ", err)
	}
	bot.Stop()

	wg.Wait()
}
