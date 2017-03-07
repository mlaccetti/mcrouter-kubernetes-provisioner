FROM golang:1.8-alpine

LABEL MAINTAINER="Michael Laccetti <michael@laccetti.com>"

RUN mkdir -p /mkp/template

ADD mcrouter-kuberentes-provisioner /mkp/mcrouter-kuberentes-provisioner

ADD ./template/mcrouter-config.tpl /mkp/template/mcrouter-config.tpl

RUN chmod +x /mkp/mcrouter-kuberentes-provisioner

ENTRYPOINT ["/mkp/mcrouter-kuberentes-provisioner"]
