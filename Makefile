OS := $(shell uname)

# Variables
NPM=../node-v10.15.3-linux-x64/bin/npm
GOPATH=$GOPATH


#::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
get_tools:
ifeq ($(OS),Linux)
	sudo apt-get install wget unzip curl
endif

ifeq ($(OS),Darwin)
	brew install wget unzip curl
endif

#::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
get_dep:

	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

#::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
get_node:

	wget -c https://nodejs.org/dist/v10.15.1/node-v10.15.1-linux-x64.tar.xz
	tar -xf node-v10.15.3-linux-x64.tar.xz

#::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

get_docker:

	curl -fsSL https://get.docker.com -o get-docker.sh
	sudo sh get-docker.sh
	sudo curl -L https://github.com/docker/compose/releases/download/1.21.2/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
	sudo chmod +x /usr/local/bin/docker-compose

#::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
dependencies:
	$(MAKE) get_tools
	$(MAKE) get_golang
	$(MAKE) get_node
	$(MAKE) get_docker

#::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
install:
	$(MAKE) dependencies

#::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
build:
	sudo docker-compose build
	cd poll-api/cmd/poll-api && $(GOPATH)/bin/dep ensure -v -update
	cd vote-consumer-service/cmd/vote-consumer-service && $(GOPATH)/bin/dep ensure -v -update

#::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
run:
	sudo docker-compose up -d
	cd poll-api/cmd/poll-api && ./main
	cd vote-consumer-service/cmd/vote-consumer-service && ./main
	cd poll-app && $(NPM) serve
