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

### BUILDING BINARIES

- GQLGen
  Build the custom gqlgen binary.

  ```go
  go build cmd/gqlgen/gqlgen.go
  ```

  Add the binary to system path and run command in a directory with gqlgen.yml file.

  ```sh
  gqlgen
  ```

### DATABASE

```sh
go run cmd/migrator/migrator.go -up
go run cmd/migrator/migrator.go -down
go run cmd/migrator/migrator.go -up-to 2
go run cmd/migrator/migrator.go -down-to 1
```

```sh
go run cmd/seeder/seeder.go -add-all
go run cmd/seeder/seeder.go -remove-all
go run cmd/seeder/seeder.go -add users
go run cmd/seeder/seeder.go -remove users
```

When you add a model, make sure to update the copy of its initial state in `/migrations/1_initial_tables.go` file
