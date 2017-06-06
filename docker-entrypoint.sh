#!/bin/sh

echo "Running: /mcrouter/mcrouter-kuberentes-provisioner --mcrouterconfig=$MCROUTERCONFIG --inputtemplate=$INPUTTEMPLATE --namespace=$NAMESPACE --incluster=$INCLUSTER --kubeconfig=$KUBECONFIG"
/mcrouter/mcrouter-kuberentes-provisioner --mcrouterconfig=$MCROUTERCONFIG --inputtemplate=$INPUTTEMPLATE --namespace=$NAMESPACE --incluster=$INCLUSTER --kubeconfig=$KUBECONFIG
