FROM alpine
COPY my-app-go /my-app-go
COPY templates /templates
CMD ["/my-app-go"]