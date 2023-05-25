package database

import (
	"TlacuachesAsesinos/constants"
	"TlacuachesAsesinos/model"
	"fmt"
    "log"
	"io/ioutil"
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
func checkError(err error, errString string ) {
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
		rg.Folio = int64(folio)
		rg.Nombre = sheet.Rows[row][1].Value
		rg.Motivo = sheet.Rows[row][2].Value
		rg.Company = sheet.Rows[row][3].Value
		rg.Calle = sheet.Rows[row][4].Value
		rg.NumeroExterior = 9999999999999999
		rg.Identificacion = sheet.Rows[row][6].Value
		rg.AutoFabricante = sheet.Rows[row][7].Value
		rg.Color = sheet.Rows[row][8].Value
		rg.FotoVehiculo = sheet.Rows[row][9].Value
		rg.FechaEntrada = time.Now()
		rg.HoraEntrada = time.Now()
		rg.FechaSalida = time.Now()
		rg.HoraSalida = time.Now()
		rg.Observaciones = sheet.Rows[row][14].Value
	}
	return
}

func Save(rg model.Registro, row int) (string, int) {
	dateS, dateE, clockS, clockE, numeroExterior := "", "", "", "", ""

	yS, mS, dS := rg.FechaSalida.Date()
	hS, minS, sS := rg.HoraSalida.Clock()

	yE, mE, dE := rg.FechaEntrada.Date()
	hE, minE, sE := rg.HoraEntrada.Clock()

	if !rg.FechaSalida.IsZero() {
		dateS = fmt.Sprintf("%02d/%v/%04d", dS, constants.PrintMes(int(mS)), yS)
	}
	if !rg.FechaSalida.IsZero() {
		dateE = fmt.Sprintf("%02d/%v/%04d", dE, constants.PrintMes(int(mE)), yE)
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

	if row == 0 {
		row = len(sheet.Data.GridData[0].RowData)
	}
	sheet.Update(row, 0, fmt.Sprintf("%05d", row))
	sheet.Update(row, 1, rg.Nombre)
	sheet.Update(row, 2, rg.Motivo)
	sheet.Update(row, 3, rg.Company)
	sheet.Update(row, 4, rg.Calle)
	sheet.Update(row, 5, numeroExterior)
	sheet.Update(row, 6, rg.Identificacion)
	sheet.Update(row, 7, rg.AutoFabricante)
	sheet.Update(row, 8, rg.Color)
	sheet.Update(row, 9, rg.FotoVehiculo)
	sheet.Update(row, 10, dateE)
	sheet.Update(row, 11, clockE)
	sheet.Update(row, 12, dateS)
	sheet.Update(row, 13, clockS)
	sheet.Update(row, 14, rg.Observaciones)
	err := sheet.Synchronize()
	checkError(err, "Error on Update/Save")
	return fmt.Sprintf("%09d - %v", row, rg.Nombre), row
}

func Connect() {
	data, err := ioutil.ReadFile("./database/token.json")
	checkError(err,"file not found")
	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	checkError(err,fmt.Sprintf("%s",data))
	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)
	spreadsheet, err := service.FetchSpreadsheet("1ytQV2XIoxR2TdiBiBIARRx5mg1v48kk7VF74eZ4LanQ")
	checkError(err,"No Sspreadsheet")
	sp = spreadsheet

	sheet, err = sp.SheetByIndex(0)
	checkError(err, "No Sheet")
}
