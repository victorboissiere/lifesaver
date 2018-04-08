#!/usr/bin/env bash

echo "NODES"
kubectl get nodes -o wide
echo -e "\nSERVICES"
kubectl get svc -o wide
echo -e "\nPODS"
kubectl get pods -o wide
echo
kubectl get event \
  --sort-by=.metadata.creationTimestamp \
  | head -n 10
