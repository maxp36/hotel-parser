############################
# STEP 1 build executable binary
############################

# golang:1.13-alpine
FROM golang@sha256:c750d6718009f2e94cb20f56a87884f601f175d43c9418ae0fa21ea00ad6a2ff as builder

RUN apk update && apk add --update --no-cache \
  build-base \
  git \
  ca-certificates \
  && update-ca-certificates

WORKDIR /hotel-parser
COPY . .

# COPY .env ./build/.env
# COPY raw ./build/raw

RUN go mod download

RUN go env

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ./build/app


############################
# STEP 2 build a small image
############################

#FROM scratch
FROM alpine

COPY --from=builder /hotel-parser/build /hotel-parser

WORKDIR /hotel-parser

ENTRYPOINT ["/bin/sh"]

# CMD [ "./app", "-dir=raw/" ]
