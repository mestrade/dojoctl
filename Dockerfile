FROM golang:buster as builder

ARG GITHUB_TOKEN

RUN mkdir /app 
ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -o dojoctl .
 


FROM busybox 

COPY --from=builder app/dojoctl dojoctl

CMD ["./dojoctl", "-setup", "dojo.conf"]

