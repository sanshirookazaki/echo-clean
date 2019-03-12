FROM mysql:5.7
LABEL maintainer "s.okazaki"

COPY ./init.sql /docker-entrypoint-initdb.d/init.sql
COPY ./my.cnf /etc/mysql/conf.d/my.cnf
