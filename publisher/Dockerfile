FROM alpine
#ADD ca-certificates.crt /etc/ssl/certs/
# for SSL
RUN apk --no-cache add ca-certificates
COPY ./config/config.json /config.json
ADD main /
# exposes port 80 to other docker containers, 
# nothing outside docker can access
#EXPOSE 80
CMD ["/main"]