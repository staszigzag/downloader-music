package telegram

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/staszigzag/downloader-music/internal/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandCmd = "cmd"
	// commandHelp = "help"
)

// To all messages send info help
func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	var commands string
	//// Exist commands text
	//for key := range b.commands {
	//	commands += fmt.Sprintf("- %s\n", key)
	//}

	text := fmt.Sprintf(b.messages.Responses.Help, strconv.Itoa(int(message.Chat.ID)), commands)

	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	_, err := b.bot.Send(msg)
	return err
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
	id, err := b.services.CreateUser(domain.User{Name: "test"})
	if err != nil {
		return err
	}

	b.logger.Debug(id)

	msg := tgbotapi.NewMessage(chatId, strconv.Itoa(id))
	if _, err = b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}
