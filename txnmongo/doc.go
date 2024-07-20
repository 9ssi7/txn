// Copyright 2024 The txnmongo Authors
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

// Package txnmongo provides a transaction management adapter designed to be used
// with the official MongoDB Go driver. This package implements the generic
// transaction interface (Txn) from the 9ssi7/txn package, leveraging MongoDB's
// transaction capabilities.
//
// Standard usage:
//
//		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
//		if err != nil {
//		    panic(err)
//		}
//
//		txManager := txnmongo.NewMongo(client)
//
//		repo := sample.NewCustomRepo(client)
//		txManager.Add(func(sc mongo.SessionContext) error {
//		    repo.WithTxn(sc)
//		    return repo.GetById(context.Background(), "some_id")
//		})
//		if err := txManager.Transaction(context.Background()); err != nil {
//	     	repo.ClearTxn()
//		    log.Fatalf("Error: %v", err)
//		}
package txnmongo
