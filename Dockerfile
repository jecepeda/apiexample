FROM gobuffalo/buffalo:v0.11.0 as builder

RUN mkdir -p $GOPATH/src/github.com/jcepedavillamayor/apiexample
WORKDIR $GOPATH/src/github.com/jcepedavillamayor/apiexample

# this will cache the npm install step, unless package.json changes
ADD . .

RUN chmod +x wait-for-it.sh

# Bind the app to 0.0.0.0 so it can be seen from outside the container
ENV ADDR=0.0.0.0

FROM alpine
RUN apk add --no-cache bash
RUN apk add --no-cache ca-certificates

ENV ADDR=0.0.0.0
