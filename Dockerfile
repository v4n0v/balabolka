#имя базового образа
FROM golang as builder

#создаем папку, где будет наша программа
RUN mkdir -p /go/src/balabolka

#идем в папку
WORKDIR /go/src/balabolka

#копируем все файлы из текущего пути к файлу Docker на вашей системе в нашу новую папку образа
COPY . /go/src/balabolka

#скачиваем зависимые пакеты через скрипт, любезно разработанный командой docker
RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM scratch
COPY --from=builder /go/src/balabolka /go/src/balabolka

#пробрасываем порт вашей программы наружу образа
EXPOSE 8080
ENTRYPOINT ["./go/src/balabolka/balabolka"]

#docker build -t balabolka_app .
#sudo docker run -itd -p 8080:8080 balabolka_app