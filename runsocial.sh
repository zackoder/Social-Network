#!/bin/bash

if [ ! -e "./frontend/.env" ]; then 
    echo "NEXT_PUBLIC_HOST=http://localhoset:8080" >> .env
fi

# cd frontend 
# npm i

# cd ../backend
# export PATH="$PATH:$HOME/go/bin"
# source ~/.bashrc
# fresh -g
# fresh

# Start frontend in a new terminal
gnome-terminal -- bash -c "
cd frontend
npm install
npm run dev
exec bash
"

# Start backend in another new terminal
gnome-terminal -- bash -c "
cd backend
go install github.com/zzwx/fresh@latest
export PATH=\"\$PATH:\$HOME/go/bin\"
source ~/.bashrc
fresh -g
fresh
exec bash
"
