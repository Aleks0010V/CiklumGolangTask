# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git mercurial gcc
ADD . /src
RUN cd /src/main && go build -o goapp

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/main /app/
EXPOSE 8888
ENTRYPOINT ./goapp