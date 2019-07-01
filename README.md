# _Google Sheets_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)

An OMG service for Google Sheets, it is for organization, analysis and storage of data in tabular form. Spreadsheets developed as computerized analogs of paper accounting worksheets.

## Direct usage in [Storyscript](https://storyscript.io/):

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### Create Spreadsheet
```sh
$ omg run createSpreadsheet -a title=<SHEET_TITLE> -e KEY=<BASE64_DATA_OF_KEY_FILE>
```
##### Find Spreadsheet
```sh
$ omg run findSpreadsheet -a spreadsheetId=<SHEET_ID> -e KEY=<BASE64_DATA_OF_KEY_FILE>
```
##### Add Sheet
```sh
$ omg run addSheet -a spreadsheetId=<SHEET_ID> -a sheetTitle=<SHEET_TITLE> -e KEY=<BASE64_DATA_OF_KEY_FILE>
```
##### Send Message By Topic
```sh
$ omg run sendMessageByTopic -a token=<TOKEN> -a topic=<TOPIC> -a body=<NOTIFICATION_BODY> -a icon=<NOTIFICATION_ICON> -a data=<DATA_OBJECT>  -e SERVER_KEY=<SERVER_KEY>
```

**Note**: the OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/firebase/blob/master/LICENSE).
