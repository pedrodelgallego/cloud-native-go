FROM alpine

EXPOSE 8080

COPY cloud-native-go /app/cloud-native-go
RUN chmod +x /app/cloud-native-go

ENTRYPOINT /app/cloud-native-go
