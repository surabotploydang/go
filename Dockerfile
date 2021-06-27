# Merge all compiled and make as a single simple docker image
FROM alpine:3.12
MAINTAINER "Internet Thailand PCL"
WORKDIR /app
COPY ./build/ /app

# Install tzdata and set docker image timezone to bangkok timezone
RUN apk add --no-cache tzdata
ENV TZ=Asia/Bangkok

# Run server version check
#RUN echo ">>> Test server version"
#RUN /app/server --version

# Expose service port
EXPOSE 5000
CMD ["/app/server", "start"]