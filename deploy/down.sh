label() {
  echo "===== $1 ====="
}

print() {
  echo "==> $1"
}

cd "$(dirname "$0")" || exit

label "stopping containers..."
docker compose down --rmi all