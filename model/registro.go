package model

import (
	"TlacuachesAsesinos/constants"
	"fmt"
	"strconv"
	"time"
)

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

func (rg *Registro) PrintEntradaMesage(msg string) string {
	return fmt.Sprintf("%s\n\n%s:\n%s\n\n%s:\n%s\n\n%s:\n%s\n\n%s:\n%s\n\n%s:\n%v", msg, constants.Textos[constants.Idioma][constants.Const_rgtNom], rg.Nombre, constants.Textos[constants.Idioma][constants.Const_rgtCom], rg.Company, constants.Textos[constants.Idioma][constants.Const_rgtMot], rg.Motivo, constants.Textos[constants.Idioma][constants.Const_rgtCll], rg.Calle, constants.Textos[constants.Idioma][constants.Const_rgtExt], rg.NumeroExterior)
}

func (rg *Registro) PrintSalidaMesage(msg string) string {
	y, m, d := rg.FechaSalida.Date()
	h, min, s := rg.HoraSalida.Clock()
	date, hour := "\n", "\n"
	if !rg.FechaSalida.IsZero() {
		date = fmt.Sprintf("%02d/%v/%04d", d, constants.PrintMes(int(m)), y)
	}
	if !rg.HoraSalida.IsZero() {
		hour = fmt.Sprintf("%02d:%02d:%02d", h, min, s)
	}
	return fmt.Sprintf("%s\n\n%s:\n%s\n\n%s:\n%s", msg, constants.Textos[constants.Idioma][constants.Const_sldFch], date, constants.Textos[constants.Idioma][constants.Const_sldHra], hour)
}

func (rg *Registro) FillRegistro(campo constants.Inline, valor interface{}) {
	switch campo {
	case constants.Const_rgtNom:
		rg.Nombre = fmt.Sprintf("%v", valor)
	case constants.Const_rgtCom:
		rg.Company = fmt.Sprintf("%v", valor)
	case constants.Const_rgtCll:
		rg.Calle = fmt.Sprintf("%v", valor)
	case constants.Const_rgtExt:
		num, err := strconv.Atoi(fmt.Sprintf("%v", valor))
		if err == nil {
			rg.NumeroExterior = num
		}
	case constants.Const_rgtMot:
		rg.Motivo = fmt.Sprintf("%v", valor)
	}
}
