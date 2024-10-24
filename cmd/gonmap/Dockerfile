FROM php:8.3.4-apache
# FROM php:8.3.12-apache
ARG DEBIAN_FRONTEND=noninteractive
RUN docker-php-ext-install mysqli
COPY init.sql /docker-entrypoint-initdb.d/
COPY ./app /var/www/html/
