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

1. **Choose a Database Adapter:** Select the appropriate adapter package for your chosen database:
   * `txngorm`: For GORM (https://github.com/9ssi7/txngorm)
   * `txnmongo`: For MongoDB (https://github.com/9ssi7/txnmongo)
   * (Add more adapters as you implement them)

2. **Create a TxnManager:**

```go
import (
    "context"
    "github.com/9ssi7/txn"
    "[github.com/9ssi7/txngorm](https://github.com/9ssi7/txngorm)" // Or your chosen adapter
    "gorm.io/gorm"
)

// ... (Database connection setup)

repo := txngorm.NewGormRepo(db) // Use the appropriate adapter's repo
txManager := txn.NewTxnManager[gorm.DB](repo)
```

3. **Execute Transactions:**

```go
err := txManager.Transaction(context.Background(), func(txDB *gorm.DB) error {
    // Perform your database operations within this transaction
    // ...
    return nil // Return nil on success, or an error to trigger rollback
})

if err != nil {
    // Handle the transaction error
}
```


## Example (Using txngorm)

```go
// ... (assuming you have a User model and a GormRepo)

func UpdateUserWithTransaction(ctx context.Context, db *gorm.DB, userID uint, updates map[string]interface{}) error {
    repo := txngorm.NewGormRepo(db)
    txManager := txn.NewTxnManager[gorm.DB](repo)

    return txManager.Transaction(ctx, func(txDB *gorm.DB) error {
        // Transactions are performed safely here
        if err := txDB.Model(&User{ID: userID}).Updates(updates).Error; err != nil {
            return err
        }
        return nil // Return nil if there is no error
    })
}
```

## Contributing

Contributions are welcome! Please feel free to submit issues, bug reports, or pull requests.

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.