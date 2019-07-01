FROM golang

RUN go get golang.org/x/net/context

RUN go get golang.org/x/oauth2/google

RUN go get gopkg.in/Iwark/spreadsheet.v2

RUN go get github.com/gorilla/mux

WORKDIR /go/src/github.com/heaptracetechnology/google-sheets

ADD . /go/src/github.com/heaptracetechnology/google-sheets

RUN go install github.com/heaptracetechnology/google-sheets

ENTRYPOINT google-sheets

EXPOSE 3000