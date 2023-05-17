package model

import (
	"TlacuachesAsesinos/constants"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Step struct {
	Name     constants.Inline
	Message  string
	Keyboard tgbotapi.InlineKeyboardMarkup
}

var (
	HomeMessage = func() Step {
		return Step{
			Name:    constants.Const_home,
			Message: constants.Textos[constants.Idioma][constants.Const_home],
			Keyboard: tgbotapi.NewInlineKeyboardMarkup(
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
					tgbotapi.NewInlineKeyboardButtonData("Ver Excel", "Excel"),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("\xF0\x9F\x87\xAA\xF0\x9F\x87\xB8", string(constants.Const_langEs)),
					tgbotapi.NewInlineKeyboardButtonData("\xF0\x9F\x87\xBA\xF0\x9F\x87\xB8", string(constants.Const_langEn)),
				)),
		}
	}

	MiniKeyboard = tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Aceptar", string(constants.Const_ok)),
		tgbotapi.NewInlineKeyboardButtonData("Cancelar", string(constants.Const_back)),
	)
	CancelKeyboard = func(input constants.Inline) Step {
		return Step{
			Name:     "",
			Message:  "",
			Keyboard: tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Cancelar", string(constants.Const_back)))),
		}
	}
	EntradaMessage = func() Step {
		return Step{
			Name:    constants.Const_entrada,
			Message: constants.Textos[constants.Idioma][constants.Const_entrada],
			Keyboard: tgbotapi.NewInlineKeyboardMarkup(
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
				MiniKeyboard),
		}
	}

	VisitanteMessage = func() Step {
		return Step{
			Name:    constants.Const_visitante,
			Message: constants.Textos[constants.Idioma][constants.Const_visitante],
			Keyboard: tgbotapi.NewInlineKeyboardMarkup(
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
				MiniKeyboard),
		}
	}

	SalidaMessage = func() Step {
		return Step{
			Name:    constants.Const_salida,
			Message: constants.Textos[constants.Idioma][constants.Const_salida],
			Keyboard: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.Textos[constants.Idioma][constants.Const_sldFch], string(constants.Const_sldFch)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.Textos[constants.Idioma][constants.Const_sldHra], string(constants.Const_sldHra)),
				),
				MiniKeyboard),
		}
	}
)
