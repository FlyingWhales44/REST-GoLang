FROM golang:1.11
WORKDIR $GOPATH/src/gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 8080
CMD ["obsluga-formularzy"]