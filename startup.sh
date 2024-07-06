#!/bin/bash

# Directory where this script is located
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Function to start backend
start_backend() {
    echo "Starting GoLang backend..."
    go run cmd/server/main.go > >(tee -a "$DIR/backend.log") 2>&1 &
    # Store the PID of the backend process
    backend_pid=$!
}

# Function to start frontend
start_frontend() {
    echo "Starting React frontend..."
    cd web/frontend
    npm start > >(tee -a "$DIR/frontend.log") 2>&1 &
    # Store the PID of the frontend process
    frontend_pid=$!
}

# Function to stop processes on script termination
cleanup() {
    echo "Stopping GoLang backend..."
    kill $backend_pid 2>/dev/null
    echo "Stopping React frontend..."
    kill $frontend_pid 2>/dev/null
    exit
}

# Trap termination signals and call cleanup function
trap cleanup INT TERM EXIT

# Start backend and frontend
start_backend
start_frontend

# Wait for script termination
wait
