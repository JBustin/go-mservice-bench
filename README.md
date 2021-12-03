# go-mservice-bench

golang API + Redis + golang Workers + Postgresql (or redis)

# Bench

Client:

- id
- firstName
- lastName

Account:

- uid
- client uid
- amount

Transaction:

- to account uid
- from account uid
- amount
- type (credit, debit)

Handle large volume of transactions.

API:

- post new transation
- get account amount by account uid
- get client total transactions + total debits + total credits + total amount by client uid
- get all client total transactions + total amount

# Architecture

- Server API restful (Gin)
- Message Broker (Redis)
- Worker(s)
  - validate transactions
  - create transactions in db
  - update accounts amounts in db
- Postgresql
  - table clients (#id, firstName, lastName)
  - table accounts (#id, clientId, amount)
  - table transactions (#id, toId, fromId, type, amount)
