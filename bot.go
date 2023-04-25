package main

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	qrcode "github.com/skip2/go-qrcode"
)

type Step struct {
	message  string
	keyboard tgbotapi.InlineKeyboardMarkup
}

type Registro struct {
	Nombre         string
	Motivo         string
	Company        string
	Calle          string
	NumeroExterior int16
	Folio          int64
	AutoFabricante string
	Color          string
	FechaEntrada   time.Time
	HoraEntrada    time.Time
	Identificacion string
	FotoVehiculo   string
	Observaciones  string
}

func (rg *Registro) printEntradaMesage() string {
	return fmt.Sprintf("Nombre:\n%s\n\nCompañia:\n%s\n\nMotivo:\n%s\n\nCalle:\n%s\n\nNumero Exterior:\n%v", rg.Nombre, rg.Company, rg.Motivo, rg.Calle, rg.NumeroExterior)
}

var (
	homeMessage = Step{
		message: textos["es"]["saludo"],
		keyboard: tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Generar QR", string(const_qrCode)),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Registrar Entrada", string(const_rgtDir)),
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
			)),
	}

	entradaMessage = Step{
		message: "Ingrese la siguiente informacion",
		keyboard: tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Nombre", string(const_qrCode)),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Compañia", string(const_rgtDir)),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Motivo", string(const_rgtDir)),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Calle", string(const_rgtVst)),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Numero Exterior", string(const_back)),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Aceptar", string(const_back)),
				tgbotapi.NewInlineKeyboardButtonData("Cancelar", string(const_back)),
			)),
	}

	visitanteMessage = Step{
		message: "Ingrese la siguiente informacion",
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
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Aceptar", string(const_back)),
				tgbotapi.NewInlineKeyboardButtonData("Cancelar", string(const_back)),
			)),
	}

	salidaMessage = Step{
		message: "Ingrese la siguiente informacion",
		keyboard: tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Fecha Salida", string(const_qrCode)),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Hora Salida", string(const_rgtDir)),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Aceptar", string(const_back)),
				tgbotapi.NewInlineKeyboardButtonData("Cancelar", string(const_back)),
			)),
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
	msg = tgbotapi.NewMessage(Id, homeMessage.message)
	msg.ReplyMarkup = homeMessage.keyboard
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
		break
	case const_rgtDir:
		dir := Registro{}
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, entradaMessage.message+"\n"+dir.printEntradaMesage(), entradaMessage.keyboard)
		bot.Send(msg)
		break
	case const_back:
		msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, homeMessage.message, homeMessage.keyboard)
		bot.Send(msg)
		break
	}
	return msg

}

func delaySecond(msgs []tgbotapi.Message, bot *tgbotapi.BotAPI) {
	for _ = range time.Tick(time.Duration(const_delay) * time.Second) {
		for _, msx := range msgs {
			delete := tgbotapi.NewDeleteMessage(msx.Chat.ID, msx.MessageID)
			bot.Send(delete)
		}
	}
}
