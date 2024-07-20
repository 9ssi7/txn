// Copyright 2024 The tx Authors
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
// Package txn provides a generic framework for managing distributed transactions
// across multiple data sources in Go applications. It offers a simple and flexible way
// to coordinate transaction operations, ensuring data consistency and atomicity.
//
// Key Concepts:
//
//   - Tx: The core interface for transaction management. It provides methods to
//     begin, commit, rollback, and cancel transactions.
//   - Adapter: An interface that defines the contract for interacting with a specific
//     data source within a transaction. Each adapter is responsible for implementing
//     the necessary operations (begin, commit, rollback) for its respective data source.
package tx
