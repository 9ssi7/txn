// Copyright 2024 The txn Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Package txn provides a generic framework for managing database transactions in Go applications.
// By leveraging Go's generics, this package enables a clean and database-agnostic approach to
// transaction handling, making it ideal for architectures like Clean Architecture or Hexagonal Architecture.
//
// Key Benefits:
//
//   - Database Independence: The `Txn` interface abstracts away the underlying database technology,
//     allowing you to switch between different databases (e.g., PostgreSQL, MongoDB) without
//     modifying your core business logic.
//   - Clean Architecture: This package encourages a clear separation of concerns, keeping your
//     business logic clean and decoupled from database-specific implementation details.
//   - Data Consistency: The transaction mechanism ensures that a series of database operations are
//     executed atomically, maintaining data consistency even in the face of errors or failures.
//   - Flexibility: The generic `Txn` interface can be easily adapted to work with various database
//     drivers or ORMs by implementing the provided interfaces.
//
// Example (Using txngorm):
//
//	func UpdateUserWithTransaction(ctx context.Context, db *gorm.DB, userID uint, updates map[string]interface{}) error {
//	    repo := txngorm.NewGormRepo(db)
//	    txManager := txn.NewTxnManager[gorm.DB](repo)
//
//	    return txManager.Transaction(ctx, func(txDB *gorm.DB) error {
//	        if err := txDB.Model(&User{ID: userID}).Updates(updates).Error; err != nil {
//	            return err
//	        }
//	        return nil // Return nil if there is no error
//	    })
//	}
//
// Refer to the documentation of the specific adapter package you choose for more detailed examples and usage instructions.
package txn
