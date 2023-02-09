#! /bin/sh

mkdir -p ./build

cd client

yarn build

rm -rf ../build/*

mv dist/* ../build/

cd ..

go build .

chmod +x tracka

./tracka
