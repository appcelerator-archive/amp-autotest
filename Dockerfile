FROM appcelerator/alpine:20160928

ENV GOPATH /go
ENV PATH $PATH:/go/bin

COPY ./ /go/src/github.com/appcelerator/amp-autotest
RUN apk update && apk upgrade && \
    apk -v add git make musl-dev go@community && \
    # package pinning doesn't work with virtual packages
    go version && \
    cd /go/src/github.com/appcelerator/amp-autotest && \
    go get -u github.com/Masterminds/glide/... && \
    glide install && \
    rm -f ./amp-autotest && \
    make install && \
    rm /go/bin/glide && \

CMD ["/go/bin/amp-autotest"]