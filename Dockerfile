FROM debian:latest
MAINTAINER Volodymyr Tkach <vladimirok5959@gmail.com>

ENV ENV_HOST=127.0.0.1 ENV_PORT=8080 ENV_DATA_DIR=/app/data

COPY ./bin/ip2location /app/app

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get -y update && \
    apt-get -y upgrade && \
    apt-get install -y curl ca-certificates && \
    dpkg-reconfigure -p critical ca-certificates && \
    echo "" >> /root/.profile && \
    echo "TIME_ZONE=\$(cat /etc/timezone)" >> /root/.profile && \
    echo "export TZ=\"\${TIME_ZONE}\"" >> /root/.profile && \
    echo "" >> /root/.bashrc && \
    echo "TIME_ZONE=\$(cat /etc/timezone)" >> /root/.bashrc && \
    echo "export TZ=\"\${TIME_ZONE}\"" >> /root/.bashrc && \
    mkdir /app/data && \
    mkdir /app/logs && \
    chmod +x /app/app

HEALTHCHECK --interval=30s --timeout=5s --start-period=5s CMD curl --fail http://localhost:$ENV_PORT/api/v1/app/health || exit 1

EXPOSE 8080
VOLUME /app/data

CMD export TZ="$(cat /etc/timezone)" && /app/app
