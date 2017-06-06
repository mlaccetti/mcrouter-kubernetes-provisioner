#!/bin/bash

while true ; do
  touch /mcrouter/.foo >> /dev/null 2>&1
  if [ $? != 0 ]; then
    echo "/mcrouter isn't available yet"
    sleep 1;
  else
    break;
  fi
done

echo "Running: /mcr/bin/mcrouter-kuberentes-provisioner --mcrouterconfig=$MCROUTERCONFIG --inputtemplate=$INPUTTEMPLATE --namespace=$NAMESPACE --incluster=$INCLUSTER --kubeconfig=$KUBECONFIG"

/mcr/bin/mcrouter-kuberentes-provisioner \
  --mcrouterconfig=$MCROUTERCONFIG \
  --inputtemplate=$INPUTTEMPLATE \
  --namespace=$NAMESPACE \
  --incluster=$INCLUSTER \
  --kubeconfig=$KUBECONFIG
