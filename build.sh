#!/bin/bash

help(){
    echo "Usage : bash build.sh [OPTIONS]"
    echo ""
    echo "Build the project"
    echo ""
    echo "Options: "
    echo -e "\t-h display help"
    echo -e "\t-d generate docker image"
    echo -e "\t-s build without running test"
    echo ""
    echo "Sample : bash build.sh -d"
    exit 0
}

#Check if script is running throught /bin/sh (SHLVL=1) or /bin/bash (SHLVL=2)
if [ "$SHLVL" -lt "2" ] ; then
    echo "ERROR : Consider running this script using /bin/bash not /bin/sh. \"/bin/bash build.sh -h\" for help"
    exit 1
fi

declare docker=false
declare skip_tests=false

declare red='\033[0;31m'
declare yellow='\033[0;33m'
declare default='\033[0m'
declare cyan='\033[0;36m'

while getopts ":hds" option; do
   case $option in
      h) # display Help
         help
         exit;;
      d) # docker
         echo -e "${yellow}App and Docker image will be generated"${default}""
         docker=true
         ;;
      n) # native
         echo -e "${yellow}Native Docker image will be generated locally"${default}""
         native=true
         ;;
      s) #skip test
        echo -e "${yellow}Using skip tests mode.${default}"
        skip_tests=true
        ;;
      \?) # exclude
         echo -e "${red}Error: Invalid option. Use -h for help${default}"
         exit;;
   esac
done

echo -e "${cyan}Building project...${default}"

echo "Generate package folder..."
# clean package
rm -dr package/

mkdir package

# copying resources
cp -r docker/* package/

#Go version needed : 1.20
go mod tidy
#build linux binary
GOOS=linux GOARCH=amd64 go build -o bin/go-api-amd64-linux main.go

# possible to build for windows
#GOOS=windows GOARCH=386 go build -o bin/app-386.exe src/main.go
# possible to build for macOS on different architectures
# 64-bit
#GOOS=darwin GOARCH=amd64 go build -o bin/app-amd64-darwin src/main.go
# 32-bit
#GOOS=darwin GOARCH=386 go build -o bin/app-386-darwin src/main.go
# possible to build for linux on different architectures
# 64-bit
#GOOS=linux GOARCH=amd64 go build -o bin/app-amd64-linux src/main.go
# 32-bit
#GOOS=linux GOARCH=386 go build -o bin/app-386-linux src/main.go

if [ "$skip_tests" = false ]; then
    go test ./...
fi

if [ "$docker" = true ]; then
    echo -e "${yellow}Building docker image...${default}"

    docker build --no-cache --build-arg VERSION=master -t valentinconan/go-server-compare:master .
    docker build --build-arg VERSION=master -f Dockerfile.test -t valentinconan/go-bombardier:master .
fi

echo -e "${cyan}Build project done${default}"





