FROM golang:1.15-buster

RUN apt-get update \
    && apt-get install -y unzip \
    && wget -q -O promtail.zip "https://github.com/grafana/loki/releases/download/v2.1.0/promtail-linux-amd64.zip" \
    && unzip ./promtail.zip -d /go/bin \
    && rm -rf promtail.zip \ 
    && apt-get clean

WORKDIR /go/src/app

COPY ./deployments/entrypoint.sh /tmp/entrypoint.sh
COPY ./deployments/promtail-config.yml /tmp/promtail-config.yml
COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/app ./main.go

EXPOSE 8080
ENTRYPOINT [ "/bin/bash", "/tmp/entrypoint.sh" ]
USER root

CMD [ "/go/bin/app" ]