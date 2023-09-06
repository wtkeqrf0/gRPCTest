FROM mysql:8.1
MAINTAINER matvey-sizov@mail.ru

ADD files/grpc_test.sql /docker-entrypoint-initdb.d

ENV MYSQL_ROOT_PASSWORD secret
ENV MYSQL_DATABASE grpc_test

EXPOSE 3306

HEALTHCHECK --interval=5m --timeout=3s \
  CMD curl -f http://localhost/ || exit 1