############################
# STEP 1 build executable binary
############################

# golang:1.12-alpine
FROM golang@sha256:c750d6718009f2e94cb20f56a87884f601f175d43c9418ae0fa21ea00ad6a2ff as builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --update --no-cache \
  build-base \
  git \
  nodejs nodejs-npm \
  ca-certificates \
  && update-ca-certificates

# Create unprivileged appuser
# RUN adduser -D -g '' appuser

WORKDIR $GOPATH/src/bitbucket.org/gorastudio/arr-back-messenger
COPY . .

RUN npm install apidoc -g
RUN rm -Rf /go/bin/messenger/build/doc \
  && apidoc -i messenger -o /go/bin/messenger/build/doc

COPY .env /go/bin/messenger/build/.env

# Fetch dependencies by dep
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -v

RUN go env

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/messenger/build/service


############################
# STEP 2 build a small image
############################

#FROM scratch
FROM alpine

#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Import the user and group files from the builder.
# COPY --from=builder /etc/passwd /etc/passwd

# Use an unprivileged user.
# USER appuser
RUN apk add nginx curl vim &&  sed -i 's/return 404/root \/messenger\/build\/doc/' /etc/nginx/conf.d/default.conf && mkdir /run/nginx

COPY --from=builder /go/bin/messenger/build /messenger/build

WORKDIR /messenger/build

# # Redis

# #Postgres

ENTRYPOINT ["./service"]
#CMD ["./service"}
