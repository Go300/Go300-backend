It's Go300-backend. Requirements:
1) Docker CE (I use: Docker version 18.03.1-ce on Ubuntu 18.04)
2) docker-compose - we recommend to use python script (install via pip), not to install by apt-get (system package manager)

If you want to see what we did, follow this commands:
git clone https://github.com/Go300/Go300-backend.git (or git@github.com:Go300/Go300-backend.git, but you need SSH key)

docker-compose up - to run in terminal
docker-compose up -d - to run process independently from terminal
docker-compose exec app bash - get inside project container's bash
go test ./... -cover - just run global testing

We are developing it now, coming soon!