#this is only for development
FROM golang:1.9.2-alpine

ENV ENV_API development

#install git
RUN apk add --no-cache git mercurial

# install dependency tool
RUN go get -u github.com/golang/dep/cmd/dep

# Fresh for rebuild on code change, no need for production
RUN go get -u github.com/pilu/fresh

COPY . /go/src/github.com/monstar-lab/goweb1

WORKDIR /go/src/github.com/monstar-lab/goweb1

# dep ensure to ensure the availability of required libraries used by the go source
# for development, pilu/fresh is used to automatically build the application everytime you save a Go or template file

RUN dep ensure
CMD fresh

# for production, it just builds and runs the binary
# CMD dep ensure && go build && ./goweb1

EXPOSE 8080