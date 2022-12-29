git submodule update --init --recursive
cd ./api

go mod tidy
go $1 -v 
