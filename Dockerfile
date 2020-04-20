FROM alpine:3.9.4

# Provide APPNAME variable when build this Dockerfile
ENV APPNAME antrian-bucket

# Preparation
RUN mkdir -p /app
ADD antrian-bucket /app
WORKDIR /app
RUN chmod +x /app/$APPNAME \
    && mv /app/$APPNAME /app/antrian

# Entry Point
ENTRYPOINT ["/app/antrian"]
