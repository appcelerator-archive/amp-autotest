FROM appcelerator/alpine:20160928

ENV GOPATH /go
ENV PATH $PATH:/go/bin

COPY ./ /go/src/github.com/appcelerator/amp-test
RUN apk update
RUN apk -v add git make go && \
    go version && \
    cd /go/src/github.com/appcelerator/amp-test && \
    go get -u github.com/Masterminds/glide/... && \
    rm -f glide.lock && \
    glide install && \
    rm -f ./amp-test && \
    make install


CMD ["/go/bin/amp-test", "--service-swarm"]
