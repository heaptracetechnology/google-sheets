package spreadsheets

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/cloudevents/sdk-go"
	"github.com/heaptracetechnology/google-sheets/result"
	"golang.org/x/oauth2/google"
	driveV3 "google.golang.org/api/drive/v3"
	sheetsV4 "google.golang.org/api/sheets/v4"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//ArgsData struct
type ArgsData struct {
	RowLength    *sheetsV4.Request `json:"rowLength"`
	ColumnLength *sheetsV4.Request `json:"columnLength"`
	Title        string            `json:"title"`
	ID           string            `json:"spreadsheetId"`
	SheetID      int64             `json:"sheetId"`
	SheetIndex   int               `json:"sheetIndex"`
	SheetTitle   string            `json:"sheetTitle"`
	Row          int64             `json:"row"`
	Column       int64             `json:"column"`
	Content      string            `json:"content"`
	Start        int               `json:"start"`
	End          int               `json:"end"`
	EmailAddress string            `json:"emailAddress"`
	Role         string            `json:"role"`
	Type         string            `json:"type"`
	CellNumber   string            `json:"cellNumber"`
}

//Subscribe struct
type Subscribe struct {
	Data      RequestParam `json:"data"`
	Endpoint  string       `json:"endpoint"`
	ID        string       `json:"id"`
	IsTesting bool         `json:"istesting"`
}

//SubscribeReturn struct
type SubscribeReturn struct {
	SpreadsheetID string `json:"spreadsheetID"`
	SheetTitle    string `json:"sheetTitle"`
	TwitterCell   string `json:"twitterCell"`
	EmailAddress  string `json:"emailAddress"`
}

//RequestParam struct
type RequestParam struct {
	SpreadsheetID string `json:"spreadsheetID"`
	SheetTitle    string `json:"sheetTitle"`
}

//Message struct
type Message struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

//SheetScope Spreadsheet
const (
	SheetScope = "https://www.googleapis.com/auth/spreadsheets"
	DriveScope = "https://www.googleapis.com/auth/drive.file"
)

//Global Variables
var (
	Listener        = make(map[string]Subscribe)
	rtmStarted      bool
	sheetService    *sheetsV4.Service
	sheetServiceErr error
	oldRowCount     int
	twitterIndex    int
	subReturn       SubscribeReturn
	count           int
)

//HealthCheck Google-Sheets
func HealthCheck(responseWriter http.ResponseWriter, request *http.Request) {

	bytes, _ := json.Marshal("OK")
	result.WriteJSONResponse(responseWriter, bytes, http.StatusOK)
}

//CreateSpreadsheet func
func CreateSpreadsheet(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("CREDENTIAL_JSON")

	decodedJSON, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		result.WriteErrorResponseString(responseWriter, err.Error())
		return
	}

	decoder := json.NewDecoder(request.Body)

	var argsdata ArgsData
	decodeErr := decoder.Decode(&argsdata)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	sheetConf, sheetConfErr := google.JWTConfigFromJSON(decodedJSON, SheetScope)
	if sheetConfErr != nil {
		result.WriteErrorResponseString(responseWriter, sheetConfErr.Error())
		return
	}

	sheetClient := sheetConf.Client(context.TODO())

	sheetService, sheetServiceErr := sheetsV4.New(sheetClient)
	if sheetServiceErr != nil {
		result.WriteErrorResponseString(responseWriter, sheetServiceErr.Error())
		return
	}

	sheetProperties := sheetsV4.Spreadsheet{
		Properties: &sheetsV4.SpreadsheetProperties{
			Title: argsdata.Title,
		},
	}

	newSpreadsheet := sheetService.Spreadsheets.Create(&sheetProperties)
	spreadsheet, sheetErr := newSpreadsheet.Do()
	if sheetErr != nil {
		result.WriteErrorResponseString(responseWriter, sheetErr.Error())
		return
	}

	spreadsheetID := spreadsheet.SpreadsheetId

	driveConf, driveConfErr := google.JWTConfigFromJSON(decodedJSON, DriveScope)
	if driveConfErr != nil {
		result.WriteErrorResponseString(responseWriter, driveConfErr.Error())
		return
	}

	driveClient := driveConf.Client(context.TODO())

	driveService, driveServiceErr := driveV3.New(driveClient)
	if driveServiceErr != nil {
		result.WriteErrorResponseString(responseWriter, driveServiceErr.Error())
		return
	}

	driveProperties := driveV3.Permission{
		EmailAddress: argsdata.EmailAddress,
		Role:         argsdata.Role,
		Type:         argsdata.Type,
	}

	if spreadsheetID != "" {
		permission := driveService.Permissions.Create(spreadsheetID, &driveProperties)
		_, doErr := permission.Do()
		if doErr != nil {
			result.WriteErrorResponseString(responseWriter, doErr.Error())
			return
		}
	} else {
		result.WriteErrorResponseString(responseWriter, "SpreadSheet ID not found")
		return
	}

	bytes, _ := json.Marshal(spreadsheet)
	result.WriteJSONResponse(responseWriter, bytes, http.StatusOK)
}

//FindSpreadsheet func
func FindSpreadsheet(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("CREDENTIAL_JSON")

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
	sheetConf, sheetConfErr := google.JWTConfigFromJSON(decodedJSON, SheetScope)
	if sheetConfErr != nil {
		result.WriteErrorResponseString(responseWriter, sheetConfErr.Error())
		return
	}

	sheetClient := sheetConf.Client(context.TODO())

	sheetService, sheetServiceErr := sheetsV4.New(sheetClient)
	if sheetServiceErr != nil {
		result.WriteErrorResponseString(responseWriter, sheetServiceErr.Error())
		return
	}

	getSpreadsheet := sheetService.Spreadsheets.Get(argsdata.ID)
	spreadsheet, sheetErr := getSpreadsheet.Do()
	if sheetErr != nil {
		result.WriteErrorResponseString(responseWriter, sheetErr.Error())
		return
	}

	bytes, _ := json.Marshal(spreadsheet)
	result.WriteJSONResponse(responseWriter, bytes, http.StatusOK)
}

//AddSheet func
func AddSheet(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("CREDENTIAL_JSON")

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

	sheetConf, sheetConfErr := google.JWTConfigFromJSON(decodedJSON, SheetScope)
	if sheetConfErr != nil {
		result.WriteErrorResponseString(responseWriter, sheetConfErr.Error())
		return
	}

	sheetClient := sheetConf.Client(context.TODO())

	sheetService, sheetServiceErr := sheetsV4.New(sheetClient)
	if sheetServiceErr != nil {
		result.WriteErrorResponseString(responseWriter, sheetServiceErr.Error())
		return
	}

	addSheet := sheetsV4.BatchUpdateSpreadsheetRequest{
		Requests: []*sheetsV4.Request{
			&sheetsV4.Request{
				AddSheet: &sheetsV4.AddSheetRequest{
					Properties: &sheetsV4.SheetProperties{
						Title: argsdata.SheetTitle,
					},
				},
			},
		},
	}

	addSpreadsheet := sheetService.Spreadsheets.BatchUpdate(argsdata.ID, &addSheet)
	spreadsheet, sheetErr := addSpreadsheet.Do()
	if sheetErr != nil {
		result.WriteErrorResponseString(responseWriter, sheetErr.Error())
		return
	}

	bytes, _ := json.Marshal(spreadsheet)
	result.WriteJSONResponse(responseWriter, bytes, http.StatusOK)
}

//FindSheet func
func FindSheet(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("CREDENTIAL_JSON")

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
		result.WriteJSONResponse(responseWriter, bytes, http.StatusBadRequest)
		return
	}

	decodedJSON, decodeErr := base64.StdEncoding.DecodeString(key)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	sheetConf, sheetConfErr := google.JWTConfigFromJSON(decodedJSON, SheetScope)
	if sheetConfErr != nil {
		result.WriteErrorResponseString(responseWriter, sheetConfErr.Error())
		return
	}

	sheetClient := sheetConf.Client(context.TODO())

	sheetService, sheetServiceErr := sheetsV4.New(sheetClient)
	if sheetServiceErr != nil {
		result.WriteErrorResponseString(responseWriter, sheetServiceErr.Error())
		return
	}

	getSheet := sheetService.Spreadsheets.Values.Get(argsdata.ID, argsdata.SheetTitle)
	sheet, sheetErr := getSheet.Do()
	if sheetErr != nil {
		result.WriteErrorResponseString(responseWriter, sheetErr.Error())
		return
	}

	bytes, _ := json.Marshal(sheet)
	result.WriteJSONResponse(responseWriter, bytes, http.StatusOK)
}

//UpdateSheetSize func
func UpdateSheetSize(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("CREDENTIAL_JSON")

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

	sheetConf, sheetConfErr := google.JWTConfigFromJSON(decodedJSON, SheetScope)
	if sheetConfErr != nil {
		result.WriteErrorResponseString(responseWriter, sheetConfErr.Error())
		return
	}

	sheetClient := sheetConf.Client(context.TODO())

	sheetService, sheetServiceErr := sheetsV4.New(sheetClient)
	if sheetServiceErr != nil {
		result.WriteErrorResponseString(responseWriter, sheetServiceErr.Error())
		return
	}

	resizeValues := sheetsV4.BatchUpdateSpreadsheetRequest{
		Requests: []*sheetsV4.Request{
			&sheetsV4.Request{
				AppendDimension: &sheetsV4.AppendDimensionRequest{
					Length:    argsdata.Row,
					Dimension: "ROWS",
					SheetId:   argsdata.SheetID,
				},
			},
			&sheetsV4.Request{
				AppendDimension: &sheetsV4.AppendDimensionRequest{
					Length:    argsdata.Column,
					Dimension: "COLUMNS",
					SheetId:   argsdata.SheetID,
				},
			},
		},
	}

	resizeSheet := sheetService.Spreadsheets.BatchUpdate(argsdata.ID, &resizeValues)
	sheet, sheetErr := resizeSheet.Do()
	if sheetErr != nil {
		result.WriteErrorResponseString(responseWriter, sheetErr.Error())
		return
	}

	bytes, _ := json.Marshal(sheet)
	result.WriteJSONResponse(responseWriter, bytes, http.StatusOK)
}

//UpdateCell func
func UpdateCell(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("CREDENTIAL_JSON")

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
	sheetConf, sheetConfErr := google.JWTConfigFromJSON(decodedJSON, SheetScope)
	if sheetConfErr != nil {
		result.WriteErrorResponseString(responseWriter, sheetConfErr.Error())
		return
	}

	sheetClient := sheetConf.Client(context.TODO())

	sheetService, sheetServiceErr := sheetsV4.New(sheetClient)
	if sheetServiceErr != nil {
		result.WriteErrorResponseString(responseWriter, sheetServiceErr.Error())
		return
	}

	writeProp := sheetsV4.ValueRange{
		MajorDimension: "ROWS",
		Values:         [][]interface{}{{argsdata.Content}},
	}

	writeSheet := sheetService.Spreadsheets.Values.Update(argsdata.ID, argsdata.SheetTitle+"!"+argsdata.CellNumber, &writeProp)
	writeSheet.ValueInputOption("USER_ENTERED")
	sheet, sheetErr := writeSheet.Do()
	if sheetErr != nil {
		result.WriteErrorResponseString(responseWriter, sheetErr.Error())
		return
	}

	bytes, _ := json.Marshal(sheet)
	result.WriteJSONResponse(responseWriter, bytes, http.StatusOK)
}

//DeleteSheet func
func DeleteSheet(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("CREDENTIAL_JSON")

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

	sheetConf, sheetConfErr := google.JWTConfigFromJSON(decodedJSON, SheetScope)
	if sheetConfErr != nil {
		result.WriteErrorResponseString(responseWriter, sheetConfErr.Error())
		return
	}

	sheetClient := sheetConf.Client(context.TODO())

	sheetService, sheetServiceErr := sheetsV4.New(sheetClient)
	if sheetServiceErr != nil {
		result.WriteErrorResponseString(responseWriter, sheetServiceErr.Error())
		return
	}

	deleteProperties := sheetsV4.BatchUpdateSpreadsheetRequest{
		Requests: []*sheetsV4.Request{
			&sheetsV4.Request{
				DeleteSheet: &sheetsV4.DeleteSheetRequest{
					SheetId: argsdata.SheetID,
				},
			},
		},
	}

	deleteSheet := sheetService.Spreadsheets.BatchUpdate(argsdata.ID, &deleteProperties)
	_, sheetErr := deleteSheet.Do()
	if sheetErr != nil {
		result.WriteErrorResponseString(responseWriter, sheetErr.Error())
		return
	}

	message := Message{true, "Sheet deleted successfully", http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJSONResponse(responseWriter, bytes, http.StatusOK)
}

//SheetSubscribe func
func SheetSubscribe(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("CREDENTIAL_JSON")

	decodedJSON, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		result.WriteErrorResponseString(responseWriter, err.Error())
		return
	}

	decoder := json.NewDecoder(request.Body)

	var sub Subscribe
	decodeError := decoder.Decode(&sub)
	if decodeError != nil {
		result.WriteErrorResponseString(responseWriter, decodeError.Error())
		return
	}

	sheetConf, sheetConfErr := google.JWTConfigFromJSON(decodedJSON, SheetScope)
	if sheetConfErr != nil {
		result.WriteErrorResponseString(responseWriter, sheetConfErr.Error())
		return
	}

	sheetClient := sheetConf.Client(context.TODO())

	sheetService, sheetServiceErr = sheetsV4.New(sheetClient)
	if sheetServiceErr != nil {
		result.WriteErrorResponseString(responseWriter, sheetServiceErr.Error())
		return
	}

	Listener[sub.Data.SpreadsheetID] = sub
	if !rtmStarted {
		go SheetRTM()
		rtmStarted = true
	}

	bytes, _ := json.Marshal("Subscribed")
	result.WriteJSONResponse(responseWriter, bytes, http.StatusOK)
}

//SheetRTM func
func SheetRTM() {
	isTest := false
	for {
		if len(Listener) > 0 {
			for k, v := range Listener {
				go getNewRowUpdate(k, v)
				isTest = v.IsTesting
			}
		} else {
			rtmStarted = false
			break
		}
		time.Sleep(10 * time.Second)
		if isTest {
			break
		}
	}
}

func getNewRowUpdate(spreadsheetID string, sub Subscribe) {

	subReturn.SpreadsheetID = spreadsheetID
	subReturn.SheetTitle = sub.Data.SheetTitle

	readSheet := sheetService.Spreadsheets.Values.Get(spreadsheetID, sub.Data.SheetTitle)
	sheet, readSheetErr := readSheet.Do()
	if readSheetErr != nil {
		fmt.Println("Read Sheet error: ", readSheetErr)
		return
	}

	currentRowCount := len(sheet.Values)

	if currentRowCount > 0 {

		sheetData := sheet.Values
		columnHeading := sheetData[0]

		for index, value := range columnHeading {
			columnContent := fmt.Sprintf("%v", value)

			if strings.EqualFold(columnContent, "twitter") {
				twitterIndex = index + 1
				letter := toCharStr(twitterIndex)
				subReturn.TwitterCell = letter + strconv.FormatInt(int64(currentRowCount), 10)
			}
		}

		if currentRowCount >= 2 {

			list := sheet.Values
			extractedList := list[currentRowCount-1]

			for _, v := range extractedList {
				columnContent := fmt.Sprintf("%v", v)
				match, _ := regexp.MatchString("^\\w+([-+.']\\w+)*@[A-Za-z\\d]+\\.com$", columnContent)
				if match {
					subReturn.EmailAddress = columnContent
				}
			}
		}

		contentType := "application/json"

		transport, err := cloudevents.NewHTTPTransport(cloudevents.WithTarget(sub.Endpoint), cloudevents.WithStructuredEncoding())
		if err != nil {
			fmt.Println("failed to create transport : ", err)
			return
		}

		client, err := cloudevents.NewClient(transport, cloudevents.WithTimeNow())
		if err != nil {
			fmt.Println("failed to create client : ", err)
			return
		}

		source, err := url.Parse(sub.Endpoint)
		event := cloudevents.Event{
			Context: cloudevents.EventContextV01{
				EventID:     sub.ID,
				EventType:   "listener",
				Source:      cloudevents.URLRef{URL: *source},
				ContentType: &contentType,
			}.AsV01(),
			Data: subReturn,
		}

		if (oldRowCount == 0 || oldRowCount < currentRowCount) && currentRowCount >= 2 {

			oldRowCount = currentRowCount
			_, resp, err := client.Send(context.Background(), event)
			if err != nil {
				log.Printf("failed to send: %v", err)
			}

			subReturn.EmailAddress = ""
			subReturn.SheetTitle = ""
			subReturn.SpreadsheetID = ""
			subReturn.TwitterCell = ""

			fmt.Printf("Response: \n%s\n", resp)
		}
	}
}

func toCharStr(i int) string {
	return string('A' - 1 + i)
}
