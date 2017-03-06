FROM golang:1.8-alpine

LABEL MAINTAINER="Michael Laccetti <michael@laccetti.com>"

RUN mkdir /mkp

WORKDIR /mkp

ADD mcrouter-kuberentes-provisioner /mkp

ENTRYPOINT ["/mkp/mcrouter-kubernetes-provisioner"]
