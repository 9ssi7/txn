# txn: Generic Distributed Transaction Management for Go

[![GoDoc](https://godoc.org/github.com/9ssi7/txn?status.svg)](https://pkg.go.dev/github.com/9ssi7/txn)
![Project status](https://img.shields.io/badge/version-1.0.2-green.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/9ssi7/txn)](https://goreportcard.com/report/github.com/9ssi7/txn)

The `txn` package provides a robust and flexible framework for managing distributed transactions across multiple data sources in your Go applications. By harnessing the power of Go generics, `txn` enables a clean, database-agnostic approach to transaction handling, ensuring data consistency and atomicity even in complex, distributed environments.

## Before of all, Check the [Real World Example](https://github.com/9ssi7/teknasyon.banking/blob/main/apps/banking/internal/app/commands/money_transfer.go#L41)

## Key Features

* **Distributed Transactions:** Coordinate transactions across multiple data sources seamlessly.
* **Clean Architecture:** Maintain a clear separation of concerns, keeping your business logic decoupled from data access details.
* **Atomicity:** Ensure that all operations within a transaction either succeed or fail together, maintaining data integrity.
* **Flexibility:** Easily extend the framework by creating custom adapters for your specific data sources.

Note:
Database independency not possible with this package. You need to use the same database for all adapters. For example, if you are using GORM, you need to use GORM for all adapters. If you are using MongoDB, you need to use MongoDB for all adapters. If GORM throws an error but MongoDB doesn't, you need to handle it yourself. This package doesn't handle that. It only provides a way to manage transactions across multiple data sources.

## Installation

```bash
go get github.com/9ssi7/txn

go get github.com/9ssi7/txn/txngorm // For GORM Adapter
go get github.com/9ssi7/txn/txnmongo // For MongoDB Adapter
go get github.com/9ssi7/txn/txnsql // For Native SQL Adapter
```

## Usage

1. **Create a Tx Instance:**

```go
tx := txn.New()
```

2. **Register Adapters:**

```go
tx.Register(txngorm.New(gormDB), txnmongo.New(mongoClient), txnsql.New(sqlDB))

// Register more adapters as needed...
```

3. **Manage Transactions:**

```go
func transaction() (err error) {
    defer func() {
        if err != nil {
            tx.Rollback(context.Background())
        }
    }()
    if err := tx.Begin(context.Background()); err != nil {
        return err
    }
    // Perform operations on each data source using their respective adapters
    // ...

    if err := tx.Commit(context.Background()); err != nil {
    return err
    }
    return nil
}


```

## Adapters

The `txn` package supports multiple database adapters:

* **txngorm:** [GORM](./txngorm) 
* **txnmongo:** [MongoDB](./txnmongo)
* **txnsql:** [Native SQL](./txnsql)

## Contributing

Contributions are welcome! Please feel free to submit issues, bug reports, or pull requests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
