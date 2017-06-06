#!/bin/bash

# wait for the template to avoid possible race

while [ ! -e $INPUTTEMPLATE ]; do
  echo "waiting for $INPUTTEMPLATE";
  sleep 1;
done

echo "Running: /mcr/bin/mcrouter-kuberentes-provisioner --mcrouterconfig=$MCROUTERCONFIG --inputtemplate=$INPUTTEMPLATE --namespace=$NAMESPACE --incluster=$INCLUSTER --kubeconfig=$KUBECONFIG"

/mcr/bin/mcrouter-kuberentes-provisioner \
  --mcrouterconfig=$MCROUTERCONFIG \
  --inputtemplate=$INPUTTEMPLATE \
  --namespace=$NAMESPACE \
  --incluster=$INCLUSTER \
  --kubeconfig=$KUBECONFIG
