package database

import (
	"fmt"
	"io/ioutil"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

func main() {
	data, err := ioutil.ReadFile("keys.json")
	checkError(err)
	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	checkError(err)
	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)
	spreadsheet, err := service.FetchSpreadsheet("1ytQV2XIoxR2TdiBiBIARRx5mg1v48kk7VF74eZ4LanQ")
	checkError(err)
	sheet, err := spreadsheet.SheetByIndex(0)
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
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func getCell(x, y int, sheet spreadsheet.Sheet)(cell spreadsheet.Cell){
	return sheet.Rows[x][y]
}