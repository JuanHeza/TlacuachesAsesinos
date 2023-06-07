package main

import (
	"TlacuachesAsesinos/constants"
	"TlacuachesAsesinos/database"
	"TlacuachesAsesinos/model"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	qrlogo "github.com/divan/qrlogo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

/*
https://www.manybooks.net/titles/doyleartetext94lostw10.html
https://www.manybooks.net/titles/baumlfraetext93wizoz10.html
https://www.manybooks.net/titles/wellshgetext92timem11.html
https://www.manybooks.net/titles/russian-roulette
https://www.manybooks.net/titles/first-magic
https://www.manybooks.net/titles/marked
https://www.manybooks.net/titles/grave-mistake
*/

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
		//bot.Request(tgbotapi.NewChatAction(update.Message.Chat.ID, "typing"))
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
			//fmt.Println("UNKNOWN")
			msg.Text = constants.Textos[constants.Idioma][constants.Const_error]
		}
	}
	return bot
}

func handleMessage(update tgbotapi.Update) (tgbotapi.Chattable, bool) {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	//fmt.Println(update.Message.From.LanguageCode)
	switch update.Message.Text {
	case "/start":
		msg = startCommand(update.Message.Chat.ID)
	default:
		if actualQuery != "" {
			var msgCallback tgbotapi.EditMessageTextConfig
			if update.Message.Photo != nil {
				if actualQuery != "" && actualForm == constants.Const_rgtVst {
					switch constants.ToInline(actualQuery) {
					case constants.Const_entFid:
						actualReg.Identificacion = update.Message.Photo[len(update.Message.Photo)-1].FileID
					case constants.Const_entFvh:
						actualReg.FotoVehiculo = update.Message.Photo[len(update.Message.Photo)-1].FileID
					default:
						return nil, false
					}

					data := model.EntradaMessage()
					msgCallback = tgbotapi.NewEditMessageTextAndMarkup(chatId, int(messageId), actualReg.PrintEntradaMesage(data.Message), data.Keyboard)

				} else {
					return nil, false
				}
			} else {
				var data model.Step
				var message string
				actualReg.FillRegistro(constants.ToInline(actualQuery), update.Message.Text)
				switch actualForm {
				case constants.Const_rgtDir:
					data = model.VisitanteMessage()
					message = actualReg.PrintRegistroMesage(data.Message)
				case constants.Const_rgtVst:
					data = model.EntradaMessage()
					message = actualReg.PrintEntradaMesage(data.Message)
				case constants.Const_rgtSld:
					data = model.SalidaMessage()
					message = actualReg.PrintSalidaMesage(data.Message)
				}
				msgCallback = tgbotapi.NewEditMessageTextAndMarkup(chatId, int(messageId), message, data.Keyboard)
			}
			msgCallback.ParseMode = "Markdown"
			return msgCallback, true
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
	//fmt.Println(update.CallbackData())
	msgCallback := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, "", tgbotapi.NewInlineKeyboardMarkup(model.MiniKeyboard(false)))
	msgCallback.ParseMode = "Markdown"

	switch constants.ToInline(strings.Split(update.CallbackData(), " - ")[0]) {
	case constants.Const_qrCode:
		now := fmt.Sprintf("%v@%s %s", time.Now().Unix(), update.CallbackQuery.From.UserName, constants.Session_key)
		encText, err := Encrypt(now, key)
		//fmt.Println("==========",now, "=================", encText, "===========")
		if err != nil {
			msg.Text = fmt.Sprint("error encrypting your classified text: ", err)
		} else {

			bot.Request(tgbotapi.NewChatAction(update.CallbackQuery.From.ID, "upload_photo"))
			image, err := downloadFile(bot, "bot.jpg")
			//fmt.Println("decode", err)
			logo, _ := qrlogo.Encode(encText, image, 2048)
			file := tgbotapi.FileBytes{
				Name:  "QR_CODE.jpg",
				Bytes: logo.Bytes(),
			}
			pic := tgbotapi.NewInputMediaPhoto(file)
			pic.Caption = fmt.Sprintf("Este mensaje se eliminara en %v segundos", constants.Const_delay)
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
			actualForm = constants.Const_rgtVst
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
		if actualReg.CheckSalida() && !actualReg.CheckEntrada() {
			actualForm = constants.Const_rgtSld
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
		pic := tgbotapi.FileID(actualReg.Identificacion)
		bot.Send(tgbotapi.NewPhoto(update.CallbackQuery.Message.Chat.ID, pic))

	case constants.Const_verFvh:
		pic := tgbotapi.FileID(actualReg.FotoVehiculo)
		bot.Send(tgbotapi.NewPhoto(update.CallbackQuery.Message.Chat.ID, pic))

	case constants.Const_back:
		data := model.HomeMessage()
		msgCallback.Text = data.Message
		msgCallback.ReplyMarkup = &data.Keyboard

	case constants.Const_ok, constants.Const_okEnt, constants.Const_okSal:
		mensajeGuardar(update.CallbackData(), bot, update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID)

	case constants.Const_langEs, constants.Const_langEn:
		changeLenguage(update, *bot)

	case constants.ToInline("Excel"):
		bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, constants.Const_file_url))

	case constants.ToInline("buscar"):
		data := model.BusquedaMessage()
		msgCallback.Text = "Seleccione su metodo de busqueda"
		msgCallback.ReplyMarkup = &data.Keyboard

	case constants.ToInline("Nombre"):
		msgCallback.Text = "Ingrese el nombre"

	case constants.ToInline("Folio"):
		folioBuscar := strings.Split(update.CallbackData(), " - ")[1]
		if folioBuscar != "0" {
			actualReg = database.Search(folioBuscar, bot.Self.UserName)
			constants.Print(actualReg)

			newMsg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, actualReg.Print())
			newMsg.ReplyMarkup = model.FolioKeyboard().Keyboard
			newMsg.ParseMode = "Markdown"
			bot.Send(newMsg)
		} else {
			msgCallback.Text = "Ingrese el folio"
		}
	case constants.ToInline("Lista"):
		data := model.ListaMessage(database.SearchPage(0, bot.Self.UserName))
		msgCallback.Text = data.Message
		msgCallback.ReplyMarkup = &data.Keyboard

	case constants.ToInline("Page"):
		pagina, _ := strconv.Atoi(strings.Split(update.CallbackData(), " - ")[1])
		data := model.ListaMessage(database.SearchPage(pagina, bot.Self.UserName))
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

func mensajeGuardar(query string, bot *tgbotapi.BotAPI, chatId int64, messageId int) {
	step := constants.GetQuerry(query)
	if step != -1 {
		row, auxFolio := database.Save(actualReg, step)
		msgNew := tgbotapi.NewMessage(chatId, fmt.Sprint(constants.GetMessage(step), row))
		msgNew.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(model.SearchKeyboard(auxFolio))

		bot.Send(tgbotapi.NewDeleteMessage(chatId, messageId))
		bot.Send(msgNew)
		bot.Send(startCommand(chatId))
	}
}

func downloadFile(bot *tgbotapi.BotAPI, filepath string) (image.Image, error) {

	me, _ := bot.GetMe()
	ph, _ := bot.GetUserProfilePhotos(tgbotapi.UserProfilePhotosConfig{
		UserID: me.ID,
		Offset: 0,
		Limit:  1,
	})
	phs := ph.Photos[0]
	tf, _ := bot.GetFile(tgbotapi.FileConfig{
		FileID: phs[len(phs)-1].FileID,
	})
	url := fmt.Sprintf("https://api.telegram.org/file/bot%v/%v", constants.Telegram_token, tf.FilePath)
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return jpeg.Decode(resp.Body)
}
