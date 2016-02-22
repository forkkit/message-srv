FROM alpine:3.2
ADD message-srv /message-srv
ENTRYPOINT [ "/message-srv" ]
