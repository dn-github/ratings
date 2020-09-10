FROM alpine:3.4

COPY ratings .
ENTRYPOINT ./ratings
EXPOSE 3003