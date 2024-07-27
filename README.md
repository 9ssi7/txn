# txn: Generic Distributed Transaction Management for Go

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![GoDoc](https://godoc.org/github.com/9ssi7/txn?status.svg)](https://pkg.go.dev/github.com/9ssi7/txn)
![Project status](https://img.shields.io/badge/version-1.0.2-green.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/9ssi7/txn)](https://goreportcard.com/report/github.com/9ssi7/txn)

The `txn` package provides a robust and flexible framework for managing distributed transactions across multiple data sources in your Go applications. By harnessing the power of Go generics, `txn` enables a clean, database-agnostic approach to transaction handling, ensuring data consistency and atomicity even in complex, distributed environments.

## Before of all, Check the [Real World Example](https://github.com/9ssi7/teknasyon.banking/blob/main/apps/banking/internal/app/commands/money_transfer.go#L41)

## Key Features

* **Distributed Transactions:** Coordinate transactions across multiple data sources seamlessly.
* **Database Independence:** Work with various databases (PostgreSQL, MongoDB etc.) using specialized adapters.
* **Clean Architecture:** Maintain a clear separation of concerns, keeping your business logic decoupled from data access details.
* **Atomicity:** Ensure that all operations within a transaction either succeed or fail together, maintaining data integrity.
* **Flexibility:** Easily extend the framework by creating custom adapters for your specific data sources.

## Installation

```bash
go get github.com/9ssi7/txn

go get github.com/9ssi7/txn/txngorm // For GORM Adapter
go get github.com/9ssi7/txn/txnmongo // For MongoDB Adapter
```

## Usage

1. **Create a Tx Instance:**

```go
tx := txn.New()
```

2. **Register Adapters:**

```go
gormAdapter := txngorm.New(gormDB)
tx.Register(gormAdapter)

mongoAdapter := txnmongo.New(mongoClient)
tx.Register(mongoAdapter)

// Register more adapters as needed...
```

3. **Manage Transactions:**

```go
err := tx.Begin(context.Background())
if err != nil {
    // Handle error
}
defer tx.End(context.Background()) // Ensure resources are cleaned up

// Perform operations on each data source using their respective adapters
// ...

if err := tx.Commit(context.Background()); err != nil {
   tx.Rollback(context.Background())
    // Handle commit error
}
```

## Adapters

The `txn` package supports multiple database adapters:

* **txngorm:** [GORM](./txngorm) 
* **txnmongo:** [MongoDB](./txnmongo)

## Contributing

Contributions are welcome! Please feel free to submit issues, bug reports, or pull requests.

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.
