git submodule update --init --recursive

cd api/thirdparty/go-sqlite3
git checkout v1.14.16

cd ../../
go mod tidy
go $1 -v 