FROM golang:1.8-alpine

LABEL MAINTAINER="Michael Laccetti <michael@laccetti.com>"

# You probably want to override this - either via CLI, or via k8s variables
ENV MCROUTERCONFIG=/srv/nfs/mcrouter-config.json
ENV INPUTTEMPLATE=/mkp/template/mcrouter-config.tpl
ENV NAMESPACE=
ENV INCLUSTER=true
ENV KUBECONFIG=

RUN mkdir -p /mkp/template

ADD docker-entrypoint.sh /mkp/docker-entrypoint.sh
ADD mcrouter-kuberentes-provisioner /mkp/mcrouter-kuberentes-provisioner
ADD ./template/mcrouter-config.tpl /mkp/template/mcrouter-config.tpl

RUN chmod +x /mkp/mcrouter-kuberentes-provisioner && \
  chmod +x /mkp/docker-entrypoint.sh

#ENTRYPOINT [ "/mkp/docker-entrypoint.sh" ]
CMD tail -f /dev/null
