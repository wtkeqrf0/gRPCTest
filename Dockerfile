FROM mysql
MAINTAINER matvey-sizov@mail.ru

ADD files/grpc_test.sql /docker-entrypoint-initdb.d

ENV MYSQL_ROOT_PASSWORD secret
ENV MYSQL_DATABASE grpc_test

EXPOSE 3306