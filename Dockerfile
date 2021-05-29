FROM golang:1.16.4-alpine as build-env

RUN echo http://dl-cdn.alpinelinux.org/alpine/edge/main >> /etc/apk/repositories
RUN echo http://dl-cdn.alpinelinux.org/alpine/edge/testing >> /etc/apk/repositories
RUN apk update

RUN apk --no-cache add gcc g++ make git

## Update CA Certificates
RUN apk add --no-cache git ca-certificates && update-ca-certificates

# Download all timezone data
RUN apk add --no-cache tzdata
RUN apk add tzdata

# Add Maintainer Info
LABEL maintainer="Supratick Dey"

RUN mkdir /app

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY ./ .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/go_gin_backend

# Run the bin
FROM scratch

COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /go/bin/go_gin_backend /go/bin/go_gin_backend
COPY --from=build-env /app/jwtsecret.key .
COPY --from=build-env /app/.env .

## Copy all timezone data inside the container
COPY --from=build-env /usr/share/zoneinfo/ /usr/share/zoneinfo/

ENTRYPOINT ["/go/bin/go_gin_backend"]