# go-mservice-bench

golang API (Gin) + Redis (go-redis) + golang Workers + Sqlite (Gorm)

## Skills

- Test Gin
- Test Gorm + Sqlite
- Message broker based on Redis
- Dotenv config
- DI pattern
- Unit tests

## How to use it

```sh
# clone the project
# go to root directory
mkdir .data tmp
go build -o tmp/goms
make start-redis
./tmp/goms -x server
# stop with ctrl+c
./tmp/goms -x worker
# stop with ctrl+c
DATA=account make post
DATA=account make post
ID=1 DATA=account make patch
COUNT=100 CONCURRENT=100 make attack
# stop with ctrl+c
make stop-redis
```
