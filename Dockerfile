FROM golang

RUN go get google.golang.org/api/drive/v3

RUN go get google.golang.org/api/sheets/v4

RUN go get github.com/cloudevents/sdk-go

RUN go get golang.org/x/oauth2/google

RUN go get github.com/gorilla/mux

WORKDIR /go/src/github.com/heaptracetechnology/google-sheets

ADD . /go/src/github.com/heaptracetechnology/google-sheets

RUN go install github.com/heaptracetechnology/google-sheets

ENTRYPOINT google-sheets

EXPOSE 3000