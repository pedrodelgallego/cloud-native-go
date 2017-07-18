FROM alpine

EXPOSE 8080

COPY cloud-native-go /

CMD ["/cloud-native-go"]
