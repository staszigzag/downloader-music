FROM  alpine:latest

#RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY ./.bin .
COPY ./configs .

EXPOSE 80

CMD ["./app"]