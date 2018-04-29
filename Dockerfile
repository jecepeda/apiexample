FROM golang:1.10 as builder


# Install postgresql (for wait-for-it.sh)
RUN apt-get update
RUN apt-get install postgresql postgresql-contrib -y

# Install buffalo
RUN wget  https://github.com/gobuffalo/buffalo/releases/download/v0.11.0/buffalo_0.11.0_linux_amd64.tar.gz
RUN tar -xvzf buffalo_0.11.0_linux_amd64.tar.gz
RUN mv buffalo-no-sqlite /usr/local/bin/buffalo

# Install dep (could be improved)
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN mkdir -p $GOPATH/src/github.com/jcepedavillamayor/apiexample
WORKDIR $GOPATH/src/github.com/jcepedavillamayor/apiexample
ADD . .

# dep
RUN dep ensure

RUN chmod +x wait-for-it.sh

ENV ADDR=0.0.0.0
