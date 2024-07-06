#!/bin/bash

# Directory where this script is located
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Backend
echo "Starting GoLang backend..."
go run cmd/server/main.go > >(tee -a "$DIR/backend.log") 2>&1 &

# Frontend
echo "Starting React frontend..."
cd web/frontend
npm start > >(tee -a "$DIR/frontend.log") 2>&1 &
