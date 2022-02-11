FROM golang:latest
LABEL maintainer="Quique <hello@pragmaticreviews.com"
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY  . .
ENV PORT 9002
RUN go build
CMD [ "./basic-api" ]