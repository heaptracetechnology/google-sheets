package spreadsheet

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/heaptracetechnology/google-sheets/result"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
	"io/ioutil"
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
}

type Message struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

//CreateSpreadsheet func
func CreateSpreadsheet(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("KEY")

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	defer request.Body.Close()

	var argsdata ArgsData
	er := json.Unmarshal(body, &argsdata)
	if er != nil {
		result.WriteErrorResponse(responseWriter, er)
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

	newSpreadsheet, sheetsErr := service.CreateSpreadsheet(spreadsheet.Spreadsheet{
		Properties: spreadsheet.Properties{
			Title: argsdata.Title,
		},
	})
	if sheetsErr != nil {
		result.WriteErrorResponse(responseWriter, sheetsErr)
		return
	}

	bytes, _ := json.Marshal(newSpreadsheet)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//FindSpreadsheet func
func FindSpreadsheet(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("KEY")

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	defer request.Body.Close()

	var argsdata ArgsData
	er := json.Unmarshal(body, &argsdata)
	if er != nil {
		result.WriteErrorResponse(responseWriter, er)
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

	bytes, _ := json.Marshal(spreadsheet)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//AddSheet func
func AddSheet(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("KEY")

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	defer request.Body.Close()

	var argsdata ArgsData
	er := json.Unmarshal(body, &argsdata)
	if er != nil {
		result.WriteErrorResponse(responseWriter, er)
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

	var sheetProperties spreadsheet.SheetProperties
	sheetProperties.Title = argsdata.Title

	fmt.Println("sheetProperties :::", sheetProperties)

	res, _ := json.Marshal(sheetProperties)
	fmt.Println("Response ::", string(res))

	sheetErr := service.AddSheet(&currentSpreadsheet, sheetProperties)
	if sheetErr != nil {
		message := Message{false, sheetErr.Error(), http.StatusBadRequest}
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

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	defer request.Body.Close()

	var argsdata ArgsData
	er := json.Unmarshal(body, &argsdata)
	if er != nil {
		result.WriteErrorResponse(responseWriter, er)
		return
	}
	fmt.Println("argsdata.SheetIndex ::", argsdata.SheetIndex)

	if argsdata.SheetID < 0 && argsdata.SheetIndex < 0 && argsdata.SheetTitle == "" {
		message := Message{false, "Please provide valid data", http.StatusBadRequest}
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

	var sheet *spreadsheet.Sheet
	var sheetErr error
	if argsdata.SheetID >= 0 {
		sheet, sheetErr = currentSpreadsheet.SheetByID(uint(argsdata.SheetID))
	} else if argsdata.SheetIndex >= 0 {
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
