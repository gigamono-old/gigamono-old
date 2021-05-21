<p align="center">
:warning:  The instructions below are not final. They change often and may be outdated  :warning:
</p>

### CONFIGURATION

Add a `gigamono.yaml` file. See [here](#) for instructions.

### DOCKER SETUP

Clone the following repos under the same directory.

```sh
git clone github.com/gigamono/gigamono
git clone github.com/gigamono/gigamono-api
git clone github.com/gigamono/gigamono-auth
```

Follow directions in respective repo to set the up properly.

Change directory.

```sh
cd gigamono
```

Add a `.env.docker` file following the template `.env.docker.sample`

Run the following docker-compose command.

```sh
docker-compose --env-file ./.env.docker -f docker/compose/local.yaml up
```

### GRAPHQL GENERATION

Build the custom gqlgen binary.

```sh
go build cmd/gqlgen/gqlgen.go
```

Add the binary to system path and run command in a directory with gqlgen.yml file.

Run command in project root directory.

```sh
gqlgen
```

### PROTO GENERATION

Make sure $GOPATH is set to $HOME/go

```fish
set -gx $GOPATH $HOME/go
```

Make sure $GOPATH/bin/ is added to $PATH

```fish
set -gx $PATH $GOPATH/bin $PATH
```

Install protoc

```sh
brew install protobuf
```

Install protoc-gen-go

```sh
go get github.com/golang/protobuf/protoc-gen-go
```

Install protoc-gen-govalidators plugin globally

```sh
GO111MODULE="off" go get github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
```

Compile proto files with this command.

```sh
protoc \
	--proto_path=./pkg/services/proto \
	--proto_path=$GOPATH/src \
	--go_out=plugins=grpc:./pkg/services/proto/generated \
	--govalidators_out=./pkg/services/proto/generated \
	--go_opt=paths=source_relative \
	--govalidators_opt=paths=source_relative \
	$(find ./pkg/services/proto -iname "*.proto")
```

### DATABASE

Install postgres

Create postgres databases for the Resource database and the Auth database. For example:

```sql
CREATE DATABASE "resourcedb";
CREATE DATABASE "authdb";
```

Add extensions:

```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "citext";
```

Add a `.env` files in respective following the template `.env.docker.sample`

##### Migration

Install goose

```
go get -u github.com/pressly/goose/cmd/goose
```

Or follow the instructions [here](https://github.com/pressly/goose#install)

Run `goose` command in a directory with migrations or specify with seed directory with `--dir` flag

```sh
goose postgres "postgres://appcypher@localhost:5432/resourcedb?sslmode=disable" up
goose postgres "postgres://appcypher@localhost:5432/resourcedb?sslmode=disable" reset
```

Check [here](https://github.com/pressly/goose) for more instructions on how to use goose.

##### Seeding

Clone pgseeder repo and cd into the created folder

```sh
git clone https://girhub.com/gigamono/pgseeder
cd pgseeder
```

Build the binary.

```sh
go build cmd/pgseeder.go
```

Add the binary to system path and run `pgseeder` command in a directory with seeds or specify with seed directory with `-d` flag

```sh
pgseeder -c "postgres://appcypher@localhost:5432/resourcedb?sslmode=disable" --add users
pgseeder -c "postgres://appcypher@localhost:5432/resourcedb?sslmode=disable" --add-all -d internal/db/seeds/resource
pgseeder -c "postgres://appcypher@localhost:5432/resourcedb?sslmode=disable" --remove-all
```

Check [here](https://github.com/gigamono/pgseeder) for more instructions on how to use goose.

### GENERATING KEYS

Generate private-public key pair with openssl

```sh
openssl ecparam -genkey -name secp521r1 -noout -out private.pem
openssl ec -in private.pem -pubout -out public.pem
```

### INSTALLATION

##### Simple Install [WIP]

```sh
curl https://www.gigamono.com/get -sSfL | sh
```

...

##### Simple Install [WIP]

```sh
git clone github.com/gigamono/gigamono
```

```sh
cd gigamono
```

```sh
sh ./install.sh
```

...
