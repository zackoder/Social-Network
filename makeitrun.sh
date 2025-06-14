#!/bin/bash

# Create .env file for frontend if it doesn't exist
if [ ! -e "./frontend/.env" ]; then 
    echo "NEXT_PUBLIC_HOST=http://localhost:8080" >> ./frontend/.env
fi

# Function to open a new Terminal window and run a command
run_in_new_terminal_mac() {
  local CMD=$1
  osascript <<EOF
tell application "Terminal"
    activate
    do script "${CMD}"
end tell
EOF
}

# Build absolute paths
FRONTEND_PATH=$(cd frontend && pwd)
BACKEND_PATH=$(cd backend && pwd)

# Start frontend in new terminal
run_in_new_terminal_mac "cd '${FRONTEND_PATH}'; npm install; npm run dev"

# Start backend in another new terminal
run_in_new_terminal_mac "cd '${BACKEND_PATH}'; go install github.com/zzwx/fresh@latest; export PATH=\$PATH:\$HOME/go/bin; fresh"
