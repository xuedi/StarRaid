
# Install
 - cd assets
 - cp config.init.in config.ini
 - ./seed.bash

## re-generate api
 - protoc --proto_path=api/ --go_out=. --go-grpc_out=. api/api.proto