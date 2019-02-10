FROM mysql:5.7
LABEL maintainer "s.okazaki"

COPY ./sample.sql /docker-entrypoint-initdb.d/echo-sample.sql

