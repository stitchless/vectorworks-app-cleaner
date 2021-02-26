FROM golang:latest

LABEL authors="Josh Peizer <github@sent.sh>"

ENV GOOS $os
ENV FLAGS $flags
ENV GOARCH amd6
ENV RUNNING true

VOLUME ["~/temp", "/go/src/app"]

WORKDIR /go/src/app

COPY ./dist.sh .

RUN apt -y update && \
    apt -y install \
    webkit2gtk-4.0 \
    gtk+-3.0 && \
    mkdir ~/temp

# Run our build
CMD ["dist.sh"]