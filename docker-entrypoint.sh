#!/bin/bash

echo "Running: /mcr/bin/mcrouter-kuberentes-provisioner --mcrouterconfig=$MCROUTERCONFIG --inputtemplate=$INPUTTEMPLATE --namespace=$NAMESPACE --incluster=$INCLUSTER --kubeconfig=$KUBECONFIG"

/mcr/bin/mcrouter-kuberentes-provisioner \
  --mcrouterconfig=$MCROUTERCONFIG \
  --inputtemplate=$INPUTTEMPLATE \
  --namespace=$NAMESPACE \
  --incluster=$INCLUSTER \
  --kubeconfig=$KUBECONFIG
