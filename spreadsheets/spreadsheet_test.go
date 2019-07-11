package spreadsheet

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
)

var (
	key                     = os.Getenv("GOOGLE_SHEETS_KEY")
	spreadsheetTitle        = os.Getenv("GOOGLE_SPREADSHEET_TITLE")
	spreadsheetID           = os.Getenv("GOOGLE_SPREADSHEET_ID")
	addsheettitle           = os.Getenv("GOOGLE_ADD_SHEET_TITLE")
	updateSheetRow          = os.Getenv("GOOGLE_SHEET_ROW")
	updateSheetColumn       = os.Getenv("GOOGLE_SHEET_COLUMN")
	updateContentSheetTitle = os.Getenv("GOOGLE_SHEET_UPDATE_CONTENT_SHEET_TITLE")
)

var _ = Describe("Create Spreadsheet with invalid base64 KEY", func() {

	//invalid key
	os.Setenv("KEY", "mockKey")

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

	os.Setenv("KEY", key)

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

var _ = Describe("Create Spreadsheet with valid params", func() {

	os.Setenv("KEY", key)

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
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Find Spreadsheet with invalid base64 KEY", func() {

	os.Setenv("KEY", "mockKey")

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

	os.Setenv("KEY", key)

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

	os.Setenv("KEY", key)

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

var _ = Describe("Find Spreadsheet with valid params", func() {

	os.Setenv("KEY", key)

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

	os.Setenv("KEY", "mockKey")

	sheet := ArgsData{ID: spreadsheetID, SheetTitle: spreadsheetTitle}
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

	os.Setenv("KEY", key)

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

	os.Setenv("KEY", key)

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

	os.Setenv("KEY", key)

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
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Find Sheet with invalid base64 KEY", func() {

	os.Setenv("KEY", "mockKey")

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

	os.Setenv("KEY", key)

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

	os.Setenv("KEY", key)

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

var _ = Describe("Find Sheet with valid params", func() {

	os.Setenv("KEY", key)

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

	os.Setenv("KEY", key)

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

	os.Setenv("KEY", "mockKey")

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

	os.Setenv("KEY", key)

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

	os.Setenv("KEY", key)

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

var _ = Describe("Update sheet size with valid params", func() {

	os.Setenv("KEY", key)

	row, _ := strconv.Atoi(updateSheetRow)
	column, _ := strconv.Atoi(updateSheetColumn)

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

	os.Setenv("KEY", key)

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

var _ = Describe("Update sheet size with invalid base64 KEY", func() {

	os.Setenv("KEY", "mockKey")

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

var _ = Describe("Update sheet size with invalid spreadsheet ID", func() {

	os.Setenv("KEY", key)

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

var _ = Describe("Delete sheet with valid params", func() {

	os.Setenv("KEY", key)

	sheet := ArgsData{ID: spreadsheetID, SheetID: 2019}
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

var _ = Describe("Delete sheet invalid param", func() {

	os.Setenv("KEY", key)

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

var _ = Describe("Update sheet size with invalid base64 KEY", func() {

	os.Setenv("KEY", "mockKey")

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

var _ = Describe("Update sheet size with invalid spreadsheet ID", func() {

	os.Setenv("KEY", key)

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

var _ = Describe("Delete sheet with valid params", func() {

	os.Setenv("KEY", key)

	sheet := ArgsData{ID: spreadsheetID, SheetID: 2019}
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

	os.Setenv("KEY", key)

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

	os.Setenv("KEY", "mockKey")

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

var _ = Describe("Update sheet size with invalid spreadsheet ID", func() {

	os.Setenv("KEY", key)

	sheet := ArgsData{ID: "mockSpreadsheetID"}
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

var _ = Describe("Update sheet size with invalid spreadsheet ID", func() {

	os.Setenv("KEY", key)

	sheet := ArgsData{ID: "mockSpreadsheetID", SheetTitle: "mockTempTitle"}
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

var _ = Describe("Update cell with valid params", func() {

	os.Setenv("KEY", key)

	sheet := ArgsData{ID: spreadsheetID, SheetTitle: updateContentSheetTitle, Row: 2, Column: 3, Content: "Test content"}
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

var _ = Describe("Get cell invalid param", func() {

	os.Setenv("KEY", key)

	sheet := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/getCell", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCell)
	handler.ServeHTTP(recorder, request)

	Describe("Get cell", func() {
		Context("get cell", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Get cell with invalid base64 KEY", func() {

	os.Setenv("KEY", "mockKey")

	sheet := ArgsData{ID: spreadsheetID}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/getCell", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCell)
	handler.ServeHTTP(recorder, request)

	Describe("Get cell", func() {
		Context("get cell", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Get Cell with invalid spreadsheet ID", func() {

	os.Setenv("KEY", key)

	sheet := ArgsData{ID: "mockSpreadsheetID"}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/getCell", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCell)
	handler.ServeHTTP(recorder, request)

	Describe("Get cell", func() {
		Context("get cell", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Get cell with valid params", func() {

	os.Setenv("KEY", key)

	sheet := ArgsData{ID: spreadsheetID, SheetTitle: updateContentSheetTitle, Row: 2, Column: 3}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(sheet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/getCell", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCell)
	handler.ServeHTTP(recorder, request)

	Describe("Get cell", func() {
		Context("get cell", func() {
			It("Should result http.StatusStatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})
