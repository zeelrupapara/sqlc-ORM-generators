## sqlc
Tutorials for how to use sqlc for your golang project with a structured way

### Installation
Arch Linux
```bash
sudo pacman -S sqlc
```
Golang
```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

### How to use the sqlc tool?
- For use, the sqlc needs to create a table schema for the database (to generate structs)
- Next, You need to write a query to generate a required function that can do some action 📍(to generate a model function)
- Finally, need to create sqlc configuration file called `sqlc.yaml`

After doing all these things,
execute,
```bash
sqlc generate
```
> - In this repo, I do the testing with my structure. To create an API using that generated model function,
> - Also include migration and many things in one pack

#### Help Menu
```bash
go run main.go
```
