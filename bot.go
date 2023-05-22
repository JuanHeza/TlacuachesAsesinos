package main

import (
	"TlacuachesAsesinos/constants"
	"TlacuachesAsesinos/database"
	"TlacuachesAsesinos/model"
	"fmt"
	"log"
	"strings"
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
	folio       = 0
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
			messageId = update.CallbackQuery.Message.MessageID

			chatId = update.CallbackQuery.Message.Chat.ID
			msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
			handleCallback(update, bot)
		case update.Message != nil:
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Dummy")
			msg, status := handleMessage(update)

			if msg != nil {

				delete := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
				bot.Send(delete)
				sended, ok := bot.Send(msg)
				fmt.Println(ok)
				if !status {
					chatId = sended.Chat.ID
					messageId = sended.MessageID
					actualQuery = ""
				}
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
			if update.Message.Photo != nil {
				fmt.Println(update.Message.Photo)
				data := model.VisitanteMessage()
				return tgbotapi.NewEditMessageTextAndMarkup(chatId, int(messageId), actualReg.PrintRegistroMesage(data.Message), data.Keyboard), true
			} else {
				actualReg.FillRegistro(constants.ToInline(actualQuery), update.Message.Text)
				data := model.VisitanteMessage()
				return tgbotapi.NewEditMessageTextAndMarkup(chatId, int(messageId), actualReg.PrintRegistroMesage(data.Message), data.Keyboard), true
			}
		}
		return nil, false
	}
	return msg, false
}

func startCommand(Id int64) tgbotapi.MessageConfig {
	home := model.HomeMessage()
	msgLocal := tgbotapi.NewMessage(Id, home.Message)
	msgLocal.ReplyMarkup = home.Keyboard
	return msgLocal
}

func handleCallback(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	fmt.Println(update.CallbackData())
	msgCallback := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, "", tgbotapi.NewInlineKeyboardMarkup(model.MiniKeyboard(false)))
	msgCallback.ParseMode = "Markdown"

	switch constants.ToInline(strings.Split(update.CallbackData(), " - ")[0]) {
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
		actualReg = model.Registro{}
		actualForm = constants.Const_rgtDir
		data := model.VisitanteMessage()
		msgCallback.Text = actualReg.PrintRegistroMesage(data.Message)
		msgCallback.ReplyMarkup = &data.Keyboard

	case constants.Const_rgtVst:
		if actualReg.CheckEntrada() {
			data := model.EntradaMessage()
			msgCallback.Text = actualReg.PrintEntradaMesage(data.Message)
			msgCallback.ReplyMarkup = &data.Keyboard
		} else {
			noModify(bot, update.CallbackQuery.From.ID)
		}

	case constants.Const_entHra:
		actualReg.HoraEntrada = time.Now()
		data := model.EntradaMessage()
		msgCallback.Text = actualReg.PrintEntradaMesage(data.Message)
		msgCallback.ReplyMarkup = &data.Keyboard

	case constants.Const_entFch:
		actualReg.FechaEntrada = time.Now()
		data := model.EntradaMessage()
		msgCallback.Text = actualReg.PrintEntradaMesage(data.Message)
		msgCallback.ReplyMarkup = &data.Keyboard

	case constants.Const_rgtSld:
		if actualReg.CheckSalida() {
			data := model.SalidaMessage()
			msgCallback.Text = actualReg.PrintSalidaMesage(data.Message)
			msgCallback.ReplyMarkup = &data.Keyboard
		} else {
			noModify(bot, update.CallbackQuery.From.ID)
		}

	case constants.Const_sldHra:
		actualReg.HoraSalida = time.Now()
		data := model.SalidaMessage()
		msgCallback.Text = actualReg.PrintSalidaMesage(data.Message)
		msgCallback.ReplyMarkup = &data.Keyboard

	case constants.Const_sldFch:
		actualReg.FechaSalida = time.Now()
		data := model.SalidaMessage()
		msgCallback.Text = actualReg.PrintSalidaMesage(data.Message)
		msgCallback.ReplyMarkup = &data.Keyboard

	case constants.Const_verFid:
		pic := tgbotapi.File{FileUniqueID: "AQADtqcxG7xLWE9-", FileID: "AgACAgEAAxkBAAICIWRq_W07eDYR5RQvJbMETm1EBiTUAAK2pzEbvEtYTztoL_lFzerQAQADAgADeQADLwQ"}
		bot.SendMediaGroup(tgbotapi.NewMediaGroup(update.CallbackQuery.From.ID, []interface{}{pic}))

	case constants.Const_verFvh:
		pic := tgbotapi.File{FileUniqueID: "AQADtqcxG7xLWE9-", FileID: "AgACAgEAAxkBAAICIWRq_W07eDYR5RQvJbMETm1EBiTUAAK2pzEbvEtYTztoL_lFzerQAQADAgADeQADLwQ"}
		bot.SendMediaGroup(tgbotapi.NewMediaGroup(update.CallbackQuery.From.ID, []interface{}{pic}))

	case constants.Const_back:
		data := model.HomeMessage()
		msgCallback.Text = data.Message
		msgCallback.ReplyMarkup = &data.Keyboard

	case constants.Const_ok:
		row, auxFolio := database.Save(actualReg, folio)
		msgNew := tgbotapi.NewMessage(update.CallbackQuery.From.ID, fmt.Sprint("Se han guardado los datos con el siguiente folio\n", row))
		msgNew.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(model.SearchKeyboard(auxFolio))

		bot.Send(tgbotapi.NewDeleteMessage(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID))
		bot.Send(msgNew)
		bot.Send(startCommand(update.CallbackQuery.Message.Chat.ID))

	case constants.Const_langEs:
		changeLenguage(update, *bot)

	case constants.Const_langEn:
		changeLenguage(update, *bot)

	case constants.ToInline("Excel"):
		bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "https://docs.google.com/spreadsheets/d/1ytQV2XIoxR2TdiBiBIARRx5mg1v48kk7VF74eZ4LanQ/edit#gid=0"))

	case constants.ToInline("buscar"):
		data := model.BusquedaMessage()
		msgCallback.Text = "Seleccione su metodo de busqueda"
		msgCallback.ReplyMarkup = &data.Keyboard

	case constants.ToInline("Folio"):
		folioBuscar := strings.Split(update.CallbackData(), " - ")[1]
		if folioBuscar != "0" {
			actualReg = database.Search(folioBuscar)
			constants.Print(actualReg)

			newMsg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, actualReg.Print())
			newMsg.ReplyMarkup = model.FolioKeyboard().Keyboard
			newMsg.ParseMode = "Markdown"
			bot.Send(newMsg)
		} else {
			msgCallback.Text = "Ingrese el folio"
		}

	case constants.ToInline("Nombre"):
		msgCallback.Text = "Ingrese el nombre"

	case constants.ToInline("Lista"):
		data := model.ListaMessage()
		msgCallback.Text = data.Message
		msgCallback.ReplyMarkup = &data.Keyboard

	default:
		actualQuery = update.CallbackData()
		msgCallback.Text = constants.InputMessage(constants.ToInline(actualQuery))
	}

	if msgCallback.Text != "" {
		bot.Send(msgCallback)
	}
}

func delaySecond(msgs []tgbotapi.Message, bot *tgbotapi.BotAPI) {
	time.Sleep(time.Duration(constants.Const_delay) * time.Second)
	for _, msx := range msgs {
		bot.Send(tgbotapi.NewDeleteMessage(msx.Chat.ID, msx.MessageID))
	}
}

func changeLenguage(update tgbotapi.Update, bot tgbotapi.BotAPI) {
	constants.Idioma = constants.ToInline(update.CallbackData())
	home := model.HomeMessage()
	bot.Send(tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, home.Message, home.Keyboard))
}

func noModify(bot *tgbotapi.BotAPI, chatId int64) {
	msgNew := tgbotapi.NewMessage(chatId, "No se pueden modificar estos datos")
	sended, ok := bot.Send(msgNew)
	if ok == nil {
		go delaySecond([]tgbotapi.Message{sended}, bot) // very useful for interval polling
	}
}
