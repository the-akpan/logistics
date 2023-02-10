FROM node:16-alpine as frontend

WORKDIR /app

COPY client/package*.json .

RUN npm install

COPY client/ .

RUN npm run build

FROM golang:1.19.2-alpine as base

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build . && chmod +x tracka

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=base /app/tracka .

COPY --from=frontend /app/dist ./build

EXPOSE 8000

CMD [ "./tracka" ]
