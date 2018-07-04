#!/usr/bin/env bash

APP="example"
SHAPP="statup_plugin"
REPO="hunterlong/statup-plugin"

bash <(curl -s https://assets.statup.io/install.sh)

rm -rf statup
mkdir statup
mkdir statup/plugins

# BUILD STATUP GOLANG BINS
rm -rf build
mkdir build
xgo -go 1.10.x -buildmode=plugin --targets=darwin/amd64 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x -buildmode=plugin --targets=darwin/386 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x -buildmode=plugin --targets=linux/amd64 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x -buildmode=plugin --targets=linux/386 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x -buildmode=plugin --targets=linux/arm-7 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x -buildmode=plugin --targets=linux/arm64 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x -buildmode=plugin --targets=linux/amd64 --dest=build -ldflags="-X main.VERSION=$VERSION -linkmode external -extldflags -static" -out alpine ./

cd build
ls

mv alpine-linux-amd64 $APP.so
tar -czvf $APP-linux-alpine.tar.gz $APP.so && rm -f $APP.so

mv $SHAPP-darwin-10.6-amd64 $APP.so
tar -czvf $APP-osx-x64.tar.gz $APP.so && mv $APP.so ../statup/plugins/$APP.so

mv $SHAPP-darwin-10.6-386 $APP.so
tar -czvf $APP-osx-x32.tar.gz $APP.so && rm -f $APP.so

mv $SHAPP-linux-amd64 $APP.so
tar -czvf $APP-linux-x64.tar.gz $APP.so && rm -f $APP.so

mv $SHAPP-linux-386 $APP.so
tar -czvf $APP-linux-x32.tar.gz $APP.so && rm -f $APP.so

mv $SHAPP-linux-arm-7 $APP.so
tar -czvf $APP-linux-arm7.tar.gz $APP.so && rm -f $APP.so

mv $SHAPP-linux-arm64 $APP.so
tar -czvf $APP-linux-arm64.tar.gz $APP && rm -f $APP.so


cd ../statup

statup test plugins

