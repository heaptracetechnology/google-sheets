package spreadsheets

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
)

var (
	key               = os.Getenv("GOOGLE_SHEETS_CREDENTIAL_JSON")
	spreadsheetTitle  = os.Getenv("GOOGLE_SPREADSHEET_TITLE")
	spreadsheetID     = os.Getenv("GOOGLE_SPREADSHEET_ID")
	addsheettitle     = os.Getenv("GOOGLE_ADD_SHEET_TITLE")
	updateSheetRow    = os.Getenv("GOOGLE_SHEET_ROW")
	updateSheetColumn = os.Getenv("GOOGLE_SHEET_COLUMN")
	emailAddress      = os.Getenv("GOOGLE_SHEET_DRIVE_PERMISSION_EMAIL_ADDRESS")
	role              = os.Getenv("GOOGLE_SHEET_DRIVE_ROLE")
	accessType        = os.Getenv("GOOGLE_SHEET_DRIVE_TYPE")
	cellNumber        = os.Getenv("GOOGLE_SHEET_CELL_NUMBER")
	cellContent       = os.Getenv("GOOGLE_SHEET_CELL_CONTENT")
)

var _ = Describe("HealthCheck", func() {

	sheet := ArgsData{}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/health", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheck)
	handler.ServeHTTP(recorder, request)

	Describe("Health Check", func() {
		Context("health check", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create Spreadsheet with invalid base64 KEY", func() {

	//invalid key
	os.Setenv("CREDENTIAL_JSON", "mockKey")

	sheet := ArgsData{Title: spreadsheetTitle}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/createSpreadsheet", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateSpreadsheet)
	handler.ServeHTTP(recorder, request)

	Describe("Create Spreadsheet", func() {
		Context("create spreadsheet", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create Spreadsheet without Sheet title", func() {

	os.Setenv("CREDENTIAL_JSON", key)

	sheet := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/createSpreadsheet", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateSpreadsheet)
	handler.ServeHTTP(recorder, request)

	Describe("Create Spreadsheet", func() {
		Context("create spreadsheet", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

//********************************************************************************************
var _ = Describe("Create Spreadsheet with valid params", func() {

	os.Setenv("CREDENTIAL_JSON", key)

	sheet := ArgsData{Title: spreadsheetTitle, IsTesting: true, EmailAddress: emailAddress, Role: role, Type: accessType}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/createSpreadsheet", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateSpreadsheet)
	handler.ServeHTTP(recorder, request)

	Describe("Create Spreadsheet", func() {
		Context("create spreadsheet", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Find Spreadsheet with invalid base64 KEY", func() {

	os.Setenv("CREDENTIAL_JSON", "mockKey")

	sheet := ArgsData{ID: spreadsheetID}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/findSpreadsheet", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FindSpreadsheet)
	handler.ServeHTTP(recorder, request)

	Describe("Find Spreadsheet", func() {
		Context("find spreadsheet", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Find Spreadsheet invalid param", func() {

	os.Setenv("CREDENTIAL_JSON", key)

	sheet := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/findSpreadsheet", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FindSpreadsheet)
	handler.ServeHTTP(recorder, request)

	Describe("Find Spreadsheet", func() {
		Context("find spreadsheet", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Find Spreadsheet with invalid spreadsheet ID", func() {

	os.Setenv("CREDENTIAL_JSON", key)

	sheet := ArgsData{ID: "mockSpreadsheetID"}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/findSpreadsheet", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FindSpreadsheet)
	handler.ServeHTTP(recorder, request)

	Describe("Find Spreadsheet", func() {
		Context("find spreadsheet", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

//*******************************************************************************************
var _ = Describe("Find Spreadsheet with valid params", func() {

	os.Setenv("CREDENTIAL_JSON", key)

	sheet := ArgsData{ID: spreadsheetID}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/findSpreadsheet", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FindSpreadsheet)
	handler.ServeHTTP(recorder, request)

	Describe("Find Spreadsheet", func() {
		Context("find spreadsheet", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Add Sheet with invalid base64 KEY", func() {

	os.Setenv("CREDENTIAL_JSON", "mockKey")

	sheet := ArgsData{ID: spreadsheetID, SheetTitle: addsheettitle}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/addSheet", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(AddSheet)
	handler.ServeHTTP(recorder, request)

	Describe("Add Sheet", func() {
		Context("add sheet", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Add Sheet without Sheet title", func() {

	os.Setenv("CREDENTIAL_JSON", key)

	sheet := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/addSheet", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(AddSheet)
	handler.ServeHTTP(recorder, request)

	Describe("Add Sheet", func() {
		Context("add sheet", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Add Sheet with invalid spreadsheet ID", func() {

	os.Setenv("CREDENTIAL_JSON", key)

	sheet := ArgsData{ID: "mockID", Title: "mockTitle"}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/addSheet", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(AddSheet)
	handler.ServeHTTP(recorder, request)

	Describe("Add Sheet", func() {
		Context("add sheet", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Add Sheet with valid params", func() {

	os.Setenv("CREDENTIAL_JSON", key)

	sheet := ArgsData{ID: spreadsheetID, SheetTitle: addsheettitle}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/addSheet", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(AddSheet)
	handler.ServeHTTP(recorder, request)

	Describe("Add Sheet", func() {
		Context("add sheet", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Find Sheet with invalid base64 KEY", func() {

	os.Setenv("CREDENTIAL_JSON", "mockKey")

	sheet := ArgsData{ID: spreadsheetID}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/findSheet", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FindSheet)
	handler.ServeHTTP(recorder, request)

	Describe("Find Sheet", func() {
		Context("find sheet", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Find Sheet invalid param", func() {

	os.Setenv("CREDENTIAL_JSON", key)

	sheet := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/findSheet", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FindSheet)
	handler.ServeHTTP(recorder, request)

	Describe("Find Sheet", func() {
		Context("find sheet", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Find Sheet with invalid spreadsheet ID", func() {

	os.Setenv("CREDENTIAL_JSON", key)

	sheet := ArgsData{ID: "mockSpreadsheetID"}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/findSheet", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FindSheet)
	handler.ServeHTTP(recorder, request)

	Describe("Find Spreadsheet", func() {
		Context("find spreadsheet", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

//******************************************************************************************
var _ = Describe("Find Sheet with valid params", func() {

	os.Setenv("CREDENTIAL_JSON", key)

	sheet := ArgsData{ID: spreadsheetID, SheetTitle: addsheettitle}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/findSheet", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FindSheet)
	handler.ServeHTTP(recorder, request)

	Describe("Find Sheet", func() {
		Context("find sheet", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Update sheet size invalid param", func() {

	os.Setenv("CREDENTIAL_JSON", key)

	sheet := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/updateSheetSize", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateSheetSize)
	handler.ServeHTTP(recorder, request)

	Describe("Update sheet size", func() {
		Context("update sheet size", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Update sheet size with invalid base64 KEY", func() {

	os.Setenv("CREDENTIAL_JSON", "mockKey")

	sheet := ArgsData{ID: spreadsheetID}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/updateSheetSize", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateSheetSize)
	handler.ServeHTTP(recorder, request)

	Describe("Update sheet size", func() {
		Context("update sheet size", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Update sheet size with invalid spreadsheet ID", func() {

	os.Setenv("CREDENTIAL_JSON", key)

	sheet := ArgsData{ID: "mockSpreadsheetID"}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/updateSheetSize", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateSheetSize)
	handler.ServeHTTP(recorder, request)

	Describe("Update sheet size", func() {
		Context("update sheet size", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Update sheet size with invalid sheet title", func() {

	os.Setenv("CREDENTIAL_JSON", key)

	sheet := ArgsData{ID: spreadsheetID, SheetTitle: "mockInvalidTitle"}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/updateSheetSize", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateSheetSize)
	handler.ServeHTTP(recorder, request)

	Describe("Update sheet size", func() {
		Context("update sheet size", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

//*************************************************************************************************
var _ = Describe("Update sheet size with valid params", func() {

	os.Setenv("CREDENTIAL_JSON", key)

	row, _ := strconv.ParseInt(updateSheetRow, 10, 64)
	column, _ := strconv.ParseInt(updateSheetColumn, 10, 64)

	sheet := ArgsData{ID: spreadsheetID, SheetTitle: addsheettitle, Row: row, Column: column}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/updateSheetSize", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateSheetSize)
	handler.ServeHTTP(recorder, request)

	Describe("Update sheet size", func() {
		Context("update sheet size", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Delete sheet invalid param", func() {

	os.Setenv("CREDENTIAL_JSON", key)

	sheet := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/deleteSheet", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteSheet)
	handler.ServeHTTP(recorder, request)

	Describe("Delete sheet", func() {
		Context("delete sheet", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Delete sheet with invalid base64 KEY", func() {

	os.Setenv("CREDENTIAL_JSON", "mockKey")

	sheet := ArgsData{ID: spreadsheetID}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/deleteSheet", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteSheet)
	handler.ServeHTTP(recorder, request)

	Describe("Delete sheet", func() {
		Context("delete sheet", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Delete sheet with invalid spreadsheet ID", func() {

	os.Setenv("CREDENTIAL_JSON", key)

	sheet := ArgsData{ID: "mockSpreadsheetID"}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/deleteSheet", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteSheet)
	handler.ServeHTTP(recorder, request)

	Describe("Delete sheet", func() {
		Context("delete sheet", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Update cell invalid param", func() {

	os.Setenv("CREDENTIAL_JSON", key)

	sheet := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/updateCell", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateCell)
	handler.ServeHTTP(recorder, request)

	Describe("Update cell", func() {
		Context("update cell", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Update cell with invalid base64 KEY", func() {

	os.Setenv("CREDENTIAL_JSON", "mockKey")

	sheet := ArgsData{ID: spreadsheetID}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/updateCell", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateCell)
	handler.ServeHTTP(recorder, request)

	Describe("Update cell", func() {
		Context("update cell", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

//*************************************************************************************************
var _ = Describe("Update cell with valid params", func() {

	os.Setenv("CREDENTIAL_JSON", key)

	sheet := ArgsData{ID: spreadsheetID, SheetTitle: addsheettitle, CellNumber: cellNumber, Content: cellContent}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/updateCell", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateCell)
	handler.ServeHTTP(recorder, request)

	Describe("Update cell", func() {
		Context("update cell", func() {
			It("Should result http.StatusStatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

//********************************************************************************************
var _ = Describe("Subscribe google sheet for new row update", func() {

	os.Setenv("CREDENTIAL_JSON", key)

	data := RequestParam{SpreadsheetID: spreadsheetID, SheetTitle: addsheettitle}
	sub := Subscribe{Endpoint: "https://webhook.site/3cee781d-0a87-4966-bdec-9635436294e9",
		ID:        "1",
		IsTesting: true,
		Data:      data,
	}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(sub)
	if err != nil {
		fmt.Println(" request err :", err)
	}
	req, err := http.NewRequest("POST", "/subscribe", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SheetSubscribe)
	handler.ServeHTTP(recorder, req)

	Describe("Subscribe", func() {
		Context("Subscribe", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})
