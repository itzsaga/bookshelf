
# Base our image on an official, minimal image of our preferred golang
FROM golang:1.9

# Note: The default golang docker image, already has the GOPATH env variable set.
# GOPATH is located at /go
# ENV GO_SRC $GOPATH/src
ENV PONZU_GITHUB github.com/ponzu-cms/ponzu
ENV PROJECT_ROOT $GOPATH/src/github.com/itzsaga/bookshelf

RUN go get $PONZU_GITHUB/...

# Consider updating package in the future. For instance ca-certificates etc.
# RUN apt-get update -qq && apt-get install -y build-essential

# Make the ponzu root directory
RUN mkdir -p $PROJECT_ROOT

# All commands will be run inside of ponzu root
WORKDIR $PROJECT_ROOT

# Copy the ponzu source into ponzu root.
COPY . .

# Define the scripts we want run once the container boots
RUN ponzu build
CMD ponzu run --port=$PORT
