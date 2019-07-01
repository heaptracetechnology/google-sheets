package route

import (
    "github.com/gorilla/mux"
    spreadsheet "github.com/heaptracetechnology/google-sheets/spreadsheets"
    "log"
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "CreateSpreadsheet",
        "POST",
        "/createSpreadsheet",
        spreadsheet.CreateSpreadsheet,
    },
    Route{
        "FindSpreadsheet",
        "POST",
        "/findSpreadsheet",
        spreadsheet.FindSpreadsheet,
    },
    Route{
        "FindSheet",
        "POST",
        "/findSheet",
        spreadsheet.FindSheet,
    },
    Route{
        "AddSheet",
        "POST",
        "/addSheet",
        spreadsheet.AddSheet,
    },
    Route{
        "ExpandSheet",
        "POST",
        "/expandSheet",
        spreadsheet.ExpandSheet,
    },
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        log.Println(route.Name)
        handler = route.HandlerFunc

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}
