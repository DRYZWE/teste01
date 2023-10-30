printf "Installing Building Artefacts\n\n"

sudo apt -y update

sudo apt -y install build-essential

printf "\n\nInstalling GoLang\n"

sudo curl -fsSL https://golang.org/dl/go1.15.12.linux-amd64.tar.gz --output go1.15.12.linux-amd64.tar.gz

sudo rm -rf /opt/go

sudo tar -C /opt -xvzf go1.15.12.linux-amd64.tar.gz

mkdir -p $HOME/go

printf "\n\nInstalling NodeJs\n"
echo "Instalando Node.js..."

# Adicione o repositório do Node.js e a chave GPG
curl -fsSL https://deb.nodesource.com/gpgkey/nodesource.gpg.key | sudo gpg --dearmor -o /usr/share/keyrings/nodesource-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/nodesource-archive-keyring.gpg] https://deb.nodesource.com/node_14.x $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/nodesource.list

# Atualize o cache do apt
sudo apt update

# Instale o Node.js e o npm
sudo apt install -y nodejs

# Exiba a versão do Node.js e do npm instalados
echo "Node.js $(node -v)"
echo "npm $(npm -v)"

echo "Node.js instalado com sucesso!"

printf "\n\nInstalling Docker\n"

curl -fsSL https://get.docker.com -o get-docker.sh && sudo sh get-docker.sh

printf "\n\nReseting Docker\n"

sudo usermod -aG docker $(whoami)

sudo grpck

sudo grpconv

newgrp docker << END

sudo systemctl restart docker.service

printf "\n\nInstalling Docker-Compose\n"

sudo curl -L "https://github.com/docker/compose/releases/download/1.29.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

sudo chmod +x /usr/local/bin/docker-compose

rm -f go1.15.12.linux-amd64.tar.gz nodesource_setup.sh get-docker.sh

cd $HOME

printf "\n\nCustomizing enviroment variables\n"

echo "export GOPATH=$HOME/go" >> ~/.bashrc

echo "export GOROOT=/opt/go" >> ~/.bashrc

source ~/.bashrc

printf "\n\nEnviroment configured\n"

END

newgrp docker
