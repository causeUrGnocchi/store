## Requirements

- go
- docker

## Run Locally

Clone the project

```bash
  git clone git@github.com:causeUrGnocchi/store.git
```

Go to the project directory

```bash
  cd store
```

Start the database

```bash
  docker compose up
```

Start the server (Available at port 8080 by default)

```bash
  go run main.go
```
