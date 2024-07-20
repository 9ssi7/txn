# txn: Generic Transaction Management for Go

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

The `txn` package provides a powerful and flexible framework for managing database transactions in your Go applications. By harnessing the capabilities of Go generics, `txn` enables a clean, database-agnostic approach to transaction handling, making it a perfect fit for modern architectures like Clean Architecture and Hexagonal Architecture.

## Key Features

* **Database Independence:** Seamlessly switch between different database technologies (PostgreSQL, MongoDB, etc.) without altering your core business logic.
* **Clean Architecture:** Maintain a clear separation of concerns, keeping your business logic pristine and decoupled from database-specific implementation details.
* **Data Consistency:** Ensure data integrity by executing multiple database operations atomically within transactions.
* **Flexibility:** Easily adapt the `txn` framework to work with various database drivers or ORMs by implementing the provided interfaces.

## Installation

```bash
go get github.com/9ssi7/txn
```

## Usage

1. **Before Use, Choose a Database Adapter:** Select the appropriate adapter package for your chosen database:
   * `txngorm`: For GORM (https://github.com/9ssi7/txn/txngorm)
   * `txnmongo`: For MongoDB (https://github.com/9ssi7/txn/txnmongo)
   * (Add more adapters as you implement them)


## Example (Using txngorm)

This example in the `examples` directory demonstrates how to use the `txn` package with GORM.

Included examples:

- [`Basic Gorm Example`](examples/basic-gorm/main.go): Demonstrates basic usage of the `txn` package with GORM.
- [`Basic Mongo Example`](examples/basic-mongo/main.go): Demonstrates basic usage of the `txn` package with MongoDB.

## Contributing

Contributions are welcome! Please feel free to submit issues, bug reports, or pull requests.

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.