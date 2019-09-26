# _Google Sheets_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.com/omg-services/google-sheets.svg?branch=master)](https://travis-ci.com/omg-services/google-sheets)
[![codecov](https://codecov.io/gh/omg-services/google-sheets/branch/master/graph/badge.svg)](https://codecov.io/gh/omg-services/google-sheets)

An OMG service for Google Sheets, it is for organization, analysis and storage of data in tabular form. Spreadsheets developed as computerized analogs of paper accounting worksheets.

## Direct usage in [Storyscript](https://storyscript.io/):

##### Create Spreadsheet
```coffee
google-sheets createSpreadsheet title:'Spreadsheet title' emailAddress:'email address for drive permission' role:'role of access' type:'type of access'
```
##### Find Spreadsheet
```coffee
google-sheets findSpreadsheet spreadsheetId:'Spreadsheet Id'
```
##### Add Sheet
```coffee
google-sheets addSheet spreadsheetId:'Spreadsheet Id' sheetTitle:'Sheet title'
```
##### Find Sheet
```coffee
google-sheets findSheet spreadsheetId:'Spreadsheet Id' sheetTitle:'Sheet title'
```
##### Update Sheet Size
```coffee
google-sheets updateSheetSize spreadsheetId:'Spreadsheet Id' sheetId:'Sheet Id' row:1 column:2
```
##### Update Cell
```coffee
google-sheets updateCell spreadsheetId:'Spreadsheet Id' sheetTitle:'Sheet title' cellNumber:'A1' content:'any content'
```
##### Delete Sheet
```coffee
google-sheets deleteSheet spreadsheetId:'Spreadsheet Id' sheetId:'sheet Id'
```
##### Subscribe Sheet
```coffee
google-sheets listener newRowUpdate spreadsheetID:'Spreadsheet Id' sheetTitle:'sheet title'
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### Create Spreadsheet
```shell
$ omg run createSpreadsheet -a title=<SPREADSHEET_TITLE> -a emailAddress=<EMAIL_ADDRESS> -a role=<ROLE_OF_ACCESS> -a type=<TYPE_OF_ACCESS> -e CREDENTIAL_JSON=<BASE64_DATA_OF_CREDENTIAL_JSON_FILE>
```
##### Find Spreadsheet
```shell
$ omg run findSpreadsheet -a spreadsheetId=<SPREADSHEET_ID> -e CREDENTIAL_JSON=<BASE64_DATA_OF_CREDENTIAL_JSON_FILE>
```
##### Add Sheet
```shell
$ omg run addSheet -a spreadsheetId=<SPREADSHEET_ID> -a sheetTitle=<SHEET_TITLE> -e CREDENTIAL_JSON=<BASE64_DATA_OF_CREDENTIAL_JSON_FILE>
```
##### Find Sheet
```shell
$ omg run findSheet -a spreadsheetId=<SPREADSHEET_ID> -a sheetTitle=<SHEET_TITLE> -e CREDENTIAL_JSON=<BASE64_DATA_OF_CREDENTIAL_JSON_FILE>
```
##### Update Sheet Size
```shell
$ omg run updateSheetSize -a spreadsheetId=<SPREADSHEET_ID> -a sheetId=<SHEET_ID> -a row=<ROW_LENGTH> -a column=<COLUMN_LENGTH> -e CREDENTIAL_JSON=<BASE64_DATA_OF_CREDENTIAL_JSON_FILE>
```
##### Update Cell
```shell
$ omg run updateCell -a spreadsheetId=<SPREADSHEET_ID> -a sheetTitle=<SHEET_TITLE> -a cellNumber=<CELL_NUMBER> -a content=<CELL_CONTENT> -e CREDENTIAL_JSON=<BASE64_DATA_OF_CREDENTIAL_JSON_FILE>
```
##### Delete Sheet
```shell
$ omg run deleteSheet -a spreadsheetId=<SPREADSHEET_ID> -a sheetId=<SHEET_ID> -e CREDENTIAL_JSON=<BASE64_DATA_OF_CREDENTIAL_JSON_FILE>
```
##### Subscribe Sheet
```shell
omg subscribe listener newRowUpdate -a spreadsheetID=<SPREADSHEET_ID> -a sheetTitle=<SHEET_TITLE> -e CREDENTIAL_JSON=<BASE64_DATA_OF_CREDENTIAL_JSON_FILE>
```

**Note**: the OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/heaptracetechnology/google-sheets/blob/master/LICENSE).
