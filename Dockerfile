FROM golang:1.8-alpine

LABEL MAINTAINER="Michael Laccetti <michael@laccetti.com>"

RUN mkdir /mkp

ADD mcrouter-kuberentes-provisioner /mkp/mcrouter-kuberentes-provisioner

RUN chmod +x /mkp/mcrouter-kuberentes-provisioner

ENTRYPOINT ["/mkp/mcrouter-kuberentes-provisioner"]
