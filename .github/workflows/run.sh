git submodule update --init --recursive

cd thirdparty/go-sqlite3
git checkout v1.14.16
cd ../..

cd ./api

go mod tidy
go $1 -v 
