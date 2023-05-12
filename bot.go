package main

import (
	"TlacuachesAsesinos/constants"
	"TlacuachesAsesinos/database"
	"TlacuachesAsesinos/model"
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	qrcode "github.com/skip2/go-qrcode"
)

var (
	actualQuery = ""
	actualReg   = model.Registro{}
	actualForm  constants.Inline
	messageId   int
	chatId      int64
	msg         tgbotapi.MessageConfig
)

func botInit() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(constants.Telegram_token)
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
			bot.Send(handleCallback(update, bot))
		case update.Message != nil:
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg, status := handleMessage(update)

			sended, _ := bot.Send(msg)
			if !status {
				chatId = sended.Chat.ID
				messageId = sended.MessageID
				actualQuery = ""
			}
		default:
			fmt.Println("UNKNOWN")
			msg.Text = constants.Textos[constants.Idioma][constants.Const_error]
		}
	}
	return bot
}

func handleMessage(update tgbotapi.Update) (tgbotapi.Chattable, bool) {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	fmt.Println(update.Message.From.LanguageCode)
	switch update.Message.Text {
	case "/start":
		msg = startCommand(update.Message.Chat.ID)
	default:
		if actualQuery != "" {
			//msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%v \xF0\x9F\x87\xBA\xF0\x9F\x87\xB8", update.Message.Text))
			//msg.ReplyToMessageID = update.Message.MessageID
			//} else {
			actualReg.FillRegistro(constants.ToInline(actualQuery), update.Message.Text)
			data := model.EntradaMessage()
			return tgbotapi.NewEditMessageTextAndMarkup(chatId, int(messageId), actualReg.PrintEntradaMesage(data.Message), data.Keyboard), true
		}
	}
	return msg, false
}

func startCommand(Id int64) tgbotapi.MessageConfig {
	home := model.HomeMessage()
	msg = tgbotapi.NewMessage(Id, home.Message)
	msg.ReplyMarkup = home.Keyboard
	return msg
}

func handleCallback(update tgbotapi.Update, bot *tgbotapi.BotAPI) tgbotapi.MessageConfig {
	fmt.Println(update.CallbackData())
	switch constants.ToInline(update.CallbackData()) {
	case constants.Const_qrCode:
		now := fmt.Sprintf("%v%s", time.Now().Unix(), constants.Session_key)
		fmt.Println(now)
		encText, err := Encrypt(now, key)
		fmt.Println(encText)
		if err != nil {
			msg.Text = fmt.Sprint("error encrypting your classified text: ", err)
		} else {
			encoded, _ := qrcode.Encode(encText, qrcode.Medium, 256)
			file := tgbotapi.FileBytes{
				Name:  "QR_CODE.jpg",
				Bytes: encoded,
			}
			pic := tgbotapi.NewInputMediaPhoto(file)
			pic.Caption = "Este mensaje se eliminara en 5 segundos"
			msgPic, err := bot.SendMediaGroup(tgbotapi.NewMediaGroup(update.CallbackQuery.From.ID, []interface{}{pic}))
			if err == nil {
				go delaySecond(msgPic, bot) // very useful for interval polling
			}
		}
	case constants.Const_rgtDir:
		actualReg := model.Registro{}
		actualForm = constants.Const_rgtDir
		data := model.EntradaMessage()
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, actualReg.PrintEntradaMesage(data.Message), data.Keyboard)
		bot.Send(msg)
	case constants.Const_rgtVst:
		data := model.VisitanteMessage()
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, actualReg.PrintEntradaMesage(data.Message), data.Keyboard)
		bot.Send(msg)
	case constants.Const_rgtSld:
		data := model.SalidaMessage()
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, actualReg.PrintSalidaMesage(data.Message), data.Keyboard)
		bot.Send(msg)
	case constants.Const_back:
		home := model.HomeMessage()
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, home.Message, home.Keyboard)
		bot.Send(msg)
	case constants.Const_ok:
		var data model.Step
		switch actualForm {
		case constants.Const_rgtDir:
			data = model.EntradaMessage()
			data.Message = actualReg.PrintEntradaMesage(data.Message)
		case constants.Const_rgtVst:
			data = model.VisitanteMessage()
		case constants.Const_rgtSld:
			data = model.SalidaMessage()
		default:
			data = model.HomeMessage()
		}
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, data.Message, data.Keyboard)
		database.Save(actualReg)
		bot.Send(msg)
	case constants.Const_langEs:
		changeLenguage(update, *bot)
	case constants.Const_langEn:
		changeLenguage(update, *bot)
	case constants.Const_sldHra:
		actualReg.HoraSalida = time.Now()
		data := model.SalidaMessage()
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, actualReg.PrintSalidaMesage(data.Message), data.Keyboard)
		bot.Send(msg)
	case constants.Const_sldFch:
		actualReg.FechaSalida = time.Now()
		data := model.SalidaMessage()
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, actualReg.PrintSalidaMesage(data.Message), data.Keyboard)
		bot.Send(msg)
	default:
		actualQuery = update.CallbackData()
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, constants.InputMessage(constants.ToInline(actualQuery)), tgbotapi.NewInlineKeyboardMarkup(model.MiniKeyboard))
		bot.Send(msg)
	}
	return msg
}

func delaySecond(msgs []tgbotapi.Message, bot *tgbotapi.BotAPI) {
	for range time.Tick(time.Duration(constants.Const_delay) * time.Second) {
		for _, msx := range msgs {
			delete := tgbotapi.NewDeleteMessage(msx.Chat.ID, msx.MessageID)
			bot.Send(delete)
		}
	}
}

func changeLenguage(update tgbotapi.Update, bot tgbotapi.BotAPI) {
	constants.Idioma = constants.ToInline(update.CallbackData())
	home := model.HomeMessage()
	msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, home.Message, home.Keyboard)
	bot.Send(msg)
}
