// Copyright 2024 The txnredis Authors
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
// Package txnredis provides a transaction adapter for Redis, seamlessly integrating with the txn package
// to enable distributed transaction management across multiple data sources, including Redis.
//
// Key Features:
// * Transactional Support:  Leverages Redis transactions (MULTI/EXEC) to ensure atomicity of operations.
// * Seamless Integration:  Easily integrates with the txn package for managing transactions across different databases.
// * Convenient API:  Provides a simple interface (RAdapter) for interacting with Redis within a transaction.
package txnredis