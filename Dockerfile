FROM golang:1.11.4

RUN apt-get update && apt-get install -y libglu1-mesa-dev libgles2-mesa-dev libxrandr-dev libxcursor-dev libxinerama-dev libxi-dev libasound2-dev && apt-get clean && rm -rf /var/lib/apt/lists/*

RUN go get github.com/hajimehoshi/ebiten/... github.com/gopherjs/gopherjs github.com/gopherjs/gopherwasm/js

WORKDIR /go/src/app
COPY app .

RUN gopherjs build -o yourgame.js app/src/

EXPOSE 8080

CMD gopherjs serve
