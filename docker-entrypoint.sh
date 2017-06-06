#!/bin/bash

# wait for the template to avoid possible race

while [ ! -d /mcrouter ]; do
  echo "waiting for /mcrouter"
  sleep 10;
done

echo "Running: /mcr/bin/mcrouter-kuberentes-provisioner --mcrouterconfig=$MCROUTERCONFIG --inputtemplate=$INPUTTEMPLATE --namespace=$NAMESPACE --incluster=$INCLUSTER --kubeconfig=$KUBECONFIG"

/mcr/bin/mcrouter-kuberentes-provisioner \
  --mcrouterconfig=$MCROUTERCONFIG \
  --inputtemplate=$INPUTTEMPLATE \
  --namespace=$NAMESPACE \
  --incluster=$INCLUSTER \
  --kubeconfig=$KUBECONFIG
