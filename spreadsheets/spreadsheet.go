package spreadsheet

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/heaptracetechnology/google-sheets/result"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
	"net/http"
	"os"
)

//ArgsData struct
type ArgsData struct {
	Title      string `json:"title"`
	ID         string `json:"spreadsheetId"`
	SheetID    int    `json:"sheetId"`
	SheetIndex int    `json:"sheetIndex"`
	SheetTitle string `json:"sheetTitle"`
	Row        int    `json:"row"`
	Column     int    `json:"column"`
	Content    string `json:"content"`
	Start      int    `json:"start"`
	End        int    `json:"end"`
}

//Message struct
type Message struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

//CreateSpreadsheet func
func CreateSpreadsheet(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("KEY")

	decodedJSON, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	decoder := json.NewDecoder(request.Body)

	var argsdata ArgsData
	decodeErr := decoder.Decode(&argsdata)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	conf, confErr := google.JWTConfigFromJSON(decodedJSON, spreadsheet.Scope)
	if confErr != nil {
		result.WriteErrorResponse(responseWriter, confErr)
		return
	}
	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)

	newSpreadsheet, _ := service.CreateSpreadsheet(spreadsheet.Spreadsheet{
		Properties: spreadsheet.Properties{
			Title: argsdata.Title,
		},
	})

	bytes, _ := json.Marshal(newSpreadsheet)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//FindSpreadsheet func
func FindSpreadsheet(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("KEY")

	decoder := json.NewDecoder(request.Body)

	var argsdata ArgsData
	decodeErr := decoder.Decode(&argsdata)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	decodedJSON, decodeErr := base64.StdEncoding.DecodeString(key)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}
	conf, confErr := google.JWTConfigFromJSON(decodedJSON, spreadsheet.Scope)
	if confErr != nil {
		result.WriteErrorResponse(responseWriter, confErr)
		return
	}
	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)

	spreadsheet, err := service.FetchSpreadsheet(argsdata.ID)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	bytes, _ := json.Marshal(spreadsheet)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//AddSheet func
func AddSheet(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("KEY")

	decoder := json.NewDecoder(request.Body)

	var argsdata ArgsData
	decodeErr := decoder.Decode(&argsdata)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	decodedJSON, decodeErr := base64.StdEncoding.DecodeString(key)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}
	conf, confErr := google.JWTConfigFromJSON(decodedJSON, spreadsheet.Scope)
	if confErr != nil {
		result.WriteErrorResponse(responseWriter, confErr)
		return
	}
	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)

	currentSpreadsheet, err := service.FetchSpreadsheet(argsdata.ID)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	var sheetProperties spreadsheet.SheetProperties
	sheetProperties.Title = argsdata.SheetTitle

	addSheetErr := service.AddSheet(&currentSpreadsheet, sheetProperties)
	if addSheetErr != nil {
		message := Message{false, addSheetErr.Error(), http.StatusBadRequest}
		bytes, _ := json.Marshal(message)
		result.WriteJsonResponse(responseWriter, bytes, http.StatusBadRequest)
		return
	}

	message := Message{true, "Sheet added successfully", http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//FindSheet func
func FindSheet(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("KEY")

	decoder := json.NewDecoder(request.Body)

	var argsdata ArgsData
	decodeErr := decoder.Decode(&argsdata)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	if argsdata.SheetID <= 0 && argsdata.SheetIndex <= 0 && argsdata.SheetTitle == "" {
		message := Message{false, "Please provide at least one argument(sheet Id, title or index)", http.StatusBadRequest}
		bytes, _ := json.Marshal(message)
		result.WriteJsonResponse(responseWriter, bytes, http.StatusBadRequest)
		return
	}

	decodedJSON, decodeErr := base64.StdEncoding.DecodeString(key)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}
	conf, confErr := google.JWTConfigFromJSON(decodedJSON, spreadsheet.Scope)
	if confErr != nil {
		result.WriteErrorResponse(responseWriter, confErr)
		return
	}
	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)

	currentSpreadsheet, err := service.FetchSpreadsheet(argsdata.ID)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	var sheet *spreadsheet.Sheet
	var sheetErr error
	if argsdata.SheetID > 0 {
		sheet, sheetErr = currentSpreadsheet.SheetByID(uint(argsdata.SheetID))
	} else if argsdata.SheetIndex > 0 {
		sheet, sheetErr = currentSpreadsheet.SheetByIndex(uint(argsdata.SheetIndex))
	} else if argsdata.SheetTitle != "" {
		sheet, sheetErr = currentSpreadsheet.SheetByTitle(argsdata.SheetTitle)
	} else if sheetErr != nil {
		result.WriteErrorResponse(responseWriter, confErr)
		return
	}

	bytes, _ := json.Marshal(sheet)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//UpdateSheetSize func
func UpdateSheetSize(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("KEY")

	decoder := json.NewDecoder(request.Body)

	var argsdata ArgsData
	decodeErr := decoder.Decode(&argsdata)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	decodedJSON, decodeErr := base64.StdEncoding.DecodeString(key)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}
	conf, confErr := google.JWTConfigFromJSON(decodedJSON, spreadsheet.Scope)
	if confErr != nil {
		result.WriteErrorResponse(responseWriter, confErr)
		return
	}
	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)

	currentSpreadsheet, err := service.FetchSpreadsheet(argsdata.ID)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	sheet, sheetErr := currentSpreadsheet.SheetByTitle(argsdata.SheetTitle)
	if sheetErr != nil {
		result.WriteErrorResponse(responseWriter, sheetErr)
		return
	}

	service.ExpandSheet(sheet, uint(argsdata.Row), uint(argsdata.Column))

	message := Message{true, "Sheet expanded successfully", http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//UpdateCell func
func UpdateCell(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("KEY")

	decoder := json.NewDecoder(request.Body)

	var argsdata ArgsData
	decodeErr := decoder.Decode(&argsdata)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	decodedJSON, decodeErr := base64.StdEncoding.DecodeString(key)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}
	conf, confErr := google.JWTConfigFromJSON(decodedJSON, spreadsheet.Scope)
	if confErr != nil {
		result.WriteErrorResponse(responseWriter, confErr)
		return
	}
	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)

	currentSpreadsheet, err := service.FetchSpreadsheet(argsdata.ID)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	sheet, sheetErr := currentSpreadsheet.SheetByTitle(argsdata.SheetTitle)
	if sheetErr != nil {
		result.WriteErrorResponse(responseWriter, sheetErr)
		return
	}

	sheet.Update(argsdata.Row, argsdata.Column, argsdata.Content)

	syncErr := sheet.Synchronize()
	if syncErr != nil {
		result.WriteErrorResponse(responseWriter, syncErr)
		return
	}

	contentMessage := fmt.Sprintf("Cell row[%d]column[%d] updated successfully with content '%s'", argsdata.Row, argsdata.Column, argsdata.Content)

	message := Message{true, contentMessage, http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//GetCell func
func GetCell(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("KEY")

	decoder := json.NewDecoder(request.Body)

	var argsdata ArgsData
	decodeErr := decoder.Decode(&argsdata)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	decodedJSON, decodeErr := base64.StdEncoding.DecodeString(key)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}
	conf, confErr := google.JWTConfigFromJSON(decodedJSON, spreadsheet.Scope)
	if confErr != nil {
		result.WriteErrorResponse(responseWriter, confErr)
		return
	}
	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)

	currentSpreadsheet, err := service.FetchSpreadsheet(argsdata.ID)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	sheet, sheetErr := currentSpreadsheet.SheetByTitle(argsdata.SheetTitle)
	if sheetErr != nil {
		result.WriteErrorResponse(responseWriter, sheetErr)
		return
	}

	value := sheet.Rows[argsdata.Row][argsdata.Column].Value

	syncErr := sheet.Synchronize()
	if syncErr != nil {
		result.WriteErrorResponse(responseWriter, syncErr)
		return
	}

	contentMessage := fmt.Sprintf("Cell row[%d]column[%d] contains '%s'", argsdata.Row, argsdata.Column, value)

	message := Message{true, contentMessage, http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//DeleteSheet func
func DeleteSheet(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("KEY")

	decoder := json.NewDecoder(request.Body)

	var argsdata ArgsData
	decodeErr := decoder.Decode(&argsdata)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	decodedJSON, decodeErr := base64.StdEncoding.DecodeString(key)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}
	conf, confErr := google.JWTConfigFromJSON(decodedJSON, spreadsheet.Scope)
	if confErr != nil {
		result.WriteErrorResponse(responseWriter, confErr)
		return
	}
	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)

	currentSpreadsheet, err := service.FetchSpreadsheet(argsdata.ID)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	deleteErr := service.DeleteSheet(&currentSpreadsheet, uint(argsdata.SheetID))
	if deleteErr != nil {
		message := Message{false, deleteErr.Error(), http.StatusBadRequest}
		bytes, _ := json.Marshal(message)
		result.WriteJsonResponse(responseWriter, bytes, http.StatusBadRequest)
		return
	}

	message := Message{true, "Sheet deleted successfully", http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
