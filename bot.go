package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	qrcode "github.com/skip2/go-qrcode"
)

var (
	startKeyboard tgbotapi.InlineKeyboardMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Generar QR", string(const_qrCode)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Registrar Direccion", string(const_rgtDir)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Registrar Visitante", string(const_rgtVst)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Registrar Salida", string(const_rgtSld)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("\xF0\x9F\x87\xAA\xF0\x9F\x87\xB8", string(const_langEs)),
			tgbotapi.NewInlineKeyboardButtonData("\xF0\x9F\x87\xBA\xF0\x9F\x87\xB8", string(const_langEn)),
		))
	msg tgbotapi.MessageConfig
)

func botInit() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(telegram_token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		print(update.Message)
		switch true {
		case update.CallbackQuery != nil:
			msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
			msg = handleCallback(update, bot)
			break
		case update.Message != nil:
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg = handleMessage(update)
			break
		default:
			fmt.Println("UNKNOWN")
			msg.Text = textos["es"]["error"]
		}
		bot.Send(msg)
	}
	return bot
}

func handleMessage(update tgbotapi.Update) tgbotapi.MessageConfig {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	fmt.Println(update.Message.From.LanguageCode)
	switch update.Message.Text {
	case "/start":
		msg = startCommand(update.Message.Chat.ID)
		break
	default:
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%v \xF0\x9F\x87\xBA\xF0\x9F\x87\xB8", update.Message.Text))
		msg.ReplyToMessageID = update.Message.MessageID
	}
	return msg
}

func startCommand(Id int64) tgbotapi.MessageConfig {
	msg = tgbotapi.NewMessage(Id, textos["es"]["saludo"])
	msg.ReplyMarkup = startKeyboard
	return msg
}

func handleCallback(update tgbotapi.Update, bot *tgbotapi.BotAPI) tgbotapi.MessageConfig {
	fmt.Println(update.CallbackData())
	switch toInline(update.CallbackData()) {
	case const_qrCode:
		encoded, _ := qrcode.Encode("https://example.org", qrcode.Medium, 256)
		file := tgbotapi.FileBytes{
			Name:  "image.jpg",
			Bytes: encoded,
		}
		bot.SendMediaGroup(tgbotapi.NewMediaGroup(update.CallbackQuery.From.ID, []interface{}{tgbotapi.NewInputMediaPhoto(file)}))

	}
	return msg

}
