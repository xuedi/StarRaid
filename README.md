
### Install ###
cd assets
cp config.init.in config.ini
./seed.bash

### run ###


### generate api ###
protoc --proto_path=api/proto/ --go_out=. api/proto/*.proto