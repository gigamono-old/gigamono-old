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

### DATABASE

##### Migration

??

##### Seeding

Build the seeder binary.

```sh
go build cmd/seeder/seeder.go
```

Seed database or roll back.

```sh
./seeder -k auth -add-all
./seeder -k auth -remove-all
./seeder -k auth -add users
./seeder -k resource -remove users
```
