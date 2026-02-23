#!/bin/bash

set -e

CLUSTER_NAME="ride-cluster"

USER_LOCAL="ridesimulator-user-app-v2:latest"
RIDE_LOCAL="ridesimulator-ride-app-v2:latest"

USER_IMAGE="montashir/ride-simulator:user-app-v2"
RIDE_IMAGE="montashir/ride-simulator:ride-app-v2"

if ! command -v kind &> /dev/null; then
    echo "kind not installed"
    exit 1
fi

if ! command -v kubectl &> /dev/null; then
    echo "kubectl not installed"
    exit 1
fi

if ! docker info &> /dev/null; then
    echo "Docker not running"
    exit 1
fi

# Recreate KIND cluster
if kind get clusters | grep -q "$CLUSTER_NAME"; then
    kind delete cluster --name "$CLUSTER_NAME"
fi

echo "Building fresh Docker images..."
docker build -t $USER_LOCAL ./User
docker build -t $RIDE_LOCAL ./Ride

echo "Tagging images for Docker Hub..."
docker tag $USER_LOCAL $USER_IMAGE
docker tag $RIDE_LOCAL $RIDE_IMAGE

echo "Logging into Docker Hub (if needed)..."
docker login

echo "Pushing images to Docker Hub..."
docker push $USER_IMAGE
docker push $RIDE_IMAGE



echo "Creating KIND cluster..."
kind create cluster --config kind-config.yaml --wait 3m

echo "Deploying services..."
kubectl apply -f k8s/postgres.yaml
kubectl apply -f k8s/redis.yaml

kubectl apply -f k8s/user-service.yaml
kubectl apply -f k8s/rider-service.yaml

echo "Waiting for deployments..."
kubectl wait --for=condition=available --timeout=200s deployment/user-service || true
kubectl wait --for=condition=available --timeout=200s deployment/ride-service || true

echo "Setup complete!"
kubectl get pods
kubectl get svc

