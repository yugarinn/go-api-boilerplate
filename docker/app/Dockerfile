FROM ubuntu:22.04

WORKDIR /var/www/html

ENV TZ UTC
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN set -x \
    && echo 'debconf debconf/frontend select Noninteractive' | debconf-set-selections \
    && apt-get update \
    && apt-get install -y gnupg \
    && echo "deb http://ppa.launchpad.net/nginx/development/ubuntu bionic main" > /etc/apt/sources.list.d/ppa_nginx_mainline.list \
    && apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys E5267A6C \
    && apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys C300EE8C \
    && apt-get update \
    && apt-get install -y vim wget curl zip unzip git gcc sqlite3 openssh-server mysql-client  \
    && apt-get -y autoremove \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Install Dockerize
ENV DOCKERIZE_VERSION v0.6.1
RUN wget -q https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

# Install go 1.18
RUN curl -OL https://golang.org/dl/go1.18.6.linux-amd64.tar.gz \
    && tar -C /usr/local -xvf go1.18.6.linux-amd64.tar.gz \
    && echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.profile \
    && echo "export PATH=\$PATH:/root/go/bin" >> ~/.profile \
    && . ~/.profile \
    && go version

# Install go system dependencies
RUN . ~/.profile && go install github.com/codegangsta/gin@latest
RUN . ~/.profile && go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN . ~/.profile && go install github.com/go-delve/delve/cmd/dlv@latest

EXPOSE 8080

RUN chown -R www-data:www-data /var/www/html
WORKDIR /var/www/html
