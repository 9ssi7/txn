// Copyright 2024 The txngorm Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package txngorm provides a transaction (Tx) management adapter designed
// to be used with Go's popular ORM library, GORM. This package implements
// the generic transaction interface (Txn) from the 9ssi7/txn package, leveraging
// GORM's transaction capabilities.
//
// Standard usage:
//
//	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
//	if err != nil {
//	    panic("failed to connect database")
//	}
//
//	txManager := txngorm.NewGorm(db)
//
//	repo := sample.NewCustomRepo(db)
//	txManager.Add(func(txDB *gorm.DB) error {
//	    repo.WithTxn(txDB)
//	    return repo.GetById("some_id")
//	})
//	if err := txManager.Transaction(context.Background()); err != nil {
//	    log.Fatalf("Error: %v", err)
//	}
package txngorm
