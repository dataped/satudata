FROM golang:latest
#LABEL maintainer="Hiromu Ochiai <otiai10@gmail.com>"
RUN cat /etc/os-release
# Output: Debian GNU/Linux 10 (buster)
RUN apt-get -qy update

RUN apt-get install -qy libleptonica-dev libtesseract-dev
RUN apt-get install -qy libtool m4 automake cmake pkg-config
RUN apt-get install -qy libicu-dev libpango1.0-dev libcairo-dev

RUN cd /opt && git clone https://github.com/tesseract-ocr/tesseract
WORKDIR /opt/tesseract
RUN git reset --hard 5.3.3
RUN ./autogen.sh &&\
    ./configure --enable-debug LDFLAGS="-L/usr/local/lib" CFLAGS="-I/usr/local/include"
RUN make -j 8
RUN make install && ldconfig
RUN tesseract --version

ENV TESSDATA_PREFIX=/usr/local/share/tessdata
ENV TESSDATA_REPO=https://github.com/tesseract-ocr/tessdata_best
WORKDIR ${TESSDATA_PREFIX}
RUN wget -q ${TESSDATA_REPO}/raw/4.1.0/eng.traineddata