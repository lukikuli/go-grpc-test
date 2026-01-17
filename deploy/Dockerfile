FROM alpine:latest
ARG ARCH

RUN addgroup -S apps && \
    adduser -S apps -G apps && \
    apk --no-cache add bash=5.2.21-r0  && \
    apk --no-cache add git=2.43.6-r0 && \
    apk --no-cache add curl=8.11.1-r0 && \
    apk --no-cache add tzdata=2024b-r0

WORKDIR /app
ENV SERVICENAME go-grpc
ENV APP_HOME /app

COPY --chown=apps:apps bin/main-$ARCH /app/main
RUN chmod +x /app/main
USER apps
CMD ["/app/main"]
