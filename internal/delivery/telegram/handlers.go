package telegram

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandCmd = "cmd"
	// commandHelp = "help"
)

// To all messages send info help
func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	filepath, err := b.services.Downloader.Download(context.TODO(), message.Text)
	if err != nil {
		return err
	}

	err = b.sendAudioFile(message.Chat.ID, filepath)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	command := message.Command()
	chatId := message.Chat.ID

	switch command {
	// Command for run live scripts
	case commandCmd:
		// Get payload
		script := strings.Replace(message.Text, fmt.Sprintf("/%s ", commandCmd), "", 1)
		return b.executeCommand(chatId, script)
	// case commandHelp:
	// return b.handleHelpCommand(message)
	default:
		return errUnknownCommand

	}
}

func (b *Bot) executeCommand(chatId int64, script string) error {
	id := 42

	b.logger.Debug(id)

	msg := tgbotapi.NewMessage(chatId, strconv.Itoa(id))
	if _, err := b.bot.Send(msg); err != nil {
		return fmt.Errorf("error execute cmd command : %v", err)
	}
	return nil
}

func (b *Bot) sendAudioFile(chatID int64, filename string) error {
	path := "./" + filename

	// defer os.Remove(path)

	msgAudio := tgbotapi.NewAudioUpload(chatID, path)
	msgAudio.Caption = "Downloaded via @"

	_, err := b.bot.Send(msgAudio)
	if err != nil {
		return fmt.Errorf("error sending audio message: %v", err)
	}
	return nil
}

func (b *Bot) sendInfoSudoChat(msg string) {
	if b.sudoChatId <= 0 {
		b.logger.Warn(errNotFoundSudoChatId)
		return
	}

	t := time.Now().Format("01.02.2006 15:04:05")

	m := tgbotapi.NewMessage(b.sudoChatId, msg+" "+t)
	_, err := b.bot.Send(m)
	if err != nil {
		b.logger.Error(err)
	}
}
