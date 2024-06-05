FROM golang:1.22.4-alpine

RUN apk update
RUN apk add && apk add --no-cache \
  g++ \
  musl-dev \
  tesseract-ocr-dev
  build-base \
  cmake \
  git \
  libjpeg-turbo-dev \
  libpng-dev \
  tiff-dev \
  opencv-dev \
  && rm -rf /var/cache/apk/*

# Set environment variable untuk GoCV
ENV CGO_CPPFLAGS="-I/usr/include"
ENV CGO_LDFLAGS="-L/usr/lib -lopencv_core -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_videoio -lopencv_objdetect"

RUN apk add tesseract-ocr-data-eng
# Unduh dan instal GoCV
RUN go get -u -d gocv.io/x/gocv

#ENV GOPATH=/go

WORKDIR $GOPATH/src/satudata
COPY go.mod ./
RUN mkdir -p $GOPATH/src/satudata

RUN apk add build-base
COPY . $GOPATH/src/satudata

#ENV GOSSERACT_CPPSTDERR_NOT_CAPTURED=1

RUN go build -o satudata .
EXPOSE 8080

CMD ["/go/src/satudata/satudata"]
