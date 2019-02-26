FROM golang:1.9

# GOPATH is located at /go
# ENV GO_SRC $GOPATH/src
ENV PONZU_GITHUB github.com/ponzu-cms/ponzu
ENV PROJECT_GITHUB https://github.com/itzsaga/bookshelf.git
ENV PROJECT_ROOT $GOPATH/src/github.com/itzsaga/bookshelf

RUN go get $PONZU_GITHUB/...

RUN mkdir -p $PROJECT_ROOT

WORKDIR $PROJECT_ROOT
RUN git clone $PROJECT_GITHUB .

RUN ponzu build
CMD ponzu run --port=${PORT} --bind=0.0.0.0
