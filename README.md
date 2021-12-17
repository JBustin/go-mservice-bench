# go-mservice-bench

golang API (Gin) + Redis (go-redis) + golang Workers + Sqlite (Gorm)

## Skills

- Test Gin
- Test Gorm + Sqlite
- Message broker based on Redis
- Dotenv config
- DI pattern
- Unit tests
- Real load tests

## How to use it

```sh
# clone the project
# go to root directory
mkdir .data tmp
go build -o tmp/goms
make start-redis
# stop with Ctrl+C
./tmp/goms -x server
# stop with Ctrl+C
./tmp/goms -x worker
DATA=account make post
DATA=account make post
ID=1 DATA=account make patch
# stop with Ctrl+C
COUNT=100 CONCURRENT=100 make attack
make stop-redis
```
