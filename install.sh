#!/bin/bash 
exist=$(which go)
id=$(id -u)
if [[ $id != 0 ]]; then 
   echo "debes de ser root"
   exit 1
fi

if [[ $exist == "" ]];then 
   echo -ne "1. pacman | 2. apt \n-> "; read opt
   if [[ $opt == 1 ]]; then
     pacman -Sy
     pacman -S go
   else
     apt update -y
     apt install go -y 
   fi
fi

go build neteye.go
cp neteye /usr/bin/
echo "Listo para usar"
