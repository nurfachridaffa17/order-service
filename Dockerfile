##
## STEP 1 - BUILD
##

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.22.3-alpine AS build

RUN apk --no-cache add tzdata

# create a working directory inside the image
WORKDIR /order-service

# copy all to image
COPY . /order-service

# download Go modules and dependencies
RUN go mod download

# compile application
RUN go build  -o /app cmd/main.go

# remove app
#RUN rm -rf /amna

##
## STEP 2 - DEPLOY
##

FROM scratch

WORKDIR /

COPY --from=build /app /app

COPY --from=build /order-service/cmd/.env.development /.env.development

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo

# tells Docker that the container listens on specified network ports at runtime
EXPOSE 8111

# command to be used to execute when the image is used to start a container
CMD ["/app"]