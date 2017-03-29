FROM alpine:edge

ENV GOPATH /go
ENV GOREPO github.com/allen13/kube-source
RUN mkdir -p $GOPATH/src/$GOREPO
COPY . $GOPATH/src/$GOREPO
WORKDIR $GOPATH/src/$GOREPO

RUN set -ex \
	&& apk add --no-cache --virtual .build-deps \
		git \
		go \
		build-base \
	&& go build kube-source.go \
	&& apk del .build-deps \
	&& rm -rf $GOPATH/pkg

EXPOSE 5606

ENV GIN_MODE release

CMD ./kube-source
