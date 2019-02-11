FROM golang:1.9

# GOPATH is located at /go
# ENV GO_SRC $GOPATH/src
ENV PONZU_GITHUB github.com/ponzu-cms/ponzu
ENV PROJECT_GITHUB github.com/itzsaga/bookshelf
ENV PROJECT_ROOT $GOPATH/src/github.com/itzsaga/bookshelf

RUN go get $PONZU_GITHUB/...
RUN go get $PROJECT_GITHUB

# RUN mkdir -p $PROJECT_ROOT
# WORKDIR $PROJECT_ROOT

# RUN git clone $PROJECT_GITHUB
WORKDIR $PROJECT_ROOT

RUN ponzu build
CMD ponzu run --port=$PORT
