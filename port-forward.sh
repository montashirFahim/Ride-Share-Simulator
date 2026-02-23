#!/bin/bash

echo "Starting port forwarding for services..."
echo "User service will be available at: http://localhost:8080"
echo "Ride service will be available at: http://localhost:8000"
echo ""
echo "Press Ctrl+C to stop"
echo ""

kubectl port-forward svc/user-service 8080:8080 &
USER_PID=$!

kubectl port-forward svc/ride-service 8000:8000 &
RIDE_PID=$!

trap "kill $USER_PID $RIDE_PID 2>/dev/null" EXIT

wait
