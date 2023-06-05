package model

import (
	"TlacuachesAsesinos/constants"
	//"TlacuachesAsesinos/database"
	"fmt"

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
			Message: constants.GetValue(constants.Const_home),
			Keyboard: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_qrCode), string(constants.Const_qrCode)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_rgtDir), string(constants.Const_rgtDir)),
				),
				//				tgbotapi.NewInlineKeyboardRow(
				//					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_rgtVst), string(constants.Const_rgtVst)),
				//				),
				//				tgbotapi.NewInlineKeyboardRow(
				//					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_rgtSld), string(constants.Const_rgtSld)),
				//				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Buscar Registro", "buscar"),
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

	MiniKeyboard = func(accept bool) []tgbotapi.InlineKeyboardButton {
		if accept {
			return tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Aceptar", string(constants.Const_ok)),
				tgbotapi.NewInlineKeyboardButtonData("Cancelar", string(constants.Const_back)),
			)
		} else {
			return tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Cancelar", string(constants.Const_back)),
			)
		}
	}

	SearchKeyboard = func(folio string) []tgbotapi.InlineKeyboardButton {
		return tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Ver Registro", fmt.Sprint("Folio - ", folio)),
		)
	}

	FolioKeyboard = func() Step {
		return Step{
			Name:    "",
			Message: "",
			Keyboard: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_rgtVst), string(constants.Const_rgtVst)),
					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_rgtSld), string(constants.Const_rgtSld)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_entFvh), string(constants.Const_verFvh)),
					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_entFid), string(constants.Const_verFid)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_back), string(constants.Const_back)),
				),
			),
		}
	}

	CancelKeyboard = func() Step {
		return Step{
			Name:    "",
			Message: "",
			Keyboard: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_back), string(constants.Const_back)),
				),
			),
		}
	}

	VisitanteMessage = func() Step {
		return Step{
			Name:    constants.Const_entrada,
			Message: constants.GetValue(constants.Const_entrada),
			Keyboard: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_rgtNom), string(constants.Const_rgtNom)),
					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_rgtCom), string(constants.Const_rgtCom)),
					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_rgtMot), string(constants.Const_rgtMot)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_rgtCll), string(constants.Const_rgtCll)),
					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_rgtExt), string(constants.Const_rgtExt)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Aceptar", string(constants.Const_ok)),
					tgbotapi.NewInlineKeyboardButtonData("Cancelar", string(constants.Const_back)),
				),
			),
		}
	}

	EntradaMessage = func() Step {
		return Step{
			Name:    constants.Const_visitante,
			Message: constants.GetValue(constants.Const_visitante),
			Keyboard: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_entFab), string(constants.Const_entFab)),
					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_entCol), string(constants.Const_entCol)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_entFch), string(constants.Const_entFch)),
					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_entHra), string(constants.Const_entHra)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_entFid), string(constants.Const_entFid)),
					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_entFvh), string(constants.Const_entFvh)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_entObs), string(constants.Const_entObs)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Aceptar", string(constants.Const_okEnt)),
					tgbotapi.NewInlineKeyboardButtonData("Cancelar", string(constants.Const_back)),
				),
			),
		}
	}

	SalidaMessage = func() Step {
		return Step{
			Name:    constants.Const_salida,
			Message: constants.GetValue(constants.Const_salida),
			Keyboard: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_sldFch), string(constants.Const_sldFch)),
					tgbotapi.NewInlineKeyboardButtonData(constants.GetValue(constants.Const_sldHra), string(constants.Const_sldHra)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Aceptar", string(constants.Const_okSal)),
					tgbotapi.NewInlineKeyboardButtonData("Cancelar", string(constants.Const_back)),
				),
			),
		}
	}

	BusquedaMessage = func() Step {
		return Step{
			Name:    constants.Const_salida,
			Message: constants.GetValue(constants.Const_salida),
			Keyboard: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Folio", "Folio - 0"),
					tgbotapi.NewInlineKeyboardButtonData("Nombre", "Nombre"),
					tgbotapi.NewInlineKeyboardButtonData("Lista", "Lista"),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Cancelar", string(constants.Const_back)),
				),
			),
		}
	}

	ListaMessage = func(prev int, next int, data [][]string) Step {
		var da = tgbotapi.NewInlineKeyboardMarkup()
		for _, button := range data {
			da.InlineKeyboard = append(da.InlineKeyboard,
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(button[0], button[1]),
				),
			)
		}
		pages := tgbotapi.NewInlineKeyboardRow()
		if prev != -1 {
			pages = append(pages, tgbotapi.NewInlineKeyboardButtonData("Anterior", fmt.Sprint("Page - ", prev)))
		}
		if next != -1 {
			pages = append(pages, tgbotapi.NewInlineKeyboardButtonData("Siguente", fmt.Sprint("Page - ", next)))
		}
		da.InlineKeyboard = append(da.InlineKeyboard, pages)
		da.InlineKeyboard = append(da.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Cancelar", string(constants.Const_back)),
		))
		return Step{
			Name:     "lista",
			Message:  "Seleccione un folio de la siguiente lista",
			Keyboard: da,
		}
	}
)
