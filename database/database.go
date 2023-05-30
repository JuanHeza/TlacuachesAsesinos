package database

import (
	"TlacuachesAsesinos/constants"
	"TlacuachesAsesinos/model"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

var (
	sp    spreadsheet.Spreadsheet
	sheet *spreadsheet.Sheet
)

/*
func main() {
	//connect()
	sheet, err := sp.SheetByIndex(0)
	checkError(err)
	for y, row := range sheet.Rows {
		for x, cell := range row {
			fmt.Println(x, y, cell.Value)
		}
	}

	// Update cell content (row, column)
	sheet.Update(1, 0, "hogehoge")

	// Make sure call Synchronize to reflect the changes
	err = sheet.Synchronize()
	checkError(err,"")
}
*/
func checkError(err error, errString string) {
	if err != nil {
		log.Println(errString)
		panic(err.Error())
	}
}

func getCell(x, y int, sheet spreadsheet.Sheet) (cell spreadsheet.Cell) {
	return sheet.Rows[x][y]
}

func Search(index string) (rg model.Registro) {
	row, err := strconv.Atoi(index)
	if err != nil {
		return
	}
	folio, err := strconv.Atoi(sheet.Rows[row][0].Value)
	if err == nil {
		rg.Folio = folio
		rg.Nombre = sheet.Rows[row][1].Value
		rg.Motivo = sheet.Rows[row][2].Value
		rg.Company = sheet.Rows[row][3].Value
		rg.Calle = sheet.Rows[row][4].Value
		rg.NumeroExterior, _ = strconv.Atoi(sheet.Rows[row][5].Value)
		rg.Identificacion = sheet.Rows[row][6].Value
		rg.AutoFabricante = sheet.Rows[row][7].Value
		rg.Color = sheet.Rows[row][8].Value
		rg.FotoVehiculo = sheet.Rows[row][9].Value
		rg.FechaEntrada, _ = time.Parse(constants.Const_date_template, sheet.Rows[row][10].Value)
		rg.HoraEntrada, _ = time.Parse(constants.Const_time_template, sheet.Rows[row][11].Value)
		rg.FechaSalida, _ = time.Parse(constants.Const_date_template, sheet.Rows[row][12].Value)
		rg.HoraSalida, _ = time.Parse(constants.Const_time_template, sheet.Rows[row][13].Value)
		rg.Observaciones = sheet.Rows[row][14].Value
	}
	return
}

func Save(rg model.Registro, pantalla int) (string, int) {
	dateS, dateE, clockS, clockE, numeroExterior := "", "", "", "", ""

	yS, mS, dS := rg.FechaSalida.Date()
	hS, minS, sS := rg.HoraSalida.Clock()

	yE, mE, dE := rg.FechaEntrada.Date()
	hE, minE, sE := rg.HoraEntrada.Clock()

	if !rg.FechaSalida.IsZero() {
		dateS = fmt.Sprintf("%02d/%02d/%04d", dS, mS, yS)
	}
	if !rg.FechaSalida.IsZero() {
		dateE = fmt.Sprintf("%02d/%02d/%04d", dE, mE, yE)
	}
	if !rg.HoraSalida.IsZero() {
		clockS = fmt.Sprintf("%02d:%02d:%02d", hS, minS, sS)
	}
	if !rg.HoraSalida.IsZero() {
		clockE = fmt.Sprintf("%02d:%02d:%02d", hE, minE, sE)
	}
	if rg.NumeroExterior != 0 {
		numeroExterior = fmt.Sprint(rg.NumeroExterior)
	}

	if rg.Folio == 0 {
		rg.Folio = len(sheet.Data.GridData[0].RowData)
	}
	sheet.Update(rg.Folio, 0, fmt.Sprintf("%05d", rg.Folio))
	switch pantalla {
	case 0:
		sheet.Update(rg.Folio, 1, rg.Nombre)
		sheet.Update(rg.Folio, 2, rg.Motivo)
		sheet.Update(rg.Folio, 3, rg.Company)
		sheet.Update(rg.Folio, 4, rg.Calle)
		sheet.Update(rg.Folio, 5, numeroExterior)
	case 1:
		sheet.Update(rg.Folio, 6, rg.Identificacion)
		sheet.Update(rg.Folio, 7, rg.AutoFabricante)
		sheet.Update(rg.Folio, 8, rg.Color)
		sheet.Update(rg.Folio, 9, rg.FotoVehiculo)
		sheet.Update(rg.Folio, 10, dateE)
		sheet.Update(rg.Folio, 11, clockE)
		sheet.Update(rg.Folio, 14, rg.Observaciones)
	case 2:
		sheet.Update(rg.Folio, 12, dateS)
		sheet.Update(rg.Folio, 13, clockS)
	}
	err := sheet.Synchronize()
	checkError(err, "Error on Update/Save")
	return fmt.Sprintf("%09d - %v", rg.Folio, rg.Nombre), rg.Folio
}

func Connect() {
	data, err := ioutil.ReadFile("./database/token.json")
	checkError(err, "file not found")
	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	checkError(err, fmt.Sprintf("%s", data))
	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)
	spreadsheet, err := service.FetchSpreadsheet("1ytQV2XIoxR2TdiBiBIARRx5mg1v48kk7VF74eZ4LanQ")
	checkError(err, "No Sspreadsheet")
	sp = spreadsheet

	sheet, err = sp.SheetByIndex(0)
	checkError(err, "No Sheet")
}
