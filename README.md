<div align="center">
    <a href="#" target="_blank">
        <img src="https://raw.githubusercontent.com/sageflow/sageflow/main/media/logo.png" alt="Sageflow Logo" width="140" height="140"></img>
    </a>
</div>

<h1 align="center">SAGEFLOW</h1>

<p align="center">
:warning:  This project is experimental and in active development  :warning:
</p>

### SETTING UP PROJECT

```sh
git clone --recursive github.com/sageflow/sageflow
```

### DATABASE

```go
go run cmd/migrator/migrator.go -up
go run cmd/migrator/migrator.go -down
go run cmd/migrator/migrator.go -up-to 2
go run cmd/migrator/migrator.go -down-to 1
```

```go
go run cmd/seeder/seeder.go -add-all
go run cmd/seeder/seeder.go -remove-all
go run cmd/seeder/seeder.go -add users
go run cmd/seeder/seeder.go -remove users
```

When you add a model, make sure to update the copy of its initial state in `/migrations/1_initial_tables.go` file

