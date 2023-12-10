FROM golang:1.21.5-alpine

RUN apk update
RUN apk add \
  g++ \
  git \
  musl-dev \
  tesseract-ocr-dev
RUN apk add tesseract-ocr-data-eng

#ENV GOPATH=/go

WORKDIR $GOPATH/src/satudata
COPY go.mod ./
RUN mkdir -p $GOPATH/src/satudata

RUN apk add build-base
COPY . $GOPATH/src/satudata

ENV GOSSERACT_CPPSTDERR_NOT_CAPTURED=1

RUN go build -o satudata .
EXPOSE 8080

CMD ["/go/src/satudata/satudata"]