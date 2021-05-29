FROM golang:1.16
MAINTAINER Jonel Mawirat

RUN mkdir /app
WORKDIR /app
COPY ./app/ /app