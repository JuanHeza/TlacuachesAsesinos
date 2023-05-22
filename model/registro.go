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

func (rg *Registro) PrintRegistroMesage(msg string) string {
	num := fmt.Sprint(rg.NumeroExterior)
	if rg.NumeroExterior == 0 {
		num = " "
	}
	return fmt.Sprintf("%s\n\n*%s*:   %s\n\n*%s*:   %s\n\n*%s*:   %s\n\n*%s*:   %s\n\n*%s*:   %v",
		msg,
		constants.GetValue(constants.Const_rgtNom), rg.Nombre,
		constants.GetValue(constants.Const_rgtCom), rg.Company,
		constants.GetValue(constants.Const_rgtMot), rg.Motivo,
		constants.GetValue(constants.Const_rgtCll), rg.Calle,
		constants.GetValue(constants.Const_rgtExt), num)
}

func (rg *Registro) Print() string {
	return fmt.Sprintf("*Folio:* _%09d_ %v %v %v", rg.Folio, rg.PrintRegistroMesage(""), rg.PrintEntradaMesage(""), rg.PrintSalidaMesage(""))
}

func (rg *Registro) PrintSalidaMesage(msg string) string {
	y, m, d := rg.FechaSalida.Date()
	h, min, s := rg.HoraSalida.Clock()
	date, hour := " ", " "
	if !rg.FechaSalida.IsZero() {
		date = fmt.Sprintf("%02d/%v/%04d", d, constants.PrintMes(int(m)), y)
	}
	if !rg.HoraSalida.IsZero() {
		hour = fmt.Sprintf("%02d:%02d:%02d", h, min, s)
	}
	return fmt.Sprintf("%s\n\n*%s*:   %s\n\n*%s*:   %s",
		msg,
		constants.GetValue(constants.Const_sldFch), date,
		constants.GetValue(constants.Const_sldHra), hour)
}

func (rg *Registro) PrintEntradaMesage(msg string) string {
	obser := ""
	if rg.Observaciones != "" {
		obser = rg.Observaciones
	}
	y, m, d := rg.FechaEntrada.Date()
	h, min, s := rg.HoraEntrada.Clock()
	date, hour := " ", " "
	if !rg.FechaEntrada.IsZero() {
		date = fmt.Sprintf("%02d/%v/%04d", d, constants.PrintMes(int(m)), y)
	}
	if !rg.HoraEntrada.IsZero() {
		hour = fmt.Sprintf("%02d:%02d:%02d", h, min, s)
	}
	return fmt.Sprintf("%s\n\n*%s*:   %s\n\n*%s*:   %s\n\n*%s*:   %s\n\n*%s*:   %s\n\n*%s*:  \n%s",
		msg,
		constants.GetValue(constants.Const_entFab), rg.AutoFabricante,
		constants.GetValue(constants.Const_entCol), rg.Color,
		constants.GetValue(constants.Const_entFch), date,
		constants.GetValue(constants.Const_entHra), hour,
		constants.GetValue(constants.Const_entObs), obser)
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

func (rg *Registro) CheckEntrada() bool{
	if !(checkDate(rg.HoraEntrada) && checkDate(rg.FechaEntrada)) {
		return false
	}
	return rg.AutoFabricante == "" && rg.Color == "" && rg.Observaciones == "" && rg.Identificacion == "" && rg.FotoVehiculo == ""
}
func (rg *Registro) CheckSalida() bool{
	return checkDate(rg.HoraSalida) && checkDate(rg.FechaSalida)
}

func checkDate(input time.Time) bool {
	return input.IsZero()
}
