package main

import (
	"TlacuachesAsesinos/constants"
	"TlacuachesAsesinos/model"
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	qrcode "github.com/skip2/go-qrcode"
)

type Step struct {
	name     constants.Inline
	message  string
	keyboard tgbotapi.InlineKeyboardMarkup
}

var (
	actualQuery = ""
	actualReg   = model.Registro{}
	actualForm  constants.Inline
	messageId   int
	chatId      int64
	homeMessage = func() Step {
		return Step{
			name:    constants.Const_home,
			message: constants.Textos[constants.Idioma][constants.Const_home],
			keyboard: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.Textos[constants.Idioma][constants.Const_qrCode], string(constants.Const_qrCode)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.Textos[constants.Idioma][constants.Const_rgtDir], string(constants.Const_rgtDir)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.Textos[constants.Idioma][constants.Const_rgtVst], string(constants.Const_rgtVst)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.Textos[constants.Idioma][constants.Const_rgtSld], string(constants.Const_rgtSld)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("\xF0\x9F\x87\xAA\xF0\x9F\x87\xB8", string(constants.Const_langEs)),
					tgbotapi.NewInlineKeyboardButtonData("\xF0\x9F\x87\xBA\xF0\x9F\x87\xB8", string(constants.Const_langEn)),
				)),
		}
	}

	miniKeyboard = tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Aceptar", string(constants.Const_ok)),
		tgbotapi.NewInlineKeyboardButtonData("Cancelar", string(constants.Const_back)),
	)

	entradaMessage = func() Step {
		return Step{
			name:    constants.Const_entrada,
			message: constants.Textos[constants.Idioma][constants.Const_entrada],
			keyboard: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.Textos[constants.Idioma][constants.Const_rgtNom], string(constants.Const_rgtNom)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.Textos[constants.Idioma][constants.Const_rgtCom], string(constants.Const_rgtCom)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.Textos[constants.Idioma][constants.Const_rgtMot], string(constants.Const_rgtMot)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.Textos[constants.Idioma][constants.Const_rgtCll], string(constants.Const_rgtCll)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.Textos[constants.Idioma][constants.Const_rgtExt], string(constants.Const_rgtExt)),
				),
				miniKeyboard),
		}
	}

	visitanteMessage = func() Step {
		return Step{
			name:    constants.Const_visitante,
			message: constants.Textos[constants.Idioma][constants.Const_visitante],
			keyboard: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Marca del vehiculo", string(constants.Const_qrCode)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Color", string(constants.Const_rgtDir)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Fecha Entrada", string(constants.Const_rgtDir)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Hora Entrada", string(constants.Const_rgtVst)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Identificacion", string(constants.Const_back)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Vehiculo", string(constants.Const_back)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Observaciones", string(constants.Const_back)),
				),
				miniKeyboard),
		}
	}

	salidaMessage = func() Step {
		return Step{
			name:    constants.Const_salida,
			message: constants.Textos[constants.Idioma][constants.Const_salida],
			keyboard: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.Textos[constants.Idioma][constants.Const_sldFch], string(constants.Const_sldFch)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.Textos[constants.Idioma][constants.Const_sldHra], string(constants.Const_sldHra)),
				),
				miniKeyboard),
		}
	}

	msg tgbotapi.MessageConfig
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
			data := entradaMessage()
			return tgbotapi.NewEditMessageTextAndMarkup(chatId, int(messageId), actualReg.PrintEntradaMesage(data.message), data.keyboard), true
		}
	}
	return msg, false
}

func startCommand(Id int64) tgbotapi.MessageConfig {
	home := homeMessage()
	msg = tgbotapi.NewMessage(Id, home.message)
	msg.ReplyMarkup = home.keyboard
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
		data := entradaMessage()
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, actualReg.PrintEntradaMesage(data.message), data.keyboard)
		bot.Send(msg)
	case constants.Const_rgtVst:
		data := visitanteMessage()
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, actualReg.PrintEntradaMesage(data.message), data.keyboard)
		bot.Send(msg)
	case constants.Const_rgtSld:
		data := salidaMessage()
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, actualReg.PrintSalidaMesage(data.message), data.keyboard)
		bot.Send(msg)
	case constants.Const_back:
		home := homeMessage()
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, home.message, home.keyboard)
		bot.Send(msg)
	case constants.Const_ok:
		var data Step
		switch actualForm {
		case constants.Const_rgtDir:
			data = entradaMessage()
			data.message = actualReg.PrintEntradaMesage(data.message)
		case constants.Const_rgtVst:
			data = visitanteMessage()
		case constants.Const_rgtSld:
			data = salidaMessage()
		default:
			data = homeMessage()
		}
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, data.message, data.keyboard)
		bot.Send(msg)
	case constants.Const_langEs:
		changeLenguage(update, *bot)
	case constants.Const_langEn:
		changeLenguage(update, *bot)
	case constants.Const_sldHra:
		actualReg.HoraSalida = time.Now()
		data := salidaMessage()
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, actualReg.PrintSalidaMesage(data.message), data.keyboard)
		bot.Send(msg)
	case constants.Const_sldFch:
		actualReg.FechaSalida = time.Now()
		data := salidaMessage()
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, actualReg.PrintSalidaMesage(data.message), data.keyboard)
		bot.Send(msg)
	default:
		actualQuery = update.CallbackData()
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, constants.InputMessage(constants.ToInline(actualQuery)), tgbotapi.NewInlineKeyboardMarkup(miniKeyboard))
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
	home := homeMessage()
	msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, home.message, home.keyboard)
	bot.Send(msg)
}
