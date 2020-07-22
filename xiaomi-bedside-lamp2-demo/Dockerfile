FROM ubuntu:latest

CMD mkdir -p light-mapper

COPY . light-mapper/

WORKDIR light-mapper

ENTRYPOINT ["/light-mapper/lamp-device-mapper/light-mapper","-logtostderr=true"]
