FROM golang:1.8-alpine

LABEL MAINTAINER="Michael Laccetti <michael@laccetti.com>"

# You probably want to override this - either via CLI, or via k8s variables
ENV MCROUTERCONFIG=/mcrouter/mcrouter-config.json
ENV INPUTTEMPLATE=/mcrouter/template/mcrouter-config.tpl
ENV NAMESPACE=
ENV INCLUSTER=true
ENV KUBECONFIG=

RUN mkdir -p /mcrouter/template
RUN mkdir -p /mcrouter/bin

ADD docker-entrypoint.sh /mcrouter/bin/docker-entrypoint.sh
ADD mcrouter-kuberentes-provisioner /mcrouter/bin/mcrouter-kuberentes-provisioner
ADD ./template/mcrouter-config.tpl /mcrouter/template/mcrouter-config.tpl

RUN chmod +x /mcrouter/bin/mcrouter-kuberentes-provisioner && \
  chmod +x /mcrouter/docker-entrypoint.sh

#ENTRYPOINT [ "/mcrouter/bin/docker-entrypoint.sh" ]
CMD tail -f /dev/null
