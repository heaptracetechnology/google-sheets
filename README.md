# _Google Sheets_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.com/omg-services/google-sheets.svg?branch=master)](https://travis-ci.com/omg-services/google-sheets)
[![codecov](https://codecov.io/gh/omg-services/google-sheets/branch/master/graph/badge.svg)](https://codecov.io/gh/omg-services/google-sheets)

An OMG service for Google Sheets, it is for organization, analysis and storage of data in tabular form. Spreadsheets developed as computerized analogs of paper accounting worksheets.

## Direct usage in [Storyscript](https://storyscript.io/):

##### Create Spreadsheet
```coffee
google-sheets createSpreadsheet title:'Spreadsheet Title'
```
##### Find Spreadsheet
```coffee
google-sheets findSpreadsheet spreadsheetId:'Spreadsheet Id'
```
##### Add Sheet
```coffee
google-sheets addSheet spreadsheetId:'Spreadsheet Id' sheetTitle:'Sheet title'
```
##### Find Sheet By Title
```coffee
google-sheets findSheet spreadsheetId:'Spreadsheet Id' sheetTitle:'Sheet title'
```
##### Find Sheet By Index
```coffee
google-sheets findSheet spreadsheetId:'Spreadsheet Id' sheetIndex:'Sheet index'
```
##### Find Sheet By ID
```coffee
google-sheets findSheet spreadsheetId:'Spreadsheet Id' sheetId:'Sheet ID'
```
##### Update Sheet Size
```coffee
google-sheets updateSheetSize spreadsheetId:'Spreadsheet Id' sheetTitle:'Sheet title' row:100 column:200
```
##### Update Cell
```coffee
google-sheets updateCell spreadsheetId:'Spreadsheet Id' sheetTitle:'Sheet title' row:1 column:2 content:'any content'
```
##### Get Cell
```coffee
google-sheets getCell spreadsheetId:'Spreadsheet Id' sheetTitle:'Sheet title' row:1 column:2
```
##### Delete Sheet
```coffee
google-sheets deleteSheet spreadsheetId:'Spreadsheet Id' sheetId:'sheet Id'
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### Create Spreadsheet
```shell
$ omg run createSpreadsheet -a title=<SPREADSHEET_TITLE> -e KEY=<BASE64_DATA_OF_KEY_FILE>
```
##### Find Spreadsheet
```shell
$ omg run findSpreadsheet -a spreadsheetId=<SPREADSHEET_ID> -e KEY=<BASE64_DATA_OF_KEY_FILE>
```
##### Add Sheet
```shell
$ omg run addSheet -a spreadsheetId=<SPREADSHEET_ID> -a sheetTitle=<SHEET_TITLE> -e KEY=<BASE64_DATA_OF_KEY_FILE>
```
##### Find Sheet By Title
```shell
$ omg run findSheet -a spreadsheetId=<SPREADSHEET_ID> -a sheetTitle=<SHEET_TITLE> -e KEY=<BASE64_DATA_OF_KEY_FILE>
```
##### Find Sheet By Index
```shell
$ omg run findSheet -a spreadsheetId=<SPREADSHEET_ID> -a sheetIndex=<SHEET_INDEX> -e KEY=<BASE64_DATA_OF_KEY_FILE>
```
##### Find Sheet By ID
```shell
$ omg run findSheet -a spreadsheetId=<SPREADSHEET_ID> -a sheetId=<SHEET_ID> -e KEY=<BASE64_DATA_OF_KEY_FILE>
```
##### Update Sheet Size
```shell
$ omg run updateSheetSize -a spreadsheetId=<SPREADSHEET_ID> -a sheetTitle=<SHEET_TITLE> -a row=<UPDATE_ROW> -a column=<UPDATE_COLUMN> -e KEY=<BASE64_DATA_OF_KEY_FILE>
```
##### Update Cell
```shell
$ omg run updateCell -a spreadsheetId=<SPREADSHEET_ID> -a sheetTitle=<SHEET_TITLE> -a row=<ROW_NUMBER> -a column=<COLUMN_NUMBER> -a content=<CELL_CONTENT> -e KEY=<BASE64_DATA_OF_KEY_FILE>
```
##### Get Cell
```shell
$ omg run getCell -a spreadsheetId=<SPREADSHEET_ID> -a sheetTitle=<SHEET_TITLE> -a row=<ROW_NUMBER> -a column=<COLUMN_NUMBER> -e KEY=<BASE64_DATA_OF_KEY_FILE>
```
##### Delete Sheet
```shell
$ omg run deleteSheet -a spreadsheetId=<SPREADSHEET_ID> -a sheetId=<SHEET_ID> -e KEY=<BASE64_DATA_OF_KEY_FILE>
```

**Note**: the OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/heaptracetechnology/google-sheets/blob/master/LICENSE).
