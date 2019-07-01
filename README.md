# _Google Sheets_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)

An OMG service for Google Sheets, it is for organization, analysis and storage of data in tabular form. Spreadsheets developed as computerized analogs of paper accounting worksheets.

## Direct usage in [Storyscript](https://storyscript.io/):

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### Create Spreadsheet
```shell
$ omg run createSpreadsheet -a title=<SHEET_TITLE> -e KEY=<BASE64_DATA_OF_KEY_FILE>
```
##### Find Spreadsheet
```shell
$ omg run findSpreadsheet -a spreadsheetId=<SHEET_ID> -e KEY=<BASE64_DATA_OF_KEY_FILE>
```
##### Add Sheet
```shell
$ omg run addSheet -a spreadsheetId=<SHEET_ID> -a sheetTitle=<SHEET_TITLE> -e KEY=<BASE64_DATA_OF_KEY_FILE>
```
##### Find Sheet
```shell
$ omg run findSheet -a spreadsheetId=<SHEET_ID> -a sheetTitle=<SHEET_TITLE>/ -a sheetIndex=<SHEET_INDEX>/ -a sheetId=<SHEET_ID> -e KEY=<BASE64_DATA_OF_KEY_FILE>
```
##### Expand Sheet
```shell
$ omg run expandSheet -a spreadsheetId=<SHEET_ID> -a sheetTitle=<SHEET_TITLE> -a row=<EXPAND_ROW> -a column=<EXPAND_COLUMN> -e KEY=<BASE64_DATA_OF_KEY_FILE>
```

**Note**: the OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/heaptracetechnology/google-sheets/blob/master/LICENSE).
