FROM alpine
COPY app /app
COPY templates /templates
CMD ["/app"]