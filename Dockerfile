FROM golang:1.12

LABEL maintainer="Jainam | Dilip"

RUN mkdir /hospital
ADD . /hospital/
WORKDIR /hospital

COPY . .

RUN go build

CMD ["./hospital"]
