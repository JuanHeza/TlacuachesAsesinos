package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	qrcode "github.com/skip2/go-qrcode"
)

type Step struct {
	name     Inline
	message  string
	keyboard tgbotapi.InlineKeyboardMarkup
}

type Registro struct {
	Nombre         string
	Motivo         string
	Company        string
	Calle          string
	NumeroExterior int
	Folio          int64
	AutoFabricante string
	Color          string
	FechaEntrada   time.Time
	HoraEntrada    time.Time
	FechaSalida    time.Time
	HoraSalida     time.Time
	Identificacion string
	FotoVehiculo   string
	Observaciones  string
}

func (rg *Registro) printEntradaMesage(msg string) string {
	return fmt.Sprintf("%s\n\n%s:\n%s\n\n%s:\n%s\n\n%s:\n%s\n\n%s:\n%s\n\n%s:\n%v", msg, textos[idioma][const_rgtNom], rg.Nombre, textos[idioma][const_rgtCom], rg.Company, textos[idioma][const_rgtMot], rg.Motivo, textos[idioma][const_rgtCll], rg.Calle, textos[idioma][const_rgtExt], rg.NumeroExterior)
}

func (rg *Registro) printSalidaMesage(msg string) string {
	y, m, d := rg.FechaSalida.Date()
	h, min, s := rg.HoraSalida.Clock()
	date, hour := "\n", "\n"
	if !rg.FechaSalida.IsZero() {
		date = fmt.Sprintf("%02d/%v/%04d", d, printMes(int(m)), y)
	}
	if !rg.HoraSalida.IsZero() {
		hour = fmt.Sprintf("%02d:%02d:%02d", h, min, s)
	}
	return fmt.Sprintf("%s\n\n%s:\n%s\n\n%s:\n%s", msg, textos[idioma][const_sldFch], date, textos[idioma][const_sldHra], hour)
}

func (rg *Registro) fillRegistro(campo Inline, valor interface{}) {
	switch campo {
	case const_rgtNom:
		rg.Nombre = fmt.Sprintf("%v", valor)
	case const_rgtCom:
		rg.Company = fmt.Sprintf("%v", valor)
	case const_rgtCll:
		rg.Calle = fmt.Sprintf("%v", valor)
	case const_rgtExt:
		num, err := strconv.Atoi(fmt.Sprintf("%v", valor))
		if err == nil {
			rg.NumeroExterior = num
		}
	case const_rgtMot:
		rg.Motivo = fmt.Sprintf("%v", valor)
	}
}

var (
	actualQuery = ""
	actualReg   = Registro{}
	actualForm  Inline
	messageId   int
	chatId      int64
	idioma      = const_langEs
	homeMessage = func() Step {
		return Step{
			name:    const_home,
			message: textos[idioma][const_home],
			keyboard: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(textos[idioma][const_qrCode], string(const_qrCode)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(textos[idioma][const_rgtDir], string(const_rgtDir)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(textos[idioma][const_rgtVst], string(const_rgtVst)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(textos[idioma][const_rgtSld], string(const_rgtSld)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("\xF0\x9F\x87\xAA\xF0\x9F\x87\xB8", string(const_langEs)),
					tgbotapi.NewInlineKeyboardButtonData("\xF0\x9F\x87\xBA\xF0\x9F\x87\xB8", string(const_langEn)),
				)),
		}
	}

	miniKeyboard = tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Aceptar", string(const_ok)),
		tgbotapi.NewInlineKeyboardButtonData("Cancelar", string(const_back)),
	)

	entradaMessage = func() Step {
		return Step{
			name:    const_entrada,
			message: textos[idioma][const_entrada],
			keyboard: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(textos[idioma][const_rgtNom], string(const_rgtNom)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(textos[idioma][const_rgtCom], string(const_rgtCom)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(textos[idioma][const_rgtMot], string(const_rgtMot)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(textos[idioma][const_rgtCll], string(const_rgtCll)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(textos[idioma][const_rgtExt], string(const_rgtExt)),
				),
				miniKeyboard),
		}
	}

	visitanteMessage = func() Step {
		return Step{
			name:    const_visitante,
			message: textos[idioma][const_visitante],
			keyboard: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Marca del vehiculo", string(const_qrCode)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Color", string(const_rgtDir)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Fecha Entrada", string(const_rgtDir)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Hora Entrada", string(const_rgtVst)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Identificacion", string(const_back)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Vehiculo", string(const_back)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Observaciones", string(const_back)),
				),
				miniKeyboard),
		}
	}

	salidaMessage = func() Step {
		return Step{
			name:    const_salida,
			message: textos[idioma][const_salida],
			keyboard: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(textos[idioma][const_sldFch], string(const_sldFch)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(textos[idioma][const_sldHra], string(const_sldHra)),
				),
				miniKeyboard),
		}
	}

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
			msg.Text = textos[idioma][const_error]
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
			actualReg.fillRegistro(toInline(actualQuery), update.Message.Text)
			data := entradaMessage()
			return tgbotapi.NewEditMessageTextAndMarkup(chatId, int(messageId), actualReg.printEntradaMesage(data.message), data.keyboard), true
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
	switch toInline(update.CallbackData()) {
	case const_qrCode:
		now := fmt.Sprintf("%v%s", time.Now().Unix(), session_key)
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
	case const_rgtDir:
		actualReg := Registro{}
		actualForm = const_rgtDir
		data := entradaMessage()
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, actualReg.printEntradaMesage(data.message), data.keyboard)
		bot.Send(msg)
	case const_rgtVst:
		data := visitanteMessage()
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, actualReg.printEntradaMesage(data.message), data.keyboard)
		bot.Send(msg)
	case const_rgtSld:
		data := salidaMessage()
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, actualReg.printSalidaMesage(data.message), data.keyboard)
		bot.Send(msg)
	case const_back:
		home := homeMessage()
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, home.message, home.keyboard)
		bot.Send(msg)
	case const_ok:
		var data Step
		switch actualForm {
		case const_rgtDir:
			data = entradaMessage()
			data.message = actualReg.printEntradaMesage(data.message)
		case const_rgtVst:
			data = visitanteMessage()
		case const_rgtSld:
			data = salidaMessage()
		default:
			data = homeMessage()
		}
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, data.message, data.keyboard)
		bot.Send(msg)
	case const_langEs:
		changeLenguage(update, *bot)
	case const_langEn:
		changeLenguage(update, *bot)
	case const_sldHra:
		actualReg.HoraSalida = time.Now()
		data := salidaMessage()
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, actualReg.printSalidaMesage(data.message), data.keyboard)
		bot.Send(msg)
	case const_sldFch:
		actualReg.FechaSalida = time.Now()
		data := salidaMessage()
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, actualReg.printSalidaMesage(data.message), data.keyboard)
		bot.Send(msg)
	default:
		actualQuery = update.CallbackData()
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, inputMessage(toInline(actualQuery)), tgbotapi.NewInlineKeyboardMarkup(miniKeyboard))
		bot.Send(msg)
	}
	return msg
}

func delaySecond(msgs []tgbotapi.Message, bot *tgbotapi.BotAPI) {
	for range time.Tick(time.Duration(const_delay) * time.Second) {
		for _, msx := range msgs {
			delete := tgbotapi.NewDeleteMessage(msx.Chat.ID, msx.MessageID)
			bot.Send(delete)
		}
	}
}

func changeLenguage(update tgbotapi.Update, bot tgbotapi.BotAPI) {
	idioma = toInline(update.CallbackData())
	home := homeMessage()
	msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, home.message, home.keyboard)
	bot.Send(msg)
}
