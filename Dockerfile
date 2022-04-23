FROM ubuntu:18.04

RUN mkdir /home/apps
COPY /build/server /home/apps/

CMD [ "/home/apps/server" ]