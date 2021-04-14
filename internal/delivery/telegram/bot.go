package telegram

import (
	"github.com/staszigzag/downloader-music/internal/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/staszigzag/downloader-music/internal/config"
	"github.com/staszigzag/downloader-music/pkg/logger"
)

type Bot struct {
	services        *service.Services
	bot             *tgbotapi.BotAPI
	messages        *config.Messages
	token           string
	sudoChatId      int64
	logger          logger.Logger
	isDebug         bool
	shutdownChannel chan struct{}
}

func NewBot(services *service.Services, config *config.Config, logger logger.Logger) *Bot {
	return &Bot{
		services:        services,
		messages:        &config.Bot.Messages,
		token:           config.Bot.TelegramToken,
		sudoChatId:      config.Bot.SudoChatId,
		logger:          logger,
		isDebug:         config.Debug,
		shutdownChannel: make(chan struct{}, 1),
	}
}

func (b *Bot) Start() error {
	botApi, err := tgbotapi.NewBotAPI(b.token)
	if err != nil {
		return err
	}
	botApi.Debug = b.isDebug
	b.bot = botApi

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	// Send info start for sudo chat
	b.sendInfoSudoChat(b.messages.Responses.Start)

	b.logger.Info("Bot is running...")

LOOP:
	for {
		select {
		case update := <-updates:
			// Ignore any non-Message Updates
			if update.Message == nil {
				continue
			}

			var err error
			if update.Message.IsCommand() {
				// Handle commands
				err = b.handleCommand(update.Message)
			} else {
				// Handle regular messages
				err = b.handleMessage(update.Message)
			}

			if err != nil {
				b.handleError(update.Message.Chat.ID, err)
			}
		// Shutdown bot
		case <-b.shutdownChannel:
			// Send info finish for sudo chat
			b.sendInfoSudoChat(b.messages.Responses.Finish)
			break LOOP

		}
	}
	b.logger.Info("Bot is stopped!")
	return nil
}

func (b *Bot) Stop() {
	b.shutdownChannel <- struct{}{}
}
