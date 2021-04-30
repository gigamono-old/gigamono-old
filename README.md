[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fgigamono%2Fgigamono.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fgigamono%2Fgigamono?ref=badge_shield)

<div align="center">
    <a href="#" target="_blank">
        <img src="https://raw.githubusercontent.com/gigamono/gigamono/main/media/logo.png" alt="Gigamono Logo" width="140" height="140"></img>
    </a>
</div>

<h1 align="center">GIGAMONO</h1>

<p align="center">
:warning:  This project is experimental and in active development  :warning:
</p>

### SETTING UP PROJECT

```sh
git clone --recursive github.com/gigamono/gigamono
```

### GRAPHQL GENERATION

Build the custom gqlgen binary.

```sh
go build cmd/gqlgen/gqlgen.go
```

Add the binary to system path and run command in a directory with gqlgen.yml file.

```sh
gqlgen
```

### PROTO GENERATION

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


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fgigamono%2Fgigamono.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fgigamono%2Fgigamono?ref=badge_large)