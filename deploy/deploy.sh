#!/bin/bash

label() {
  echo "===== $1 ====="
}

print() {
  echo "==> $1"
}

function clone() {
    local repo_url=$1
    local folder_name=$2

    if [ -d "$folder_name" ]; then
        cd "$folder_name" || return 1
        if git pull; then
            print "pull successful"
            cd ..
        else
            print "pull failed, try re-cloning"
            cd ..
            rm -rf "$folder_name"
            git clone "$repo_url" "$folder_name"
        fi
    else
        git clone "$repo_url" "$folder_name" || return 1
    fi
}

cd "$(dirname "$0")" || exit

label "clone resonance-microservice"
clone git@github.com:kmou424/resonance-microservice.git resonance-microservice || exit 1

label "clone resonance-dataserver"
clone git@github.com:kmou424/resonance-dataserver.git resonance-dataserver || exit 1

label "deploying apps..."
docker compose up --force-recreate --build -d

label "cleaning old build cache..."
docker builder prune --force --filter "until=24h"

read -r -p "Press any key to continue..."